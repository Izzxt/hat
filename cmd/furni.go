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
	"github.com/Izzxt/hat/furnitures"
	"github.com/spf13/cobra"
)

var furniName string
var furniRevision string

// furniCmd represents the furni command
var furniCmd = &cobra.Command{
	Use:   "furni",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
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
				if d.GetOutput() != "" {
					d.SetOutput(d.GetOutput())
				} else {
					d.SetOutput("resource/dcr/hof_furni")
				}
				f := furnitures.NewFurnitures(*d, &wg, &mu)
				i := f.GetFurnis()
				d.SetOther()

				for _, v := range i {
					wg.Add(1)
					go func(v furnitures.Furni) {
						defer wg.Done()
						d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", v.Revision))
						d.SetFileName(v.Name)
						d.Download()
						fmt.Println(fmt.Sprintf("Download %s/%s", v.Revision, v.Name))
					}(v)
					time.Sleep(150 * time.Millisecond)
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
						f := furnitures.NewFurnitures(*d, &wg, &mu)
						i := f.GetIcons()
						d.SetOther()

						for _, v := range i {
							wg.Add(1)
							go func(v furnitures.Furni) {
								defer wg.Done()
								d.SetPath(fmt.Sprintf("/dcr/hof_furni/%s", v.Revision))
								d.SetFileName(v.Name)
								d.Download()
								fmt.Println(fmt.Sprintf("Download %s/%s", v.Revision, v.Name))
							}(v)
							time.Sleep(150 * time.Millisecond)
						}
						wg.Wait()
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(furniCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	furniCmd.PersistentFlags().StringVarP(&furniName, "name", "n", "", "A help for foo")
	furniCmd.PersistentFlags().StringVarP(&furniRevision, "revision", "r", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// furniCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
