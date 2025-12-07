package files

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// replaceWithSymlink replaces source file tree with symlink of destination file tree
func replaceWithSymlink(src string, dest string) error {
	err := copyFileTree(src, dest)
	if err != nil {
		return err
	}

	err = os.RemoveAll(src)
	if err != nil {
		return err
	}

	err = os.Symlink(dest, src)
	if err != nil {
		return err
	}

	return nil
}

// copyFileTree copies regular files and directories from source to destination
func copyFileTree(src string, dest string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// compute relative path from source to current file (as WalkDir walks)
		// see https://pkg.go.dev/io/fs#WalkDirFunc for more
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// compute destination path
		destPath := filepath.Join(dest, relPath)

		// if file is not regular then continue to next file (if its dir then create dir before
		// continuing to next file)
		if !d.Type().IsRegular() {
			if d.IsDir() {
				err = os.Mkdir(destPath, 0777)
				if err != nil {
					return err
				}
			}

			return nil
		}

		err = copyFile(path, destPath)
		if err != nil {
			return err
		}

		return nil
	})
}

// copyFile copies file from source to destination
func copyFile(src string, dest string) error {
	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, in, 0666)
	if err != nil {
		return err
	}

	return nil
}

// createSymlink creates symlink of source at destination
func createSymlink(src string, dest string) error {
	return os.Symlink(src, dest)
}

// moveFile moves file from source to destination
func moveFile(src string, dest string) error {
	// ensure dest does not exist
	err := os.RemoveAll(dest)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	err = copyFile(src, dest)
	if err != nil {
		return err
	}

	err = os.RemoveAll(src)
	if err != nil {
		return err
	}

	return nil
}
