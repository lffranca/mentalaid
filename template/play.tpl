{{define "play.tpl"}}
    {{template "layout/page-content-top.tpl" .}}

        <div class="page-play">
            <h2>Vídeos recomendados:</h2>
            <!-- INIT CONTENT PAGE -->
            <div class="seletor video">
                <div class="video1"></div>
                <div class="video2"></div>
                <div class="video3"></div>
            </div>

            <h2>Podcast do dia:</h2>
            <div class="seletor podcast">
                <div class="spotify"></div>
                <div class="podcast"></div>
            </div>
        </div>
    {{template "layout/page-content-bottom.tpl" .}}
{{end}}