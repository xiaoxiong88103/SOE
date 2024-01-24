package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/disk"
)

func main() {
	// 获取所有磁盘分区
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatalf("Error retrieving partitions: %v", err)
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			log.Printf("Error getting usage for partition %s: %v", p.Mountpoint, err)
			continue
		}

		remainingGB := float64(usage.Free) / 1024 / 1024 / 1024
		fmt.Printf("Partition: %s, Remaining space: %.2f GB\n", p.Mountpoint, remainingGB)
	}
}
