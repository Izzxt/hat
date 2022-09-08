package clothes

import (
	"fmt"
	"sync"

	"github.com/Izzxt/hat/downloader"
)

type Clothes struct {
	mu         *sync.Mutex
	wg         *sync.WaitGroup
	downloader downloader.Downloader
}

type Result struct {
	Response []byte
	Code     int
}

func NewClothes(d downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Clothes {
	return &Clothes{
		downloader: d,
		wg:         wg,
		mu:         mu,
	}
}

func (c *Clothes) Download(fileName string) {
	defer c.wg.Done()

	d := c.downloader

	d.SetDomain("com")
	d.SetDomain("com")
	d.SetGordon()
	p := d.GetCurrentProduction()
	d.SetProduction(p)
	d.SetPath("/")
	d.SetFileName(fmt.Sprintf("%s.swf", fileName))

	d.Download()
}
