package config

import (
	"errors"
	"os"
	"strconv"
)

func GetOrDefault(key string, setvalue int) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return strconv.Itoa(setvalue), nil
	}
	return value, nil
}

func Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("key not found in env")
	}
	return value, nil
}

func GetAsInt64(key string) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, errors.New("key not found in env")
	}
	v, _ := strconv.Atoi(value)

	return int64(v), nil
}
