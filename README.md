# Meminfo

Simple go package that return the out put of the `/proc/meminfo` file as a go
struct.

## How to use ?

There is one function `NewMeminfo() (Meminfo, error)` that return you a new
Meminfo struct and possible errors.

The Meminfo struct represent what the `/proc/meminfo` file give you.

## How it work ?

The program will read the file `/proc/meminfo`.

Exemple of what contain the `/proc/meminfo` file:

```console
MemTotal:        7845016 kB
MemFree:         3366128 kB
MemAvailable:    5233460 kB
Buffers:          229556 kB
Cached:          2079316 kB
SwapCached:            0 kB
Active:          2671580 kB
Inactive:        1299716 kB
Active(anon):    1801276 kB
Inactive(anon):   149252 kB
Active(file):     870304 kB
Inactive(file):  1150464 kB
Unevictable:      120184 kB
Mlocked:              16 kB
SwapTotal:       2097148 kB
SwapFree:        2097148 kB
Dirty:               420 kB
Writeback:             0 kB
AnonPages:       1782732 kB
Mapped:           649312 kB
Shmem:            290492 kB
KReclaimable:     131724 kB
Slab:             258564 kB
SReclaimable:     131724 kB
SUnreclaim:       126840 kB
KernelStack:       13120 kB
PageTables:        38840 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     6019656 kB
Committed_AS:    9425388 kB
VmallocTotal:   34359738367 kB
VmallocUsed:       31160 kB
VmallocChunk:          0 kB
Percpu:             7008 kB
HardwareCorrupted:     0 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
CmaTotal:              0 kB
CmaFree:               0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:      307392 kB
DirectMap2M:     7770112 kB
```

Then the program convert it as a struct like this (this not result of the
previous file):

```go
Meminfo{
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
```

Then you can use this struct as you want.
