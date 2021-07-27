package boot

import (
	"api-testing/internal/creating"
	"api-testing/internal/platform/server"
	"api-testing/internal/platform/server/storage/psql"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

const (
	DbConn = "user=<your_user> password=<your_password> dbname=<your_database> sslmode=disable"
)

func Run() error {
	//Set DB connection
	conn, err := pgx.Connect(context.Background(), DbConn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	m := psql.NewMovieRepository(conn, 10)
	s := creating.NewMovieService(m)

	srv := server.New(s)
	return srv.Start(":8080")
}
