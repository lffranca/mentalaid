{{define "layout/sidebar-menu.tpl"}}
    <div class="menu-content">
        {{range .Routes}}
            <li class="row">
                <a href="{{.Path}}" class="{{if .Active}}button-active{{else}}button{{end}}">
                    <i class="{{.ClassIcon}}"></i>
                    <div class="title">{{.Title}}</div>
                </a>
            </li>
        {{end}}
    </div>
{{end}}
