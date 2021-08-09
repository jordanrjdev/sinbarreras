package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/JGurus/sin-barreras-backend/config"
	"github.com/JGurus/sin-barreras-backend/models"
)

var (
	once sync.Once
	db   *sql.DB
)

func GetConnection() *sql.DB {
	once.Do(connection)
	return db
}

func connection() {
	var err error
	configDB := getConfig()
	uri := fmt.Sprintf("%s:%s@/%s?tls=false&autocommit=true&parseTime=true", configDB.User, configDB.Password, configDB.Database)
	db, err = sql.Open(configDB.Engine, uri)
	if err != nil {
		log.Fatalf("No se pudo conectar a la bd: %v", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping: %v", err.Error())
	}
	fmt.Println("Conectado a la BD")
}

func getConfig() models.ConfigDB {
	var err error
	configDB, err := config.GetConfigDB()
	if err != nil {
		log.Fatal(err.Error())
		return configDB
	}
	return configDB
}
