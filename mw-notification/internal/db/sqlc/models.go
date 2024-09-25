// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type NotificationChannel string

const (
	NotificationChannelMail   NotificationChannel = "mail"
	NotificationChannelWebapp NotificationChannel = "webapp"
)

func (e *NotificationChannel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationChannel(s)
	case string:
		*e = NotificationChannel(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationChannel: %T", src)
	}
	return nil
}

type NullNotificationChannel struct {
	NotificationChannel NotificationChannel `json:"notification_channel"`
	Valid               bool                `json:"valid"` // Valid is true if NotificationChannel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationChannel) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationChannel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationChannel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationChannel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NotificationChannel), nil
}

type NotificationNature string

const (
	NotificationNaturePrivateChat      NotificationNature = "private_chat"
	NotificationNatureGroupChat        NotificationNature = "group_chat"
	NotificationNatureOwnWay           NotificationNature = "own_way"
	NotificationNatureMentoringWay     NotificationNature = "mentoring_way"
	NotificationNatureMentoringRequest NotificationNature = "mentoring_request"
	NotificationNatureFavoriteWay      NotificationNature = "favorite_way"
)

func (e *NotificationNature) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationNature(s)
	case string:
		*e = NotificationNature(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationNature: %T", src)
	}
	return nil
}

type NullNotificationNature struct {
	NotificationNature NotificationNature `json:"notification_nature"`
	Valid              bool               `json:"valid"` // Valid is true if NotificationNature is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationNature) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationNature, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationNature.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationNature) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NotificationNature), nil
}

type EnabledNotification struct {
	Uuid      pgtype.UUID         `json:"uuid"`
	UserUuid  pgtype.UUID         `json:"user_uuid"`
	Nature    NotificationNature  `json:"nature"`
	Channel   NotificationChannel `json:"channel"`
	IsEnabled bool                `json:"is_enabled"`
}

type Notification struct {
	Uuid        pgtype.UUID        `json:"uuid"`
	UserUuid    pgtype.UUID        `json:"user_uuid"`
	IsRead      bool               `json:"is_read"`
	Description pgtype.Text        `json:"description"`
	Url         pgtype.Text        `json:"url"`
	Nature      NotificationNature `json:"nature"`
	CreatedAt   pgtype.Timestamp   `json:"created_at"`
}