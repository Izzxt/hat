/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/effects"
	"github.com/Izzxt/hat/xml"
	"github.com/spf13/cobra"
)

var effectName string

// effectsCmd represents the effects command
var effectsCmd = &cobra.Command{
	Use:   "effects",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetOutput(Output)
		d.SetDomain(Domain)

		if Prod != "" {
			Prod = d.GetCurrentProduction()
		}

		d.SetProduction(Prod)

		e := effects.NewEffects(*d, &wg, &mu)
		if effectName != "" {
			d.SetFileName(fmt.Sprintf("%s.swf", effectName))
			d.Download()
		} else {
			if d.GetOutput() != "" {
				d.SetOutput(d.GetOutput())
			} else {
				d.SetOutput(fmt.Sprintf("gordon/%s", d.GetProduction()))
			}
			var effect xml.EffectMap

			eBtye := e.GetAllEffectLib()
			xml.Parse(&effect, strings.NewReader(string(eBtye)))
			for _, v := range effect.Effect {
				wg.Add(1)
				go func(v xml.EffectAttr) {
					defer wg.Done()
					d.SetFileName(fmt.Sprintf("%s.swf", v.Lib))
					d.Download()
				}(v)
				time.Sleep(100 * time.Millisecond)
			}
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(effectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	effectsCmd.PersistentFlags().StringVarP(&effectName, "name", "n", "", "Effect name to download single")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// effectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
