/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"fmt"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/version"
	"github.com/spf13/cobra"
)

var habboswfCmd = &cobra.Command{
	Use:   "habboswf",
	Short: "Download Habbo.swf",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		version.StartupMessage(c)
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		p := d.GetCurrentProduction()

		if Output != "" {
			d.SetOutput(Output)
		} else {
			d.SetOutput(fmt.Sprintf("resource/gordon/%s", p))
		}

		d.SetGordon()
		d.SetProduction(p)
		// d.SetPath("/")
		d.SetFileName("Habbo.swf")
		d.Download()
	},
}

func init() {
	rootCmd.AddCommand(habboswfCmd)
}
