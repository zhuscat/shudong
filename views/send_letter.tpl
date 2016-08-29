<html>
<head>
	
<title>私信对话</title>

</head>
<div class="container">	
<div class="letters">
<p>与 {{.ToUser.Name}} 的对话</p>
<form action="/privateletter/{{.ToUser.Id}}" method="post">
<div class="content">
<input type="text" class="form-control" name="content"style="width:400px;" >&nbsp;&nbsp;&nbsp;&nbsp;
<input type="submit" value=" 发送 ">
</div>
</form>
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

{{range .Letters}}
<div class="privateletters-div">
	<p>
	{{if eq .FromId $.User.Id}}
	<img src="/static/www/avatar/{{$.User.Avatar}}"width="32" height="32">{{$.User.Name}}
	{{end}}
	{{if eq .FromId $.ToUser.Id}}
	<img src="/static/www/avatar/{{$.ToUser.Avatar}}"width="32" height="32">{{$.ToUser.Name}}
	{{end}}:
	{{.Content}}</p><p style="font-size:12px;">{{.SendTime}}</p>
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
</body></html>