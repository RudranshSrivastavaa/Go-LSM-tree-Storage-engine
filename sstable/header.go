package sstable

type Header struct {
	Version     uint32
	RecordCount uint64
	IndexOffset  int64
	BloomOffset int64
}

type IndexEntry struct {
    Key    string
    Offset int64
}

const IndexInterval = 4