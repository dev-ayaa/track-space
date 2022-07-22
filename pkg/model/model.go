package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      string             `form:"first_name" json:"first_name" Usage:"required max=32 min=3" binding:"required max=32 min=3"`
	LastName       string             `form:"last_name" json:"last_name" Usage:"required max=32 min=3" binding:"required max=32 min=3"`
	Email          string             `form:"email" json:"email" Usage:"email required" binding:"required max=32 min=3" `
	Password       string             `form:"password" json:"password" Usage:"required_with=Email alphanum" binding:"required max=32 min=3"`
	YrsOfExp       string             `form:"yrs_of_exp" json:"yrs_of_exp" Usage:"numeric omitempty"`
	Country        string             `form:"country" json:"country" Usage:"required" binding:"required max=32 min=3"`
	PhoneNumber    string             `form:"phone_number" json:"phone_number" Usage:"required max=15 min=8" binding:"required max=32 min=3"`
	IPAddress      string             `form:"ip_address" json:"ip_address"`
	Address        string             `form:"address" json:"address" Usage:"required" binding:"required max=32 min=3"`
	UserType       []string           `form:"user_type" json:"user_type" Usage:"omitempty"`
	Stack          []string           `form:"stack" json:"stack" Usage:"omitempty"`
	ProjectDetails []Project          `form:"project_details" json:"project_details" bson:"project_details"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id"`
	ProjectName string             `json:"project_name"`
	ToolsUse    []string           `json:"tools_use"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Duration    time.Duration      `json:"duration"`
	ProjectTask string             `json:"project_tasks"`
}

type Email struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Message      string             `json:"message"`
	Receiver     string             `json:"receiver" validate:"required"`
	Sender       string             `json:"sender" validate:"required"`
	MailTemplate string             `json:"mail_template"`
}
