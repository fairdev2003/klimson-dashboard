package models

import "time"

type ListRecord struct {
	Name         string    `json:"name"`
	IsDir        bool      `json:"is_dir"`
	ModifiedTime time.Time `json:"modified"`
	Size         int64     `json:"file_size"`
}
