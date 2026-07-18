local url = "ws://mojprojekt.test:8080/ws" -- lub Twój adres
local ws, err = http.websocket(url)

if not ws then
    error("Nie można połączyć: " .. err)
end

while true do
    local event, url_event, message = os.pullEvent("websocket_message")

    if url_event == url then
        print("Odebrano: " .. message)
    end
end
