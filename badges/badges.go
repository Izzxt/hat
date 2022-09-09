package badges

import (
	"regexp"
	"sync"

	"github.com/Izzxt/hat/downloader"
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

func (b *Badges) GetAllCode() {
	dw := b.downloader

	domain := []string{
		"com.br", "com.tr", "com",
		"de", "es", "fi",
		"fr", "it", "nl",
	}

	for _, v := range domain {
		defer b.wg.Done()
		go func(v string) {
			dw.SetDomain(v)
			dw.SetPath("/gamedata/external_flash_texts/0")
		}(v)
	}
}

func matchRegex(findString string) [][]string {

	m := regexp.MustCompile(`(?m)badge_(desc|name)_*(.[\w+\-$#!@#$%^&*()].*)=`)

	return m.FindAllStringSubmatch(findString, -1)
}
