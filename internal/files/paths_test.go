package files

import (
	"path/filepath"
	"testing"
)

func TestDotfilesDir(t *testing.T) {
	tmp := t.TempDir()

	got := dotfilesDir(tmp)
	want := filepath.Join(tmp, "dotfiles")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDofDir(t *testing.T) {
	tmp := t.TempDir()

	got := dofDir(tmp)
	want := filepath.Join(tmp, "dotfiles", ".dof")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMetadataFile(t *testing.T) {
	tmp := t.TempDir()

	got := metadataFile(tmp, "metadata.json")
	want := filepath.Join(tmp, "dotfiles", ".dof", "metadata.json")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestNameDir(t *testing.T) {
	tmp := t.TempDir()

	got := nameDir(tmp, "namedir")
	want := filepath.Join(tmp, "dotfiles", "namedir")

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
