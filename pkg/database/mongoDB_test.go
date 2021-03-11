package database

import "testing"

// Should be an integration test
func TestMongoDBConnection(t *testing.T) {
	mp := NewMongoProducts()
	if mp == nil {
		t.Fail()
	}
}

// Temporary test
func TestMongoDBConnectionAndShutdown(t *testing.T) {
	mp := NewMongoProducts()
	mp.CloseDB()
}
