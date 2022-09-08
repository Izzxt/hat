/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/Izzxt/hat/articles"
	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var (
	name  string
	wg    sync.WaitGroup
	mu    sync.Mutex
	data  []string
	after []string
)

// articlesCmd represents the articles command
var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan articles.Result)
		c := client.NewClient()
		d := downloader.NewDownloader(c)
		a := articles.NewArticles(&wg, *d, &mu)

		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetOther()
		d.SetPath("/web_images/habbo-web-articles")

		if name != "" {
			d.SetFileName(fmt.Sprintf("%s.png", name))
			d.Download()
		} else {
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

			defer wg.Wait()

			rg, _ := regexp.Compile(`([\w!@#$%^&*+-])*\.png`)
			rgmt := regexp.MustCompile(`_thumb.png`)

			for _, d := range data {
				s := rg.FindAllString(d, -1)
				for _, tr := range s {
					r := rgmt.ReplaceAllString(tr, ".png")
					after = append(after, r)
				}
			}

			bar := progressbar.Default(int64(len(after)))
			for _, v := range after {
				bar.Add(1)
				bar.Describe(v)
				d.SetFileName(v)
				d.Download()
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
