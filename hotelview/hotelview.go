package hotelview

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

func NewHotelView(d downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Promo {
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

	dw.SetDomain(dw.GetDomain())
	dw.SetPath("/external_variables/0")
	byte, _ := dw.Fetch()
	match := matchRegex(string(byte))
	for _, entry := range match {
		if _, value := keys[entry[3]]; !value {
			keys[entry[3]] = true
			c := strings.TrimSpace(entry[3])
			images = append(images, c)
		}
	}

	return images

}

func matchRegex(findString string) [][]string {

	m := regexp.MustCompile(`(?m)(landing\.view\.background.+=).+(reception\/)(.+)`)

	return m.FindAllStringSubmatch(findString, -1)
}
