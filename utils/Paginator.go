package utils

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Paginator struct {
	Request     *http.Request
	PerPageNums int

	totalNums int64
	pageRange []int
	pageNums  int
	page      int
}

// 返回页数
func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}

	pageNums := math.Ceil(float64(p.totalNums) / float64(p.PerPageNums))

	p.pageNums = int(pageNums)

	return p.pageNums
}

func (p *Paginator) Nums() int64 {
	return p.totalNums
}

// 设置显示项目总的个数 如 所有商品的个数
func (p *Paginator) SetTotalNums(nums int64) {
	p.totalNums = nums
}

func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

// 此函数返回显示页数的信息 如显示的页数（按钮）为 6, 7, 8, 9 , 10
func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.totalNums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i, _ := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i, _ := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i, _ := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}

	fmt.Println("hahahaah")
	return p.pageRange
}

func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	// 取得URL后面?的东西
	// 转换成一个结构
	values := link.Query()
	// if page == 1 {
	// 	// 删除p
	// 	values.Del("p")
	// } else {
	// 	values.Set("p", strconv.Itoa(page))
	// }
	values.Set("p", strconv.Itoa(page))
	link.RawQuery = values.Encode()
	return link.String()
}

func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return
}

func (p *Paginator) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

func (p *Paginator) PageLinkLast() (link string) {
	return p.PageLink(p.PageNums())
}

func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

func NewPaginator(req *http.Request, per int, nums int64) *Paginator {
	p := Paginator{}
	p.Request = req
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetTotalNums(nums)
	return &p
}
