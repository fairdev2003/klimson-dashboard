package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/models"
)

func (gc GlobalController) GetFile(c *gin.Context) {
	filePath := c.Param("filepath")

	fullPath := "./static/uploads" + filePath

	c.File(fullPath)
}

func (gc GlobalController) NewFile(ctx *gin.Context) {
	filePath := ctx.Param("filepath")

	fullPath := "./static/uploads" + filePath

	var requestBody struct {
		Content string `json:"content"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		api.BadRequestResponse(ctx, nil, err.Error())
		return
	}

	err := os.WriteFile(fullPath, []byte(requestBody.Content), 0644)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
		return
	}

	api.SuccessResponse(ctx, nil, "Success!")

}

func (gc GlobalController) GetSecuredFile(c *gin.Context) {
	filePath := c.Param("filepath")

	cleanPath := filepath.Clean(filePath)

	if strings.Contains(cleanPath, "passwords.txt") {
		api.UnauthorizedResponse(c, nil, "Access denied to config files")
		return
	}

	fullPath := filepath.Join("./secured", cleanPath)

	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) || info.IsDir() {
		api.NotFoundResponse(c)
		return
	}

	c.File(fullPath)
}

func (gc *GlobalController) PushChangedTextFile(c *gin.Context) {
	filePath := c.Param("filepath")

	var requestBody struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		api.BadRequestResponse(c, nil, "Błędne dane")
		return
	}

	fullPath := filepath.Join("./static/uploads", filePath)

	err := os.WriteFile(fullPath, []byte(requestBody.Content), 0644)
	if err != nil {
		api.InternalServerErrorResponse(c, nil, err.Error())
		return
	}

	api.SuccessResponse(c, nil, "Plik zapisany pomyślnie")
}

func (gc GlobalController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		api.BadRequestResponse(c, nil, "Brak pliku w żądaniu")
		return
	}

	folderPath := c.Param("folder")
	dst := filepath.Join("./static/uploads", folderPath, file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		api.InternalServerErrorResponse(c, nil, "Błąd zapisu pliku")
		return
	}

	api.SuccessResponse(c, nil, "Plik przesłany pomyślnie")
}

func (gc GlobalController) CreateFolder(c *gin.Context) {
	var input struct {
		FolderName string `json:"folder_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		api.BadRequestResponse(c, nil, "Nieprawidłowe dane")
		return
	}

	folderPath := c.Param("folder")
	fullPath := filepath.Join("./static/uploads", folderPath, input.FolderName)

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		api.BadRequestResponse(c, nil, "Folder o tej nazwie już istnieje")
		return
	}

	if err := os.MkdirAll(fullPath, 0755); err != nil {
		api.InternalServerErrorResponse(c, nil, "Błąd tworzenia folderu")
		return
	}

	api.StatusCreatedResponse(c, nil, "Folder stworzony")
}

func (gc GlobalController) DeleteFileOrFolder(c *gin.Context) {
	targetPath := c.Param("folder")
	fullPath := filepath.Join("./static/uploads", targetPath)

	if !strings.HasPrefix(fullPath, filepath.Clean("./static/uploads")) {
		api.BadRequestResponse(c, nil, "Błędna ścieżka")
		return
	}

	err := os.RemoveAll(fullPath)
	if err != nil {
		api.InternalServerErrorResponse(c, nil, "Nie udało się usunąć")
		return
	}

	api.SuccessResponse(c, nil, "Usunięto pomyślnie")
}

func (gc GlobalController) RenameItem(c *gin.Context) {
	oldPath := c.Param("folder")
	var req struct {
		NewName string `json:"newName"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		api.BadRequestResponse(c, nil, "Błędne dane")
		return
	}

	oldFullPath := filepath.Join("./static/uploads", oldPath)

	dir := filepath.Dir(oldFullPath)
	newFullPath := filepath.Join(dir, req.NewName)

	err := os.Rename(oldFullPath, newFullPath)
	if err != nil {
		api.InternalServerErrorResponse(c, nil, "Nie udało się zmienić nazwy")
		return
	}

	api.SuccessResponse(c, nil, "Nazwa zmieniona")
}

func (gc GlobalController) ListFiles(c *gin.Context) {
	folderPath := c.Param("folder")
	fullPath := "./static/uploads" + folderPath

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		api.InternalServerErrorResponse(c, nil, "Błąd odczytu folderu")
		return
	}

	var fileList []models.ListRecord
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		fileList = append(fileList, models.ListRecord{
			Name:         entry.Name(),
			IsDir:        entry.IsDir(),
			ModifiedTime: info.ModTime(),
			Size:         info.Size(),
		})
	}

	api.SuccessResponse(c, fileList)
}

func (gc GlobalController) Interface(c *gin.Context) {
	subFolder := c.Param("folder")
	basePath := "./static/uploads"
	fullPath := filepath.Join(basePath, subFolder)

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		api.NotFoundResponse(c)
		return
	}

	if !fileInfo.IsDir() {
		c.File(fullPath)
		return
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		api.InternalServerErrorResponse(c, nil, "Błąd odczytu folderu")
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
