/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/files"
	"github.com/Izzxt/hat/hotelview"
	"github.com/Izzxt/hat/version"
	"github.com/spf13/cobra"
)

var hotelViewName string

var hotelviewCmd = &cobra.Command{
	Use:   "hotelview",
	Short: "Download habbo hotel view",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
		version.StartupMessage(c)
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)

		p := hotelview.NewHotelView(*d, &wg, &mu)

		if hotelViewName != "" {
			d.SetFileName(hotelViewName)
			d.Download()
		} else {
			images := p.GetAllImages()
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/c_images/reception")
			}
			d.SetOther()
			d.SetPath("/c_images/reception")
			for _, v := range images {
				exts := files.IsFileExists(d.GetOutput(), v)
				if !exts {
					go func(v string) {
						wg.Add(1)
						defer wg.Done()
						d.SetFileName(v)
						d.Download()
					}(v)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hotelviewCmd)

	hotelviewCmd.PersistentFlags().StringVarP(&hotelViewName, "name", "n", "", "Image name without file extension")
}
