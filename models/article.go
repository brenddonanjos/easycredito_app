package models

import "time"

type Article struct {
	Id          int    `json:"id" db:"id"`
	Featured    bool   `json:"featured" db:"featured"`
	Title       string `json:"title" db:"title"`
	Url         string `json:"url" db:"url"`
	ImageUrl    string `json:"imageUrl" db:"imageUrl"`
	NewsSite    string `json:"newsSite" db:"newsSite"`
	Summary     string `json:"summary" db:"summary"`
	PublishedAt string `json:"publishedAt" db:"publishedAt"`
	Launches    Launches
	Events      Events
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Articles []Article
