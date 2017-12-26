package tegra

import "github.com/pkg/errors"

type KindType string

const (
	TX1Kind     KindType = "tx1"
	TX2Kind     KindType = "tx2"
	UnknownKind KindType = "unknown"
)

func Kind() (KindType, error) {
	soc, err := SOC()
	if err != nil {
		return UnknownKind, errors.Wrap(err, "unable to determine kind because of error in soc determination")
	}
	switch soc {
	case Tegra210SOC:
		return TX1Kind, nil
	case Tegra186SOC:
		return TX2Kind, nil
	default:
		return UnknownKind, nil
	}
}
