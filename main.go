//
// Copyright zerjioang. 2023 All Rights Reserved.
// Licensed under the MIT
// you may not use this file except in compliance with the License.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"

	"github.com/zerjioang/flights/cmd"
	"github.com/zerjioang/flights/embed"
)

// main is the CLI app main entry point
func main() {
	fmt.Println("Running version: ", embed.Version)
	cmd.Execute()
}
