package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

)

const (
	DbConn = "user=<your_user> password=<your_password> dbname=<your_database> sslmode=disable"
)

type Repository struct {
	db *sql.DB
}

func Start() Repository {
	// Create DB pool
	db,err := sql.Open("postgres", DbConn)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	return Repository{db: db}
}

func (r *Repository) Save(m Movie) error {
	sqlStatement := " INSERT INTO public.movies (id, movie_name, genre, duration) VALUES ($1, $2, $3, $4) RETURNING id"
	id := m.id
	err := r.db.QueryRow(sqlStatement, m.id, m.name, m.genre, m.duration).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New movie added and the ID is:", id)
	return err
}