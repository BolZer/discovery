package discovery

import (
	"strconv"
	"strings"
)

func FormatString(value string) string {
	return strings.TrimSpace(value)
}

func FormatStringToInt(value string) int {
	castedValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return castedValue
}
