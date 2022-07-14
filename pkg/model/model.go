package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" validate:"required"`
	LastName       *string            `json:"last_name" validate:"required"`
	Email          *string            `json:"email" validate:"required"`
	Password       *string            `json:"password" validate:"required"`
	YrsOfExp       *int32             `json:"yrs_of_exp"`
	Location       *string            `json:"location" validate:"required"`
	PhoneNumber    *string            `json:"phone_number" validate:"required"`
	IPAddress      *string            `json:"ip_address"`
	UserType       []string           `json:"user_type"`
	Stack          []string           `json:"stack"`
	ProjectDetails []Project          `json:"project_details" bson:"project_details"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type Project struct {
	ID          primitive.ObjectID `bson:"_id"`
	ProjectName *string            `json:"project_name"`
	ToolsUse    []string           `json:"tools_use"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Duration    time.Duration      `json:"duration"`
	ProjectTask *string            `json:"project_tasks"`
}

type Email struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Message      *string            `json:"message"`
	Receiver     *string            `json:"receiver" validate:"required"`
	Sender       *string            `json:"sender" validate:"required"`
	MailTemplate *string            `json:"mail_template"`
}
