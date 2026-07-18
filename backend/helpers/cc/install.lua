-- Konfiguracja
local USER = "fairdev2003"
local REPO = "klimson-dashboard"
local PATH = "backend/helpers/cc"

local api_url = string.format("https://api.github.com/repos/%s/%s/contents/%s", USER, REPO, PATH)

print("Łączenie z GitHub API...")

local headers = {
    ["User-Agent"] = "CC-Tweaked-Fetch-Script"
}

local response = http.get(api_url, headers)

if not response then
    error("Błąd: Nie można pobrać danych z API. Sprawdź konfigurację HTTP.")
end

local body = response.readAll()
response.close()

local files = textutils.unserializeJSON(body)

if not files or type(files) ~= "table" then
    error("Błąd: Nie udało się odczytać listy plików. Sprawdź nazwę folderu.")
end

print("Znaleziono pliki. Pobieranie...")

for _, file in ipairs(files) do
    if file.type == "file" then
        print("Pobieranie: " .. file.name)

        local file_response = http.get(file.download_url)
        if file_response then
            local data = file_response.readAll()
            file_response.close()

            local local_file = fs.open(file.name, "w")
            if local_file then
                local_file.write(data)
                local_file.close()
            else
                print("❌ Błąd: Brak uprawnień do zapisu: " .. file.name)
            end
        else
            print("❌ Błąd: Pobieranie nie powiodło się: " .. file.name)
        end
    elseif file.type == "dir" then
        print("Pominięto podfolder: " .. file.name)
    end
end

print("\nSukces! Pliki zaktualizowane.")
