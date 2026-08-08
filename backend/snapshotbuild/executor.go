package snapshotbuild

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CommandExecutor struct{}

func (CommandExecutor) Version(ctx context.Context, executable string) (string, error) {
	command := exec.CommandContext(ctx, executable, "--version")
	command.Env = boundedEnvironment("", "")
	output, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(string(output)))
	}
	version := strings.TrimSpace(string(output))
	if version == "" {
		return "", fmt.Errorf("DuckDB returned an empty version")
	}
	return version, nil
}

func (CommandExecutor) Run(ctx context.Context, invocation Invocation) ([]byte, error) {
	command := exec.CommandContext(
		ctx,
		invocation.Executable,
		invocation.DatabasePath,
		"-json",
		"-c",
		invocation.SQL,
	)
	command.Dir = invocation.WorkingDirectory
	command.Env = boundedEnvironment(invocation.WorkingDirectory, invocation.ExtensionDirectory)
	output, err := command.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, strings.TrimSpace(string(output)))
	}
	return output, nil
}

func boundedEnvironment(home, extensionDirectory string) []string {
	keep := map[string]bool{
		"SystemRoot": true,
		"WINDIR":     true,
		"TMP":        true,
		"TEMP":       true,
		"TMPDIR":     true,
		"LANG":       true,
		"LC_ALL":     true,
	}
	result := make([]string, 0, len(keep)+3)
	for _, entry := range os.Environ() {
		name, _, ok := strings.Cut(entry, "=")
		if ok && keep[name] {
			result = append(result, entry)
		}
	}
	if home != "" {
		result = append(result, "HOME="+filepath.Clean(home), "USERPROFILE="+filepath.Clean(home))
	}
	if extensionDirectory != "" && extensionDirectory != "." {
		result = append(result, "DUCKDB_EXTENSION_DIRECTORY="+filepath.Clean(extensionDirectory))
	}
	return result
}
