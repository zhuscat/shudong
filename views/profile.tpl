<div class="container">
<div class="left-area-wrapper">
<ul class="top-menu-bar">
{{str2html .Menubar}}
<!-- <li class="search-item">
	<form>
		<i class="fa fa-search"></i><input type="text" name="wd" placeholder="搜索你发布的书">
	</form>
</li> -->
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
		<li><a href="/book/edit/{{.Id}}" class="radius-button radius-button-small">编辑</a></li>
		{{if eq .Onsale true}}
		<li><a href="/book/change/{{.Id}}" class="radius-button radius-button-small change-btn" data-bookid="{{.Id}}">下架</a></li>
		{{else}}
		<li><a href="/book/change/{{.Id}}" class="radius-button radius-button-small change-btn" data-bookid="{{.Id}}">上架</a></li>
		{{end}}
		<li><a href="" class="radius-button radius-button-small radius-button-warning">删除</a></li>
	</ul>
	<div class="seprate"></div>
</div>
{{end}}
{{range .Orders}}
<div class="publish-book-div order-attr" data-orderid="{{.Id}}">
<div class="publish-book-top-wrapper">
	<img src="/static/www/book/{{.Book.Picture}}">
	<ul class="book-info">
		<li class="title"><a href="">{{.Book.Title}}</a></li>
		<li>订单号：{{.Id}}</li>
		<li>成交价：{{.Price}}元</li>
		<li>下单日期：{{.CreatedDate}}</li>
		<li>ISBN: {{.Book.Isbn}}</li>
	</ul>
	<ul class="edit">
		{{if eq .Vendor.Id $.User.Id}}
			{{if eq .Status 0}}
			<li><a href="javascript:void(0)" class="manage-order-button radius-button radius-button-small" data-orderid="{{.Id}}" data-url="/order/accept">接受</a></li>
			<li><a href="javascript:void(0)" class="manage-order-button radius-button radius-button-small" data-orderid="{{.Id}}" data-url="/order/close">关闭</a></li>
			{{end}}
			{{if eq .Status 1}}
			<li><a href="javascript:void(0)" class="radius-button radius-button-small radius-button-disable">已接受</a></li>
			{{end}}
			{{if eq .Status 2}}
			<li><a href="javascript:void(0)" class="radius-button radius-button-small radius-button-disable">已完成</a></li>
			{{end}}
			{{if eq .Status 3}}
			<li><a href="javascript:void(0)" class="radius-button radius-button-small radius-button-disable">已关闭</a></li>
			{{end}}
		{{else}}
			{{if eq .Status 0}}
			<li><a href="javascript:void(0)" class="manage-order-button radius-button radius-button-small" data-orderid="{{.Id}}" data-url="/order/close">关闭</a></li>
			{{end}}
			{{if eq .Status 1}}
			<li><a href="javascript:void(0)" class="manage-order-button radius-button radius-button-small" data-orderid="{{.Id}}" data-url="/order/complete">完成</a></li>
			{{end}}
			{{if eq .Status 2}}
			<li><a href="javascript:void(0)" class="radius-button radius-button-small radius-button-disable">已完成</a></li>
			{{end}}
			{{if eq .Status 3}}
			<li><a href="javascript:void(0)" class="radius-button radius-button-small radius-button-disable">已关闭</a></li>
			{{end}}
		{{end}}
	</ul>
	<div class="seprate"></div>
</div>
	<div class="bottom-info">
	{{if eq .Vendor.Id $.User.Id}}
		<img src="/static/www/avatar/{{.Customer.Avatar}}">
		<a href="/user/{{.Customer.Id}}">{{.Customer.Name}}</a>
	{{else}}
		<img src="/static/www/avatar/{{.Vendor.Avatar}}">
		<a href="/user/{{.Vendor.Id}}">{{.Vendor.Name}}</a>
	{{end}}
	</div>
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
		<img src="/static/www/avatar/{{.User.Avatar}}">
		<div class="basic-info">
			<p class="username">{{.User.Name}}</p>
			<p class="address">{{.User.Place}}</p>
		</div>
	</div>
	<ul>
		<li>邮箱：{{.User.Email}}</li>
		<li>手机号：{{.User.PhoneNumber}}</li>
		<li>QQ号：{{.User.Qq}}</li>
		<li>微信：{{.User.Weixin}}</li>
	</ul>
	<a class="radius-button radius-button-small" href="/edit-profile">编辑</a>
	{{if not .User.Active}}
	<a class="radius-button radius-button-green radius-button-small" href="/user/active">激活</a>
	{{end}}
	</div>
	<ul class="menu-area">
	{{str2html .RightMenubar}}
	</ul>
</div>
</div>
<script type="text/javascript">
	$(".manage-order-button").click(function() {
		$.ajax({
			url: $(this).data("url"),
			type: "POST",
			dataType: "json",
			data: {orderid: $(this).data("orderid")},
			success: function(data) {
				// 现在只是简单的重载
				// 以后再改吧
				location.reload();
			}
		})
	});
	$(".change-btn").click(function(e) {
		$btn = $(this);
		$.ajax({
			url: $btn.attr("href"),
			type: "GET",
			dataType: "json",
			success: function(data) {
				if (data["success"] == true) {
					if ($btn.html() == "上架") {
						$btn.html("下架");
					} else {
						$btn.html("上架");
					}
				}
				alert(data["msg"]);
			}
		});
		e.preventDefault();
	});
</script>
