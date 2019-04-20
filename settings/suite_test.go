package settings

import (
	"bufio"
	"errors"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type Suitor struct {
	suite.Suite

	loader *Loader
	JSONFile string
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suitor *Suitor) SetupSuite() {
	suitor.JSONFile = "test.dev.json"

	suitor.loader = &Loader{
		fs: afero.NewOsFs(),
		paths: defaultPaths,
		fileName: suitor.JSONFile,
	}

	initConfig(suitor.JSONFile, jsonExample)
}

// The SetupTest method will be run before every test in the suite.
func (suitor *Suitor) SetupTest() {
}

// The TearDownTest method will be run after every test in the suite.
func (suitor *Suitor) TearDownTest() {
}

// TearDownAllSuite has a TearDownSuite method, which will run after
// all the tests in the suite have been run.
func (suitor *Suitor) TearDownSuite() {
	os.Remove(suitor.JSONFile)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suitor))
}

type AppConfig struct {
	AppName  string         `json:"app_name" yaml:"app_name" toml:"app_name"`
	AppEnv   string         `json:"app_env" yaml:"app_env" toml:"app_env" env:"APP_ENV"`
	Debug    bool           `json:"debug" yaml:"debug" toml:"debug" env:"APP_DEBUG"`
	Database DatabaseConfig `json:"database" yaml:"database" toml:"database"`
}

type DatabaseConfig struct {
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn" env:"DB_DSN"`
}


func initConfig(file, config string) error {
	outputFile, outputError := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		return errors.New("an error occurred with file opening or creation")
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(config)
	return outputWriter.Flush()
}

var jsonExample = `
{
	"app_name": "test-app",
	"app_env": "test",
	"debug": true,
	"database": {
	  "dsn": "root@tcp(localhost:3306)/test"
	}
  }
`