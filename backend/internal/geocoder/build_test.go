package geocoder

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReplaceSnapshotRestoresExistingSnapshotWhenInstallFails(t *testing.T) {
	directory := t.TempDir()
	outputPath := filepath.Join(directory, "geocoder.sqlite")
	temporaryPath := filepath.Join(directory, "replacement.sqlite")
	if err := os.WriteFile(outputPath, []byte("existing snapshot"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(temporaryPath, []byte("replacement snapshot"), 0o600); err != nil {
		t.Fatal(err)
	}

	renameCalls := 0
	rename := func(oldPath, newPath string) error {
		renameCalls++
		switch renameCalls {
		case 1:
			return errors.New("direct replacement is unsupported")
		case 3:
			return errors.New("replacement installation failed")
		default:
			return os.Rename(oldPath, newPath)
		}
	}

	if err := replaceSnapshotWithRename(temporaryPath, outputPath, rename); err == nil {
		t.Fatal("expected replacement failure")
	}
	contents, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "existing snapshot" {
		t.Fatalf("output contents=%q", contents)
	}
	if _, err := os.Stat(temporaryPath); err != nil {
		t.Fatalf("replacement should remain available for cleanup: %v", err)
	}
}
