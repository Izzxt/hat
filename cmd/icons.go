/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/files"
	"github.com/Izzxt/hat/version"
	"github.com/spf13/cobra"
)

var iconName string

var iconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "Download habbo catalogue icons",
	Run: func(cmd *cobra.Command, args []string) {

		c := client.NewClient()
		version.StartupMessage(c)
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
					exts := files.IsFileExists(d.GetOutput(), fmt.Sprintf("icon_%d.png", i))

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

	iconsCmd.PersistentFlags().StringVarP(&iconName, "name", "n", "", "Icon name without file extension")
}
