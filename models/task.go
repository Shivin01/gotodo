package models

import "gorm.io/gorm"

// Status of Enum type
// TODO:
// 1. Comments
// 2. History
// 3. Attachments

type STATUS int

const (
	TODO STATUS = iota
	INPROGRESS
	DONE
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Status      STATUS
	Reporter    User
	Assignee    User
}
