<!-- 其实这个是一个消息中心，有时间改一下名字 -->
<div class="container">
<div class="left-area-wrapper">
<ul class="top-menu-bar">
<li class="bar-item {{if eq .Tab ""}}bar-item-selected{{end}}"><a href="/message">所有</a></li>
<li class="bar-item {{if eq .Tab "read"}}bar-item-selected{{end}}"><a href="/message/read">已读</a></li>
<li class="bar-item {{if eq .Tab "unread"}}bar-item-selected{{end}}"><a href="/message/unread">未读</a></li>`
</ul>
<button id="read-all" class="radius-button radius-button-small" style="margin-top: 16px;">全部标记为已读</button>
{{range .Messages}}
<div class="message-div">
	<p>{{.Content}}</p>
	<p>{{.SendTime}}</p>
	{{if not .Read}}
	<a href="javascript:void(0)" class="radius-button radius-button-small read-button" data-messageid="{{.Id}}">已读</a>
	{{end}}
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
	</div>
	<ul class="menu-area">
	</ul>
</div>
</div>
<script type="text/javascript">
	$(".read-button").click(function() {
		$this = $(this);
		$.ajax({
			url: "/message/confirm-read",
			type: "POST",
			data: {"messageid": $(this).data("messageid")},
			dataType: "json",
			success: function(data) {
				if (data["success"] == true) {
					$this.css("display", "none");
				}
			}
		});
	});

	$("#read-all").click(function() {
		$.ajax({
			url: "/message/read-all",
			type: "GET",
			dataType: "json",
			success: function(data) {
				if (data["success"] == true) {
					$(".read-button").css("display", "none");
				}
			}
		});
		e.preventDefault();
	});
</script>