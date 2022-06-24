package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PSQLRepository struct {
	db *gorm.DB
}

func InitPostgresConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=gorm password=gorm sslmode=disable dbname=ddgf_dev"),
		&gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to db: %s.\n", err)
	} else {
		log.Println("database connection successful.")
	}

	db.AutoMigrate(&Page{}, &Advice{}, &Tag{})

	pgDB, _ := db.DB()
	if err = pgDB.Ping(); err != nil {
		log.Fatalf("unable to ping database: %s.\n", err)
	} else {
		log.Println("database pinged successfully.")
	}

	return db
}

func NewPSQLRepository(db *gorm.DB) PSQLRepository {
	return PSQLRepository{db: db}
}
