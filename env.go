package main

import (
	"fmt"
	"os"
)

func getEnv(name string, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return defaultValue
}

func getEnvPanic(name string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	panic(fmt.Sprintf("%s not provided", name))
}
