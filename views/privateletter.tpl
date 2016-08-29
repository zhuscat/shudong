<div class="container">
<div class="letterlist">
<p>对话列表</p>
{{range .ToUsers}}
	<div class="private-letter-dialog">
		<div class="private-letter-avatar-wrapper">
			<img class="private-letter-avatar" src="/static/www/avatar/{{.Avatar}}">
			<a class="private-letter-detail-link" href="/privateletter/{{.Id}}">{{.Name}}</a>
		</div>
	{{$var:=.Id}}
	{{range $.LastLetters}}
		<div>
		{{if eq .FromId $var}}
		<div class="private-letter-content">
			{{.Content}}
		</div>
		<div class="private-letter-time">
			{{.SendTime}}
		</div>
		{{end}}
		{{if eq .ToId $var}}
		<div class="private-letter-content">
			{{.Content}}
		</div>
		<div class="private-letter-time">
			{{.SendTime}}
		</div>
		{{end}}
	</div>
	{{end}}
	</div>
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
