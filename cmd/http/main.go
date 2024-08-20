package http

import (
	"pari/database"
	"pari/pkg/inits"
)

func main() {

	cfg := inits.InitializeConfig()

	database.ConnectDB()

	db := inits.InitializeDatabase(cfg.Env.Db, cfg.Env.Tz)

	router := InitializeGin(cfg.Env.Server.Name, cfg.Env.Environment)

	InitializeRepositories(db)
	InitializeUsecases(cfg)
	InitializeControllers(router, cfg)
}
