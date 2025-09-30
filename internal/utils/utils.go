package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Encode[T any](w http.ResponseWriter, v *T, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(*v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil

}

func Decode[T any](r *http.Request) (T, error) {

	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil

}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
