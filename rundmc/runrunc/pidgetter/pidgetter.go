package pidgetter

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"code.cloudfoundry.org/guardian/rundmc"
	"code.cloudfoundry.org/lager"
)

type FilePidGetter struct {
	Depot rundmc.Depot
}

func (f *FilePidGetter) GetPid(logger lager.Logger, containerHandle string) (int, error) {
	bundlePath, err := f.Depot.Lookup(logger, containerHandle)
	if err != nil {
		return 0, err
	}

	pidfileContent, err := ioutil.ReadFile(filepath.Join(bundlePath, "pidfile"))
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(pidfileContent))
	if err != nil {
		return 0, err
	}
	return pid, nil
}
