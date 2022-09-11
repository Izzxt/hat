/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/spf13/cobra"
)

// habboswfCmd represents the habboswf command
var habboswfCmd = &cobra.Command{
	Use:   "habboswf",
	Short: "Download latest habbo swf",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		p := d.GetCurrentProduction()
		if d.GetOutput() != "" {
			d.SetOutput(d.GetOutput())
		} else {
			d.SetOutput(fmt.Sprintf("gordon/%s", p))
		}

		d.SetGordon()
		d.SetProduction(p)
		d.SetPath("/")
		d.SetFileName("/Habbo.swf")
		d.Download()
	},
}

func init() {
	rootCmd.AddCommand(habboswfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// habboswfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// habboswfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
