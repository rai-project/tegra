package tegra

import (
	"runtime"

	"github.com/pkg/errors"

	"github.com/Unknwon/com"
)

var (
	IsSupported       bool
	ErrorNotSupported = errors.New("tegra is not supported on current system")
)

func init() {
	IsSupported = runtime.GOARCH == "arm64" && runtime.GOOS == "linux" && com.IsFile("/etc/nv_tegra_release")
}
