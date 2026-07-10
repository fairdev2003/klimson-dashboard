package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) GetTables(ctx *gin.Context) {
	type TableResponse struct {
		Table string `json:"table"` // e.g. "quizzes"
		Name  string `json:"name"`  // e.g. "Quizzes"
		Icon  string `json:"icon"`  // e.g. "quiz-icon"
	}

	var tables []TableResponse

	for _, m := range models.MigratableModels {
		stmt := &gorm.Statement{DB: controller.db}
		if err := stmt.Parse(m.Model); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to parse model schema",
				"details": err.Error(),
			})
			return
		}

		tables = append(tables, TableResponse{
			Table: stmt.Schema.Table,
			Name:  m.Name,
			Icon:  m.Icon,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tables": tables,
	})
}

func (controller GlobalController) GetTableData(ctx *gin.Context) {
	tableName := ctx.Param("table_name")

	columnTypes, err := controller.db.Migrator().ColumnTypes(tableName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve column structure for table " + tableName,
			"details": err.Error(),
		})
		return
	}

	var currentModel any
	stmt := &gorm.Statement{DB: controller.db}

	for _, m := range models.MigratableModels {
		if err := stmt.Parse(m.Model); err == nil {
			if stmt.Schema.Table == tableName {
				currentModel = m.Model
				break
			}
		}
	}

	if currentModel == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Table " + tableName + " is not registered in models",
		})
		return
	}

	if err := stmt.Parse(currentModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var columns []map[string]string
	for _, col := range columnTypes {
		isForeignKey := "false"

		for _, rel := range stmt.Schema.Relationships.Relations {
			for _, reference := range rel.References {
				if reference.ForeignKey != nil && reference.ForeignKey.DBName == col.Name() {
					isForeignKey = "true"
					break
				}
			}
			if isForeignKey == "true" {
				break
			}
		}

		columns = append(columns, map[string]string{
			"name":        col.Name(),
			"type":        col.DatabaseTypeName(),
			"slug":        col.Name(),
			"foreign_key": isForeignKey,
		})
	}

	var results []map[string]interface{}
	err = controller.db.Table(tableName).Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve data from table " + tableName,
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"table":   tableName,
		"columns": columns,
		"count":   len(results),
		"data":    results,
	})
}
