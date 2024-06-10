package db

import (
	"context"
	"fmt"

	"products-observability/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewPostgresDB(pgDsn PostgresDsn) *sqlx.DB {
	db, err := otelsqlx.Open(
		"postgres", pgDsn.ToString(),
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL))
	if err != nil {
		logger.Fatal(context.Background(), err.Error())
	}

	return db
}

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func (p PostgresDsn) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", p.Host, p.User, p.Password, p.Db, p.Port)
}
