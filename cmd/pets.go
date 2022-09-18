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
	"github.com/spf13/cobra"
)

var petName string

var petsCmd = &cobra.Command{
	Use:   "pets",
	Short: "Download habbo pets",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetGordon()

		if Prod == "" {
			Prod = d.GetCurrentProduction()
		}

		d.SetProduction(Prod)
		// d.SetPath("/")
		if petName != "" {
			d.SetFileName(fmt.Sprintf("%s.swf", petName))
			d.Download()
		} else {
			if Output != "" {
				d.SetOutput(Output)
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
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(petsCmd)

	petsCmd.PersistentFlags().StringVarP(&petName, "name", "n", "", "Pet name without file extension")
}
