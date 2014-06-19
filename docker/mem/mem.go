package mem

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	LXC_MEM_DIR        = "/sys/fs/cgroup/memory/docker"
	LXC_MEM_USAGE_FILE = "memory.usage_in_bytes"
)

func GetUsage(id string) (int64, error) {
	path := fmt.Sprintf("%s/%s/%s", LXC_MEM_DIR, id, LXC_MEM_USAGE_FILE)
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error while opening : ", err)
		return 0, err
	}
	defer f.Close()

	buffer := make([]byte, 16)
	n, err := f.Read(buffer)
	if err != nil {
		log.Println("Error while reading ", path, " : ", err)
		return 0, err
	}

	buffer = buffer[:n-1]
	val, err := strconv.ParseInt(string(buffer), 10, strconv.IntSize)
	if err != nil {
		log.Println("Error while parsing ", string(buffer), " : ", err)
		return 0, err
	}

	return val, nil
}
