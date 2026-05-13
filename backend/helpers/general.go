package helpers

import (
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/logger"
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
