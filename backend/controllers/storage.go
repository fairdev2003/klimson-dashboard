package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (gc GlobalController) GetFile(c *gin.Context) {
	filePath := c.Param("filepath")

	fullPath := "./static/uploads" + filePath

	c.File(fullPath)
}

func (gc GlobalController) GetSecuredFile(c *gin.Context) {
	filePath := c.Param("filepath")

	cleanPath := filepath.Clean(filePath)

	if strings.Contains(cleanPath, "passwords.txt") {
		c.AbortWithStatusJSON(403, gin.H{"error": "Access denied to config files"})
		return
	}

	fullPath := filepath.Join("./secured", cleanPath)

	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) || info.IsDir() {
		c.AbortWithStatusJSON(404, gin.H{"error": "File not found"})
		return
	}

	c.File(fullPath)
}

type ListRecord struct {
	Name         string    `json:"name"`
	IsDir        bool      `json:"is_dir"`
	ModifiedTime time.Time `json:"modified"`
	Size         int64     `json:"file_size"`
}

func (gc GlobalController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Brak pliku w żądaniu", "success": false})
		return
	}

	folderPath := c.Param("folder")
	dst := filepath.Join("./static/uploads", folderPath, file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Błąd zapisu pliku", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plik przesłany pomyślnie", "success": true})
}

func (gc GlobalController) CreateFolder(c *gin.Context) {
	var input struct {
		FolderName string `json:"folder_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane", "success": false})
		return
	}

	folderPath := c.Param("folder")
	fullPath := filepath.Join("./static/uploads", folderPath, input.FolderName)

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		c.JSON(http.StatusConflict, gin.H{"message": "Folder o tej nazwie już istnieje", "success": false})
		return
	}

	if err := os.MkdirAll(fullPath, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Błąd tworzenia folderu", "success": false})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Folder stworzony", "success": true})
}

func (gc GlobalController) ListFiles(c *gin.Context) {
	folderPath := c.Param("folder")
	fullPath := "./static/uploads" + folderPath

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Błąd odczytu folderu")
		return
	}

	var fileList []ListRecord
	for _, entry := range entries {
		// Pobieramy pełne informacje o pliku
		info, err := entry.Info()
		if err != nil {
			continue // Jeśli nie da się pobrać info, pomijamy ten plik
		}

		fileList = append(fileList, ListRecord{
			Name:         entry.Name(),
			IsDir:        entry.IsDir(),
			ModifiedTime: info.ModTime(),
			Size:         info.Size(),
		})
	}

	c.JSON(http.StatusOK, fileList)
}

func (gc GlobalController) Interface(c *gin.Context) {
	subFolder := c.Param("folder")
	basePath := "./static/uploads"
	fullPath := filepath.Join(basePath, subFolder)

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		c.String(http.StatusNotFound, "Błąd: Nie znaleziono folderu lub pliku")
		return
	}

	if !fileInfo.IsDir() {
		c.File(fullPath)
		return
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Błąd odczytu folderu")
		return
	}

	htmlStyles := `
    <style>
        body { 
            background-color: #121212; 
            color: #e0e0e0; 
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; 
            padding: 40px; 
            line-height: 1.6;
        }
        h2 { color: #bb86fc; border-bottom: 1px solid #333; padding-bottom: 10px; margin-bottom: 0; }
        .path { color: #757575; font-size: 0.9em; margin-bottom: 20px; display: block; }
        ul { list-style: none; padding: 0; }
        li { margin: 4px 0; padding: 8px 12px; border-radius: 6px; transition: background 0.2s; display: flex; align-items: center; }
        li:hover { background-color: #2c2c2c; }
        a { color: #03dac6; text-decoration: none; font-weight: 500; flex-grow: 1; }
        a:hover { color: #66fff1; }
        hr { border: 0; border-top: 1px solid #333; margin: 20px 0; }
        .icon { margin-right: 12px; width: 24px; text-align: center; font-size: 1.2em; }
    </style>`

	html := fmt.Sprintf("<html><head><title>Index of %s</title>%s</head><body>", subFolder, htmlStyles)
	html += fmt.Sprintf("<h2>Storage Explorer</h2><span class='path'>Lokalizacja: /%s</span><hr><ul>", strings.Trim(subFolder, "/"))

	if subFolder != "" && subFolder != "/" {
		html += "<li><span class='icon'>⬅️</span><a href='..'>.. (W górę)</a></li>"
	}

	for _, entry := range entries {
		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		icon := "📄"

		if entry.IsDir() {
			icon = "📁"
			name += "/"
		} else {
			switch ext {
			case ".mp3", ".wav", ".ogg", ".flac", ".m4a":
				icon = "🔊"
			case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg":
				icon = "🖼️"
			case ".mp4", ".mov", ".avi", ".webm":
				icon = "🎬"
			case ".zip", ".rar", ".7z", ".tar", ".gz":
				icon = "📦"
			case ".txt", ".md", ".log":
				icon = "📝"
			case ".sfm":
				icon = "🔴"
			}

		}

		html += fmt.Sprintf("<li><span class='icon'>%s</span><a href='%s'>%s</a></li>", icon, name, name)
	}

	html += "</ul></body></html>"

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
