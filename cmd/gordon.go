/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
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
			d.SetOutput(Output)
			d.SetDomain(Domain)

			d.SetGordon()
			p := d.GetCurrentProduction()
			d.SetProduction(p)
			// d.SetPath("/")

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
				d.SetFileName("PlaceHolderPet.xml")
				d.Download()
			case "PlaceHolderWallItem":
				d.SetFileName("PlaceHolderWallItem.xml")
				d.Download()
			case "SelectionArrow":
				d.SetFileName("SelectionArrow.xml")
				d.Download()
			case "TileCursor":
				d.SetFileName("TileCursor.xml")
				d.Download()
			default:
				// TODO: Download all
				fmt.Print("Gordon")
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
