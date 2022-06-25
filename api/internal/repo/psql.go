package repo

import (
	"ddgf-new/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLRepository struct {
	db *gorm.DB
}

func InitPostgresConn(config *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(MakeDSNString(config)), &gorm.Config{})
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

func MakeDSNString(config *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		fmt.Sprint(config.DbPort),
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
}

func NewPSQLRepository(db *gorm.DB) PSQLRepository {
	return PSQLRepository{db: db}
}
