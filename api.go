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
	"encoding/json"
	"fmt"
	"os/exec"
)

var (
	ErrItemNotFound  = fmt.Errorf("1password item not found")
	ErrVaultNotFound = fmt.Errorf("1password vault not found")
)

// SectionField is a field in the free-form Section chapter of an item
type SectionField struct {
	Label string `json:"t"`
	Kind  string `json:"k"`
	Value string `json:"v"`
}

type Section struct {
	Name   string         `json:"name"`
	Title  string         `json:"title"`
	Fields []SectionField `json:"fields"`
}

type Details struct {
	Password string    `json:"password"`
	Sections []Section `json:"sections"`
}

// PasswordItem is the result of get item on a "password" item.
type PasswordItem struct {
	Details  Details `json:"details"`
	Overview struct {
		Title string `json:"title"`
	} `json:"overview"`
}

type ExecCommand func(name string, arg ...string) *exec.Cmd
type Printer func(fmt string, arg ...interface{})

type Api struct {
	VaultUUID string
	// Allow mocking in tests
	execCmd ExecCommand
	debugf  Printer
}

func NewApi() *Api {
	return &Api{
		execCmd: exec.Command,
		// Default debug: NOP
		debugf: func(fmt string, arg ...interface{}) {},
	}
}

func (a *Api) SetDebugf(f Printer) {
	a.debugf = f
}

func (a *Api) SetVault(uuid string) {
	a.VaultUUID = uuid
}

// Section gets the section with this title (!)
//
// Technically, 1p item sections have a name and a title, and the name seems
// more like a machine readable, unique string to use exactly for this kind of
// purpose. However, it's not accessible from the stock 1 Password GUI, which
// makes this very hard to debug if you ever accidentally change something
// through there (e.g. delete a section and recreate one with the same name;
// "why isn't it working?!"). Basically, using the title, while less elegant,
// follows the principle of least surprise. User > Developer.
func (it *PasswordItem) Section(title string) *Section {
	for i := range it.Details.Sections {
		sec := &it.Details.Sections[i]
		if sec.Title == title {
			return sec
		}
	}
	return nil
}

func (sec *Section) Field(label string) *SectionField {
	for i := range sec.Fields {
		f := &sec.Fields[i]
		if f.Label == label {
			return f
		}
	}
	return nil
}

func (a *Api) Item(name string) (*PasswordItem, error) {
	args := []string{"get", "item"}
	if a.VaultUUID != "" {
		args = append(args, "--vault="+a.VaultUUID)
	}
	args = append(args, name)
	out, err := a.call(args...)
	if err != nil {
		return nil, fmt.Errorf("getItem: %v", err)
	}
	item := &PasswordItem{}
	err = json.Unmarshal(out, item)
	if err != nil {
		a.debugf("undecodable json from op tool: %s", out)
		return nil, fmt.Errorf("getItem: decoding payload from `op` tool failed: %v", err)
	}
	return item, err
}
