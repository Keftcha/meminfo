package meminfo

import (
	"io/ioutil"
	"path/filepath"
)

// ReadFile read a file and return his content as a string
func readFile(filePath string) (string, error) {
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

// ReadMeminfo read the `/proc/meminfo` file and return it as a string
func readMeminfo() (string, error) {
	return readFile("/proc/meminfo")
}
