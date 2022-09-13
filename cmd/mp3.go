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

var mp3Name string

// mp3Cmd represents the mp3 command
var mp3Cmd = &cobra.Command{
	Use:   "mp3",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		c := client.NewClient()
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
					exts := fs.IsFileExists(d.GetOutput(), fmt.Sprintf("sound_machine_sample_%d.mp3", i))
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	mp3Cmd.PersistentFlags().StringVarP(&mp3Name, "name", "n", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mp3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
