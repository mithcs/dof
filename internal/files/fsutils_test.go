package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	tmp := t.TempDir()

	src := filepath.Join(tmp, "src")
	dest := filepath.Join(tmp, "dest")

	err := os.WriteFile(src, []byte("hello.world"), 0666)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	err = copyFile(src, dest)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	got, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	want := "hello.world"
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCopyFileTree(t *testing.T) {
	tmp := t.TempDir()

	src := filepath.Join(tmp, "src")
	file := filepath.Join(src, "testfile")
	dest := filepath.Join(tmp, "dest")

	err := os.Mkdir(src, 0777)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	err = os.WriteFile(file, []byte("hello.world"), 0666)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	// actual func call is here
	err = copyFileTree(src, dest)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	want := "hello.world"
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func TestMoveFile(t *testing.T) {
	tmp := t.TempDir()

	src := filepath.Join(tmp, "src")
	dest := filepath.Join(tmp, "dest")

	err := os.WriteFile(src, []byte("hello.world"), 0666)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	err = moveFile(src, dest)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	_, err = os.ReadFile(src)
	if err == nil {
		t.Error("got nil, expected err")
	}

	got, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	want := "hello.world"
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
