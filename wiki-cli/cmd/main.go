package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/soyart/far"
)

type cli struct {
	RenameAsset *cmdRenameAsset `arg:"subcommand:rename-asset" help:"Rename 1 asset file and replace all references to the asset with the new filename"`
}

type cmdRenameAsset struct {
	Force bool   `arg:"force" help:"Ignore warning such as mismatch file extensions"`
	Old   string `arg:"required,positional" help:"Old asset filename"`
	New   string `arg:"required,positional" help:"New asset filename"`
}

const basePath = "./assets"

func main() {
	c := cli{}
	arg.MustParse(&c)
	switch {
	case c.RenameAsset != nil:
		err := c.RenameAsset.run()
		if err != nil {
			panic(err)
		}
	}
}

func (c *cmdRenameAsset) run() error {
	extOld, extNew := filepath.Ext(c.Old), filepath.Ext(c.New)
	if extOld != extNew && !c.Force {
		return fmt.Errorf("mismatch extension '%s' vs '%s' from src '%s' and dst '%s'", extOld, extNew, c.Old, c.New)
	}

	pathOld := filepath.Join(basePath, c.Old)
	pathNew := filepath.Join(basePath, c.New)

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
	err = far.FindAndReplace("pages", c.Old, c.New)
	if err != nil {
		return fmt.Errorf("failed to find and replace in pages: %w", err)
	}
	err = far.FindAndReplace("journals", c.Old, c.New)
	if err != nil {
		return fmt.Errorf("failed to find and replace in journals: %w", err)
	}
	err = os.Remove(pathOld)
	if err != nil {
		return fmt.Errorf("failed to remove old file '%s': %w", pathOld, err)
	}
	return nil
}
