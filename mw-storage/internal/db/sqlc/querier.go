// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateFile(ctx context.Context, arg CreateFileParams) (CreateFileRow, error)
	RegenerateDbData(ctx context.Context) error
	RemoveEverything(ctx context.Context) error
}

var _ Querier = (*Queries)(nil)