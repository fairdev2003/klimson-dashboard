package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Wewnętrzna funkcja budująca ustandaryzowany obiekt JSON z bazowymi polami oraz dynamicznymi danymi.
func build_response(ctx *gin.Context, statusCode int, isError bool, message string, extraData any) {
	response := gin.H{
		"error":   isError,
		"message": message,
	}

	if dataMap, ok := extraData.(gin.H); ok {
		for key, value := range dataMap {
			if key == "error" || key == "message" {
				continue
			}
			response[key] = value
		}
	} else if dataMap, ok := extraData.(map[string]any); ok {
		for key, value := range dataMap {
			if key == "error" || key == "message" {
				continue
			}
			response[key] = value
		}
	}

	ctx.JSON(statusCode, response)
}

// Wewnętrzna funkcja pomocnicza - wyciąga pierwszą wiadomość z tablicy wariadycznej lub zwraca domyślną.
func getMessage(provided []string, defaultMsg string) string {
	if len(provided) > 0 && provided[0] != "" {
		return provided[0]
	}
	return defaultMsg
}

// Returns a JSON which is made for response that is successful (HTTP 200)
func SuccessResponse(ctx *gin.Context, data any, m ...string) {
	message := getMessage(m, "Request is successful")
	build_response(ctx, http.StatusOK, false, message, data)
}

// Returns a JSON which is made for response that is successful (HTTP 201)
func StatusCreatedResponse(ctx *gin.Context, data any, m ...string) {
	message := getMessage(m, "Content is successfully created!")
	build_response(ctx, http.StatusCreated, false, message, data)
}

func NotFoundResponse(ctx *gin.Context) {
	build_response(ctx, http.StatusNotFound, true, "Requested resource is not found.", nil)
}

// Returns a JSON which is made for response that is unauthorized (HTTP 401)
func UnauthorizedResponse(ctx *gin.Context, data any, m ...string) {
	message := getMessage(m, "Unauthorized Access")
	build_response(ctx, http.StatusUnauthorized, true, message, data)
}

// Returns a JSON which is made for response that has too many request registered (HTTP 429)
func TooManyRequestsResponse(ctx *gin.Context, cooldown string, m ...string) {
	message := getMessage(m, "Too many requests! Try again in: "+cooldown)
	build_response(ctx, http.StatusTooManyRequests, true, message, nil)
}

// Returns a JSON which is made for response that is bad (HTTP 400)
func BadRequestResponse(ctx *gin.Context, data any, m ...string) {
	message := getMessage(m, "Request is bad")
	build_response(ctx, http.StatusBadRequest, true, message, data)
}

// Returns a JSON which is made for response that indicates an internal server error (HTTP 500)
func InternalServerErrorResponse(ctx *gin.Context, data any, m ...string) {
	message := getMessage(m, "Internal Server Error")
	build_response(ctx, http.StatusInternalServerError, true, message, data)
}
