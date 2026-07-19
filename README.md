## Zamysł tego projektu

Od zawsze chciałem zrobić swoje własne "centrum dowodzenia"
gdzie zrobię wszystko tak jakbym chciał. Od jakiegoś czasu marzyłem o takim moim
miejscu gdzie wszystkie "mikro serwisy" bedą ze soba spójne i bede mógł nimi łatwo zarządzać.
Z tego tez tytułu stworzyłem coś co pozwala połączyć moje wszystkie zajawki w jedno miejsce.

![image](https://api.klimson.dev/interface/bucket/klimson.dev/projects/dashboard/assets/hub2.png)

_Strona główna mojego panelu_

## Moje własne Google Drive

Jednym z pierwszych moich taskach w mojej todo liście bylo stworzenie swojego własnego systemu
plików. Na serwerze gin-gonic (framework do golanga) mam statycznie ustawioną scieżke na głowny folder na moje pliki.
W tym folderze tworze cały moj kontent na swoja stronę https://www.klimson.dev.
Rzeczy typu: **zdjęcia ze znajomymi**, **fotografie na blogi** czy po prostu coś co sie przyda na stronie trzymam właśnie w tym miejscu.
Nawet ten artykuł o tym projekcie jest napędzany moją chmurą na pliki ponieważ wykorzystuje plik **markdown** który tez jest zwracany z mojego serwera.
Chmura plików zatem stała sie jedną z wazniejszych funkcji tego projektu 😁

## Fajne dodatki

- Możliwość przełączanie między serwerem produkcyjnym a serwerem lokalnym. Funkcja ta nie dziala gdy frontend panelu jest na produkcji.

![image](https://api.klimson.dev/interface/bucket/klimson.dev/projects/dashboard/assets/flexible.png)

- Opcja zmiany motywu wbudowanego edytora markdowna/tekstu który bazuje na **Monaco** - ten sam silnik który napędza **Visual Studio Code**

![image](https://api.klimson.dev/interface/bucket/klimson.dev/projects/dashboard/assets/code_theme.png)
![image](https://api.klimson.dev/interface/bucket/klimson.dev/projects/dashboard/assets/built-in_editor.png)

- Opcja customizacji panelu pod własne preferencje

![image](https://api.klimson.dev/interface/bucket/klimson.dev/projects/dashboard/assets/customization.png)

# Redis Writable

Mam w planach stworzenia coś typu: [useState](https://react.dev/reference/react/useState) (z reacta) lub [writable](https://svelte.dev/docs/svelte/stores) (ze svelte), gdzie zmiany sa odświeżane na **kliencie** wraz ze **zmiana zmiennej na serwerze**. Zmienna będzie reaktywna i bedzie wykorzystywać mechanizm PUBSUB z redisa. Łącząc to z Websocketem jestem w stanie zrobić reaktywny stan łączony z oddzielnym serwerem.

```go
func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_URL"),
		Password: os.Getenv("RDB_PASS"),
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Ping(ctx).Result()
	if err != nil {
		logger.ErrorLog("Error during connecting to redis database: %v", err)
	}

	logger.GreenServerLog("✓ Connected to redis storage successfully with message: ", response)
	return client
}
```
