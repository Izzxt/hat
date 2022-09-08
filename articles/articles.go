package articles

import (
	"fmt"
	"sync"

	"github.com/Izzxt/hat/downloader"
)

type Article struct {
	wg         *sync.WaitGroup
	mu         *sync.Mutex
	downloader downloader.Downloader
}

type Result struct {
	Response []byte
	Code     int
}

func NewArticles(wg *sync.WaitGroup, d downloader.Downloader, mu *sync.Mutex) *Article {
	return &Article{
		wg:         wg,
		downloader: d,
		mu:         mu,
	}
}

func (ar *Article) FetchAll(name string, ch chan Result) {
	defer ar.wg.Done()

	d := ar.downloader

	d.SetDomain("com")
	d.SetOther()
	d.SetPath("/habbo-web-news/en/production")
	d.SetFileName(name)

	bytes, c := d.Fetch()

	r := new(Result)
	r.Response = bytes
	r.Code = c
	ch <- *r
}

func (ar *Article) GetMaxPage() int {
	i := 1
	d := ar.downloader
	d.SetDomain("com")
	d.SetOther()
	d.SetPath("/habbo-web-news/en/production")

	for {
		d.SetFileName(fmt.Sprintf("all_%d.html", i))

		_, c := d.Fetch()

		i++

		if c == 404 {
			break
		}
	}

	return i
}
