/*
Copyright © 2022 Izzat
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
	"github.com/Izzxt/hat/files"
	"github.com/Izzxt/hat/version"
	"github.com/Izzxt/hat/xml"
	"github.com/spf13/cobra"
)

var effectName string

var effectsCmd = &cobra.Command{
	Use:   "effects",
	Short: "Download habbo effects",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex
		var effect xml.EffectMap

		keys := make(map[string]bool)
		c := client.NewClient()
		version.StartupMessage(c)
		d := downloader.NewDownloader(c)
		d.SetOutput(Output)
		d.SetDomain(Domain)

		if Prod == "" {
			Prod = d.GetCurrentProduction()
		}

		d.SetProduction(Prod)

		e := effects.NewEffects(*d, &wg, &mu)
		if effectName != "" {
			d.SetFileName(fmt.Sprintf("%s.swf", effectName))
			d.Download()
		} else {
			d.SetGordon()
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput(fmt.Sprintf("resource/gordon/%s", d.GetProduction()))
			}

			eBtye := e.GetAllEffectLib()
			d.SetPath("")
			xml.Parse(&effect, strings.NewReader(string(eBtye)))
			for _, entry := range effect.Effect {
				if _, value := keys[entry.Lib]; !value {
					keys[entry.Lib] = true
					exts := files.IsFileExists(d.GetOutput(), fmt.Sprintf("%s.swf", entry.Lib))
					if !exts {
						wg.Add(1)
						go func(v xml.EffectAttr) {
							defer wg.Done()
							d.SetFileName(fmt.Sprintf("%s.swf", v.Lib))
							d.Download()
						}(entry)
						time.Sleep(100 * time.Millisecond)
					}
				}
			}
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(effectsCmd)

	effectsCmd.PersistentFlags().StringVarP(&effectName, "name", "n", "", "Effects name without file extension")
}
