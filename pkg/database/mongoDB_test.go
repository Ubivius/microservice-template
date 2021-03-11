package database

import "testing"

func TestMongoDBConnection(t *testing.T) {
	mp := NewMongoProducts()
	if mp == nil {
		t.Fail()
	}
}
