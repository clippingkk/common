package settings

import (
	"path/filepath"

	"github.com/spf13/afero"
)

var (
	defaultPaths = []string{"", "./conf", "/conf", "../", "../conf", "../..", "../../conf"}
)

// Loader operator to load settings
type Loader struct {
	fs afero.Fs

	container    interface{}
	paths []string
	fileName  string
	file  string
}

// Initialize load settings
func Initialize(fileName string, settingsPtr interface{}) error {
	// TODO
	return nil
}

func (l *Loader) fetchFile() (string, error) {
	for _, path := range l.paths {
		cfgFile := filepath.Join(path, l.fileName)
		res, err := afero.Exists(l.fs, cfgFile)
		if err != nil {
			return "", err
		}

		if res {
			return cfgFile, nil
		}
	}

	return "", nil
}

func (l *Loader) loadSettings() error {
	// TODO
	return nil
}

func (l *Loader) unmarshalReader() (interface{}, error) {
	// TODO
	return nil, nil
}
