package promo

import (
	"regexp"
	"strings"
	"sync"

	"github.com/Izzxt/hat/downloader"
)

type Promo struct {
	downloader downloader.Downloader
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

func NewPromo(d downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Promo {
	return &Promo{
		downloader: d,
		wg:         wg,
		mu:         mu,
	}
}

func (h *Promo) GetAllImages() []string {
	var images []string
	keys := make(map[string]bool)
	dw := h.downloader

	domain := []string{
		"com.br", "com.tr", "com",
		"de", "es", "fi",
		"fr", "it", "nl",
	}

	for _, d := range domain {
		dw.SetDomain(d)
		dw.SetPath("/external_flash_texts/0")
		bytes, _ := dw.Fetch()
		match := matchRegex(string(bytes))
		for _, entry := range match {
			if _, value := keys[entry[1]]; !value {
				keys[entry[1]] = true
				c := strings.TrimSpace(entry[1])
				images = append(images, c)
			}
		}
	}

	return images

}

func matchRegex(findString string) [][]string {

	m := regexp.MustCompile(`(?m)web_promo_small\/([\w]*(.+?)png)`)

	return m.FindAllStringSubmatch(findString, -1)
}
