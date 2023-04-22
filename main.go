package main

import (
	"Event-Planning-App/config"
	"log"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg)
	config.Migrate(db)

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
