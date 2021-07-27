package psql

import (
	mmdb "api-testing/internal"
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type MovieRepository struct {
	db        *pgx.Conn
	dbTimeOut time.Duration
}

func NewMovieRepository(conn *pgx.Conn, dbTimeout time.Duration) *MovieRepository {
	return &MovieRepository{
		db:        conn,
		dbTimeOut: dbTimeout,
	}
}

func (r *MovieRepository) Save(ctx context.Context, m mmdb.Movie) error {
	sqlStatement := "INSERT INTO public.movies (id, movie_name, genre, duration) VALUES ($1, $2, $3, $4)"
	//Will cancel sql query after timout value
	ctxTimeOut, cancel := context.WithTimeout(ctx, r.dbTimeOut)
	defer cancel()
	_, err := r.db.Exec(ctxTimeOut, sqlStatement, m.ID().String(), m.Name().String(), m.Genre().String(), m.Duration().String())

	if err != nil {
		panic(err)
	}

	return nil
}
