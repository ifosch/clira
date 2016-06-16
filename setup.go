// Copyright © 2016 Ignasi Fosch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clira

import (
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
)

// Config contains Configuration parameters.
type Config struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func validConfigFileContents() bool {
	valid := viper.InConfig("url")
	valid = valid && viper.InConfig("username")
	valid = valid && viper.InConfig("password")
	return valid
}

var validConfig = validConfigFileContents

func goOnWithSetup() bool {
	if validConfig() {
		fmt.Println(viper.ConfigFileUsed(), "already contains valid configuration.")
		return confirm("Do you want me to discard it? ")
	}
	fmt.Println("No valid configuration found in", viper.ConfigFileUsed())
	return true
}

func getConfigFromUser() Config {
	fmt.Print("What's the URL of your JIRA instance? ")
	url := ask()
	fmt.Print("What's your username? ")
	username := ask()
	fmt.Print("What's your password? ")
	password := ask()
	return Config{url, username, password}
}

func getConfigFromViper() Config {
	url := viper.GetString("url")
	username := viper.GetString("username")
	password := viper.GetString("password")
	return Config{url, username, password}
}

func marshalConfig(config Config) ([]byte, error) {
	content, err := yaml.Marshal(config)
	return content, err
}

// ConfigStore is a config interface.
type ConfigStore interface {
	WriteString(string) (int, error)
	Close()
}

// ConfigFile is a config file wrapper.
type ConfigFile struct {
	file *os.File
}

// WriteString enables os.File.WriteString mocking
func (cf ConfigFile) WriteString(content string) (int, error) {
	written, err := cf.file.WriteString(content)
	return written, err
}

// Close enables os.File.Close mocking
func (cf ConfigFile) Close() {
	cf.file.Close()
}

func createFile(path string) (ConfigStore, error) {
	file, err := os.Create(path)
	configFile := ConfigFile{file}
	return configFile, err
}

var toYAML = marshalConfig
var openFile = createFile

func saveConfigToFile(config Config) error {
	YAMLConfig, err := toYAML(config)
	if err != nil {
		return err
	}
	f, err2 := openFile(viper.ConfigFileUsed())
	if err2 != nil {
		return err2
	}
	defer f.Close()
	_, err3 := f.WriteString(string(YAMLConfig))
	return err3
}

var askConfig = getConfigFromUser
var saveConfig = saveConfigToFile

// Setup sets up new or existing config file.
func Setup() error {
	if goOnWithSetup() {
		saveConfig(askConfig())
	}
	return nil
}
