package meminfo

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Given one line of the `/proc/meminfo` fole,
// we return the name and the value of the information given
func nameAndValue(line string) (string, uint64, error) {
	// Remove white space in the line
	line = strings.ReplaceAll(line, " ", "")
	// Split the name and the value
	lineInfo := strings.Split(line, ":")
	name, value := lineInfo[0], lineInfo[1]

	// Remove the unit at the end of the value if present
	nb, err := removeUnit(value)
	if err != nil {
		return "", 0, err
	}

	return name, nb, nil
}

// Given waht is after the `:` without any space of a line in the
// `/proc/meminfo` we remove the unit if present
func removeUnit(value string) (uint64, error) {
	// If the number finish by `kB` remove it
	value = strings.TrimSuffix(value, "kB")
	nb, err := strconv.ParseUint(value, 10, 64)

	return nb, err
}

// Given the `/proc/meminfo` file, we convert it as map with key is the name
// and value is his value
func namesAndValues(file string) (map[string]uint64, error) {
	// Split the file in lines
	lines := strings.Split(file, "\n")
	infos := map[string]uint64{} // Infos of `/proc/meminfo` as a map

	for _, line := range lines {
		if line != "" {
			name, value, err := nameAndValue(line)
			if err != nil {
				return nil, err
			}
			infos[name] = value
		}
	}

	return infos, nil
}

// Given the map with names and values of `/proc/meminfo` return the Meminfo
// struct newly created
func convertNamesAndValuesInMeminfo(infos map[string]uint64) (Meminfo, error) {
	jsn, err := json.Marshal(infos)
	if err != nil {
		return Meminfo{}, err
	}

	var meinfo Meminfo
	err = json.Unmarshal([]byte(jsn), &meinfo)
	if err != nil {
		return Meminfo{}, err
	}

	return meinfo, nil
}
