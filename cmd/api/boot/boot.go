package boot

import (
	"api-testing/internal/platform/server"
)

func Run() error{
	srv := server.New("8080")
	return srv.Start(":8080")
}