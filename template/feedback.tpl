{{define "feedback.tpl"}}
    {{template "layout/page-content-top.tpl" .}}
    <div class="page-feedback">
        <form class="content" method="get">
            <h1 class="title">
                Nos envie seu feedback
            </h1>

            <label for="nome">Nome</label>
            <input type="text" name="nome" id="nome" placeholder="Nome" required>

            <label for="email">E-mail</label>
            <input type="email" name="email" id="email" placeholder="E-mail" required>

            <label for="message">Mensagem</label>
            <textarea name="message" id="message" cols="30" rows="10" placeholder="Digite sua mensagem..." required></textarea>

            <div class="button-content">
                <button type="submit">Enviar</button>
            </div>
        </form>
    </div>
    {{template "layout/page-content-bottom.tpl" .}}
{{end}}