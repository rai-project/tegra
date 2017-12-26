package tegra

func Clocks() (string, error) {
	if !IsSupported {
		return nil, ErrorNotSupported
	}

	sh := MustAsset("_fixtures/jetson_clocks.sh")
	return runShellScript(sh)
}
