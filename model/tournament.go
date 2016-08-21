package model

type Tournament struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
