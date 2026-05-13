package config

import "github.com/zgierz/harc_quiz/backend/models"

var PreGeneratedConfig = &models.Config{
	DisplayName:      "HarcQuiz Klimson Api",
	Desc:             "Serwer największej bazy quizów harcerskich",
	InternalName:     "harcquiz-api",
	ServerConfigPath: "./",
	Port:             8080,
	Debug:            true,
	Api: models.Api{
		Path:    "/harc-api",
		Version: "v1",
	}}
