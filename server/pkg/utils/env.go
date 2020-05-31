package utils

import (
	stringUtil "github.com/fedoratipper/bitkiosk/server/pkg/utils/string"
	"log"
	"os"
	"strconv"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}

	v = *stringUtil.RemoveEmptyQuote(&v)

	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicln("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}

func MustGetInt(k string) int {
	v := os.Getenv(k)
	if v == "" {
		log.Panicln("ENV missing, key: " + k)
	}

	i, err := strconv.Atoi(v)

	if err != nil {
		log.Panicln("ENV err: [" + k + "]\n" + err.Error())
	}

	return i
}