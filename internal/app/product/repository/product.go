package repository

import(
    "github.com/jackc/pgx/v5"
)

type ProductPostgreSQLItf interface {}

type ProductPostgreSQL struct {
    db *pgx.Conn
}

func NewProductPostgreSQL(db *pgx.Conn) ProductPostgreSQLItf {
    return &ProductPostgreSQL{db}
}
