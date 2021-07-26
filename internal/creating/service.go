package creating

import (
	mmdb "api-testing/internal"
	"context"
)

type MovieService struct {
	movieRepository mmdb.MovieRepository
}

func NewMovieService(movieRepository mmdb.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

func (s MovieService) CreateMovie(ctx context.Context, id, name, genre, duration string) error {
	movie, err := mmdb.NewMovie(id, name, genre, duration)
	if err != nil {
		return err
	}
	return s.movieRepository.Save(ctx, movie)
}
