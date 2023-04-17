package main

import (
	"Event-Planning-App/config"
	"Event-Planning-App/config/database"
	"log"
)

func main() {
	// Database connection
	cfg, err := config.InitConfiguration()
	db, err := database.GetConnection(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// Check database connection
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("cannot get sql.DB instance: %v", err)
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("cannot ping database: %v", err)
	}
	log.Println("Connected with database!")
}
