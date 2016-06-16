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
	"strings"
)

func isYes(reply string) bool {
	return strings.ToLower(reply)[0] == 'y'
}

func getAnswer() string {
	var answer string
	fmt.Scanln(&answer)
	return answer
}

var ask = getAnswer

func getConfirmation(message string) bool {
	fmt.Print(message)
	return isYes(ask())
}

var confirm = getConfirmation
