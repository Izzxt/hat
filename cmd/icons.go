/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/spf13/cobra"
)

var iconName string

// iconsCmd represents the icons command
var iconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetOther()
		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetPath("/c_images/catalogue")

		if iconName != "" {
			d.SetFileName(fmt.Sprintf("%s.png", iconName))
			d.Download()
		} else {
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/c_images/catalogue")
			}
			i := 1
			run := true
			attempt := 0

			for run {
				d.SetFileName(fmt.Sprintf("icon_%d.png", i))
				byte, _ := d.Fetch()
				mimeType := http.DetectContentType(byte)

				if mimeType == "image/png" {
					exts := fs.IsFileExists(d.GetOutput(), fmt.Sprintf("icon_%d.png", i))

					if !exts {
						d.Download()
						time.Sleep(100 * time.Millisecond)
					}
				} else {
					attempt++
				}

				if attempt > 5 {
					run = false
				}
				i++
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(iconsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	iconsCmd.PersistentFlags().StringVarP(&iconName, "name", "n", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iconsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
