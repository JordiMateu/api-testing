# API WITH POSTGRESQL
## _Firts attempts with Golang_

![N|Solid](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/320px-Go_Logo_Blue.svg.png)

A small test to learn about go, go modules and the standard go project layout

## Pourpose

The idea was to do some code able to receive a GET and a POST saving some data into a Database without geeting depper into it and skipping testing.
Also no validations are done in the Db Repository like checks for duplicated IDS and so on...

### Second iteration
- Implement more or less a data 'repository'
- Working with injection
- Endpoint testing using echo and gomock
- Repo testing with gomock and pgxmock
- Avoid data mutability and parameter validation (Not go compilant)


### Third iteration
- Service Layer

### Fourth iteration
- Using context to define DB query timeout value
- Middleware implemented to keep track stats

## Installation and SetUp

Download and install [Golang](//https://golang.org/).
Donwload and install [PostgreSQL](https://www.postgresql.org/download/).
On PostgreSQL up and running use this script to create and add a new table into a database
```sql
CREATE TABLE IF NOT EXISTS public.movies
(
    id integer NOT NULL DEFAULT nextval('movies_id_seq'::regclass),
    movie_name text COLLATE pg_catalog."default",
    genre text COLLATE pg_catalog."default",
    duration text COLLATE pg_catalog."default",
    CONSTRAINT movies_pkey PRIMARY KEY (id)
)
```

Replace in repository.go ln 12 the <values> with your PostgreSQL data

```go
const (
	DbConn = "user=<your_user> password=<your_password> dbname=<your_database> sslmode=disable"
)
```
## URLS and Inspiration
[https://github.com/golang-standards/project-layout]
	
[https://blog.golang.org/using-go-modules]
	
[https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/]

[https://duythhuynh.medium.com/gomock-unit-testing-made-easy-b59a0e947ba7]

[https://echo.labstack.com/guide/testing/]

[https://blog.golang.org/context]

## Modules Used

| Module | Url |
| ------ | ------ |
| Echo | [https://github.com/labstack/echo] |
| PGX | [https://github.com/jackc/pgx] |
| GoMock | [https://github.com/golang/mock] |
| PGXMock | [https://github.com/pashagolub/pgxmock] |
| Testify   |[https://github.com/stretchr/testify] |
	
## License

MIT License

Copyright (c) [2021] [Jordi Mateu]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
