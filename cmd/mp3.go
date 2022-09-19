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

var mp3Name string

var mp3Cmd = &cobra.Command{
	Use:   "mp3",
	Short: "Download habbo mp3 songs",
	Run: func(cmd *cobra.Command, args []string) {

		c := client.NewClient()
		version.StartupMessage(c)
		d := downloader.NewDownloader(c)
		d.SetOther()
		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetPath("/dcr/hof_furni/mp3")

		if mp3Name != "" {
			d.SetFileName(fmt.Sprintf("%s.mp3", mp3Name))
			d.Download()
		} else {
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/dcr/hof_furni/mp3")
			}
			i := 1
			run := true
			attempt := 0

			for run {
				d.SetFileName(fmt.Sprintf("sound_machine_sample_%d.mp3", i))
				byte, _ := d.Fetch()
				mimeType := http.DetectContentType(byte)

				if mimeType == "application/octet-stream" {
					exts := files.IsFileExists(d.GetOutput(), fmt.Sprintf("sound_machine_sample_%d.mp3", i))
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
	rootCmd.AddCommand(mp3Cmd)

	mp3Cmd.PersistentFlags().StringVarP(&mp3Name, "name", "n", "", "MP3 name without file extension")
}
