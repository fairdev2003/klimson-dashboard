package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zgierz/klimson/backend/api"
	"github.com/zgierz/klimson/backend/models"
	"gorm.io/gorm"
)

func (controller GlobalController) GetTables(ctx *gin.Context) {
	type TableResponse struct {
		Table string `json:"table"` 
		Name  string `json:"name"`  
		Icon  string `json:"icon"` 
	}

	var tables []TableResponse

	for _, m := range models.MigratableModels {
		stmt := &gorm.Statement{DB: controller.db}
		if err := stmt.Parse(m.Model); err != nil {
			api.InternalServerErrorResponse(ctx, nil, "Failed to parse model schema: "+err.Error())
			return
		}

		tables = append(tables, TableResponse{
			Table: stmt.Schema.Table,
			Name:  m.Name,
			Icon:  m.Icon,
		})
	}

	api.SuccessResponse(ctx, gin.H{
		"tables": tables,
	})
}

func (controller GlobalController) GetTableData(ctx *gin.Context) {
	tableName := ctx.Param("table_name")

	columnTypes, err := controller.db.Migrator().ColumnTypes(tableName)
	if err != nil {
		api.InternalServerErrorResponse(ctx, nil, "Failed to retrieve column structure for table "+tableName+": "+err.Error())
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
		api.NotFoundResponse(ctx)
		return
	}

	if err := stmt.Parse(currentModel); err != nil {
		api.InternalServerErrorResponse(ctx, nil, err.Error())
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
		api.InternalServerErrorResponse(ctx, nil, "Failed to retrieve data from table "+tableName+": "+err.Error())
		return
	}

	api.SuccessResponse(ctx, gin.H{
		"table":   tableName,
		"columns": columns,
		"count":   len(results),
		"data":    results,
	})
}
