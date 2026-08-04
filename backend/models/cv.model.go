package models

type Resume struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	Photo       string       `json:"photo"`
	LandingPage string       `json:"landing_page"`
	LinkedIn    string       `json:"linkedin"`
	Contact     Contact      `json:"contact"`
	Summary     string       `json:"summary"`
	EULA        string       `json:"eula"`
	Experience  []Experience `json:"experience"`
	Education   []Education  `json:"education"`
	Skills      []string     `json:"skills"`
	Languages   []Language   `json:"languages"`
	Hobbies     []string     `json:"hobbies"`
}

type Contact struct {
	Phone                string `json:"phone"`
	Location             string `json:"location"`
	LocationGoogleSearch string `json:"location_google_search"`
	Email                string `json:"email"`
}

type Experience struct {
	Role    string   `json:"role"`
	From    string   `json:"from"`
	To      string   `json:"to"`
	Bullets []string `json:"bullets"`
}

type Education struct {
	School string `json:"school"`
	Field  string `json:"field"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}
