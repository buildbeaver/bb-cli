package exec

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/buildbeaver/buildbeaver/runner/runtime"
)

type Config struct {
	runtime.Config
	ShellOrNil *string
}

// Runtime executes jobs directly on the host machine.
type Runtime struct {
	config Config
}

func NewRuntime(config Config) *Runtime {
	return &Runtime{config: config}
}

// Start initializes the runtime and prepares it to have commands Exec'd inside it.
func (r *Runtime) Start(ctx context.Context) error {
	return nil
}

// Stop tears down the runtime.
func (r *Runtime) Stop(ctx context.Context) error {
	return nil
}

// Exec executes a command inside the runtime.
// Start must have been called before calling Exec.
func (r *Runtime) Exec(ctx context.Context, config runtime.ExecConfig) error {
	scriptPath, err := runtime.WriteScript(r.config.StagingDir, config.Name, config.Commands)
	if err != nil {
		return err
	}
	shell := runtime.ShellOrDefault(runtime.GetHostOS(), r.config.ShellOrNil)
	cmd := exec.CommandContext(ctx, shell, scriptPath)
	cmd.Dir = r.config.WorkspaceDir
	cmd.Env = config.Env
	cmd.Stdout = config.Stdout
	cmd.Stderr = config.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running command: %w", err)
	}
	return nil
}

// StartService starts a service inside the runtime.
// The service must be resolvable by name to commands run with Exec.
// Service names are unique within the runtime - it is an error to try start service with the same name twice.
func (r *Runtime) StartService(ctx context.Context, config runtime.ServiceConfig) error {
	return fmt.Errorf("services are not supported with exec jobs")
}

// CleanUp removes any resources left over from previous commands that may not have finished cleanly.
func (r *Runtime) CleanUp(ctx context.Context) error {
	// For Exec runtimes there are no services and commands run inline, so there's nothing to do.
	return nil
}
