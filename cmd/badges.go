/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/hat/badges"
	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/spf13/cobra"
)

var badgeCode string

// badgesCmd represents the badges command
var badgesCmd = &cobra.Command{
	Use:   "badges",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)

		b := badges.NewBadges(*d, &wg, &mu)

		defer wg.Wait()

		if badgeCode != "" {
			d.SetFileName(fmt.Sprintf("%s.gif", badgeCode))
			d.Download()
		} else {
			code := b.GetAllCode()
			if d.GetOutput() != "" {
				d.SetOutput(d.GetOutput())
			} else {
				d.SetOutput("resource/c_images/album1584")
			}
			d.SetPath("/c_images/album1584")
			for _, v := range code {
				wg.Add(1)
				go func(v string) {
					defer wg.Done()
					d.SetFileName(fmt.Sprintf("%s.gif", v))
					d.Download()
				}(v)
				time.Sleep(100 * time.Millisecond)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(badgesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	badgesCmd.PersistentFlags().StringVarP(&badgeCode, "code", "C", "", "Badge code to download single")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// badgesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
