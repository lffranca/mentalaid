{{ define "layout/top.tpl" }}
<!DOCTYPE html>
<html lang="pt-br">
{{ template "layout/head.tpl" .}}
<body class="body">
    {{ template "layout/sidebar.tpl" .}}
{{ end }}