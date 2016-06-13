<div class="container">
<div class="book-card-wrapper">
{{range .Books}}
<div class="book-card">
	<img src="/static/www/book/{{.Picture}}">
	<ul>
		<li><a class="title" href="/book/detail/{{.Id}}">{{.Title}}</a></li>
		<li class="price">￥{{.Price}}</li>
		<li>作者：{{.Author}}</li>
		<li>出版社：{{.Publisher}}</li>
	</ul>
	<div class="book-card-user">
		<a href="/user/{{.Vendor.Id}}">{{.Vendor.Name}}</a>
		<img src="/static/www/avatar/{{.Vendor.Avatar}}">
	</div>
	</div>
{{end}}
</div>
<!-- 分页 -->
{{if gt .Page.PageNums 1}}
<ul class="page">
	{{if .Page.HasPrev}}
		<li><a href="{{.Page.PageLinkFirst}}">第一页</a></li>
		<li><a href="{{.Page.PageLinkPrev}}">&lt;</a></li>
	{{else}}
		<li class="disabled"><a>第一页</a></li>
		<li class="disabled"><a>&lt;</a></li>
	{{end}}
	{{range $index, $page := .Page.Pages}}
		<li {{if $.Page.IsActive .}}class="active"{{end}}>
			<a href="{{$.Page.PageLink $page}}">{{$page}}</a>
		</li>
	{{end}}
	{{if .Page.HasNext}}
        <li><a href="{{.Page.PageLinkNext}}">&gt;</a></li>
        <li><a href="{{.Page.PageLinkLast}}">最后一页</a></li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>最后一页</a></li>
    {{end}}
</ul>
{{end}}
</div>