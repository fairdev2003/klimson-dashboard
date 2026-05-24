package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type SpotifyAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type NowPlayingResponse struct {
	IsPlaying bool   `json:"is_playing"`
	TrackName string `json:"track_name,omitempty"`
	Artist    string `json:"artist,omitempty"`
	Album     string `json:"album,omitempty"`
	AlbumArt  string `json:"album_art,omitempty"`
}

type SpotifyTokenCache struct {
	mu          sync.Mutex
	accessToken string
	expiresAt   time.Time
}

var tokenCache = &SpotifyTokenCache{}

func GetOrFetchAccessToken() (string, error) {
	tokenCache.mu.Lock()
	defer tokenCache.mu.Unlock()

	if tokenCache.accessToken != "" && time.Now().Before(tokenCache.expiresAt.Add(-10*time.Second)) {
		return tokenCache.accessToken, nil
	}

	freshToken, err := FetchSpotifyAccessToken()
	if err != nil {
		return "", err
	}

	tokenCache.accessToken = freshToken
	tokenCache.expiresAt = time.Now().Add(1 * time.Hour)

	return tokenCache.accessToken, nil
}

func FetchSpotifyAccessToken() (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		return "", fmt.Errorf("missing spotify credentials in environment variables")
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	apiURL := "https://accounts.spotify.com/api/token"
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request to spotify failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("spotify returned non-200 status: %s", resp.Status)
	}

	var authResp SpotifyAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return "", fmt.Errorf("failed to decode spotify response: %v", err)
	}

	return authResp.AccessToken, nil
}

func RequestSpotify(spotify_url string, m ...string) (map[string]interface{}, error) {
	accessToken, err := GetOrFetchAccessToken()
	if err != nil {
		return nil, err
	}

	method := "GET"
	if len(m) > 0 {
		method = m[0]
	}

	req, err := http.NewRequest(method, spotify_url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var spotifyRaw map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&spotifyRaw); err != nil {
		return nil, err
	}

	return spotifyRaw, nil
}

func (controller GlobalController) GetPlaybackState(ctx *gin.Context) {
	playback_state, err := RequestSpotify("https://api.spotify.com/v1/me/player")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, playback_state)
}
