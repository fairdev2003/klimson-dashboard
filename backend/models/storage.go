package models

type Storage struct {
	BaseModel
	// you can change name of the file, default will be the filename.
	Name string `json:"name" gorm:"column:name;not null;comment:Resource name"`
	// not really needed, you can add description
	Description string `json:"description" gorm:"type:text;comment:Resource description"`
	// image.png - file name and file type
	Filename string `json:"file_name" gorm:"index;not null"`
	// image - file type will be based on extension
	Filetype string `json:"file_type" gorm:"size:50"`
	// path to file without filename for example: /klimson.dev/images
	Path string `json:"path" `
}
