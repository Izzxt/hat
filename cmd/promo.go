/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/Izzxt/hat/promo"
	"github.com/spf13/cobra"
)

var promoName string

// promoCmd represents the hotelview command
var promoCmd = &cobra.Command{
	Use:   "promo",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)

		p := promo.NewPromo(*d, &wg, &mu)

		if promoName != "" {
			d.SetFileName(fmt.Sprintf("%s.png", promoName))
			d.Download()
		} else {
			images := p.GetAllImages()
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/c_images/web_promo_small")
			}
			d.SetOther()
			d.SetPath("/c_images/web_promo_small")
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
	rootCmd.AddCommand(promoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	promoCmd.PersistentFlags().StringVarP(&promoName, "name", "n", "", "Web Promo Images Name for download single")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hotelviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
