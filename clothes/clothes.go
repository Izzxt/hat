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

func (c *Clothes) Download(fileName string, out string) {
	defer c.wg.Done()

	d := c.downloader

	d.SetDomain("com")
	d.SetGordon()
	d.SetProduction(out)
	if d.GetOutput() != "" {
		d.SetOutput(d.GetOutput())
	} else {
		d.SetOutput(fmt.Sprintf("gordon/%s", d.GetProduction()))
	}
	d.SetPath("/")
	d.SetFileName(fmt.Sprintf("%s.swf", fileName))

	d.Download()
}
