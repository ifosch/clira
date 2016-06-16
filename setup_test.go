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
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestValidConfigFileContents(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("non-existing-clira.yaml")
	if validConfigFileContents() {
		t.Error("Expected invalid config file content")
	}
	viper.SetConfigFile("fixtures/config/clira.yaml")
	if validConfigFileContents() {
		t.Error("Expected invalid config file content")
	}
	viper.SetDefault("url", "http://your.jira-instance.com")
	viper.SetDefault("username", "your_username")
	viper.SetDefault("password", "your_password")
	if validConfigFileContents() {
		t.Error("Expected invalid config file content")
	}
	viper.Set("url", "http://your.jira-instance.com")
	viper.Set("username", "your_username")
	viper.Set("password", "your_password")
	if validConfigFileContents() {
		t.Error("Expected invalid config file content")
	}
	viper.ReadInConfig()
	if !validConfigFileContents() {
		t.Error("Expected valid config file content")
	}
}

var mockedValidConfigResponse = false

func mockValidConfigFileContents() bool {
	return mockedValidConfigResponse
}

var mockedConfirmResponse = true

func mockConfirm(message string) bool {
	return mockedConfirmResponse
}

func TestGoOnWithSetup(t *testing.T) {
	validConfig = mockValidConfigFileContents
	confirm = mockConfirm
	if !goOnWithSetup() {
		t.Error("Invalid configuration should keep up with setup")
	}
	mockedValidConfigResponse = true
	if !goOnWithSetup() {
		t.Error("Valid configuration with confirmation should keep up with setup")
	}
	mockedConfirmResponse = false
	if goOnWithSetup() {
		t.Error("Valid configuration without confirmation should not go on with setup")
	}
	validConfig = validConfigFileContents
	confirm = getConfirmation
}

var mockedAskResponse [3]string
var mockedAskResponseIndex = -1

func mockAsk() string {
	mockedAskResponseIndex++
	return mockedAskResponse[mockedAskResponseIndex]
}

func TestGetConfigFromUser(t *testing.T) {
	mockedAskResponse[0] = "http://mysite.jira.com"
	mockedAskResponse[1] = "my_username"
	mockedAskResponse[2] = "my_password"
	ask = mockAsk
	config := getConfigFromUser()
	if config.URL != mockedAskResponse[0] {
		t.Error("Failed to get URL from user")
	}
	if config.Username != mockedAskResponse[1] {
		t.Error("Failed to get username from user")
	}
	if config.Password != mockedAskResponse[2] {
		t.Error("Failed to get password from user")
	}
	ask = getAnswer
}

func TestGetConfigFromViper(t *testing.T) {
	viper.Reset()
	viper.Set("url", "http://mysite.jira.com")
	viper.Set("username", "my_username")
	viper.Set("password", "my_password")
	config := getConfigFromViper()
	if config.URL != mockedAskResponse[0] {
		t.Error("Failed to get URL from viper")
	}
	if config.Username != mockedAskResponse[1] {
		t.Error("Failed to get username from viper")
	}
	if config.Password != mockedAskResponse[2] {
		t.Error("Failed to get password from viper")
	}
}

func TestMarshalConfig(t *testing.T) {
	URL := "http://mysite.jira.com"
	Username := "my_username"
	Password := "my_password"
	config := Config{URL, Username, Password}
	expected := "password: my_password\n"
	expected += "url: http://mysite.jira.com\n"
	expected += "username: my_username\n"
	result, err := marshalConfig(config)
	if err != nil {
		t.Errorf("marshalConfig failed: %v\n", err)
	}
	if expected != string(result) {
		t.Errorf(
			"Result does not satisfy expectation.\n Expected: %v\n Result: %v\n",
			expected,
			string(result))
	}
}

func TestCreateFile(t *testing.T) {
	filePath := "/tmp/clira.test"
	configFile, err := createFile(filePath)
	if err != nil {
		t.Error("Unexpected error when creating the config file")
	}
	_, err = configFile.WriteString("Hola")
	if err != nil {
		t.Error("Unexpected error when writing in the create config file")
	}
	configFile.Close()
	os.Remove(filePath)
}

type mockedConfigFile struct {
	filename string
	content  string
}

func (cf mockedConfigFile) WriteString(content string) (int, error) {
	cf.content = content
	return 0, nil
}

func (cf mockedConfigFile) Close() {}

func mockCreateFile(path string) (ConfigStore, error) {
	configFile := mockedConfigFile{path, ""}
	return configFile, nil
}

func TestSaveConfigToFile(t *testing.T) {
	openFile = mockCreateFile
	URL := "http://mysite.jira.com"
	Username := "my_username"
	Password := "my_password"
	config := Config{URL, Username, Password}
	saveConfigToFile(config)
}
