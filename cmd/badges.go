/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/hat/badges"
	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/spf13/cobra"
)

var badgeCode string

var badgesCmd = &cobra.Command{
	Use:   "badges",
	Short: "Download habbo badges",
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
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/c_images/album1584")
			}
			d.SetOther()
			d.SetPath("/c_images/album1584")
			for _, v := range code {
				exts := fs.IsFileExists(d.GetOutput(), fmt.Sprintf("%s.gif", v))
				if !exts {
					go func(v string) {
						wg.Add(1)
						defer wg.Done()
						d.SetFileName(fmt.Sprintf("%s.gif", v))
						d.Download()
					}(v)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(badgesCmd)

	badgesCmd.PersistentFlags().StringVarP(&badgeCode, "code", "C", "", "Badge code without file extension")
}
