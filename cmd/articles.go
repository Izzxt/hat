/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/Izzxt/hat/articles"
	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/spf13/cobra"
)

var (
	name  string
	data  []string
	after []string
)

// todo: use attempt instead
// articlesCmd represents the articles command
var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		var mu sync.Mutex

		keys := make(map[string]bool)
		ch := make(chan articles.Result)
		c := client.NewClient()
		d := downloader.NewDownloader(c)
		a := articles.NewArticles(&wg, *d, &mu)

		d.SetDomain(Domain)
		d.SetOther()
		d.SetPath("/web_images/habbo-web-articles")

		if name != "" {
			d.SetFileName(fmt.Sprintf("%s.png", name))
			d.Download()
		} else {
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/habbo-web-articles")
			}
			fmt.Println("Initializing...")
			p := a.GetMaxPage()

			for i := 1; i <= p-2; i++ {
				wg.Add(1)
				go a.FetchAll(fmt.Sprintf("all_%d.html", i), ch)
			}

			for i := 1; i <= p-2; i++ {
				select {
				case msg := <-ch:
					data = append(data, string(msg.Response))
				}
			}

			rg, _ := regexp.Compile(`([\w!@#$%^&*+-])*\.png`)
			rgmt := regexp.MustCompile(`_thumb.png`)

			for _, a := range data {
				s := rg.FindAllString(a, -1)
				for _, entry := range s {
					if _, value := keys[entry]; !value {
						keys[entry] = true
						r := rgmt.ReplaceAllString(entry, ".png")
						after = append(after, r)
					}
				}
			}

			defer wg.Wait()

			for _, v := range after {
				wg.Add(1)
				exts := fs.IsFileExists(d.GetOutput(), v)
				if !exts {
					go func(v string) {
						defer wg.Done()
						d.SetFileName(v)
						d.Download()
					}(v)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(articlesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	articlesCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "download single image by name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// articlesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
