// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserIntro struct {
	Uuid                       pgtype.UUID      `json:"uuid"`
	UserUuid                   pgtype.UUID      `json:"user_uuid"`
	DeviceUuid                 pgtype.UUID      `json:"device_uuid"`
	Role                       string           `json:"role"`
	PreferredInterfaceLanguage string           `json:"preferred_interface_language"`
	StudentGoals               string           `json:"student_goals"`
	StudentExperience          string           `json:"student_experience"`
	WhyRegistered              string           `json:"why_registered"`
	Source                     string           `json:"source"`
	CreatedAt                  pgtype.Timestamp `json:"created_at"`
}
