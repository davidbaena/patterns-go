package main

import (
	"sync"
	"testing"
)

func TestGetDatabaseConnectionReturnsSameInstance(t *testing.T) {
	db1 := GetDatabaseConnection()
	db2 := GetDatabaseConnection()

	if db1 != db2 {
		t.Errorf("Expected same instance, got different instances")
	}
}

func TestGetDatabaseConnectionIsThreadSafe(t *testing.T) {
	var wg sync.WaitGroup
	instances := make([]*DatabaseConnection, 100)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = GetDatabaseConnection()
		}(i)
	}

	wg.Wait()

	for i := 1; i < 100; i++ {
		if instances[i] != instances[0] {
			t.Errorf("Expected same instance, got different instances at index %d", i)
		}
	}
}
