package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zgierz/harc_quiz/backend/helpers"
	"github.com/zgierz/harc_quiz/backend/permission"
	"gorm.io/gorm"
)

type GlobalController struct {
	db         *gorm.DB
	ctx        context.Context
	publicPath *gin.RouterGroup
	adminPath  *gin.RouterGroup
	Hub        *helpers.WSHub
}

func NewQuizController(db *gorm.DB, ctx context.Context, publicPath *gin.RouterGroup, adminPath *gin.RouterGroup, hub *helpers.WSHub) GlobalController {
	return GlobalController{
		db:         db,
		ctx:        ctx,
		publicPath: publicPath,
		adminPath:  adminPath,
		Hub:        hub,
	}
}

func (controller GlobalController) RegisterRoutes() {

	controller.publicPath.GET("/quizzes", controller.GetPublicQuizes)
	controller.publicPath.GET("/questions/:question_id", controller.GetQuestion)
	controller.publicPath.GET("/quiz", controller.GetPublicQuiz)
	controller.publicPath.GET("/check/answer", controller.CheckAnswer)

	// public blog

	blog := controller.publicPath.Group("/blog")
	blog.GET("/all", controller.GetAllBlogs)

	hero := controller.publicPath.Group("/hero")
	hero.GET("all", controller.GetAllBlogs)

	// public stats
	stats := controller.publicPath.Group("/stats")
	stats.GET("/all", controller.GetAllStats)
	stats.POST("/create", controller.NewStat)
	stats.GET("/count", controller.CountCompletedQuizzes)

	stats.GET("/weekly", controller.GetWeeklyStats)

	statsAdmin := controller.adminPath.Group("/stats")
	statsAdmin.DELETE("/delete/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.STAT_DELETE,
		Name:          "Usuwanie statystyki",
		Color:         "red",
		Icon:          "ri:image-add-fill",
		Description:   "Pozwala na usuwanie statystyk",
	}), controller.DeleteStat)

	controller.adminPath.POST("/upload", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.IMAGE_UPLOAD,
		Name:          "Przesyłanie obrazów",
		Color:         "blue",
		Icon:          "ri:image-add-fill",
		Description:   "Pozwala na dodawanie nowych plików graficznych do serwera",
	}), controller.SendImage)

	controller.adminPath.GET("/images/list", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.IMAGE_GET,
		Name:          "Lista obrazów",
		Color:         "lime",
		Icon:          "ri:image-line",
		Description:   "Podgląd wszystkich dostępnych grafik",
	}), controller.ListImages)

	controller.adminPath.DELETE("/images/delete", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.IMAGE_DELETE,
		Name:          "Usuwanie obrazów",
		Color:         "red",
		Icon:          "ri:image-edit-fill",
		Description:   "Usuwanie plików graficznych z galerii",
	}), controller.DeleteImage)

	admin := controller.adminPath.Group("/quizzes")
	admin.PUT("/update/basic/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_UPDATE_BASIC,
		Name:          "Edycja podstawowa quizu",
		Color:         "rose",
		Icon:          "ri:edit-box-line",
		Description:   "Zmiana tytułu, opisu i ustawień głównych quizu",
	}), controller.UpdateBasicInfo)

	admin.PUT("/update/image/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_UPDATE_IMAGE,
		Name:          "Dodanie miniatury",
		Color:         "rose",
		Icon:          "material-symbols:thumbnail-bar-sharp",
		Description:   "Zmiana miniatury quizu",
	}), controller.UpdateQuizImage)

	admin.PUT("/save/basic", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_SAVE_BASIC,
		Name:          "Edycja podstawowa quizu",
		Color:         "rose",
		Icon:          "ri:edit-box-line",
		Description:   "Zmiana tytułu, opisu i ustawień głównych quizu",
	}), controller.SaveBasicInfo)

	admin.PUT("/update/settings/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_UPDATE_SETTINGS,
		Name:          "Edycja ustawien quizu",
		Color:         "indigo",
		Icon:          "ooui:page-settings",
		Description:   "Zmiana ustawień quizu",
	}), controller.UpdateQuizSettings)

	admin.PUT("/update/field/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_QUICK_UPDATE,
		Name:          "Szybka edycja pola",
		Icon:          "ri:edit-circle-fill",
		Color:         "text-blue-300",
		Description:   "Zmiana pojedynczych parametrów quizu",
	}), controller.UpdateOneField)

	admin.GET("/all", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_ALL,
		Name:          "Podgląd wszystkich quizów",
		Icon:          "ri:list-check",
		Color:         "rose",
		Description:   "Dostęp do pełnej listy quizów w panelu",
	}), controller.GetAdminQuizzes)

	admin.GET("/quiz/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.GET_QUIZ,
		Name:          "Szczegóły quizu",
		Color:         "blue",
		Icon:          "ri:file-search-fill",
		Description:   "Podgląd edycji konkretnego arkusza",
	}), controller.GetAdminQuiz)

	admin.POST("", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_CREATE,
		Name:          "Tworzenie quizów",
		Icon:          "ri:add-box-fill",
		Color:         "blue",
		Description:   "Dodawanie nowych quizów do bazy danych",
	}), controller.CreateQuiz)

	admin.PUT("/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_UPDATE,
		Name:          "Pełna aktualizacja quizu",
		Color:         "rose",
		Icon:          "ri:save-3-fill",
		Description:   "Główna operacja zapisu zmian w quizie",
	}), controller.UpdateQuiz)

	admin.DELETE("/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUIZ_DELETE,
		Name:          "Usuwanie quizu",
		Icon:          "ri:delete-bin-7-fill",
		Color:         "red",
		Description:   "Nieodwracalne usunięcie quizu z systemu",
	}), controller.DeleteQuiz)

	// Blog
	blogAdmin := controller.adminPath.Group("/blog")
	blogAdmin.POST("/create", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.BLOG_CREATE,
		Name:          "Tworzenie wpisów",
		Icon:          "ri:article-fill",
		Color:         "blue",
		Description:   "Dodawanie nowych artykułów na bloga",
	}), controller.CreateBlog)

	blogAdmin.DELETE("/delete/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.BLOG_DELETE,
		Name:          "Usuwanie wpisów",
		Icon:          "ri:chat-delete-fill",
		Color:         "red",
		Description:   "Usuwanie artykułów z bazy",
	}), controller.DeleteBlog)

	blogAdmin.PUT("/update/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.BLOG_UPDATE,
		Name:          "Edycja wpisów",
		Icon:          "ri:edit-2-fill",
		Color:         "rose",
		Description:   "Aktualizacja treści istniejących postów",
	}), controller.UpdateBlog)

	// Hero
	heroAdmin := controller.adminPath.Group("/hero")
	heroAdmin.POST("/create", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.HERO_CREATE,
		Name:          "Dodawanie Hero",
		Icon:          "ri:layout-top-fill",
		Color:         "blue",
		Description:   "Tworzenie nowych sekcji Hero",
	}), controller.CreateHero)

	heroAdmin.DELETE("/delete/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.HERO_DELETE,
		Name:          "Usuwanie Hero",
		Color:         "red",
		Icon:          "ri:layout-bottom-fill",
		Description:   "Usuwanie sekcji nagłówkowych",
	}), controller.DeleteHero)

	heroAdmin.PUT("/update/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.HERO_UPDATE,
		Name:          "Edycja Hero",
		Color:         "blue",
		Icon:          "material-symbols:deployed-code-update",
		Description:   "Zmiana treści i grafik w sekcjach Hero",
	}), controller.UpdateHero)

	// Questions
	questions := controller.adminPath.Group("/questions")
	questions.PUT("/update/many", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUESTION_UPDATE_MANY,
		Name:          "Masowa edycja pytań",
		Icon:          "ri:stack-fill",
		Color:         "indigo",
		Description:   "Aktualizacja wielu pytań naraz",
	}), controller.UpdateQuestions)

	questions.GET("/all", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUESTION_ALL,
		Name:          "Wszystkie pytania",
		Icon:          "ri:question-answer-fill",
		Color:         "indigo",
		Description:   "Podgląd bazy pytań",
	}), controller.GetAdminQuestions)

	questions.POST("/create/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUESTION_ASIGN_TO_QUIZ,
		Name:          "Dodawanie pytania",
		Icon:          "ri:chat-new-fill",
		Color:         "indigo",
		Description:   "Tworzenie nowego pytania w quizie",
	}), controller.CreateQuestion)

	questions.PUT("/update/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUESTION_UPDATE,
		Name:          "Edycja pytania",
		Icon:          "ri:questionnaire-fill",
		Color:         "indigo",
		Description:   "Zmiana treści lub odpowiedzi w pytaniu",
	}), controller.UpdateQuestion)

	questions.DELETE("/delete/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.QUESTION_DELETE,
		Name:          "Usuwanie pytania",
		Icon:          "ri:question-line",
		Color:         "indigo",
		Description:   "Usuwanie pytania z arkusza",
	}), controller.DeleteQuestion)

	// Contributors
	contributors := controller.adminPath.Group("/contributors")
	contributors.GET("/all", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_ALL,
		Name:          "Lista kontrybutorów",
		Icon:          "ri:team-fill",
		Color:         "indigo",
		Description:   "Podgląd osób zarządzających panelem",
	}), controller.GetContributors)

	contributors.PUT("/update/details/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_UPDATE_DETAILS,
		Name:          "Edycja danych kontrybutora",
		Icon:          "ri:user-settings-fill",
		Color:         "indigo",
		Description:   "Zmiana nazwy lub loginu użytkownika",
	}), controller.UpdateContributorDetails)

	contributors.POST("/create", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_CREATE,
		Name:          "Dodawanie kontrybutora",
		Icon:          "ri:user-settings-fill",
		Color:         "indigo",
		Description:   "Dodaj nowego kontrybutora",
	}), controller.CreateContributor)

	contributors.PUT("/update/password/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_UPDATE_PASSWORD,
		Name:          "Zmiana hasła",
		Icon:          "ri:lock-password-fill",
		Color:         "rose",
		Description:   "Możliwość zmiany hasła kontrybutora",
	}), controller.UpdateContributorPassword)

	contributors.PUT("/update/permissions/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_UPDATE_PERMISSIONS,
		Name:          "Zarządzanie uprawnieniami",
		Icon:          "ri:shield-user-fill",
		Color:         "rose",
		Description:   "Przydzielanie i odbieranie dostępów",
	}), controller.UpdateContributorPermissions)

	contributors.DELETE("/delete/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_DELETE,
		Name:          "Usuwanie kontrybutora",
		Icon:          "ri:user-unfollow-fill",
		Color:         "red",
		Description:   "Całkowite usunięcie użytkownika z systemu",
	}), controller.DeleteContributor)

	contributors.PUT("/switch/block/:id", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_BLOCK,
		Name:          "Blokowanie kontrybutora",
		Icon:          "ri:user-forbid-fill",
		Color:         "amber",
		Description:   "Tymczasowe zawieszanie dostępu do panelu",
	}), controller.SwitchContributorBlock)

	contributors.GET("/view/permissions", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.CONTRIBUTOR_PERMISSIONS,
		Name:          "Przegląd uprawnień",
		Icon:          "ri:key-2-fill",
		Color:         "cyan",
		Description:   "Dostęp do listy wszystkich tagów uprawnień",
	}), controller.GetPermissionsList)

	controller.publicPath.GET("/test/redirect", helpers.RequirePermission(helpers.PermissionsMetadata{
		PermissionTag: permission.REDIRECT,
		Name:          "Testowy redirect",
		Icon:          "ri:share-forward-fill",
		Color:         "lime",
		Description:   "Uprawnienie do testów przekierowań",
	}), controller.TestRedirect)
}
