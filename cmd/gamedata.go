package cmd

import (
	"fmt"
	"strings"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gamedataCmd represents the gamedata command
var (
	production  string
	isXml       bool
	isTxt       bool
	types       string
	gamedataCmd = &cobra.Command{
		Use:   "gamedata",
		Short: "A brief description of your command",
		Long:  `A`,
		Run: func(cmd *cobra.Command, args []string) {
			c := client.NewClient()
			d := downloader.NewDownloader(c)

			if Domain != "" {
				d.SetDomain(Domain)
			}

			if Output != "" {
				d.SetOutput(Output)
			}

			switch strings.Join(args, ",") {
			case "furnidata":
				d.SetPath("/furnidata_json/0")
				d.SetFileName("furnidata.json")

				if isXml {
					d.SetXml()
					d.SetFileName("furnidata.xml")
				}

				if isTxt {
					d.SetXml()
					d.SetFileName("furnidata.txt")
				}

				d.Download()

			case "productdata":
				d.SetPath("/productdata_json/0")
				d.SetFileName("productdata.json")

				if isXml {
					d.SetXml()
					d.SetFileName("productdata.xml")
				}

				if isTxt {
					d.SetXml()
					d.SetFileName("furnidata.txt")
				}

				d.Download()

			case "figuredata":
				d.SetPath("/figuredata/0")
				d.SetFileName("figuredata.xml")
				d.Download()

			case "external_variables":
				d.SetPath("/productdata_json/0")
				d.SetFileName("productdata.json")
				d.Download()

			case "external_flash_texts":
				d.SetPath("/productdata_json/0")
				d.SetFileName("productdata.json")

				d.Download()

			case "external_override_variables":
				d.SetPath("/productdata_json/0")
				d.SetFileName("productdata.json")

				d.Download()

			case "external_override_flash_texts":
				d.SetPath("/productdata_json/0")
				d.SetFileName("productdata.json")

				d.Download()

			case "figuremap":
				d.SetGordon()
				d.SetPath("/figuremap.xml")
				d.SetFileName("figuremap.xml")
				current := d.GetCurrentProduction()
				d.SetProduction(current)

				if production != "" {
					d.SetProduction(production)
				}

				d.Download()

			case "effectmap":
				d.SetGordon()
				d.SetPath("/effectmap.xml")
				d.SetFileName("effectmap.xml")
				current := d.GetCurrentProduction()
				d.SetProduction(current)

				if production != "" {
					d.SetProduction(production)
				}

				d.Download()

			default:
				fmt.Print(viper.Get("test"))
				fmt.Print(CfgFile)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(gamedataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	gamedataCmd.PersistentFlags().StringVarP(&production, "production", "", "", "habbo production")
	gamedataCmd.PersistentFlags().BoolVarP(&isXml, "xml", "", false, "Output fetch to xml")
	gamedataCmd.PersistentFlags().BoolVarP(&isTxt, "txt", "", false, "Output fetch to xml")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	gamedataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
