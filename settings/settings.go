package settings

import (
	"github.com/spf13/afero"
)

// Loader operator to load settings
type Loader struct {
	fs afero.Fs

	container    interface{}
	settingPaths []string
	settingName  string
	settingFile  string
}

// Initialize load settings
func Initialize(fileName string, settingsPtr interface{}) error {
	// TODO
	return nil
}

func (l *Loader) fetchFile() (string, error) {
	// TODO
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
