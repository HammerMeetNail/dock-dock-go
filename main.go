package main

import (
	"fmt"

	"github.com/HammerMeetNail/dock-dock-go/pkg/stats"
)

func main() {
	cli := stats.Cli("1.39")
	volumes := stats.GetVolumeSize(cli)

	for name, size := range volumes {
		fmt.Println(name, size)
	}
}
