// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (owner_uuid, room_uuid, text)
VALUES ($1, $2, $3)
RETURNING uuid, owner_uuid, text
`

type CreateMessageParams struct {
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	RoomUuid  pgtype.UUID `json:"room_uuid"`
	Text      string      `json:"text"`
}

type CreateMessageRow struct {
	Uuid      pgtype.UUID `json:"uuid"`
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	Text      string      `json:"text"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (CreateMessageRow, error) {
	row := q.db.QueryRow(ctx, createMessage, arg.OwnerUuid, arg.RoomUuid, arg.Text)
	var i CreateMessageRow
	err := row.Scan(&i.Uuid, &i.OwnerUuid, &i.Text)
	return i, err
}

const getMessagesByRoomUUID = `-- name: GetMessagesByRoomUUID :many
SELECT
    messages.uuid,
    messages.owner_uuid,
    messages.text,
    ARRAY(
        SELECT receiver_uuid
        FROM message_status
        WHERE messages.uuid = message_status.message_uuid
            AND message_status.is_read = true
        ORDER BY updated_at DESC
    )::UUID[] AS message_status_user_uuids,
    ARRAY(
        SELECT updated_at
        FROM message_status
        WHERE messages.uuid = message_status.message_uuid
            AND message_status.is_read = true
        ORDER BY updated_at DESC
    )::TIMESTAMP[] AS message_status_updated_at
FROM messages
WHERE room_uuid = $1
LIMIT 100
`

type GetMessagesByRoomUUIDRow struct {
	Uuid                   pgtype.UUID        `json:"uuid"`
	OwnerUuid              pgtype.UUID        `json:"owner_uuid"`
	Text                   string             `json:"text"`
	MessageStatusUserUuids []pgtype.UUID      `json:"message_status_user_uuids"`
	MessageStatusUpdatedAt []pgtype.Timestamp `json:"message_status_updated_at"`
}

func (q *Queries) GetMessagesByRoomUUID(ctx context.Context, roomUuid pgtype.UUID) ([]GetMessagesByRoomUUIDRow, error) {
	rows, err := q.db.Query(ctx, getMessagesByRoomUUID, roomUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMessagesByRoomUUIDRow{}
	for rows.Next() {
		var i GetMessagesByRoomUUIDRow
		if err := rows.Scan(
			&i.Uuid,
			&i.OwnerUuid,
			&i.Text,
			&i.MessageStatusUserUuids,
			&i.MessageStatusUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
