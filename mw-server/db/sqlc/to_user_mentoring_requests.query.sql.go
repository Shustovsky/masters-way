// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: to_user_mentoring_requests.query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createToUserMentoringRequest = `-- name: CreateToUserMentoringRequest :one
INSERT INTO to_user_mentoring_requests(
    user_uuid,
    way_uuid
) VALUES (
    $1, $2
) RETURNING user_uuid, way_uuid
`

type CreateToUserMentoringRequestParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	WayUuid  uuid.UUID `json:"way_uuid"`
}

func (q *Queries) CreateToUserMentoringRequest(ctx context.Context, arg CreateToUserMentoringRequestParams) (ToUserMentoringRequest, error) {
	row := q.queryRow(ctx, q.createToUserMentoringRequestStmt, createToUserMentoringRequest, arg.UserUuid, arg.WayUuid)
	var i ToUserMentoringRequest
	err := row.Scan(&i.UserUuid, &i.WayUuid)
	return i, err
}

const deleteToUserMentoringRequestByIds = `-- name: DeleteToUserMentoringRequestByIds :exec
DELETE FROM to_user_mentoring_requests
WHERE user_uuid = $1 AND way_uuid = $2
`

type DeleteToUserMentoringRequestByIdsParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	WayUuid  uuid.UUID `json:"way_uuid"`
}

func (q *Queries) DeleteToUserMentoringRequestByIds(ctx context.Context, arg DeleteToUserMentoringRequestByIdsParams) error {
	_, err := q.exec(ctx, q.deleteToUserMentoringRequestByIdsStmt, deleteToUserMentoringRequestByIds, arg.UserUuid, arg.WayUuid)
	return err
}

const getToMentorUserRequestsByWayId = `-- name: GetToMentorUserRequestsByWayId :many
SELECT users.uuid from ways
JOIN to_user_mentoring_requests ON to_user_mentoring_requests.way_uuid = ways.uuid
JOIN users ON users.uuid = to_user_mentoring_requests.user_uuid
WHERE way_uuid = $1
`

func (q *Queries) GetToMentorUserRequestsByWayId(ctx context.Context, wayUuid uuid.UUID) ([]uuid.UUID, error) {
	rows, err := q.query(ctx, q.getToMentorUserRequestsByWayIdStmt, getToMentorUserRequestsByWayId, wayUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var uuid uuid.UUID
		if err := rows.Scan(&uuid); err != nil {
			return nil, err
		}
		items = append(items, uuid)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}