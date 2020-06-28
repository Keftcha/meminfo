package meminfo

import (
	"io/ioutil"
	"path/filepath"
)

// Read a file and return his content as a string
func ReadFile(filePath string) (string, error) {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func ReadMeminfo() (string, error) {
	return ReadFile("/proc/meminfo")
}
