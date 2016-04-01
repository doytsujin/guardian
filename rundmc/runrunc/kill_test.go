package runrunc_test

import (
	"os/exec"

	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc"
	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc/fakes"
	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager"
	"github.com/pivotal-golang/lager/lagertest"
)

var _ = Describe("Kill", func() {
	var (
		commandRunner *fake_command_runner.FakeCommandRunner
		loggingRunner *fakes.FakeRuncCmdRunner
		runcBinary    *fakes.FakeRuncBinary
		logger        *lagertest.TestLogger

		killer *runrunc.Killer
	)

	BeforeEach(func() {
		runcBinary = new(fakes.FakeRuncBinary)
		commandRunner = fake_command_runner.New()
		loggingRunner = new(fakes.FakeRuncCmdRunner)
		logger = lagertest.NewTestLogger("test")

		killer = runrunc.NewKiller(loggingRunner, runcBinary)

		runcBinary.KillCommandStub = func(id, signal, logFile string) *exec.Cmd {
			return exec.Command("funC", "--log", logFile, "kill", id, signal)
		}

		loggingRunner.RunAndLogStub = func(_ lager.Logger, fn runrunc.LoggingCmd) error {
			return commandRunner.Run(fn("potato.log"))
		}
	})

	It("runs 'runc kill' in the container directory using the logging runner", func() {
		Expect(killer.Kill(logger, "some-container")).To(Succeed())
		Expect(commandRunner).To(HaveExecutedSerially(fake_command_runner.CommandSpec{
			Path: "funC",
			Args: []string{"--log", "potato.log", "kill", "some-container", "KILL"},
		}))
	})
})
