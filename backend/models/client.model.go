package models

// redis sync-in config
type ClientConfig struct {
	DashboardTheme     string     `json:"theme"`
	CodeTheme          string     `json:"code_theme"`
	SidebarPreferences []string   `json:"client_pills"`
	Dock               *bool      `json:"dock_on"`
	Bookmarks          []Bookmark `json:"bookmarks"`
	SidebarBehavior    string     `json:"sidebarBehavior"`
}

type Bookmark struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Href  string `json:"href"`
	Color string `json:"color"`
}
