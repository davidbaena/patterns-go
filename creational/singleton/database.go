package main

import "sync"

var instance *DatabaseConnection
var once sync.Once

// GetDatabaseConnection returns the singleton instance of DatabaseConnection.
// It ensures that only one instance of DatabaseConnection is created and provides
// a global point of access to it. The instance is initialized only once using sync.Once.
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
