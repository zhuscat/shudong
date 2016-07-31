package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

// 卖的书（商品）
type Book struct {
	Id          int64
	Title       string
	Isbn        string
	Author      string
	Publisher   string
	Price       float64 // 设置一下digits, decimal
	Onsale      bool    `orm:"default(true)"`
	VendorId    int64
	Picture     string `orm:"default(default-book.png)"`
	Description string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func NewBook() *Book {
	book := new(Book)
	book.Onsale = true
	return book
}

func BookAdd(book *Book) (int64, error) {
	if book.Title == "" {
		return 0, errors.New("书名不能为空")
	}
	if book.Isbn == "" {
		return 0, errors.New("ISBN号不能为空")
	}
	if book.Author == "" {
		return 0, errors.New("作者不能为空")
	}
	if book.Price <= 0 {
		return 0, errors.New("价格不能为负数")
	}
	if book.VendorId <= 0 {
		return 0, errors.New("卖家ID不能为空")
	}
	book.CreatedTime = time.Now()
	book.UpdatedTime = time.Now()
	return orm.NewOrm().Insert(book)
}

func (self *Book) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(self, fields...)
	return err
}

// // 添加书籍
// func AddBook(title string, isbn string, author string, publisher string, price float64, pic string, vid int64, desc string) (int64, error) {
// 	var book Book
// 	book.Title = title
// 	book.Isbn = isbn
// 	book.Author = author
// 	book.Publisher = publisher
// 	book.Price = price
// 	book.CreatedTime = time.Now()
// 	book.UpdatedTime = time.Now()
// 	book.VendorId = vid
// 	book.Picture = pic
// 	book.Description = desc
// 	return orm.NewOrm().Insert(&book)
// }

// 当出现错误的时候返回一个空用户 这样好不好
func (self *Book) Vendor() *User {
	user, err := GetUser(self.VendorId)
	if err != nil {
		return new(User)
	} else {
		return user
	}
}

// 使用 Id 获取书籍模型
func GetBook(id int64) (*Book, error) {
	var book Book
	book.Id = id
	err := orm.NewOrm().Read(&book)
	if err == nil {
		return &book, nil
	} else {
		return nil, err
	}
}

// 获取许多书籍的信息 分页使用
func GetBooks(limit int, offset int) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").OrderBy("CreatedTime", "-CreatedTime").Limit(limit, offset).All(&books)
	return books, err
}

// 获取状态是上架的书籍的信息
func GetOnsaleBooks(limit int, offset int) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").Filter("onsale", true).OrderBy("-CreatedTime").Limit(limit, offset).All(&books)
	return books, err
}

// 获取书籍的总数
func GetBookCount() (int64, error) {
	return orm.NewOrm().QueryTable("book").Count()
}

// 获取已经上架的书籍的总数
func GetOnsaleBookCount() (int64, error) {
	return orm.NewOrm().QueryTable("book").Filter("onsale", true).Count()
}

// 获取书籍的总数，并给予一个已上架还是未上架的属性
func GetBookCountWithSaleAttr(onsale bool) (int64, error) {
	return orm.NewOrm().QueryTable("book").Filter("onsale", onsale).Count()
}

// 获取许多书籍的信息 分页使用
func GetBooksWithSaleAttr(onsale bool, limit int, offset int) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").Filter("onsale", onsale).OrderBy("CreatedTime", "-CreatedTime").Limit(limit, offset).All(&books)
	return books, err
}

// 搜索书籍
func FindBooks(content string, limit int, offset int) ([]*Book, error) {
	var books []*Book
	keyword := "%" + content + "%"
	_, err := orm.NewOrm().Raw("SELECT * from book where title like ? or author like ? or isbn like ? limit ? offset ?",
		keyword, keyword, keyword, limit, offset).QueryRows(&books)
	return books, err
}

// 获取搜索书籍数量
func GetBookCountWithContent(content string) (int64, error) {
	var books []*Book
	keyword := "%" + content + "%"
	num, err := orm.NewOrm().Raw("SELECT * from book where title like ? or author like ? or isbn like ?",
		keyword, keyword, keyword).QueryRows(&books)
	return num, err
}

// 根据用户名搜索
func FindBooksWithUserId(id int64, limit int, offset int) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").Filter("vendor_id", id).Limit(limit, offset).All(&books)
	return books, err
}

// 获取给定用户发布书的数量
func GetBookCountWithUserId(id int64) (int64, error) {
	return orm.NewOrm().QueryTable("book").Filter("vendor_id", id).Count()
}

// 获取给定用户及状态，其发布书的数量
func GetBookCountWithUserIdAndStatus(id int64, onsale bool) (int64, error) {
	return orm.NewOrm().QueryTable("book").Filter("vendor_id", id).Filter("onsale", onsale).Count()
}

// 根据用户名以及是否上架
func FindBooksWithUserIdAndStatus(id int64, onsale bool, limit int, offset int) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").Filter("vendor_id", id).Filter("onsale", onsale).Limit(limit, offset).All(&books)
	return books, err
}

// 该卖家还出售...
func GetRecommendBookWithUserId(uid int64, exceptBookId int64) ([]*Book, error) {
	var books []*Book
	_, err := orm.NewOrm().QueryTable("book").Filter("vendor_id", uid).Filter("onsale", true).Exclude("id", exceptBookId).OrderBy("-CreatedTime").Limit(4).All(&books)
	return books, err
}

// 更新书
func UpdateBook(book *Book) (int64, error) {
	return orm.NewOrm().Update(book)
}

// BookGetList 接受参数 page 页数 pageSize 一页最多的个数 filters 过滤条件
// 如: ["id", 1, "onsale", true]
// 返回 Book数组和Book的总量
// 给查找书籍列表一个统一的接口
func BookGetList(page int, pageSize int, filters ...interface{}) ([]*Book, int64) {
	offset := (page - 1) * pageSize
	var books []*Book
	query := orm.NewOrm().QueryTable("book")
	if len(filters) >= 2 && len(filters)%2 == 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-CreatedTime").Limit(pageSize, offset).All(&books)
	return books, total
}

// BookDelete 删除一本书籍
func BookDelete(book *Book) (int64, error) {
	return orm.NewOrm().Delete(book)
}
