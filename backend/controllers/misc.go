package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller GlobalController) TestRedirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}

func (controller GlobalController) GetStorageLeftPercentage(ctx *gin.Context) {

}
