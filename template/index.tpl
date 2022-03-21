{{define "index.tpl"}}
{{template "layout/page-content-top.tpl" .}}
<div class="page-index">
    <div class="graph-row">
        <div class="graph-content graph-01-"></div>
    </div>
    <div class="graph-row">
        <div class="graph-content graph-01"></div>
        <div class="graph-content graph-02"></div>
        <div class="graph-content graph-03"></div>
    </div>
    <div class="graph-row">
        <div class="graph-content graph-04"></div>
        <div class="graph-content graph-05"></div>
        <div class="graph-content graph-06"></div>
    </div>
</div>


{{template "layout/page-content-bottom.tpl" .}}
{{end}}