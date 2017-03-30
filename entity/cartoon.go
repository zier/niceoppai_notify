package entity

import (
	"fmt"
)

// Cartoon ...
type Cartoon struct {
	URL          string
	Thumbnail    string
	Name         string
	ChapterTitle string
}

// NewCartoon ...
func NewCartoon() *Cartoon {
	return &Cartoon{
		URL:          "",
		Thumbnail:    "",
		Name:         "",
		ChapterTitle: "",
	}
}

// GetURL ...
func (c *Cartoon) GetURL() string {
	return fmt.Sprintf("%s?all", c.URL)
}
