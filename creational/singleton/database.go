package main

import "sync"

var instance *DatabaseConnection
var once sync.Once

func GetDatabaseConnection() *DatabaseConnection {
	once.Do(func() {
		instance = &DatabaseConnection{
			connectionString: "user:password@/dbname",
		}
	})
	return instance
}

type DatabaseConnection struct {
	connectionString string
}
