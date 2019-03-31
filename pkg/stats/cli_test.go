package stats

import (
	"testing"
)

var versionTests = []struct {
	version string
}{
	{"1.38"},
	{"1.39"},
}

func TestCliVersions(t *testing.T) {
	for _, supported := range versionTests {
		cli := Cli(supported.version)
		if cli.ClientVersion() != supported.version {
			t.Errorf("Cli(%s): expected %s, actual %s", supported.version, supported.version, cli.ClientVersion())
		}
	}
}
