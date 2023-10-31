package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Address   string             `json:"address" validate:"required"`
	City      string             `json:"city" validate:"required"`
	Country   string             `json:"country" validate:"required"`
	Pincode   string             `json:"pin_code" validate:"required"`
	Marks     int32              `json:"marks" validate:"required,min=0,max=1600"`
	Passed    bool               `json:"passed"`
	UserId    string             `json:"user_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UpdateMarksInput struct {
	Name  string `json:"name" validate:"required"`
	Marks int32  `json:"marks" validate:"required"`
}

type UserInput struct {
	Name string `json:"name" validate:"required"`
}
