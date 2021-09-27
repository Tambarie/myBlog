package models

import "time"

type Blog struct {
	ID string
	Author string
	Title string
	Body string
	Time time.Time
}
