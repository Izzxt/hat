package clothes

import (
	"fmt"
	"log"
	"sync"

	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
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
	d.SetPath("/")
	d.SetFileName(fmt.Sprintf("%s.swf", fileName))

	ext, err := fs.Exists(fmt.Sprintf("%s%s.swf", d.GetOutput(), fileName))
	if err != nil {
		log.Fatal(err)
	}

	if ext {
		fmt.Println("skipped ", fileName)
	} else {
		d.Download()
	}
}
