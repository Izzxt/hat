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
	"github.com/spf13/cobra"
)

var petName string

// petsCmd represents the pets command
var petsCmd = &cobra.Command{
	Use:   "pets",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetGordon()

		if Prod != "" {
			Prod = d.GetCurrentProduction()
		}

		d.SetProduction(Prod)
		d.SetPath("/")
		if petName != "" {
			d.SetFileName(fmt.Sprintf("%s.swf", petName))
			d.Download()
		} else {
			if d.GetOutput() != "" {
				d.SetOutput(d.GetOutput())
			} else {
				d.SetOutput(fmt.Sprintf("resource/gordon/%s", d.GetProduction()))
			}

			pets := []string{
				"bear.swf", "bearbaby.swf", "bunnydepressed.swf", "bunnyeaster.swf", "bunnyevil.swf",
				"bunnylove.swf", "cat.swf", "chicken.swf", "cow.swf", "croco.swf",
				"demonmonkey.swf", "dog.swf", "dragon.swf", "fools.swf", "frog.swf",
				"gnome.swf", "haloompa.swf", "horse.swf", "kittenbaby.swf", "lion.swf",
				"monkey.swf", "monster.swf", "monsterplant.swf", "pig.swf", "pigeonevil.swf",
				"pigeongood.swf", "pigletbaby.swf", "pterosaur.swf", "puppybaby.swf", "rhino.swf",
				"spider.swf", "terrier.swf", "terrierbaby.swf", "turtle.swf", "velociraptor.swf",
			}

			for _, v := range pets {
				wg.Add(1)
				go func(v string) {
					defer wg.Done()
					d.SetFileName(v)
					d.Download()
					fmt.Println(fmt.Sprintf("Download %s", v))
				}(v)
				time.Sleep(100 * time.Millisecond)
			}
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(petsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	petsCmd.PersistentFlags().StringVarP(&petName, "name", "n", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// petsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
