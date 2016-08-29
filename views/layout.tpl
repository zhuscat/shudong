<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>主页</title>
  <link rel="stylesheet" type="text/css" href="/static/css/font-awesome.css">
  <link rel="stylesheet" type="text/css" href="/static/css/main.css">
  <link rel="stylesheet" type="text/css" href="/static/css/tooltip.css">
  <link rel="stylesheet" type="text/css" href="/static/css/management.css">
  <link rel="stylesheet" type="text/css" href="/static/css/private-letter.css">
  <script type="text/javascript" src="/static/js/jquery-2.2.4.js"></script>
  <script type="text/javascript" src="/static/js/tooltip.js"></script>
  <script type="text/javascript" src="/static/js/shudong.js"></script>
</head>
<body>
<div class="header">
<div class="logo"><a href="/"><img src="/static/img/logo.png" alt="logo"></a></div>
<div class="search">
<i class="fa fa-search"></i>
<form action="/search">
  <input type="text" name="wd" placeholder="书名、作者、ISBN...">
</form>
</div>
<div class="menu">
{{if .Login}}
  <a class="header-menu-item" href="/privateletter" data-toggle="tooltip" data-placement="bottom" title="私信"><i class="fa fa-envelope"></i></a>
  <a class="header-menu-item" href="/book/publish" data-toggle="tooltip" data-placement="bottom" title="发布"><i class="fa fa-plus"></i></a>
  <a class="header-menu-item" href="/message" data-toggle="tooltip" data-placement="bottom" title="提醒"><i class="fa fa-bell"><div></div></i></a>
  <a class="header-menu-item" href="/profile/published/all" data-toggle="tooltip" data-placement="bottom" title="个人中心"><i class="fa fa-user"></i></a>
  {{if .IsAdmin}}
  <a class="header-menu-item" href="/management" data-toggle="tooltip" data-placement="bottom" title="管理中心"><i class="fa fa-cog"></i></a>
  {{end}}
  <a class="header-menu-item" href="/signout" data-toggle="tooltip" data-placement="bottom" title="登出"><i class="fa fa-sign-out"></i></a>
{{else}}
  <a class="header-menu-item" href="/signup" data-toggle="tooltip" data-placement="bottom" title="注册"><i class="fa fa-user"></i></a>
  <a class="header-menu-item" href="/signin" data-toggle="tooltip" data-placement="bottom" title="登录"><i class="fa fa-sign-in"></i></a>
{{end}}
</div>
</div>
{{.LayoutContent}}
<div class="footer">
© 2016 书洞, all rights reserved
</div>
</body>
</html>
