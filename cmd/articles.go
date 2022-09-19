/*
Copyright © 2022 Izzat
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
	"github.com/Izzxt/hat/files"
	"github.com/spf13/cobra"
)

var (
	name  string
	data  []string
	after []string
)

// todo: use attempt instead
var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "Download habbo web articles",
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
				exts := files.IsFileExists(d.GetOutput(), v)
				if !exts {
					wg.Add(1)
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

	articlesCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Picture name without file extension")
}
