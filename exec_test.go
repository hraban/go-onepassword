// This file is part of go-onepassword.
//
// go-onepassword is free software: you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, version 3 of the License.
//
// Foobar is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public
// License for more details.
//
// You should have received a copy of the GNU Lesser General Public
// License along with Foobar. If not, see
// <http://www.gnu.org/licenses/>.

package onepassword

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

const passwordItemJson = `
{
  "uuid": "vhudxtufw5dd5k5c6njsnpkmgy",
  "vaultUuid": "vkltlvui3cc8fifyz2s346qafq",
  "templateUuid": "005",
  "createdAt": "2018-03-17 17:35:40 +0000 UTC",
  "updatedAt": "2018-03-17 18:04:56 +0000 UTC",
  "changerUuid": "I52VYSY75VA33KMH7OOSUJLJOE",
  "overview": {
    "ainfo": "17 Mar 2018 at 18:35:40",
    "pbe": 42.484812912008536,
    "pgrng": true,
    "ps": 56,
    "title": "test password"
  },
  "details": {
    "password": "malapert.phiz.mandate",
    "sections": [
      {
        "name": "linked items",
        "title": "Related Items"
      },
      {
        "fields": [
          {
            "k": "string",
            "n": "B0B05F16C4CB433F90518FDC8025472F",
            "t": "some label",
            "v": "labelʼs value 👍"
          }
        ],
        "name": "Section_4CB2C9FD7CA84B33B10B3C0BEEB5FF97",
        "title": "my title"
      }
    ]
  }
}
`

func helper1() {
	expectedArgv := []string{
		"op", "get", "item", "--vault=vkltlvui3cc8fifyz2s346qafq", "test"}
	if !reflect.DeepEqual(expectedArgv, os.Args[3:]) {
		fmt.Fprintf(os.Stderr, "unexpected argv: %#v", os.Args)
		os.Exit(1)
	}
	fmt.Print(passwordItemJson)
	os.Exit(0)
}

func TestHelperProcess(t *testing.T) {
	switch os.Getenv("GO_WANT_HELPER_PROCESS") {
	case "test-get":
		helper1()
	}
}
