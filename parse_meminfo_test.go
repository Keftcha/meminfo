package meminfo

import (
	"encoding/json"
	"errors"
	"log"
	"testing"
)

func TestNameAndValueWithNormalLine(t *testing.T) {
	gotName, gotNb, gotErr := nameAndValue("MemTotal:        7845016 kB")
	expectedName, expectedNb := "MemTotal", uint64(7845016)
	var expectedErr error = nil

	if gotName != expectedName {
		t.Errorf("Expected: %s, but got: %s", expectedName, gotName)
	}
	if gotNb != expectedNb {
		t.Errorf("Expected: %d, but got: %d", expectedNb, gotNb)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestNameAndValueWithEmptyString(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The function did not panic")
		}
	}()

	nameAndValue("")
}

func TestNameAndValueWithNormalLineWithoutTheUnit(t *testing.T) {
	gotName, gotNb, gotErr := nameAndValue("MemTotal:        7845016")
	expectedName, expectedNb := "MemTotal", uint64(7845016)
	var expectedErr error = nil

	if gotName != expectedName {
		t.Errorf("Expected: %s, but got: %s", expectedName, gotName)
	}
	if gotNb != expectedNb {
		t.Errorf("Expected: %d, but got: %d", expectedNb, gotNb)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestNameAndValueWithOnlyColon(t *testing.T) {
	gotName, gotNb, gotErr := nameAndValue(":")
	expectedName, expectedNb := "", uint64(0)
	expectedErr := errors.New("strconv.ParseUint: parsing \"\": invalid syntax")

	if gotName != expectedName {
		t.Errorf("Expected: %s, but got: %s", expectedName, gotName)
	}
	if gotNb != expectedNb {
		t.Errorf("Expected: %d, but got: %d", expectedNb, gotNb)
	}
	if gotErr.Error() != expectedErr.Error() {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestRemoveUnitWithNormalValue(t *testing.T) {
	gotValue, gotErr := removeUnit("1932056kB")
	expectedValue := uint64(1932056)
	var expectedErr error = nil

	if gotValue != expectedValue {
		t.Errorf("Expected: %d, but got: %d", expectedValue, gotValue)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestRemoveUnitWithNoUnit(t *testing.T) {
	gotValue, gotErr := removeUnit("3169696")
	expectedValue := uint64(3169696)
	var expectedErr error = nil

	if gotValue != expectedValue {
		t.Errorf("Expected: %d, but got: %d", expectedValue, gotValue)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestRemoveUnitWithEmptyString(t *testing.T) {
	gotValue, gotErr := removeUnit("")
	expectedValue := uint64(0)
	var expectedErr error = errors.New(`strconv.ParseUint: parsing "": invalid syntax`)

	if gotValue != expectedValue {
		t.Errorf("Expected: %d, but got: %d", expectedValue, gotValue)
	}
	if gotErr.Error() != expectedErr.Error() {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestNamesAndValuesWithNormalFile(t *testing.T) {
	file, err := readFile("./meminfo_files/meminfo_0.txt")
	if err != nil {
		log.Fatal(err)
	}

	got, gotErr := namesAndValues(file)
	expected := map[string]uint64{
		"MemTotal":          7845016,
		"MemFree":           526100,
		"MemAvailable":      3570744,
		"Buffers":           492632,
		"Cached":            3572852,
		"SwapCached":        30800,
		"Active":            4476136,
		"Inactive":          2243540,
		"Active(anon)":      2671788,
		"Inactive(anon)":    879560,
		"Active(file)":      1804348,
		"Inactive(file)":    1363980,
		"Unevictable":       121304,
		"Mlocked":           32,
		"SwapTotal":         2097148,
		"SwapFree":          2055960,
		"Dirty":             3724,
		"Writeback":         0,
		"AnonPages":         2742800,
		"Mapped":            1563820,
		"Shmem":             901728,
		"KReclaimable":      179940,
		"Slab":              318660,
		"SReclaimable":      179940,
		"SUnreclaim":        138720,
		"KernelStack":       15232,
		"PageTables":        58192,
		"NFS_Unstable":      0,
		"Bounce":            0,
		"WritebackTmp":      0,
		"CommitLimit":       6019656,
		"Committed_AS":      12811736,
		"VmallocTotal":      34359738367,
		"VmallocUsed":       33516,
		"VmallocChunk":      0,
		"Percpu":            7328,
		"HardwareCorrupted": 0,
		"AnonHugePages":     0,
		"ShmemHugePages":    0,
		"ShmemPmdMapped":    0,
		"CmaTotal":          0,
		"CmaFree":           0,
		"HugePages_Total":   0,
		"HugePages_Free":    0,
		"HugePages_Rsvd":    0,
		"HugePages_Surp":    0,
		"Hugepagesize":      2048,
		"Hugetlb":           0,
		"DirectMap4k":       409792,
		"DirectMap2M":       7667712,
	}
	var expectedErr error = nil

	gotStr, _ := json.Marshal(got)
	expectedStr, _ := json.Marshal(expected)
	if string(gotStr) != string(expectedStr) {
		t.Errorf("Expected: %s, but got: %s", string(expectedStr), string(gotStr))
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestNamesAndValuesWithEmptyString(t *testing.T) {
	got, gotErr := namesAndValues("")
	expected := map[string]int{}
	var expectedErr error = nil

	gotStr, _ := json.Marshal(got)
	expectedStr, _ := json.Marshal(expected)
	if string(gotStr) != string(expectedStr) {
		t.Errorf("Expected: %s, but got: %s", string(expectedStr), string(gotStr))
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestConvertNamesAndValuesInMeminfo(t *testing.T) {
	meminfoFile, err := readFile("./meminfo_files/meminfo_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	meminfoMap, err := namesAndValues(meminfoFile)
	if err != nil {
		log.Fatal(err)
	}

	gotMeminfo, gotErr := convertNamesAndValuesInMeminfo(meminfoMap)
	expectedMeminfo := Meminfo{
		MemTotal:          16585088,
		MemFree:           10162564,
		Buffers:           34032,
		Cached:            188576,
		SwapCached:        0,
		Active:            167556,
		Inactive:          157876,
		ActiveAnon:        103104,
		InactiveAnon:      17440,
		ActiveFile:        64452,
		InactiveFile:      140436,
		Unevictable:       0,
		Mlocked:           0,
		SwapTotal:         14569852,
		SwapFree:          14340140,
		Dirty:             0,
		Writeback:         0,
		AnonPages:         102824,
		Mapped:            71404,
		Shmem:             17720,
		Slab:              13868,
		SReclaimable:      6744,
		SUnreclaim:        7124,
		KernelStack:       2848,
		PageTables:        2524,
		NFS_Unstable:      0,
		Bounce:            0,
		WritebackTmp:      0,
		CommitLimit:       515524,
		Committed_AS:      3450064,
		VmallocTotal:      122880,
		VmallocUsed:       21296,
		VmallocChunk:      66044,
		HardwareCorrupted: 0,
		AnonHugePages:     2048,
		HugePages_Total:   0,
		HugePages_Free:    0,
		HugePages_Rsvd:    0,
		HugePages_Surp:    0,
		Hugepagesize:      2048,
		DirectMap4k:       12280,
		DirectMap2M:       897024,
	}
	var expectedErr error = nil

	if gotMeminfo != expectedMeminfo {
		t.Errorf("Expected: %+v, but got: %+v", expectedMeminfo, gotMeminfo)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}

func TestConvertNamesAndValuesInMeminfoWithEmptyMap(t *testing.T) {
	meminfoMap := map[string]uint64{}

	gotMeminfo, gotErr := convertNamesAndValuesInMeminfo(meminfoMap)
	expectedMeminfo := Meminfo{}
	var expectedErr error = nil

	if gotMeminfo != expectedMeminfo {
		t.Errorf("Expected: %+v, but got: %+v", expectedMeminfo, gotMeminfo)
	}
	if gotErr != expectedErr {
		t.Errorf("Expected: %s, but got: %s", expectedErr, gotErr)
	}
}
