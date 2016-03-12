package web

var (
	layoutTemplate = `
{{define "layout"}}
<html>
	<head>
		<title>Micro Web</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
		<style>
		{{ template "style" . }}
		</style>
	</head>
	<body>
	  <nav class="navbar navbar-inverse">
	    <div class="container">
	      <div class="navbar-header">
		<a class="navbar-brand" href="/">Micro</a>
	      </div>
	    </div>
	  </nav>
          <div class="container">
            <div class="row">
	      <div class="col-sm-3">
                <h4>&nbsp;</h4>
	        <ul class="list-group">
	          <li class="list-group-item"><a href="/">Home</a></li>
	          <li class="list-group-item"><a href="registry">Registry</a></li>
	          <li class="list-group-item"><a href="query">Query</a></li>
	        </ul>
	      </div>
	      <div class="col-sm-9">
	        <h1 class="page-header">{{ template "title" . }}</h1>
                {{ template "content" . }}
              </div>
            </div>
          </div>
	  <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
	  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
	  {{template "script" . }}
	</body>
</html>
{{end}}
{{ define "style" }}{{end}}
{{ define "script" }}{{end}}
{{ define "title" }}{{end}}
`

	indexTemplate = `
{{define "title"}}Welcome to the Micro Web{{end}}
{{define "content"}}
	{{if .HasWebServices}}
		<ul class="list-group">
			{{range .WebServices}}
			<li class="list-group-item"><a href="/{{.}}">{{.}}</a></li>
			{{end}}
		</ul>
	{{else}}
		<div class="alert alert-info" role="alert">
			<strong>No web services found</strong>
		</div>
	{{end}}
{{end}}
`
	queryTemplate = `
{{define "title"}}Query a Service{{end}}
{{define "style"}}
	pre {
		word-wrap: break-word;
	}
{{end}}
{{define "content"}}
<div class="row">
	<div class="col-sm-5">
		<form id="query-form" onsubmit="return query();">
			<div class="form-group">
				<label for="service">Service</label>
				<ul class="list-group">
					<select class="form-control" type=text name=service id=service> 
					<option disabled selected> -- select a service -- </option>
					{{range $key, $value := .}}
					<option class = "list-group-item" value="{{$key}}">{{$key}}</option>
					{{end}}
					</select>
				</ul>
			</div>
			<div class="form-group">
				<label for="method">Method</label>
				<ul class="list-group">
					<select class="form-control" type=text name=method id=method>
					<option disabled selected> -- select a method -- </option>
					</select>
				</ul>
			</div>
			<div class="form-group">
				<label for="othermethod">Other Method</label>
				<ul class="list-group">
					<input class="form-control" type=text name=othermethod id=othermethod disabled placeholder="Method"/>
				</ul>
			</div>
			<div class="form-group">
				<label for="request">Request</label>
				<textarea class="form-control" name=request id=request rows=8>{}</textarea>
			</div>
			<div class="form-group">
				<button class="btn btn-default">Go!</button>
			</div>
		</form>
	</div>
	<div class="col-sm-7">
		<p><b>Response</b></p>
		<pre id="response">{}</pre>
	</div>
</div>
{{end}}
{{define "script"}}
	<script>
		$(document).ready(function(){
			//Function executes on change of first select option field 
			$("#service").change(function(){
				var select = $("#service option:selected").val();
				$("#othermethod").attr("disabled", true);
				$('#othermethod').val('');
				$("#method").empty();
				$("#method").append("<option disabled selected> -- select a method -- </option>");
				var s_map = {};
				{{ range $service, $methods := . }}
				var m_list = [];
				{{range $index, $element := $methods}}
				m_list[{{$index}}] = {{$element.Name}}
				{{end}}
				s_map[{{$service}}] = m_list
				{{ end }}
				if (select in s_map) {
				var serviceMethods = s_map[select]
				var len = serviceMethods.length;
					for(var i = 0; i < len; i++) {
						$("#method").append("<option value=\""+serviceMethods[i]+"\">"+serviceMethods[i]+"</option>");	
					}
				}
				$("#method").append("<option value=\"other\">other</option>");
			});
			//Function executes on change of second select option field 
			$("#method").change(function(){
				var select = $("#method option:selected").val();
				if (select == "other") {
					$("#othermethod").attr("disabled", false);
				} else {
					$("#othermethod").attr("disabled", true);
					$('#othermethod').val('');
				}

			});
		});
	</script>
	<script>
		function query() {
			var req = new XMLHttpRequest()
			req.onreadystatechange = function() {
				if (req.readyState == 4 && req.status == 200) {
					document.getElementById("response").innerText = JSON.stringify(JSON.parse(req.responseText), null, 2);
					console.log(req.responseText);
				}
			}
			var method = document.forms[0].elements["method"].value
			if (!($('#othermethod').prop('disabled'))) {
				method = document.forms[0].elements["othermethod"].value
			}
			var request = {
				"service": document.forms[0].elements["service"].value,
				"method": method,
				"request": JSON.parse(document.forms[0].elements["request"].value)
			}
			req.open("POST", "/rpc", true);
			req.setRequestHeader("Content-type","application/json");				
			req.send(JSON.stringify(request));

			return false;
		};	
	</script>
{{end}}
`
	registryTemplate = `
{{define "title"}}Registry{{end}}
{{define "content"}}
	<ul class="list-group">
		{{range .}}
		<li class="list-group-item"><a href="registry?service={{.Name}}">{{.Name}}</a></li>
		{{end}}
	</ul>
{{end}}
`

	serviceTemplate = `
{{define "title"}}Service {{with $svc := index . 0}}{{$svc.Name}}{{end}}{{end}}
{{define "content"}}
	<h4>Nodes</h4>
	{{range .}}
	<h5>Version {{.Version}}</h5>
	<table class="table table-bordered table-striped">
		<thead>
			<th>Id</th>
			<th>Address</th>
			<th>Port</th>
			<th>Metadata</th>
		<thead>
		<tbody>
			{{range .Nodes}}
			<tr>
				<td>{{.Id}}</td>
				<td>{{.Address}}</td>
				<td>{{.Port}}</td>
				<td>{{ range $key, $value := .Metadata }}{{$key}}={{$value}} {{end}}</td>
			</tr>
			{{end}}
		</tbody>
	</table>
	{{end}}
	{{with $svc := index . 0}}
	{{if $svc.Endpoints}}
	<h4>Endpoints</h4>
	<hr/>
	{{end}}
	{{range $svc.Endpoints}}
		<h4>{{.Name}}</h4>
		<table class="table table-bordered">
			<tbody>
				<tr>
					<th class="col-sm-2" scope="row">Metadata</th>
					<td>{{ range $key, $value := .Metadata }}{{$key}}={{$value}} {{end}}</td>
				</tr>
				<tr>
					<th class="col-sm-2" scope="row">Request</th>
					<td><pre>{{format .Request}}</pre></td>
				</tr>
				<tr>
					<th class="col-sm-2" scope="row">Response</th>
					<td><pre>{{format .Response}}</pre></td>
				</tr>
			</tbody>
		</table>
	{{end}}
	{{end}}
{{end}}

`
)
