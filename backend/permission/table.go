package permission

const (

	// image
	IMAGE_UPLOAD = "image:upload"
	IMAGE_GET    = "image:get"
	IMAGE_DELETE = "image:delete:$"

	// quiz
	QUIZ_UPDATE_BASIC    = "quiz:update:basic:$"
	QUIZ_UPDATE_IMAGE    = "quiz:update:image:$"
	QUIZ_SAVE_BASIC      = "quiz:update:basic:$"
	QUIZ_UPDATE_SETTINGS = "quiz:update:settings:$"
	QUIZ_QUICK_UPDATE    = "quiz:update:one_field"
	QUIZ_ALL             = "quiz:get:all"
	GET_QUIZ             = "quiz:get:$"
	QUIZ_CREATE          = "quiz:create"
	QUIZ_UPDATE          = "quiz:update:$"
	QUIZ_DELETE          = "quiz:delete:$"

	// blog
	BLOG_CREATE = "blog:create"
	BLOG_DELETE = "blog:delete:$"
	BLOG_UPDATE = "blog:update:$"

	// hero
	HERO_CREATE = "hero:create"
	HERO_DELETE = "hero:delete:$"
	HERO_UPDATE = "hero:update:$"

	// question
	QUESTION_UPDATE_MANY   = "question:update:$"
	QUESTION_ALL           = "question:all"
	QUESTION_ASIGN_TO_QUIZ = "question:create:$"
	QUESTION_UPDATE        = "question:update:$"
	QUESTION_DELETE        = "question:delete:$"

	STAT_DELETE = "stat:delete:$"

	// contributor

	// /admin/contributors/all - wszyscy kontrybutorzy
	//  - fetch wszystkich kontrybutorów
	CONTRIBUTOR_ALL = "contributor:all"

	// /admin/contributors/update/details/:id
	//  - zmiana podstawowych informacji o kontrybutorze
	//  - mozliwosc ustawienia id lub wildcard (*)
	CONTRIBUTOR_UPDATE_DETAILS = "contributor:update:details:$"

	// /admin/contributors/update/password/:id
	//  - zmiana hasła dla kontrybutora
	CONTRIBUTOR_UPDATE_PASSWORD = "contributor:update:password"

	// /admin/contributors/update/permissions/:id
	//  - zmiana uprawnienień dla kontrybutora
	//  - mozliwosc ustawienia id lub wildcard (*)
	CONTRIBUTOR_UPDATE_PERMISSIONS = "contributor:update:permissions:$"

	// /admin/contributors/delete/:id
	//  - usuwanie kontrybutora
	//  - mozliwosc ustawienia id lub wildcard (*)
	CONTRIBUTOR_DELETE = "contributor:delete:$"

	// /admin/contributors/switch/block/:id
	//  - zablokowanie dostępu do panelu
	//  - mozliwosc ustawienia id lub wildcard (*)
	CONTRIBUTOR_BLOCK = "contributor:switch:block:$"

	// /admin/contributors/view/permissions
	//  - usuwanie kontrybutora
	CONTRIBUTOR_PERMISSIONS = "contributor:view:permissions"

	// /admin/contributors/create
	//  - dodawanie kontrybutora
	CONTRIBUTOR_CREATE = "contributor:create"

	// /admin/test/redirect
	//  - test przeniesienia
	REDIRECT = "harc:redirect"
)
