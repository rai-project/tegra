package tegra

func Clocks() (string, error) {
	if !IsSupported {
		return "", ErrorNotSupported
	}

	sh := MustAsset("_fixtures/jetson_clocks.sh")
	return runShellScript(string(sh))
}
