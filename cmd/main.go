package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/database"
	"github.com/amirmohammadkariimi/interview-task/internal/server"
	"github.com/amirmohammadkariimi/interview-task/pkg/cache"
)

func main() {
	configFile := flag.String("config", "", "Location of config file")
	flag.Parse()
	config, err := readConfig(*configFile)
	if err != nil {
		slog.Error("cannot read config file", "error", err)
		os.Exit(1)
	}
	// create new database
	db, err := database.New(config.Database.Address, config.Database.Name, config.Database.User, config.Database.Pass)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		os.Exit(1)
	}
	database.Migrate(db)
	// create new cache
	c := cache.NewCache()
	// create new server
	s := server.New(config.Port, db, c)
	err = s.Run()
	if err != nil {
		slog.Error("error running server", "error", err)
	}
}
