package helpers

import (
	"fmt"
	"image"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/logger"
)

func NetworkLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.ServerLog(c.Request.Host)
	}
}

func CorsConf(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {

		logger.ServerLog(logger.GetMethodColor(c.Request.Method)(c.Request.Method) + " ==> " + c.Request.RequestURI)
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, UPDATE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetMostCommonColor(imageURL string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("błąd podczas pobierania zdjęcia: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("serwer zwrócił niepoprawny status: %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", fmt.Errorf("błąd podczas dekodowania obrazu: %w", err)
	}

	colors, err := prominentcolor.Kmeans(img)
	if err != nil {
		return "", fmt.Errorf("błąd podczas przetwarzania K-Means: %w", err)
	}

	if len(colors) == 0 {
		return "", fmt.Errorf("nie znaleziono dominującego koloru")
	}

	dominantColor := colors[0]

	hexColor := fmt.Sprintf("#%02X%02X%02X",
		dominantColor.Color.R,
		dominantColor.Color.G,
		dominantColor.Color.B,
	)

	return hexColor, nil
}
