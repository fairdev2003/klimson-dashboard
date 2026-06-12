package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
)

func (controller GlobalController) GetClanInfo(ctx *gin.Context) {
	clan_id := ctx.Param("clan_id")

	client := req.C()

	type ExternalResponse struct {
		Data string `json:"data"`
	}

	var extResp ExternalResponse

	resp, err := client.R().
		SetContext(ctx.Request.Context()).
		SetBody(map[string]interface{}{
			"clan_id": clan_id,
		}).
		SetSuccessResult(&extResp).
		Post("https://asteroidpg3d.xyz/api/get_clan_info")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp.IsSuccessState() {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "External API returned non-200 status"})
		return
	}

	rawString := extResp.Data
	if rawString == "" {
		rawString = resp.String()
	}

	if rawString == "" || rawString == "{}" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "empty_response",
			"message": "Clan not found or external API returned no data",
			"raw":     resp.String(),
		})
		return
	}

	var formattedData interface{}
	if err := json.Unmarshal([]byte(rawString), &formattedData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "partial_success",
			"message": "Failed to parse inner JSON string, showing raw text",
			"error":   err.Error(),
			"data":    rawString,
		})
		return
	}

	ctx.JSON(http.StatusOK, formattedData)
}

func (controller GlobalController) GetPlayerData(ctx *gin.Context) {
	player_id := ctx.Param("player_id")

	client := req.C()

	type ExternalResponse struct {
		Data string `json:"data"`
	}

	var extResp ExternalResponse

	resp, err := client.R().
		SetContext(ctx.Request.Context()).
		SetBody(map[string]interface{}{
			"player_id": player_id,
		}).
		SetSuccessResult(&extResp).
		Post("https://asteroidpg3d.xyz/api/get_player_info")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp.IsSuccessState() {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "External API returned non-200 status"})
		return
	}

	rawString := extResp.Data
	if rawString == "" {
		rawString = resp.String()
	}

	if rawString == "" || rawString == "{}" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "empty_response",
			"message": "Clan not found or external API returned no data",
			"raw":     resp.String(),
		})
		return
	}

	var formattedData interface{}
	if err := json.Unmarshal([]byte(rawString), &formattedData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "partial_success",
			"message": "Failed to parse inner JSON string, showing raw text",
			"error":   err.Error(),
			"data":    rawString,
		})
		return
	}

	ctx.JSON(http.StatusOK, formattedData)
}
