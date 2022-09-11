/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
	"github.com/Izzxt/hat/xml"
	"github.com/spf13/cobra"
)

var (
	clothesName string
)

// clothesCmd represents the clothes command
var clothesCmd = &cobra.Command{
	Use:   "clothes",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

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
			d.SetFileName(fmt.Sprintf("/%s.swf", clothesName))
			d.Download()
		} else {
			d.SetFileName("/figuremapv2.xml")

			byte, _ := d.Fetch()

			var figure xml.FigureMap

			xml.Parse(&figure, strings.NewReader(string(byte)))

			cl := clothes.NewClothes(*d, &wg, &mu)

			for _, v := range figure.Lib {
				wg.Add(1)
				go func(v xml.FigureLib) {
					cl.Download(fmt.Sprintf("/%s", v.Id), Prod)
				}(v)
				time.Sleep(100 * time.Millisecond)
			}
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(clothesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	clothesCmd.PersistentFlags().StringVarP(&clothesName, "name", "n", "", "Clothes name for single download")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clothesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
