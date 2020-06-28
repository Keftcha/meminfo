package meminfo

// Meminfo type represent all data in the /proc/meminfo file
type Meminfo struct {
	MemTotal          uint64
	MemFree           uint64
	MemAvailable      uint64
	Buffers           uint64
	Cached            uint64
	SwapCached        uint64
	Active            uint64
	Inactive          uint64
	ActiveAnon        uint64 `json:"Active(anon)"`
	InactiveAnon      uint64 `json:"Inactive(anon)"`
	ActiveFile        uint64 `json:"Active(file)"`
	InactiveFile      uint64 `json:"Inactive(file)"`
	Unevictable       uint64
	Mlocked           uint64
	SwapTotal         uint64
	SwapFree          uint64
	Dirty             uint64
	Writeback         uint64
	AnonPages         uint64
	Mapped            uint64
	Shmem             uint64
	KReclaimable      uint64
	Slab              uint64
	SReclaimable      uint64
	SUnreclaim        uint64
	KernelStack       uint64
	PageTables        uint64
	NFS_Unstable      uint64
	Bounce            uint64
	WritebackTmp      uint64
	CommitLimit       uint64
	Committed_AS      uint64
	VmallocTotal      uint64
	VmallocUsed       uint64
	VmallocChunk      uint64
	Percpu            uint64
	HardwareCorrupted uint64
	AnonHugePages     uint64
	ShmemHugePages    uint64
	ShmemPmdMapped    uint64
	CmaTotal          uint64
	CmaFree           uint64
	HugePages_Total   uint64
	HugePages_Free    uint64
	HugePages_Rsvd    uint64
	HugePages_Surp    uint64
	Hugepagesize      uint64
	Hugetlb           uint64
	DirectMap4k       uint64
	DirectMap2M       uint64
}

// NewMeminfo return you a new Meminfo
func NewMeminfo() (Meminfo, error) {
	file, err := readMeminfo()
	if err != nil {
		return Meminfo{}, err
	}

	mapInfos, err := namesAndValues(file)
	if err != nil {
		return Meminfo{}, err
	}

	return convertNamesAndValuesInMeminfo(mapInfos)
}
