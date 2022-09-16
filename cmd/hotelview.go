/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/Izzxt/hat/hotelview"
	"github.com/spf13/cobra"
)

var hotelViewName string

// hotelviewCmd represents the hotelview command
var hotelviewCmd = &cobra.Command{
	Use:   "hotelview",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
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
				exts := fs.IsFileExists(d.GetOutput(), v)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	hotelviewCmd.PersistentFlags().StringVarP(&hotelViewName, "name", "n", "", "Hotel View Images Name for download single")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hotelviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
