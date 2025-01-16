package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/quantinium03/conduit/internal/types"
)

var Settings = types.Settings{
	OutputPath:   EnvLookup("OUTPUT_DIR", "/out"),
	MetadataPath: EnvLookup("METADATA_DIR", "/metadata"),
	RoutePrefix:  EnvLookup("ROUTE_PREFIX", "/encoder"),
	SafePath:     EnvLookup("SAFE_PATH", "/video"),
    // Add hardware acc
}

func EnvLookup(env string, def string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env file")
		return def
	}

	e := os.Getenv(env)
	if e == "" {
		return def
	}
	return e
}
