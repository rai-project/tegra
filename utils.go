package tegra

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
)

func runShellScript(sh string) (string, error) {
	tmpFile, err := ioutil.TempFile("", "tegra-")
	if err != nil {
		return "", errors.Wrap(err, "cannot create temporary file")
	}
	defer os.Remove(tmpFile)
	if err := com.WriteFile(tmpFile, sh); err != nil {
		return "", errors.Wrapf(err, "cannot write to temporary file %v", tmpFile)
	}

	buf := new(bytes.Buffer)

	cmd := exec.Command("/bin/sh", tmpFile)
	cmd.Stderr = buf
	cmd.Stderr = buf
	err = cmd.Run()

	if err != nil {
		return "", errors.Wrapf(err, "failed while running temporary file %v", tmpFile)
	}

	return buf.String(), nil
}
