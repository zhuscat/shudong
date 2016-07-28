<div class="container">
<div class="left-area-wrapper">
<div class="detail-book-card">
	<div class="detail-book-card-top">
		<img src="/static/www/book/{{.Book.Picture}}">
		<ul>
			<li class="title">{{.Book.Title}}</li>
			<li class="price">￥{{.Book.Price}}</li>
			<li>作者：{{.Book.Author}}</li>
			<li>出版社：{{.Book.Publisher}}</li>
			<li>ISBN：{{.Book.Isbn}}</li>
			<li>所在地：{{.Book.Vendor.Place}}</li>
		</ul>
		{{if .Book.Onsale}}
			{{if .Login}}
				{{if eq .User.Id .Book.VendorId}}
				<a class="radius-button radius-button-middle buy-button radius-button-disable">购买</a>
				{{else}}
				<a href="/order/customer/confirm/{{.Book.Id}}" class="radius-button radius-button-middle buy-button">购买</a>
				{{end}}
			{{else}}
				<a href="/order/customer/confirm/{{.Book.Id}}" class="radius-button radius-button-middle buy-button">购买</a>
			{{end}}
		{{else}}
		<a class="radius-button radius-button-middle buy-button radius-button-disable">已下架</a>
		{{end}}
	</div>
	<div class="detail-book-card-bottom">
		<ul class="tab-menu">
		{{if .ShowComment}}
		<li data-item=".introduction">
			详情
		</li>
		<li class="item-selected" data-item=".comment">回复</li>
		{{else}}
		<li class="item-selected" data-item=".introduction">
			详情
		</li>
		<li data-item=".comment">回复</li>
		{{end}}
		</ul>
		<div class="introduction" {{if .ShowComment}}style="display: none;"{{else}}style="display: block;"{{end}}>
			{{.Book.Description}}
		</div>
		<div class="comment" {{if .ShowComment}}style="display: block;"{{else}}style="display: none;"{{end}}>
			<div id="comment-display">
				{{range .Comments}}
			<div class="comment-div">
				<img src="/static/www/avatar/{{.User.Avatar}}">
				<div class="comment-wrapper">
					<a href="/user/{{.User.Id}}">{{.User.Name}}</a>
					<div class="content">{{.Content}}</div>
				</div>
				<div class="horizontal-sep"></div>
			</div>
			{{end}}
			</div>
			{{if .Login}}
			<form method="POST" id="comment-form" action="/book/comment">
				<div class="shudong-normal-input shudong-normal-input-comment">
				<input type="text" name="bookid" value="{{.Book.Id}}" style="display:none;">
				<input id="content" type="text" name="content" placeholder="写下你的评论">
				</div>
				<input type="submit" class="radius-button radius-button-small radius-button-green">
			</form>
			{{else}}
			<p>请先 <a href="/signin?next=/book/detail/{{.Book.Id}}">登陆</a> 再进行评论</p>
			{{end}}
					{{if gt .Page.PageNums 1}}
<ul class="comment-page">
	{{range $index, $page := .Page.Pages}}
		<li {{if $.Page.IsActive .}}class="active"{{end}}>
			<a href="{{$.Page.PageLink $page}}">{{$page}}</a>
		</li>
	{{end}}
</ul>
{{end}}
		</div>
	</div>
</div>
</div>
<div class="right-area-wrapper">
	<div class="profile-div">
	<div class="profile-info">
	<div class="avatar-area">
		<img src="/static/www/avatar/{{.Book.Vendor.Avatar}}">
		<div class="basic-info">
			<p class="username"><a href="/user/{{.Book.Vendor.Id}}">{{.Book.Vendor.Name}}</a></p>
			<p class="address">{{.Book.Vendor.Place}}</p>
		</div>
	</div>
	<ul>
		<li>邮箱：{{.Book.Vendor.Email}}</li>
		<li>手机号：{{.Book.Vendor.PhoneNumber}}</li>
		<li>QQ号：{{.Book.Vendor.Qq}}</li>
		<li>微信：{{.Book.Vendor.Weixin}}</li>
	</ul>
	<a class="radius-button radius-button-small" href="">私信</a>
	</div>
	</div>
	<div class="other-book-div">
	<div>该卖家还出售...</div>
	<ul>
	{{range .OtherBooks}}
	<li><a href="/book/detail/{{.Id}}"><img src="/static/www/book/{{.Picture}}" alt="{{.Title}}"></a></li>
	{{end}}
	</ul>
	</div>
	<a class="report" href="/">给我们反馈</a>
</div>
</div>
<script type="text/javascript">
	// $("#content").focusin(function() {
	// 	$('#comment-form>input[type="submit"]').css("display", "block")
	// });
	// $("#content").focusout(function() {
	// 	$('#comment-form>input[type="submit"]').css("display", "none")
	// });
	$(".tab-menu>li").click(function() {
		$(".tab-menu>li").toggleClass("item-selected");
		$(".tab-menu>li").each(function() {
			var item = $(this).data("item");
			$(item).css("display", "none");
		});
		var thisItem = $(this).data("item");
		$(thisItem).css("display", "block");
	});

	$(".comment-page a").click(function(e) {
		// 获取链接里面的参数
		var href = $(this).attr("href");
		var re = /\/book\/detail\/(\d+)\?p=(\d+)/gi;
		var matches = re.exec(href);
		var getUrl = "/comment/get/" + matches[1] + "?p=" + matches[2];
		$(".comment-page a").parent().removeClass("active");
		$(this).parent().addClass("active");
		$.ajax({
			url: getUrl,
			type: "GET", 
			success: function(data) {
				$("#comment-display").html(data);
			}
		});
		e.preventDefault();
	});
</script>