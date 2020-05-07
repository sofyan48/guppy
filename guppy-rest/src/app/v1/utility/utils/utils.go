package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/sofyan48/guppy/guppy-cli/entity"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Utils ...
type Utils struct{}

// UtilsHandler ...
func UtilsHandler() *Utils {
	return &Utils{}
}

// UtilsInterface ..
type UtilsInterface interface {
	ConvertUnixTime(unixTime int64) time.Time

	CheckFile(path string) bool
	MakeDirs(path string) error
	FileRemove(path string) error
	CreateFile(path string) bool
	WriteFile(path string, value string, perm os.FileMode) bool
	ReadFile(path string, perm os.FileMode) string
	DeleteFile(path string) bool
	ReadHome() string
	GetCurrentPath() string

	CheckEnvironment(path string) string
	ParseJSON(data string) (map[string]interface{}, error)
	CheckTemplateFile(path string) (string, error)
	ParsingYAML(path string) (*entity.TemplatesModels, error)
}

// CheckFile function check folder
// @path : string
// return error
func (util *Utils) CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// MakeDirs fucntion create directory
// @path : string
// return error
func (util *Utils) MakeDirs(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// FileRemove Remove Files
// @path : string
// return error
func (util *Utils) FileRemove(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile function create file
// @path : string
// return bool
func (util *Utils) CreateFile(path string) bool {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, _ = os.Create(path)
		defer file.Close()
		return false
	}
	return true
}

// WriteFile func write local file
func (util *Utils) WriteFile(path string, value string, perm os.FileMode) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if err != nil {
		return false
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(value)
	if err != nil {
		return false
	}
	// save changes
	err = file.Sync()
	if err != nil {
		return false
	}

	return true
}

// ReadFile function
func (util *Utils) ReadFile(path string, perm os.FileMode) string {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if err != nil {
		return err.Error()
	}
	defer file.Close()
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			if err != nil {
				return err.Error()
			}
			break
		}
	}
	return string(text)
}

// DeleteFile Function
func (util *Utils) DeleteFile(path string) bool {
	var err = os.Remove(path)
	if err != nil {
		return false
	}
	return true
}

// ReadHome function
// return string
func (util *Utils) ReadHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// CheckEnvironment function check default env
// @path : string
// return bool, error
func (util *Utils) CheckEnvironment(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ""
	}
	return path
}

// GetCurrentPath get current path
// return string
func (util *Utils) GetCurrentPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// ConvertUnixTime ...
// @unixTime: int64
func (util *Utils) ConvertUnixTime(unixTime int64) time.Time {
	tm := time.Unix(unixTime, 0)
	return tm
}

// ParseJSON function conver json string to object
// @data: string
// return map[string]interface{}, error
func (util *Utils) ParseJSON(data string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CheckTemplateFile check template path
// @argsFile: string
func (util *Utils) CheckTemplateFile(path string) (string, error) {
	var templates string
	if path == "" {
		templates = util.GetCurrentPath() + "/guppy.yml"
	} else {
		templates = path
	}
	if !util.CheckFile(templates) {
		return "", cli.NewExitError("No Templates Parse", 1)
	}
	return templates, nil
}

// ParseYAML ...
func (util *Utils) ParseYAML(path string) (*entity.TemplatesModels, error) {
	yamlObject := &entity.TemplatesModels{}
	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return yamlObject, err
	}
	err = yaml.Unmarshal(ymlFile, yamlObject)
	if err != nil {
		return yamlObject, err
	}
	return yamlObject, nil
}
