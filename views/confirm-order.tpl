<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>个人信息</title>
	<link rel="stylesheet" type="text/css" href="/static/css/font-awesome.css">
	<link rel="stylesheet" type="text/css" href="/static/css/main.css">
	<script type="text/javascript" src="/static/js/jquery-2.2.4.js"></script>
</head>
<body>
<div class="container">
<div class="publish-book-div" style="float: left;">
	<img src="/static/www/book/{{.Book.Picture}}">
	<ul class="book-info">
		<li class="title"><a href="">{{.Book.Title}}</a></li>
		<li class="price">￥{{.Book.Price}}</li>
		<li>作者：{{.Book.Author}}</li>
		<li>出版社: {{.Book.Publisher}}</li>
		<li>ISBN: {{.Book.Isbn}}</li>
	</ul>
	<ul class="edit">
		<li><a href="javascript:void(0);" id="add-order" class="radius-button radius-button-small">下单</a></li>
	</ul>
	<div class="seprate"></div>
</div>
<div class="profile-div" style="margin-top: 32px;">
	<div class="profile-info">
	<div class="avatar-area">
		<img src="/static/www/avatar/{{.Vendor.Avatar}}">
		<div class="basic-info">
			<p class="username"><a href="/user/{{.Vendor.Id}}">{{.Vendor.Name}}</a></p>
			<p class="address">{{.Vendor.Place}}</p>
		</div>
	</div>
	<ul>
		<li>邮箱：{{.Vendor.Email}}</li>
		<li>手机号：{{.Vendor.PhoneNumber}}</li>
		<li>QQ号：{{.Vendor.Qq}}</li>
		<li>微信：{{.Vendor.Weixin}}</li>
	</ul>
	<a class="radius-button radius-button-small" href="/edit-profile">私信</a>
	</div>
	</div>
	<p style="float:left;">
	*请与卖家进行积极的沟通后再下单，右边是联系卖家的方式
</p>
<form id="order-form" method="POST">
</form>
</div>
</body>
<script type="text/javascript">
	$("#add-order").click(function() {
		$("#order-form").submit();
	});
</script>
</html>
