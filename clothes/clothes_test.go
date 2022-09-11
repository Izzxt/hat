package clothes

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/xml"
	"github.com/stretchr/testify/assert"
)

var wg sync.WaitGroup
var mu sync.Mutex

func TestFetchFigure(t *testing.T) {
	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetDomain("com")
	d.SetGordon()
	p := d.GetCurrentProduction()
	d.SetProduction(p)
	d.SetPath("/")

	d.SetFileName("/figuremapv2.xml")

	byte, _ := d.Fetch()

	assert.Regexp(regexp.MustCompile("lib"), string(byte))
}

func TestParseFigureFromFetch(t *testing.T) {
	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetDomain("com")
	d.SetGordon()
	p := d.GetCurrentProduction()
	d.SetProduction(p)
	d.SetPath("/")

	d.SetFileName("/figuremapv2.xml")

	byte, _ := d.Fetch()

	var figure xml.FigureMap

	xml.Parse(&figure, strings.NewReader(string(byte)))

	for _, v := range figure.Lib {
		if v.Id == "hh_people_pool" {
			assert.Equal("hh_people_pool", v.Id)
		}
	}
}

func TestDownload(t *testing.T) {
	data := []string{
		"hh_people_pool",
		"hh_human_shirt",
	}

	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)
	d.SetOutput("")
	cl := NewClothes(*d, &wg, &mu)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go cl.Download(fmt.Sprintf("/%s", data[i]), "")
	}

	assert.FileExists(fmt.Sprintf("%s.swf", "hh_people_pool"))

	wg.Wait()
}
