package cmd

import (
	"fmt"
	"strings"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"

	"github.com/spf13/cobra"
)

// gamedataCmd represents the gamedata command
type Gamedata struct {
	file string
	path string
}

var (
	isXml       bool
	isTxt       bool
	types       string
	gamedataCmd = &cobra.Command{
		Use:   "gamedata",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {

			c := client.NewClient()
			d := downloader.NewDownloader(c)
			d.SetDomain(Domain)

			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/gamedata")
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
					d.SetTxt()
					d.SetFileName("productdata.txt")
				}

				d.Download()

			case "figuredata":
				d.SetPath("/figuredata/0")
				d.SetFileName("figuredata.xml")
				d.Download()

			case "external_variables":
				d.SetPath("/external_variables/0")
				d.SetFileName("external_variables.txt")
				d.Download()

			case "external_flash_texts":
				d.SetPath("/external_flash_texts/0")
				d.SetFileName("external_flash_texts.txt")
				d.Download()

			case "external_override_variables":
				d.SetPath("/external_override_variables/0")
				d.SetFileName("external_override_variables.txt")
				d.Download()

			case "external_override_flash_texts":
				d.SetPath("/external_override_flash_texts/0")
				d.SetFileName("external_override_flash_texts.txt")
				d.Download()

			case "figuremap":
				d.SetGordon()
				d.SetFileName("figuremap.xml")
				current := d.GetCurrentProduction()
				d.SetProduction(current)

				if Prod != "" {
					d.SetProduction(Prod)
				}

				d.Download()

			case "effectmap":
				d.SetGordon()
				d.SetFileName("effectmap.xml")
				current := d.GetCurrentProduction()
				d.SetProduction(current)

				if Prod != "" {
					d.SetProduction(Prod)
				}

				d.Download()

			default:

				gamedata := []Gamedata{
					{
						file: "furnidata.txt",
						path: "furnidata",
					},
					{
						file: "furnidata.json",
						path: "furnidata_json",
					},
					{
						file: "furnidata.xml",
						path: "furnidata_xml",
					},
					{
						file: "productdata.txt",
						path: "productdata",
					},
					{
						file: "productdata.json",
						path: "productdata_json",
					},
					{
						file: "productdata.xml",
						path: "productdata_xml",
					},
					{
						file: "figuredata.xml",
						path: "figuredata",
					},
					{
						file: "external_variables.txt",
						path: "external_variables",
					},
					{
						file: "external_flash_texts.txt",
						path: "external_flash_texts",
					},
					{
						file: "external_override_variables.txt",
						path: "external_override_variables",
					},
					{
						file: "external_override_flash_texts.txt",
						path: "external_override_flash_texts",
					},
					{
						file: "figuremap.xml",
						path: "figuremap",
					},
					{
						file: "effectmap.xml",
						path: "effectmap",
					},
				}

				current := d.GetCurrentProduction()
				d.SetProduction(current)

				for _, g := range gamedata {
					exts := fs.IsFileExists(d.GetOutput(), g.file)
					if !exts {
						if g.path == "external_override_variables" || g.path == "external_override_flash_texts" {
							d.SetOutput("resource/gamedata/override")
						}

						if g.path != "figuremap" && g.path != "effectmap" {
							d.SetOutput("resource/gamedata")
							d.SetPath(fmt.Sprintf("/%s/0", g.path))
							d.SetFileName(g.file)
							d.Download()
						} else {
							d.SetOutput("resource/gamedata")
							d.SetPath("")
							d.SetGordon()
							d.SetFileName(g.file)
							d.Download()
						}
					}
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(gamedataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	gamedataCmd.PersistentFlags().BoolVarP(&isXml, "xml", "", false, "Output fetch to xml")
	gamedataCmd.PersistentFlags().BoolVarP(&isTxt, "txt", "", false, "Output fetch to xml")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	gamedataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
