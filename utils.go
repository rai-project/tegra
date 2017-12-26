package tegra

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
)

func runShellScript(sh string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "tegra-")
	if err != nil {
		return "", errors.Wrap(err, "cannot create temporary directory")
	}
	defer os.RemoveAll(tmpDir)
	tmpFile := filepath.Join(tmpDir, "script.sh")
	if err := com.WriteFile(tmpFile, []byte(sh)); err != nil {
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
