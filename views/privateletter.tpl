<html>
<head>
<title>私信对话</title>
</head>
<div class="container">
<div class="letterlist"> 
<p>对话列表</p>
{{range .ToUsers}}
	<p>
	<img src="/static/www/avatar/{{.Avatar}}"width="32" height="32">
	<a href="/privateletter/{{.Id}}">{{.Name}}</a>
	{{$var:=.Id}}
	{{range $.LastLetters}}
		<p style="font-size:14px;">
		{{if eq .FromId $var}}{{.Content}}&nbsp;&nbsp;&nbsp;&nbsp;{{.SendTime}}{{end}}
		{{if eq .ToId $var}}{{.Content}}&nbsp;&nbsp;&nbsp;&nbsp;{{.SendTime}}{{end}}
		</p>
	{{end}}
	</p>
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
</html>