package clients

import (
	"github.com/docker/docker/client"
)

// Cli connects to the local Docker installation and returns a client
func Cli(version string) (cli *client.Client) {
	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	if err != nil {
		panic(err)
	}
	return cli
}
