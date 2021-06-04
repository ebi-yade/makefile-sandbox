package utils

import (
	"fmt"
	"os"
)

// MustNonEmptyEnv calls panic when the environment variable is not defined or an empty string
func MustNonEmptyEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("the environment variable `%s` is not defined", key))
	}

	if len(val) == 0 {
		panic(fmt.Errorf("the environment variable `%s` is defined but an empty string", key))
	}

	return val
}
