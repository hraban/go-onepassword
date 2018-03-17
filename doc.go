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

// 1 Password API package, a simple wrapper around the `op` CLI tool
//
// This might, theoretically, be better off as a stand-alone library. However,
// realistically speaking, this is probably not actually a good abstraction of
// the 1 password CLI. I wrote this without any reference documentation, with a
// lot of guess work. It does the job for the github.com/99designs/keyring
// package, but it probably needs a lot of work before it can be properly spun
// off.
package onepassword
