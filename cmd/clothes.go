/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/clothes"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/Izzxt/hat/xml"
	"github.com/spf13/cobra"
)

var (
	clothesName string
)

var clothesCmd = &cobra.Command{
	Use:   "clothes",
	Short: "Download habbo clothes",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		keys := make(map[string]bool)
		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetOutput(Output)
		d.SetDomain(Domain)
		d.SetGordon()
		d.SetPath("/")

		if Prod == "" {
			Prod = d.GetCurrentProduction()
		}

		d.SetProduction(Prod)

		if clothesName != "" {
			d.SetFileName(fmt.Sprintf("%s.swf", clothesName))
			d.Download()
		} else {
			d.SetFileName("figuremapv2.xml")

			byte, _ := d.Fetch()

			var figure xml.FigureMap

			xml.Parse(&figure, strings.NewReader(string(byte)))

			if d.GetOutput() != "" {
				d.SetOutput(d.GetOutput())
			} else {
				d.SetOutput(fmt.Sprintf("resource/gordon/%s", d.GetProduction()))
			}

			cl := clothes.NewClothes(*d, &wg, &mu)

			for _, entry := range figure.Lib {
				if _, value := keys[entry.Id]; !value {
					keys[entry.Id] = true
					exts := fs.IsFileExists(d.GetOutput(), fmt.Sprintf("/%s.swf", entry.Id))
					if !exts {
						wg.Add(1)
						go func(v xml.FigureLib) {
							cl.Download(fmt.Sprintf("%s", v.Id), Prod)
						}(entry)
						time.Sleep(100 * time.Millisecond)
					}
				}
			}
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(clothesCmd)

	clothesCmd.PersistentFlags().StringVarP(&clothesName, "name", "n", "", "Clothes name without file extension")
}
