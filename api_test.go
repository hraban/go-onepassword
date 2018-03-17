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
	"os"
	"os/exec"
	"testing"
)

func TestSetVault(t *testing.T) {
	a := NewApi()
	a.SetVault("foobar")
	if a.VaultUUID != "foobar" {
		t.Errorf("Got back unexpected vault id")
	}
}

func buildMock(testId string) ExecCommand {
	return func(name string, args ...string) *exec.Cmd {
		argv := []string{"-test.run=TestHelperProcess", "--", name}
		argv = append(argv, args...)
		cmd := exec.Command(os.Args[0], argv...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=" + testId}
		return cmd
	}
}

func TestGet(t *testing.T) {
	a := NewApi()
	a.SetVault("vkltlvui3cc8fifyz2s346qafq")
	a.debugf = t.Logf
	a.execCmd = buildMock("test-get")
	item, err := a.Item("test")
	if err != nil {
		t.Fatalf("Item: %v", err)
	}
	sec := item.Section("my title")
	if sec == nil {
		t.Fatal("item contained no section 'my title'")
	}
	f := sec.Field("some label")
	if f == nil {
		t.Fatal("section 'my title' did not contain field 'some label'")
	}
	if f.Value != "labelʼs value 👍" {
		t.Fatalf("Unexpected value for label: %q", f.Value)
	}
}
