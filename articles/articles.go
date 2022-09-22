package articles

import (
	"fmt"
	"regexp"
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
	d := ar.downloader
	d.SetDomain("com")
	d.SetOther()
	d.SetPath("/habbo-web-news/en/production")

	total := 0
	i := 1
	run := true
	attempt := 0

	for run {
		d.SetFileName(fmt.Sprintf("all_%d.html", i))

		o, c := d.Fetch()

		match, _ := regexp.MatchString("<section>", string(o))

		if match {
			total += 1
		}

		if c == 404 {
			attempt++
		}

		if attempt > 5 {
			run = false
		}

		i++
	}

	return total
}
