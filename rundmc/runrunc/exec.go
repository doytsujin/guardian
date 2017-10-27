package runrunc

import (
	"os"
	"os/exec"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/lager"
	"github.com/opencontainers/runtime-spec/specs-go"
)

//go:generate counterfeiter . UidGenerator
//go:generate counterfeiter . UserLookupper
//go:generate counterfeiter . EnvDeterminer
//go:generate counterfeiter . Mkdirer
//go:generate counterfeiter . BundleLoader
//go:generate counterfeiter . ProcessTracker
//go:generate counterfeiter . Process

type UidGenerator interface {
	Generate() string
}

type ExecUser struct {
	Uid   int
	Gid   int
	Sgids []int
	Home  string
}

type UserLookupper interface {
	Lookup(rootFsPath string, user string) (*ExecUser, error)
}

type Mkdirer interface {
	MkdirAs(rootFSPathFile string, uid, gid int, mode os.FileMode, recreate bool, path ...string) error
}

type LookupFunc func(rootfsPath, user string) (*ExecUser, error)

func (fn LookupFunc) Lookup(rootfsPath, user string) (*ExecUser, error) {
	return fn(rootfsPath, user)
}

type EnvDeterminer interface {
	EnvFor(bndl goci.Bndl, spec ProcessSpec) []string
}

type EnvFunc func(bndl goci.Bndl, spec ProcessSpec) []string

func (fn EnvFunc) EnvFor(bndl goci.Bndl, spec ProcessSpec) []string {
	return fn(bndl, spec)
}

type BundleLoader interface {
	Load(path string) (goci.Bndl, error)
}

type Process interface {
	garden.Process
}

type ProcessTracker interface {
	Run(processID string, cmd *exec.Cmd, io garden.ProcessIO, tty *garden.TTYSpec, pidFile string) (garden.Process, error)
	Attach(processID string, io garden.ProcessIO, pidFilePath string) (garden.Process, error)
}

//go:generate counterfeiter . ExecRunner
type ExecRunner interface {
	Run(log lager.Logger, passedID string, spec *PreparedSpec, bundlePath, processesPath, handle string, tty *garden.TTYSpec, io garden.ProcessIO) (garden.Process, error)
	Attach(log lager.Logger, processID string, io garden.ProcessIO, processesPath string) (garden.Process, error)
}

type PreparedSpec struct {
	specs.Process
	ContainerRootHostUID uint32
	ContainerRootHostGID uint32
}

//go:generate counterfeiter . ProcessBuilder
type ProcessBuilder interface {
	BuildProcess(bndl goci.Bndl, processSpec ProcessSpec) *PreparedSpec
}

type ProcessSpec struct {
	garden.ProcessSpec
	ContainerUID int
	ContainerGID int
}

//go:generate counterfeiter . Waiter
//go:generate counterfeiter . Runner

type Waiter interface {
	Wait() (int, error)
}

type Runner interface {
	Run(log lager.Logger)
}

//go:generate counterfeiter . WaitWatcher

type WaitWatcher interface { // get it??
	OnExit(log lager.Logger, process Waiter, onExit Runner)
}

type Watcher struct{}

func (w Watcher) OnExit(log lager.Logger, process Waiter, onExit Runner) {
	process.Wait()
	onExit.Run(log)
}

type RemoveFiles []string

func (files RemoveFiles) Run(log lager.Logger) {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Error("cleanup-process-json-failed", err)
		}
	}
}

func intersect(l1 []string, l2 []string) (result []string) {
	for _, a := range l1 {
		for _, b := range l2 {
			if a == b {
				result = append(result, a)
			}
		}
	}

	return result
}
