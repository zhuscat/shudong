package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

const (
	CUSTOMER_REQUEST    = 0
	VENDOR_ACCEPT       = 1
	TRASACTION_COMPLETE = 2
	TRANSACTION_CLOSE   = 3
)

// 订单
// Status
// 0 买家提出了交易请求
// 1 卖家接收了交易请求
// 2 交易完成
// 3 交易关闭
type Order struct {
	Id          int64
	Status      int `orm:"default(0)"`
	VendorId    int64
	CustomerId  int64
	BookId      int64
	Price       float64 // 设置一下digits, decimal
	Address     string
	CreatedDate time.Time
}

func GetOrder(id int64) (*Order, error) {
	order := Order{Id: id}
	err := orm.NewOrm().Read(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// 传入-1获取所有所有订单的个数
func GetOrderCount(status int, vid int64, cid int64) int64 {
	qs := orm.NewOrm().QueryTable("Order")
	if status >= 0 && status <= 3 {
		qs = qs.Filter("status", status)
	}
	if vid != 0 {
		qs = qs.Filter("vendor_id", vid)
	}
	if cid != 0 {
		qs = qs.Filter("customer_id", cid)
	}
	count, err := qs.Count()
	if err != nil {
		return 0
	}
	return count
}

func (self *Order) Vendor() *User {
	user, err := GetUser(self.VendorId)
	if err != nil {
		return new(User)
	}
	return user
}

func (self *Order) Customer() *User {
	user, err := GetUser(self.CustomerId)
	if err != nil {
		return new(User)
	}
	return user
}

func (self *Order) Book() *Book {
	book, err := GetBook(self.BookId)
	if err != nil {
		return new(Book)
	}
	return book
}

func FindOrderWithVendorId(vid int64, limit int, offset int) ([]*Order, error) {
	var orders []*Order
	_, err := orm.NewOrm().QueryTable("order").Filter("vendor_id", vid).OrderBy("-CreatedDate").Limit(limit, offset).All(&orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func FindOrderWithVendorIdAndStatus(vid int64, status int, limit int, offset int) ([]*Order, error) {
	var orders []*Order
	_, err := orm.NewOrm().QueryTable("order").Filter("vendor_id", vid).Filter("status", status).OrderBy("-CreatedDate").Limit(limit, offset).All(&orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func FindOrderWithCustomerId(cid int64, limit int, offset int) ([]*Order, error) {
	var orders []*Order
	_, err := orm.NewOrm().QueryTable("order").Filter("customer_id", cid).OrderBy("-CreatedDate").Limit(limit, offset).All(&orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func FindOrderWithCustomerIdAndStatus(cid int64, status int, limit int, offset int) ([]*Order, error) {
	var orders []*Order
	_, err := orm.NewOrm().QueryTable("order").Filter("customer_id", cid).Filter("status", status).OrderBy("-CreatedDate").Limit(limit, offset).All(&orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// 买家添加一个订单，卖家添加订单之后，会给卖家发送一条站内信
func AddOrder(bid int64, cid int64, addr string) (int64, error) {
	o := orm.NewOrm()
	var order Order
	// 查找 book 是否存在
	book, err := GetBook(bid)
	if book == nil {
		return 0, err
	}
	// 查找用户存在不存在
	customer, err := GetUser(cid)
	if customer == nil {
		return 0, err
	}

	order.BookId = bid
	order.Status = 0
	order.VendorId = book.VendorId
	order.CustomerId = cid
	order.Price = book.Price
	order.CreatedDate = time.Now()
	order.Address = addr

	o.Begin()
	num, oerr := o.Insert(&order)
	content := "买家" + customer.Name + "给你的书籍 " + book.Title + " 下订单了，请尽快与买家联系"
	message := Message{UserId: book.VendorId, Content: content, SendTime: time.Now()}
	_, merr := o.Insert(&message)
	if oerr != nil || merr != nil {
		return 0, o.Rollback()
	} else {
		return num, o.Commit()
	}
}

// 卖家确认订单，卖家确认订单之后，会给买家发送一条站内信
// 此外，卖家确认订单之后会自动把书置为下架状态
func ConfirmOrder(oid, uid int64) error {
	order, _ := GetOrder(oid)
	if order != nil {
		if order.VendorId == uid && order.Status == CUSTOMER_REQUEST {
			order.Status = VENDOR_ACCEPT
			o := orm.NewOrm()
			o.Begin()
			_, oerr := o.Update(order, "Status")
			content := "你的订单 " + strconv.Itoa(int(order.Id)) + " 已经被确认了，请积极与卖家联系"
			message := Message{UserId: order.CustomerId, Content: content, SendTime: time.Now()}
			_, merr := o.Insert(&message)
			book, berr := GetBook(order.BookId)
			book.Onsale = false
			_, berr = UpdateBook(book)
			if oerr != nil || merr != nil || berr != nil {
				o.Rollback()
			} else {
				o.Commit()
			}
		} else {
			return LOGICAL_ERR
		}
	}
	return NOTFOUND_ERR
}

// 买家或者卖家在卖家未确认的情况下关闭订单
func CloseOrder(oid, uid int64) error {
	order, _ := GetOrder(oid)
	if order != nil {
		if (order.VendorId == uid || order.CustomerId == uid) && order.Status == CUSTOMER_REQUEST {
			order.Status = TRANSACTION_CLOSE
			o := orm.NewOrm()
			o.Begin()
			_, oerr := o.Update(order, "Status")
			content := "你的订单" + strconv.Itoa(int(order.Id)) + "被关闭了"
			var receiver int64
			if order.VendorId == uid {
				receiver = order.CustomerId
			} else {
				receiver = order.VendorId
			}
			message := Message{UserId: receiver, Content: content, SendTime: time.Now()}
			_, merr := o.Insert(&message)
			if oerr != nil || merr != nil {
				o.Rollback()
			} else {
				o.Commit()
			}
		} else {
			return LOGICAL_ERR
		}
	}
	return NOTFOUND_ERR
}

// 买家收到货物以后，确认订单，交易完成
func CompeleteOrder(oid, uid int64) error {
	order, _ := GetOrder(oid)
	if order != nil {
		if order.CustomerId == uid && order.Status == VENDOR_ACCEPT {
			order.Status = TRASACTION_COMPLETE
			o := orm.NewOrm()
			o.Begin()
			_, oerr := o.Update(order, "Status")
			content := "你的订单 " + strconv.Itoa(int(order.Id)) + " 已经完成"
			message := Message{UserId: order.VendorId, Content: content, SendTime: time.Now()}
			_, merr := o.Insert(&message)
			if oerr != nil || merr != nil {
				o.Rollback()
			} else {
				o.Commit()
			}
		} else {
			return LOGICAL_ERR
		}
	}
	return NOTFOUND_ERR
}
