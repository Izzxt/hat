package furnitures

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/xml"
)

type Furnitures struct {
	downloader downloader.Downloader
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

type Furni struct {
	Name     string
	Revision string
}

func NewFurnitures(d downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Furnitures {
	return &Furnitures{
		downloader: d,
		wg:         wg,
		mu:         mu,
	}
}

func (f *Furnitures) GetIcons() []Furni {
	var furni xml.FurniData
	var icons []Furni

	d := f.downloader
	d.SetDomain(d.GetDomain())
	d.SetPath("/furnidata_xml/0")

	byte, _ := d.Fetch()

	xml.Parse(&furni, strings.NewReader(string(byte)))

	keys := make(map[string]bool)
	for _, entry := range furni.RoomItemType.FurniType {
		r := regexp.MustCompile(`\*`)
		i := fmt.Sprintf("%s_icon.png", r.ReplaceAllString(entry.ClassName, "_"))
		if _, value := keys[i]; !value {
			keys[i] = true
			icons = append(icons, Furni{
				Name:     i,
				Revision: entry.Revision,
			})
		}
	}

	return icons
}

func (f *Furnitures) GetFurnis() []Furni {
	var furni xml.FurniData
	var furnis []Furni

	d := f.downloader
	d.SetDomain(d.GetDomain())
	d.SetPath("/furnidata_xml/0")

	byte, _ := d.Fetch()

	xml.Parse(&furni, strings.NewReader(string(byte)))

	for _, entry := range furni.RoomItemType.FurniType {
		keys := make(map[string]bool)
		r := regexp.MustCompile(`\*`)
		i := fmt.Sprintf("%s.swf", r.Split(entry.ClassName, -1)[0])
		if _, value := keys[i]; !value {
			keys[i] = true
			furnis = append(furnis, Furni{
				Name:     i,
				Revision: entry.Revision,
			})
		}
	}

	return furnis
}
