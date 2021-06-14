package impl_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "13.3", []string{"POSTGRES_PASSWORD=postgres"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		dns := fmt.Sprintf(
			"postgres://postgres:postgres@localhost:%s/postgres",
			resource.GetPort("5432/tcp"))
		db, err = sqlx.Open("pgx", dns)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	schema := `
	CREATE TABLE courses (
		id bigint PRIMARY KEY,
		classroom_id bigint NOT NULL,
		name text NOT NULL,
		stream text NOT NULL
	);

	CREATE TABLE lessons (
		id bigint PRIMARY KEY,
		course_id bigint NOT NULL,
		number integer,
		name text NOT NULL
	);
	`

	_, err = db.Exec(schema)

	if err != nil {
		log.Fatalf("Couldn't create table: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
