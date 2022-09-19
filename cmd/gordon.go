/*
Copyright © 2022 Izzat
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/files"
	"github.com/Izzxt/hat/version"
	"github.com/spf13/cobra"
)

var (
	gordonType string
	gordonCmd  = &cobra.Command{
		Use:   "gordon [FLAGS]",
		Short: "Download habbo gordon assets",
		Long: `
Available Types:
  - HabboConfig
  - HabboAvatarActions
  - HabboRoomContent
  - PlaceHolderFurniture
  - PlaceHolderPet
  - PlaceHolderWallItem
    `,
		Run: func(cmd *cobra.Command, args []string) {
			c := client.NewClient()
			version.StartupMessage(c)
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
					exts := files.IsFileExists(d.GetOutput(), v)
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

	gordonCmd.PersistentFlags().StringVarP(&gordonType, "type", "t", "", "Gordon type")
}
