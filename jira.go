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
	"github.com/andygrunwald/go-jira"
)

// JIRAClient is the interface to JIRA API.
var JIRAClient *jira.Client

// Login initializes JIRA client.
func Login() {
	fmt.Println("Logging in")
	config := getConfigFromViper()
	var err error
	JIRAClient, err = jira.NewClient(nil, config.URL)
	if err != nil {
		panic(fmt.Errorf("Error logging in: %v\n", err))
	}
	_, err = JIRAClient.Authentication.AcquireSessionCookie(
		config.Username,
		config.Password)
	if err != nil {
		panic(fmt.Errorf("Error authenticating in: %v\n", err))
	}
}