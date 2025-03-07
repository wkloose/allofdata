package repository

import(
    "github.com/jackc/pgx/v5"
)

type UserPostgreSQLItf interface {}

type UserPostgreSQL struct {
    db *pgx.Conn
}

func NewUserPostgreSQL(db *pgx.Conn) UserPostgreSQLItf {
    return &UserPostgreSQL{db}
}
