// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: workflow.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const workflowget = `-- name: Workflowget :one
select
    id,
    slice,
    app,
    class,
    setname as name,
    is_active,
    is_internal,
    ruleset as flowrules,
    createdat,
    createdby,
    editedat,
    editedby
from ruleset
where
    slice = $1
    and app = $2
    and class = $3
    and setname = $4
`

type WorkflowgetParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Setname string `json:"setname"`
}

type WorkflowgetRow struct {
	ID         int32           `json:"id"`
	Slice      int32           `json:"slice"`
	App        string          `json:"app"`
	Class      string          `json:"class"`
	Name       string          `json:"name"`
	IsActive   sql.NullBool    `json:"is_active"`
	IsInternal bool            `json:"is_internal"`
	Flowrules  json.RawMessage `json:"flowrules"`
	Createdat  time.Time       `json:"createdat"`
	Createdby  string          `json:"createdby"`
	Editedat   time.Time       `json:"editedat"`
	Editedby   string          `json:"editedby"`
}

func (q *Queries) Workflowget(ctx context.Context, arg WorkflowgetParams) (WorkflowgetRow, error) {
	row := q.db.QueryRowContext(ctx, workflowget,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
	)
	var i WorkflowgetRow
	err := row.Scan(
		&i.ID,
		&i.Slice,
		&i.App,
		&i.Class,
		&i.Name,
		&i.IsActive,
		&i.IsInternal,
		&i.Flowrules,
		&i.Createdat,
		&i.Createdby,
		&i.Editedat,
		&i.Editedby,
	)
	return i, err
}
