package utils

import (
	"os"
	"strings"
)

func GetDefaultPath() (defaultPath string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	defaultPath = strings.Split(dir, "manager/")[0] + "manager/"

	return
}
