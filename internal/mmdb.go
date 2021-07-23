package mmdb

import (
	"context"
	"errors"
)

// Movie is the data structure that represents a movie.
type Movie struct {
	id       MovieID
	name     MovieName
	genre    MovieGenre
	duration MovieDuration
}

var ErrEmptyID = errors.New("the field Id can not be empty")
var ErrEmptyName = errors.New("the field Name can not be empty")
var ErrEmptyGenre = errors.New("the field Genre can not be empty")
var ErrEmptyDuration = errors.New("the field Duration can not be empty")

func NewMovieDuration(value string) (MovieDuration, error) {
	if value == "" {
		return MovieDuration{}, ErrEmptyDuration
	}

	return MovieDuration{
		value: value,
	}, nil
}

func NewMovieName(value string) (MovieName, error) {
	if value == "" {
		return MovieName{}, ErrEmptyName
	}

	return MovieName{
		value: value,
	}, nil
}

func NewMovieGenre(value string) (MovieGenre, error) {
	if value == "" {
		return MovieGenre{}, ErrEmptyGenre
	}

	return MovieGenre{
		value: value,
	}, nil
}

func NewMovieID(value string) (MovieID, error) {
	if value == "" {
		return MovieID{}, ErrEmptyID
	}

	return MovieID{
		value: value,
	}, nil
}

// String type converts the MoveID into string.
func (id MovieID) String() string {
	return id.value
}

func (name MovieName) String() string {
	return name.value
}

func (genre MovieGenre) String() string {
	return genre.value
}

func (duration MovieDuration) String() string {
	return duration.value
}

// NewMovie creates a new movie.
func NewMovie(id, name, genre, duration string) (Movie, error) {

	idV, err := NewMovieID(id)
	if err != nil {
		return Movie{}, err
	}

	nameV, err := NewMovieName(name)
	if err != nil {
		return Movie{}, err
	}

	genreV, err := NewMovieGenre(genre)
	if err != nil {
		return Movie{}, err
	}

	durationV, err := NewMovieDuration(duration)
	if err != nil {
		return Movie{}, err
	}

	return Movie{
		id:       idV,
		name:     nameV,
		genre:    genreV,
		duration: durationV,
	}, nil

}

type MovieID struct {
	value string
}

type MovieName struct {
	value string
}

type MovieGenre struct {
	value string
}

type MovieDuration struct {
	value string
}

// MovieRepository defines  a movie storage.
type MovieRepository interface {
	Save(ctx context.Context, movie Movie) error
}

//Getters... I know ...
func (m Movie) ID() MovieID {
	return m.id
}

func (m Movie) Name() MovieName {
	return m.name
}

func (m Movie) Genre() MovieGenre {
	return m.genre
}

func (m Movie) Duration() MovieDuration {
	return m.duration
}
