package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/helpers/maroto"
)

func (gc GlobalController) ReturnSecurityCV(ctx *gin.Context) {
	pdfPath := "static/docs/assets/pdf/cv.security.pdf"

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		maroto.GenerateSecurityCV()
	}

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", "attachment; filename=cv.security.pdf")

	ctx.File(pdfPath)
}

func (gc GlobalController) ReturnDevCV(ctx *gin.Context) {
	pdfPath := "static/docs/assets/pdf/cv.dev.pdf"

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		err := maroto.GenerateDevCV()
		if err != nil {
			api.InternalServerErrorResponse(ctx, err, "Failed to generate CV PDF.")
			return
		}
	}

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		api.InternalServerErrorResponse(ctx, nil, "PDF file was not found after generation.")
		return
	}

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", "attachment; filename=cv.dev.pdf")

	ctx.File(pdfPath)
}
