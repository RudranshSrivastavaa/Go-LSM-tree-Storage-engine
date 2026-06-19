package wal

import (
	"bufio"
	"os"
	"strings"

	"storage-engine/skiplist"
)

func Recover(path string) (*skiplist.SkipList, error) {

	sl := skiplist.New()

	file, err := os.Open(path)

	if err != nil {

		if os.IsNotExist(err) {
			return sl, nil
		}

		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, "|")

		if len(parts) < 2 {
			continue
		}

		switch parts[0] {

		case "PUT":

			if len(parts) != 3 {
				continue
			}

			sl.Insert(
				parts[1],
				skiplist.Entry{
					Value: []byte(parts[2]),
				},
			)

		case "DEL":

			sl.Delete(parts[1])
		}
	}

	return sl, scanner.Err()
}