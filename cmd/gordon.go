/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/spf13/cobra"
)

// gordonCmd represents the gordon command
var (
	gordonType string
	gordonCmd  = &cobra.Command{
		Use:   "gordon",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			c := client.NewClient()
			d := downloader.NewDownloader(c)
			d.SetDomain(Domain)
			d.SetGordon()

			if Prod == "" {
				Prod = d.GetCurrentProduction()
			}

			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput(fmt.Sprintf("resource/gordon/%s", Prod))
			}

			d.SetProduction(Prod)

			switch gordonType {
			case "HabboConfig":
				d.SetFileName("config_habbo.xml")
				d.Download()
			case "HabboAvatarActions":
				d.SetFileName("HabboAvatarActions.xml")
				d.Download()
			case "HabboRoomContent":
				d.SetFileName("HabboRoomContent.swf")
				d.Download()
			case "PlaceHolderFurniture":
				d.SetFileName("PlaceHolderFurniture.swf")
				d.Download()
			case "PlaceHolderPet":
				d.SetFileName("PlaceHolderPet.swf")
				d.Download()
			case "PlaceHolderWallItem":
				d.SetFileName("PlaceHolderWallItem.swf")
				d.Download()
			case "SelectionArrow":
				d.SetFileName("SelectionArrow.swf")
				d.Download()
			case "TileCursor":
				d.SetFileName("TileCursor.swf")
				d.Download()
			default:

				gordon := []string{
					"config_habbo.xml", "HabboAvatarActions.xml", "HabboRoomContent.swf",
					"PlaceHolderFurniture.swf", "PlaceHolderPet.swf", "PlaceHolderWallItem.swf",
					"SelectionArrow.swf", "TileCursor.swf",
				}

				for _, v := range gordon {
					exts := fs.IsFileExists(d.GetOutput(), v)
					if !exts {
						go func(v string) {
							d.SetFileName(v)
							d.Download()
						}(v)
						time.Sleep(100 * time.Millisecond)
					}
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(gordonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	gordonCmd.PersistentFlags().StringVarP(&gordonType, "type", "t", "", "Gordon type")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gordonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
