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
<div id="header-wrapper">
	<div id="header" class="container">
		<div align="right">{{if .InSession}}
		Welcome, {{.First}} [<a href="/user/logout">Logout</a>|]
		{{else}}
		[<a href="/user/login/home">Login</a>]
		{{end}}
		</div>

		<div id="menu">
			<ul>
				<li class="active"><a href="/home">home</a></li>
				{{if .InSession}}
					<li class="active"><a href="/user/profile">Your Account</a></li>
				{{else}}
				{{end}}
				<!-- <li class="active"><a href="/user/profile">Your Applications</a></li> -->
			</ul>
		</div>
		<div id="banner" class="container">
			<div class="title">
				<h2>Klouds.io</h2>
				<span class="byline">Application Hosting Made Easy</span>
			</div>
			{{if .InSession}}
			{{else}}
				<ul class="actions">
					<li><a href="/user/login/home" class="button">Get Started</a></li>
				</ul>
			{{end}}
		</div>
	</div>
	</div>
</div>
<!-- end header -->
<!-- start page -->
<div id="page">