/*
Copyright © 2022 Izzat
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

var promoCmd = &cobra.Command{
	Use:   "promo",
	Short: "Download habbo web promo",
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

	promoCmd.PersistentFlags().StringVarP(&promoName, "name", "n", "", "Image name without file extension")
}
