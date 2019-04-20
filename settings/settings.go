package settings

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"path/filepath"
	"reflect"

	"github.com/jinzhu/copier"
	"github.com/spf13/afero"
)

var (
	defaultPaths = []string{"", "./conf", "/conf", "../", "../conf", "../..", "../../conf"}
	loader *Loader
)

// Loader operator to load settings
type Loader struct {
	fs afero.Fs

	container    interface{}
	paths []string
	fileName  string
	file  string
}

func init() {
	loader = &Loader{
		fs: afero.NewOsFs(),
		container: nil,
		paths: defaultPaths,
		fileName: "",
		file: "",
	}
}

// Initialize load settings
func Initialize(fileName string, settingsPtr interface{}) error {
	loader.fileName = fileName
	loader.container = settingsPtr

	if err := loader.fetchFile(); err != nil {
		return err
	}

	if err := loader.loadSettings(); err != nil {
		return err
	}

	copier.Copy(settingsPtr, loader.container)
	return nil
}

func (l *Loader) fetchFile() error {
	for _, path := range l.paths {
		cfgFile := filepath.Join(path, l.fileName)
		res, err := afero.Exists(l.fs, cfgFile)
		if err != nil {
			return err
		}

		if res {
			l.file = cfgFile
			return nil
		}
	}

	return errors.New("settings file not found")
}

func (l *Loader) loadSettings() error {
	file, err := afero.ReadFile(l.fs, l.fileName)
	if err != nil {
		return err
	}

	cfg, err := l.unmarshalReader(bytes.NewReader(file))
	if err != nil {
		return err
	}

	l.container = cfg
	return nil
}

// only support json now
func (l *Loader) unmarshalReader(file io.Reader) (interface{}, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	cfg := copyObject(l.container)

	if err := json.Unmarshal(buf.Bytes(), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// return inner type copy
func copyObject(ptr interface{}) interface{} {
	return reflect.New(reflect.Indirect(reflect.ValueOf(ptr)).Type()).Interface()
}
