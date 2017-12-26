package tegra

import (
	"io/ioutil"
	"strings"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
)

type SOCType string

const (
	Tegra210SOC SOCType = "tegra210"
	Tegra186SOC SOCType = "tegra186"
	UnknownSOC  SOCType = "unknown"
)

func SOC() (SOCType, error) {
	if !IsSupported {
		return UnknownSOC, ErrorNotSupported
	}
	loc := "/proc/device-tree/compatible"
	if !com.IsFile(loc) {
		return UnknownSOC, errors.Errorf("the file %s was not found", loc)
	}
	data, err := ioutil.ReadFile(loc)
	if err != nil {
		return UnknownSOC, errors.Wrapf(err, "unable to read %v", loc)
	}
	str := string(data)
	if strings.Contains(str, "nvidia,tegra210") {
		return Tegra210SOC, nil
	}
	if strings.Contains(str, "nvidia,tegra186") {
		return Tegra186SOC, nil
	}
	return UnknownSOC, errors.Errorf("unable to determine the soc from %v", str)
}
