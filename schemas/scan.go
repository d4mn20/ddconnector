package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Scan struct {
	gorm.Model
	Product    string
	Engagement string
	Test       string
	FilePath   string
	Branch     string
	RepoUrl    string
	Origin     string
}

type ScanResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	Product    string    `json:"product"`
	Engagement string    `json:"engagement"`
	Test       string    `json:"test"`
	Branch     string    `json:"branch"`
	RepoUrl    string    `json:"repoUrl"`
	Origin     string    `json:"origin"`
	FilePath   string    `json:"filepath"`
}
