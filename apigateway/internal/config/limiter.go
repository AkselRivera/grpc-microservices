package config

import (
	"time"

	"github.com/gofiber/storage/memory/v2"
)

type Config struct {
	// Time before deleting expired keys
	//
	// Default is 10 * time.Second
	GCInterva time.Duration
}

func NewStore() *memory.Storage {
	// Initialize custom config
	store := memory.New(memory.Config{
		GCInterval: 10 * time.Second,
	})

	return store
}
