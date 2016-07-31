package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

// 对某一个商品的留言
type Comment struct {
	Id          int64
	CreatedTime time.Time
	BookId      int64
	UserId      int64
	VendorId    int64
	Content     string
}

func (self *Comment) User() *User {
	user, err := GetUser(self.UserId)
	if err != nil {
		return new(User)
	}
	return user
}

func NewComment() *Comment {
	return new(Comment)
}

func CommentAdd(comment *Comment) (int64, error) {
	if comment.BookId <= 0 {
		return 0, errors.New("书籍ID不能为空")
	}
	if comment.UserId <= 0 {
		return 0, errors.New("用户ID不能为空")
	}
	if comment.VendorId <= 0 {
		return 0, errors.New("卖家ID不能为空")
	}
	if comment.Content == "" {
		return 0, errors.New("评论内容不能为空")
	}
	comment.CreatedTime = time.Now()
	return orm.NewOrm().Insert(comment)
}

// 这里的total值得是全部的数量，不是limit之后的数量
func CommentsGetByBookId(id int64, args ...int) (comments []*Comment, total int64, err error) {
	qs := orm.NewOrm().QueryTable("comment").Filter("BookId", id).OrderBy("-created_time")
	total, err = qs.Count()
	if len(args) == 0 {
		_, err = qs.All(&comments)
		return
	} else if len(args) == 2 {
		_, err = qs.Limit(args[0], args[1]).All(&comments)
		return
	}
	err = errors.New("参数错误")
	return
}

// 添加一条评论
func AddComment(uid int64, bid int64, vid int64, content string) (int64, error) {
	var comment Comment
	comment.UserId = uid
	comment.BookId = bid
	comment.VendorId = vid
	comment.Content = content
	comment.CreatedTime = time.Now()
	return orm.NewOrm().Insert(&comment)
}

// 输入商品Id获取评论
func GetComments(pid int64, limit int, offset int) ([]*Comment, error) {
	var comments []*Comment
	_, err := orm.NewOrm().QueryTable("comment").Filter("BookId", pid).OrderBy("-created_time").Limit(limit, offset).All(&comments)
	return comments, err
}

func GetCommentCount(pid int64) (int64, error) {
	return orm.NewOrm().QueryTable("comment").Filter("BookId", pid).Count()
}

// CommentGetList 根据过滤条件获取评论
func CommentGetList(page int, pageSize int, content string, filters ...interface{}) ([]*Comment, int64) {
	offset := (page - 1) * pageSize
	var comments []*Comment
	query := orm.NewOrm().QueryTable("comment")
	if content != "" {
		query = query.Filter("content__contains", content)
	}
	if len(filters) >= 2 && len(filters)%2 == 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-CreatedTime").Limit(pageSize, offset).All(&comments)
	return comments, total
}

// CommentDelete 删除一条评论
func CommentDelete(id int64) (int64, error) {
	comment := &Comment{Id: id}
	return orm.NewOrm().Delete(comment)
}

// TotalComment 获取评论总数
func TotalComment() (int64, error) {
	return orm.NewOrm().QueryTable("comment").Count()
}
