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
	"github.com/Izzxt/hat/files"
	"github.com/Izzxt/hat/furnitures"
	"github.com/Izzxt/hat/version"
	"github.com/spf13/cobra"
)

var furniName string
var furniRevision string

var furniCmd = &cobra.Command{
	Use:   "furni [icons]",
	Short: "Download habbo furni or icons",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
		version.StartupMessage(c)
		d := downloader.NewDownloader(c)
		d.SetOutput(Output)
		d.SetDomain(Domain)
		if len(args) < 1 {
			if furniName != "" {
				if furniRevision == "" {
					fmt.Println("--revision is required")
					return
				}
				cmd.MarkFlagRequired("revision")
				d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", furniRevision))
				d.SetFileName(fmt.Sprintf("%s.png", furniName))
				d.Download()
			} else {
				d.SetDomain("com")
				if Output != "" {
					d.SetOutput(Output)
				} else {
					d.SetOutput("resource/dcr/hof_furni")
				}
				f := furnitures.NewFurnitures(*d, &wg, &mu)
				i := f.GetFurnis()
				d.SetOther()

				for _, v := range i {
					exts := files.IsFileExists(d.GetOutput(), v.Name)
					if !exts {
						go func(v furnitures.Furni) {
							wg.Add(1)
							defer wg.Done()
							d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", v.Revision))
							d.SetFileName(v.Name)
							d.Download()
						}(v)
						time.Sleep(100 * time.Millisecond)
					}
				}
				wg.Wait()
			}
		} else {
			if len(args) > 1 {
				fmt.Println("Too many arguments")
			} else {
				if args[0] == "icons" {
					if furniName != "" {
						if furniRevision == "" {
							fmt.Println("--revision is required")
							return
						}
						cmd.MarkFlagRequired("revision")
						d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", furniRevision))
						d.SetFileName(fmt.Sprintf("%s.png", furniName))
						d.Download()
					} else {
						d.SetDomain("com")
						if Output != "" {
							d.SetOutput(Output)
						} else {
							d.SetOutput("resource/dcr/hof_furni/icons")
						}
						f := furnitures.NewFurnitures(*d, &wg, &mu)
						i := f.GetIcons()
						d.SetOther()

						for _, v := range i {
							exts := files.IsFileExists(d.GetOutput(), v.Name)
							if !exts {
								wg.Add(1)
								go func(v furnitures.Furni) {
									defer wg.Done()
									d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", v.Revision))
									d.SetFileName(v.Name)
									d.Download()
								}(v)
								time.Sleep(100 * time.Millisecond)
							}
						}
						wg.Wait()
					}
				} else {
					fmt.Printf("No such command : %s", args[0])
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(furniCmd)

	furniCmd.PersistentFlags().StringVarP(&furniName, "name", "n", "", "Furni name without file extension")
	furniCmd.PersistentFlags().StringVarP(&furniRevision, "revision", "r", "", "Furni revision")
}
