package badges

import (
	"fmt"
	"log"
	"regexp"
	"sync"

	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
)

type Badges struct {
	downloader downloader.Downloader
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

func NewBadges(downloader downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Badges {
	return &Badges{
		downloader: downloader,
		wg:         wg,
		mu:         mu,
	}
}

func (b *Badges) GetAllCode() []string {
	var code []string
	dw := b.downloader

	domain := []string{
		"com.br", "com.tr", "com",
		"de", "es", "fi",
		"fr", "it", "nl",
	}

	for _, d := range domain {
		dw.SetDomain(d)
		dw.SetPath("/external_flash_texts/0")
		fmt.Println(dw.GetUrl())
		byte, _ := dw.Fetch()
		match := matchRegex(string(byte))
		for _, v := range match {
			exts, err := fs.Exists(fmt.Sprintf("%s%s%s", dw.GetOutput(), v[2], ".gif"))
			if err != nil {
				log.Fatal(err)
			}

			if !exts {
				code = append(code, v[2])
			}
		}
	}

	return code
}

func matchRegex(findString string) [][]string {

	m := regexp.MustCompile(`(?m)badge_(desc|name)_*(.[\w+\-$#!@#$%^&*()\s.|]*)=`)

	return m.FindAllStringSubmatch(findString, -1)
}
