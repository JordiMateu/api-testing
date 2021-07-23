package psql

import (
	mmdb "api-testing/internal"
	"context"
	"github.com/jackc/pgx/v4"
)

type MovieRepository struct {
	db *pgx.Conn
}

func NewMovieRepository(conn *pgx.Conn) *MovieRepository {
	return &MovieRepository{
		db: conn,
	}
}

func (r *MovieRepository) Save(ctx context.Context, m mmdb.Movie) error {
	sqlStatement := "INSERT INTO public.movies (id, movie_name, genre, duration) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(ctx, sqlStatement, m.ID().String(), m.Name().String(), m.Genre().String(), m.Duration().String())

	if err != nil {
		panic(err)
	}

	return nil
}
