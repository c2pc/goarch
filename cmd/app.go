package cmd

import (
	"github.com/chincharovpc/goarch/server"
	"github.com/spf13/cobra"
	"log"
)

var appCmd = &cobra.Command{
	Use:   "goarch",
	Short: "A simple golang API server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func Execute() {
	if err := appCmd.Execute(); err != nil {
		log.Fatal(err)
		return
	}
}
