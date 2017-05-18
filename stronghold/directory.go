package stronghold

import (
	"os"
	"path/filepath"
)

var markers = map[string]string{
	"stronghold": ".stronghold",
	"git":        ".git",
	"mercurial":  ".hg",
	"svn":        ".svn",
}

// UpsertStrongholdDirectory creates the stronghold directory if need be. It
// returns the newly created path.
func UpsertStrongholdDirectory() (string, error) {
	strongholdDirectory, err := GetStrongholdDirectory()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(strongholdDirectory, 0700)
	if err != nil {
		return "", err
	}
	return strongholdDirectory, nil
}

// GetStrongholdDirectory joins the root directory with the Stronghold marker
// to forge the path in which this package should operate. Returns the path.
func GetStrongholdDirectory() (string, error) {
	rootDirectory, err := GetRootDirectory()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootDirectory, markers["stronghold"]), nil
}

// GetRootDirectory recursively crawls up the filesystem until it finds a marker
// (either a VCS marker or a Stronghold marker). Returns the directory without
// the marker.
func GetRootDirectory() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	currentDirectory := workingDirectory
	for {
		for _, marker := range markers {
			stat, err := os.Stat(filepath.Join(currentDirectory, marker))
			if err == nil && stat.IsDir() {
				return currentDirectory, nil
			}
		}
		if currentDirectory == "/" {
			return workingDirectory, nil
		}
		currentDirectory = filepath.Dir(currentDirectory)
	}
}

// WipeStrongholdDirectory deletes the Stronghold directory. Returns the
// deleted path and the error if any.
func WipeStrongholdDirectory() (string, error) {
	strongholdDirectory, err := GetStrongholdDirectory()
	if err != nil {
		return "", err
	}
	err = os.RemoveAll(strongholdDirectory)
	if err != nil {
		return "", err
	}
	return strongholdDirectory, nil
}
