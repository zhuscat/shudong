<div class="container">
<div class="left-area-wrapper">
<ul class="top-menu-bar">
<li class="bar-item {{if eq .Tab ""}}bar-item-selected{{end}}"><a href="/user/{{.OtherUser.Id}}">所有</a></li>
<li class="bar-item {{if eq .Tab "onsale"}}bar-item-selected{{end}}"><a href="/user/{{.OtherUser.Id}}/onsale">正在卖的</a></li>
<li class="bar-item {{if eq .Tab "out-of-stock"}}bar-item-selected{{end}}"><a href="/user/{{.OtherUser.Id}}/out-of-stock">已下架的</a></li>`
</ul>
{{range .Books}}
<div class="publish-book-div">
	<img src="/static/www/book/{{.Picture}}">
	<ul class="book-info">
		<li class="title"><a href="">{{.Title}}</a></li>
		<li class="price">￥{{.Price}}</li>
		<li>作者：{{.Author}}</li>
		<li>出版社: {{.Publisher}}</li>
		<li>ISBN: {{.Isbn}}</li>
	</ul>
	<ul class="edit">
		<li><a href="/book/detail/{{.Id}}" class="radius-button radius-button-small">查看</a></li>
	</ul>
	<div class="seprate"></div>
</div>
{{end}}
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
<div class="profile-div">
	<div class="profile-info">
	<div class="avatar-area">
		<img src="/static/www/avatar/{{.OtherUser.Avatar}}">
		<div class="basic-info">
			<p class="username">{{.OtherUser.Name}}</p>
			<p class="address">{{.OtherUser.Place}}</p>
		</div>
	</div>
	<ul>
		<li>邮箱：{{.OtherUser.Email}}</li>
		<li>手机号：{{.OtherUser.PhoneNumber}}</li>
		<li>QQ号：{{.OtherUser.Qq}}</li>
		<li>微信：{{.OtherUser.Weixin}}</li>
	</ul>
	<a class="radius-button radius-button-small" href="/">私信</a>
	</div>
</div>
</div>
