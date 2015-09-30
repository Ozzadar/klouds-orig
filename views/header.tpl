<!DOCTYPE html>

<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8" />
<title>Klouds IO -- Scalable Open Source Software Hosting</title>
<meta name="keywords" content="" />
<meta name="description" content="" />
<link href="/static/css/default.css" rel="stylesheet" type="text/css" />
</head>

<body>
<!-- start header -->
<div id="header-bg">
	<div id="header">
		<div align="right">{{if .InSession}}
		Welcome, {{.First}} [<a href="../user/logout">Logout</a>|<a href="../user/profile">Profile</a>]
		{{else}}
		[<a href="../login/home">Login</a>]
		{{end}}
		</div>
		<div id="logo">
			<img src="../../static/img/img03.png"></img>
		</div>
		<div id="menu">
			<ul>
				<li class="active"><a href="../../home">home</a></li>
				<li class="active"><a href="../../appLaunch">Launch an App!</a></li>
			</ul>
		</div>
	</div>
</div>
<!-- end header -->
<!-- start page -->
<div id="page">