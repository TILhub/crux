// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: ruleset.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const getApp = `-- name: GetApp :one
SELECT app
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
`

type GetAppParams struct {
	Slice int32  `json:"slice"`
	App   string `json:"app"`
	Class string `json:"class"`
	Realm int32  `json:"realm"`
}

func (q *Queries) GetApp(ctx context.Context, arg GetAppParams) (string, error) {
	row := q.db.QueryRow(ctx, getApp,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
	)
	var app string
	err := row.Scan(&app)
	return app, err
}

const getClass = `-- name: GetClass :one
SELECT class
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
`

type GetClassParams struct {
	Slice int32  `json:"slice"`
	App   string `json:"app"`
	Class string `json:"class"`
	Realm int32  `json:"realm"`
}

func (q *Queries) GetClass(ctx context.Context, arg GetClassParams) (string, error) {
	row := q.db.QueryRow(ctx, getClass,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
	)
	var class string
	err := row.Scan(&class)
	return class, err
}

const getWFActiveStatus = `-- name: GetWFActiveStatus :one
SELECT is_active
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
    AND setname = $5
`

type GetWFActiveStatusParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Realm   int32  `json:"realm"`
	Setname string `json:"setname"`
}

func (q *Queries) GetWFActiveStatus(ctx context.Context, arg GetWFActiveStatusParams) (pgtype.Bool, error) {
	row := q.db.QueryRow(ctx, getWFActiveStatus,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
		arg.Setname,
	)
	var is_active pgtype.Bool
	err := row.Scan(&is_active)
	return is_active, err
}

const getWFInternalStatus = `-- name: GetWFInternalStatus :one
SELECT is_internal
FROM ruleset
WHERE
    slice = $1
    AND app = $2
    AND class = $3
    AND realm = $4
    AND brwf = 'W'
    AND setname = $5
`

type GetWFInternalStatusParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Realm   int32  `json:"realm"`
	Setname string `json:"setname"`
}

func (q *Queries) GetWFInternalStatus(ctx context.Context, arg GetWFInternalStatusParams) (bool, error) {
	row := q.db.QueryRow(ctx, getWFInternalStatus,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Realm,
		arg.Setname,
	)
	var is_internal bool
	err := row.Scan(&is_internal)
	return is_internal, err
}

const workFlowNew = `-- name: WorkFlowNew :exec
INSERT INTO
    ruleset (
        realm, slice, app, brwf, class, setname, schemaid, is_active, is_internal, ruleset, createdat, createdby
    )
VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, $11
    )
`

type WorkFlowNewParams struct {
	Realm      int32       `json:"realm"`
	Slice      int32       `json:"slice"`
	App        string      `json:"app"`
	Brwf       BrwfEnum    `json:"brwf"`
	Class      string      `json:"class"`
	Setname    string      `json:"setname"`
	Schemaid   int32       `json:"schemaid"`
	IsActive   pgtype.Bool `json:"is_active"`
	IsInternal bool        `json:"is_internal"`
	Ruleset    []byte      `json:"ruleset"`
	Createdby  string      `json:"createdby"`
}

func (q *Queries) WorkFlowNew(ctx context.Context, arg WorkFlowNewParams) error {
	_, err := q.db.Exec(ctx, workFlowNew,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Brwf,
		arg.Class,
		arg.Setname,
		arg.Schemaid,
		arg.IsActive,
		arg.IsInternal,
		arg.Ruleset,
		arg.Createdby,
	)
	return err
}

const workFlowUpdate = `-- name: WorkFlowUpdate :exec
UPDATE ruleset
SET
    brwf = $4,
    setname = $5,
    ruleset = $6,
    editedat = CURRENT_TIMESTAMP,
    editedby = $7
WHERE
    slice = $1
    AND class = $2
    AND app = $3
`

type WorkFlowUpdateParams struct {
	Slice    int32       `json:"slice"`
	Class    string      `json:"class"`
	App      string      `json:"app"`
	Brwf     BrwfEnum    `json:"brwf"`
	Setname  string      `json:"setname"`
	Ruleset  []byte      `json:"ruleset"`
	Editedby pgtype.Text `json:"editedby"`
}

func (q *Queries) WorkFlowUpdate(ctx context.Context, arg WorkFlowUpdateParams) error {
	_, err := q.db.Exec(ctx, workFlowUpdate,
		arg.Slice,
		arg.Class,
		arg.App,
		arg.Brwf,
		arg.Setname,
		arg.Ruleset,
		arg.Editedby,
	)
	return err
}

const workflowDelete = `-- name: WorkflowDelete :execresult
DELETE from ruleset
where
    brwf = 'W'
    AND is_active = false
    and slice = $1
    and app = $2
    and class = $3
    and setname = $4
    AND realm = $5
`

type WorkflowDeleteParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Setname string `json:"setname"`
	Realm   int32  `json:"realm"`
}

func (q *Queries) WorkflowDelete(ctx context.Context, arg WorkflowDeleteParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, workflowDelete,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Realm,
	)
}

const workflowList = `-- name: WorkflowList :many
select
    id,
    slice,
    app,
    class,
    setname as name,
    is_active,
    is_internal,
    createdat,
    createdby,
    editedat,
    editedby
from ruleset
where
    brwf = 'W'
    AND realm = $1
    AND ($2::INTEGER is null OR slice = $2::INTEGER)
    AND ( $3::text[] is null OR app = any( $3::text[]))
    AND ($4::text is null OR class = $4::text)
    AND ($5::text is null OR setname = $5::text)
    AND ($6::BOOLEAN is null OR is_active = $6::BOOLEAN)
    AND ($7::BOOLEAN is null OR is_internal = $7::BOOLEAN)
`

type WorkflowListParams struct {
	Realm      int32       `json:"realm"`
	Slice      pgtype.Int4 `json:"slice"`
	App        []string    `json:"app"`
	Class      pgtype.Text `json:"class"`
	Setname    pgtype.Text `json:"setname"`
	IsActive   pgtype.Bool `json:"is_active"`
	IsInternal pgtype.Bool `json:"is_internal"`
}

type WorkflowListRow struct {
	ID         int32            `json:"id"`
	Slice      int32            `json:"slice"`
	App        string           `json:"app"`
	Class      string           `json:"class"`
	Name       string           `json:"name"`
	IsActive   pgtype.Bool      `json:"is_active"`
	IsInternal bool             `json:"is_internal"`
	Createdat  pgtype.Timestamp `json:"createdat"`
	Createdby  string           `json:"createdby"`
	Editedat   pgtype.Timestamp `json:"editedat"`
	Editedby   pgtype.Text      `json:"editedby"`
}

func (q *Queries) WorkflowList(ctx context.Context, arg WorkflowListParams) ([]WorkflowListRow, error) {
	rows, err := q.db.Query(ctx, workflowList,
		arg.Realm,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.IsActive,
		arg.IsInternal,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WorkflowListRow
	for rows.Next() {
		var i WorkflowListRow
		if err := rows.Scan(
			&i.ID,
			&i.Slice,
			&i.App,
			&i.Class,
			&i.Name,
			&i.IsActive,
			&i.IsInternal,
			&i.Createdat,
			&i.Createdby,
			&i.Editedat,
			&i.Editedby,
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
    and realm = $5
    AND brwf = 'W'
`

type WorkflowgetParams struct {
	Slice   int32  `json:"slice"`
	App     string `json:"app"`
	Class   string `json:"class"`
	Setname string `json:"setname"`
	Realm   int32  `json:"realm"`
}

type WorkflowgetRow struct {
	ID         int32            `json:"id"`
	Slice      int32            `json:"slice"`
	App        string           `json:"app"`
	Class      string           `json:"class"`
	Name       string           `json:"name"`
	IsActive   pgtype.Bool      `json:"is_active"`
	IsInternal bool             `json:"is_internal"`
	Flowrules  []byte           `json:"flowrules"`
	Createdat  pgtype.Timestamp `json:"createdat"`
	Createdby  string           `json:"createdby"`
	Editedat   pgtype.Timestamp `json:"editedat"`
	Editedby   pgtype.Text      `json:"editedby"`
}

func (q *Queries) Workflowget(ctx context.Context, arg WorkflowgetParams) (WorkflowgetRow, error) {
	row := q.db.QueryRow(ctx, workflowget,
		arg.Slice,
		arg.App,
		arg.Class,
		arg.Setname,
		arg.Realm,
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
