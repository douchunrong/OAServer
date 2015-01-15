<nav class="navbar navbar-default navbar-static-top" role="navigation">
    <div class="container-fluid">
      <div class="navbar-header">
        <a class="navbar-brand " href="#">
          <img alt="Brand"  width="30px" height="30px" src="/static/img/oauth2.png">
        </a>
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
          <span class="sr-only">Oauth2.0</span>
        </button>
        <a class="navbar-brand" href="#">Oauth2.0</a>
      </div>
      <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
      </ul>
   
  	 <ul class="nav navbar-nav navbar-right">
  		{{if .IsLogin}}
  		<li><a href="/home?exit=true">退出</a></li>
  		{{else}}
  		<li><a href="/login">登录</a></li>
      <li><a href="/register">注册</a></li>
  		{{end}}
  	</ul>
      </div>
    </div>
</nav>