package model

import (
	"time"
)

type User struct {
	FirstName      string
	LastName       string
	Email          string
	Password       string
	YrsOfExp       int32
	Location       string
	PhoneNumber    string
	IPAddress      string
	UserType       []string
	Stack          []string
	ProjectDetails []Project
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Project struct {
	ProjectName string
	ToolsUse    []string
	Duration    time.Duration
	ProjectTask string
}

type Email struct {
	Message      string
	Receiver     string
	Sender       string
	MailTemplate string
}
