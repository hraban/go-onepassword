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
	"os/exec"
	"regexp"
)

func isItemNotFoundErr(err *exec.ExitError) bool {
	return regexp.MustCompile(`Unable to find item`).Match(err.Stderr)
}

func isVaultNotFoundErr(err *exec.ExitError) bool {
	return regexp.MustCompile(`Vault .* not found`).Match(err.Stderr)
}

func (a *Api) handleCallErr(err error) error {
	if ee, ok := err.(*exec.ExitError); ok {
		// I don't think they have a system for this, so I'm just going with
		// whatever looks reliable.
		if isItemNotFoundErr(ee) {
			return ErrItemNotFound
		}
		if isVaultNotFoundErr(ee) {
			return ErrVaultNotFound
		}
		a.debugf("op stderr: %s", ee.Stderr)
	}
	return fmt.Errorf("op command: %v", err)
}

func (a *Api) call(args ...string) ([]byte, error) {
	cmd := a.execCmd("op", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, a.handleCallErr(err)
	}
	return out, nil
}
