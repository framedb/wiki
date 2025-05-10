package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/alexflint/go-arg"
	"github.com/soyart/far"
)

type cli struct {
	RenameAsset  *cmdRenameAsset  `arg:"subcommand:rename" help:"Rename 1 asset file and replace all references to the asset with the new filename"`
	RenameAssets *cmdRenameAssets `arg:"subcommand:renames" help:"Read a JSON replacement map from file and rename those assets as well as references to the asset with the new filenames"`
}

type cmdRenameAsset struct {
	Force bool   `arg:"-f,--force" help:"Ignore warning such as mismatch file extensions"`
	Old   string `arg:"required,positional" help:"Old asset filename"`
	New   string `arg:"required,positional" help:"New asset filename"`
}

type cmdRenameAssets struct {
	Force bool   `arg:"-f,--force" help:"Ignore warning such as mismatch file extensions"`
	Path  string `arg:"required,positional" help:"Path to replacement map in JSON"`
}

const basePath = "./assets"

func main() {
	c := cli{}
	arg.MustParse(&c)

	var err error
	switch {
	case c.RenameAsset != nil:
		err = c.RenameAsset.run()
	case c.RenameAssets != nil:
		err = c.RenameAssets.run()
	default:
		err = errors.New("unmatched subcommand")
	}

	if err != nil {
		panic(err)
	}
}

func (c *cmdRenameAsset) run() error {
	extOld, extNew := filepath.Ext(c.Old), filepath.Ext(c.New)
	if extOld != extNew && !c.Force {
		return fmt.Errorf("mismatch extension '%s' vs '%s' from src '%s' and dst '%s'", extOld, extNew, c.Old, c.New)
	}
	return renameWithReference(c.Old, c.New)
}

func (c *cmdRenameAssets) run() error {
	j, err := os.ReadFile(c.Path)
	if err != nil {
		return fmt.Errorf("failed to read replacement JSON from file '%s': %w", c.Path, err)
	}
	var m map[string]string
	err = json.Unmarshal(j, &m)
	if err != nil {
		return fmt.Errorf("failed to marshal replacement JSON from file '%s': %w", c.Path, err)
	}

	// Deterministic runs, hence slice
	s := make([][2]string, len(m))
	i := 0
	for old, new := range m {
		if old == "" {
			return fmt.Errorf("found empty replacement key")
		}
		if new == "" {
			return fmt.Errorf("found empty replacement value")
		}
		extK, extV := filepath.Ext(old), filepath.Ext(new)
		if extK != extV && !c.Force {
			return fmt.Errorf("mismatch extension '%s' vs '%s' from src '%s' and dst '%s'", extK, extV, old, new)
		}

		s[i] = [2]string{old, new}
		i++
	}

	sort.Slice(s, func(i, j int) bool {
		return s[0][0] < s[1][0]
	})

	for i := range s {
		old, new := s[i][0], s[i][1]
		err := renameWithReference(old, new)
		if err != nil {
			return err
		}
	}

	return nil
}

func renameWithReference(old, new string) error {
	pathOld := filepath.Join(basePath, old)
	pathNew := filepath.Join(basePath, new)

	statOld, err := os.Stat(pathOld)
	if err != nil {
		return fmt.Errorf("failed to stat old path '%s': %w", pathOld, err)
	}
	if statOld.IsDir() {
		return fmt.Errorf("path '%s' is a directory", pathOld)
	}
	_, err = os.Stat(pathNew)
	if !os.IsNotExist(err) {
		return fmt.Errorf("path '%s' already exists", pathNew)
	}
	data, err := os.ReadFile(pathOld)
	if err != nil {
		return fmt.Errorf("failed to read data for copying '%s': %w", pathOld, err)
	}
	err = os.WriteFile(pathNew, data, statOld.Mode().Perm())
	if err != nil {
		return fmt.Errorf("failed to copy data to '%s': %w", pathNew, err)
	}
	err = far.FindAndReplace("pages", old, new)
	if err != nil {
		return fmt.Errorf("failed to find and replace in pages: %w", err)
	}
	err = far.FindAndReplace("journals", old, new)
	if err != nil {
		return fmt.Errorf("failed to find and replace in journals: %w", err)
	}
	err = os.Remove(pathOld)
	if err != nil {
		return fmt.Errorf("failed to remove old file '%s': %w", pathOld, err)
	}
	return nil
}
