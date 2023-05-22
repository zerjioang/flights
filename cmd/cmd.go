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

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zerjioang/flights/server"
)

var rootCmd = &cobra.Command{
	Use:   "flights",
	Short: "flights API example",
	Long:  `flights API example`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
