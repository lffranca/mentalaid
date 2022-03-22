{{define "layout/page-content-top.tpl"}}
{{template "layout/top.tpl" .}}
    <div class="page-content">
        {{template "layout/header.tpl" .}}
        <div class="container">
            <div class="breadcrumbs">
                {{range $index, $breadcrumbs := .Route.Breadcrumbs}}
                    {{if lt (inc $index) (len $.Route.Breadcrumbs)}}
                        <a href="{{$breadcrumbs.Path}}">{{$breadcrumbs.Title}}</a> >
                    {{else}}
                        {{$breadcrumbs.Title}}
                    {{end}}
                {{end}}
            </div>
            <div class="page-container">
            <!-- INIT CONTENT PAGE -->
{{end}}