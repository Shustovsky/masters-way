// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: project.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUserToProject = `-- name: AddUserToProject :one
INSERT INTO users_projects(
    user_uuid,
    project_uuid
) VALUES (
    $1,
    $2
) RETURNING user_uuid, project_uuid
`

type AddUserToProjectParams struct {
	UserUuid    pgtype.UUID `json:"user_uuid"`
	ProjectUuid pgtype.UUID `json:"project_uuid"`
}

func (q *Queries) AddUserToProject(ctx context.Context, arg AddUserToProjectParams) (UsersProject, error) {
	row := q.db.QueryRow(ctx, addUserToProject, arg.UserUuid, arg.ProjectUuid)
	var i UsersProject
	err := row.Scan(&i.UserUuid, &i.ProjectUuid)
	return i, err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects(
    name,
    owner_uuid
) VALUES (
    $1,
    $2
) RETURNING
    uuid,
    name,
    owner_uuid,
    is_private,
    COALESCE(
            ARRAY(
                SELECT user_uuid
                FROM users_projects
                WHERE users_projects.project_uuid = uuid
            ),
            '{}'
    )::VARCHAR[] AS user_uuids
`

type CreateProjectParams struct {
	Name      string      `json:"name"`
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
}

type CreateProjectRow struct {
	Uuid      pgtype.UUID `json:"uuid"`
	Name      string      `json:"name"`
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	IsPrivate bool        `json:"is_private"`
	UserUuids []string    `json:"user_uuids"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (CreateProjectRow, error) {
	row := q.db.QueryRow(ctx, createProject, arg.Name, arg.OwnerUuid)
	var i CreateProjectRow
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.OwnerUuid,
		&i.IsPrivate,
		&i.UserUuids,
	)
	return i, err
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT
    uuid,
    name,
    owner_uuid,
    is_private,
    COALESCE(
            ARRAY(
                SELECT uuid
                FROM ways
                WHERE ways.project_uuid = uuid
            ),
            '{}'
    )::VARCHAR[] AS way_uuids,
    COALESCE(
            ARRAY(
                SELECT user_uuid
                FROM users_projects
                WHERE users_projects.project_uuid = uuid
            ),
            '{}'
    )::VARCHAR[] AS user_uuids
FROM projects
WHERE projects.uuid = $1
`

type GetProjectByIDRow struct {
	Uuid      pgtype.UUID `json:"uuid"`
	Name      string      `json:"name"`
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	IsPrivate bool        `json:"is_private"`
	WayUuids  []string    `json:"way_uuids"`
	UserUuids []string    `json:"user_uuids"`
}

func (q *Queries) GetProjectByID(ctx context.Context, projectUuid pgtype.UUID) (GetProjectByIDRow, error) {
	row := q.db.QueryRow(ctx, getProjectByID, projectUuid)
	var i GetProjectByIDRow
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.OwnerUuid,
		&i.IsPrivate,
		&i.WayUuids,
		&i.UserUuids,
	)
	return i, err
}

const updateProject = `-- name: UpdateProject :one
UPDATE projects
SET name = coalesce($1, name),
    is_private = coalesce($2, is_private),
    is_deleted = coalesce($3, is_deleted)
WHERE projects.uuid = $4
RETURNING
    uuid,
    name,
    owner_uuid,
    is_private,
    COALESCE(
            ARRAY(
                SELECT uuid
                FROM ways
                WHERE ways.project_uuid = uuid
            ),
            '{}'
    )::VARCHAR[] AS way_uuids,
    COALESCE(
            ARRAY(
                SELECT user_uuid
                FROM users_projects
                WHERE users_projects.project_uuid = uuid
            ),
            '{}'
    )::VARCHAR[] AS user_uuids
`

type UpdateProjectParams struct {
	Name        pgtype.Text `json:"name"`
	IsPrivate   pgtype.Bool `json:"is_private"`
	IsDeleted   pgtype.Bool `json:"is_deleted"`
	ProjectUuid pgtype.UUID `json:"project_uuid"`
}

type UpdateProjectRow struct {
	Uuid      pgtype.UUID `json:"uuid"`
	Name      string      `json:"name"`
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	IsPrivate bool        `json:"is_private"`
	WayUuids  []string    `json:"way_uuids"`
	UserUuids []string    `json:"user_uuids"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (UpdateProjectRow, error) {
	row := q.db.QueryRow(ctx, updateProject,
		arg.Name,
		arg.IsPrivate,
		arg.IsDeleted,
		arg.ProjectUuid,
	)
	var i UpdateProjectRow
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.OwnerUuid,
		&i.IsPrivate,
		&i.WayUuids,
		&i.UserUuids,
	)
	return i, err
}