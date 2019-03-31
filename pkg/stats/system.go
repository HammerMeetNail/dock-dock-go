package stats

import (
	"context"

	"github.com/docker/docker/client"
)

// GetVolumeSize lists name and size of all volumes
func GetVolumeSize(cli *client.Client) (volumes map[string]int64) {

	diskUsage, err := cli.DiskUsage(context.Background())
	if err != nil {
		panic(err)
	}

	volumes = make(map[string]int64)

	for _, volume := range diskUsage.Volumes {
		volumes[volume.Name] = volume.UsageData.Size
	}

	return volumes
}
