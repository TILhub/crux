// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: stepworkflow.sql

package sqlc

import (
	"context"
)

const getWorkflow = `-- name: GetWorkflow :one
SELECT workflow FROM stepworkflow
WHERE step =$1
`

func (q *Queries) GetWorkflow(ctx context.Context, step string) (string, error) {
	row := q.db.QueryRow(ctx, getWorkflow, step)
	var workflow string
	err := row.Scan(&workflow)
	return workflow, err
}
