<!DOCTYPE html>
<html>
	<head>
		{{template "../base/T.head.tpl" .}}
		{{template "head" .}}
	</head>
	<body>
		<noscript>Please enable JavaScript in your browser!</noscript>
		{{template "../base/T.navbar.tpl" .}}
		<div class="container">
			<div class="body">
				{{template "body" .}}
			</div>
	    </div>
		{{template "../base/T.footer.tpl" .}}
	</body>
</html>