{{define "layout/sidebar.tpl"}}
    <div class="sidebar">
        {{ template "layout/sidebar-header.tpl" .}}
        {{ template "layout/sidebar-menu.tpl" .}}
    </div>
{{end}}