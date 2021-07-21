package psql

type Movie struct {
	id      string
	name	string
	genre	string
	duration string
}

// Creates a new movie as struct to be used by the repository
func AddMovie(id, name, genre, duration string) Movie{
	m:= Movie{
		id:       id,
		name:     name,
		genre: 	genre,
		duration: duration,
	}

	r:= Start()
	r.Save(m)
	return m
}