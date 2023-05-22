package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zerjioang/flights/server"
	"log"
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
