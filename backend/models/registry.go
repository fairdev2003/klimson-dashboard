package models

type MigrationType struct {
	Model any
	Name  string
	Icon  string
}

var MigratableModels = []MigrationType{
	{Model: &Quiz{}, Name: "Quizzes", Icon: "lucide:help-circle"},
	{Model: &Question{}, Name: "Questions", Icon: "lucide:file-question"},
	{Model: &Answer{}, Name: "Answers", Icon: "lucide:message-square-text"},
	{Model: &Blog{}, Name: "Blogs", Icon: "lucide:book-open-text"},
	{Model: &Hero{}, Name: "Heroes", Icon: "lucide:shield"},
	{Model: &Stat{}, Name: "Statistics", Icon: "lucide:bar-chart-3"},
	{Model: &Contributor{}, Name: "Contributors", Icon: "lucide:users-2"},
	{Model: &Log{}, Name: "Logs", Icon: "lucide:terminal"},
	{Model: &Role{}, Name: "Roles", Icon: "lucide:key-round"},
	{Model: &ContextStorage{}, Name: "Context Storage", Icon: "material-symbols:contextual-token"},
}
