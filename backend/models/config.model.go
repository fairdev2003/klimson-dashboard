package models

type Config struct {
	DisplayName      string `json:"display_name"`
	InternalName     string `json:"internal_name"`
	Desc             string `json:"description"`
	Port             int    `json:"port"`
	Debug            bool   `json:"debug"`
	Api              Api    `json:"api_config"`
	ServerConfigPath string `json:"config_path"`
}

type Api struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}
