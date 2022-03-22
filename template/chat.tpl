{{define "chat.tpl"}}
    {{template "layout/page-content-top.tpl" .}}
    <div class="page-chat">
        <div class="chat-header">
            To: <span>ROBOT</span>
        </div>
        <div class="chat-body">
            <div class="message-row received">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Robot
                    </div>
                </div>
                <div class="message">
                    It was popularised in the 1960s with the release of Letraset sheets
                    containing Lorem Ipsum passages, and more recently
                </div>
                <div class="datetime">07:30</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
            <div class="message-row sent">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Fulano
                    </div>
                </div>
                <div class="message">
                    Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                    when an unknown printer took a galley of type and scrambled it to make a type specimen book
                </div>
                <div class="datetime">08:55</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
            <div class="message-row received">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Robot
                    </div>
                </div>
                <div class="message">
                    It was popularised in the 1960s with the release of Letraset sheets
                    containing Lorem Ipsum passages, and more recently
                </div>
                <div class="datetime">07:30</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
            <div class="message-row sent">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Fulano
                    </div>
                </div>
                <div class="message">
                    Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                    when an unknown printer took a galley of type and scrambled it to make a type specimen book
                </div>
                <div class="datetime">08:55</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
            <div class="message-row received">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Robot
                    </div>
                </div>
                <div class="message">
                    It was popularised in the 1960s with the release of Letraset sheets
                    containing Lorem Ipsum passages, and more recently
                </div>
                <div class="datetime">07:30</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
            <div class="message-row sent">
                <div class="profile">
                    <div class="avatar"></div>
                    <div class="name">
                        Fulano
                    </div>
                </div>
                <div class="message">
                    Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                    when an unknown printer took a galley of type and scrambled it to make a type specimen book
                </div>
                <div class="datetime">08:55</div>
                <i class="fa-solid fa-check-circle" aria-hidden="true"></i>
            </div>
        </div>
        <div class="chat-footer">
            <textarea placeholder="Message..." rows="4"></textarea>
        </div>
    </div>
    {{template "layout/page-content-bottom.tpl" .}}
{{end}}