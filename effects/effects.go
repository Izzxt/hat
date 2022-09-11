package effects

import (
	"sync"

	"github.com/Izzxt/hat/downloader"
)

type Effect struct {
	downloader downloader.Downloader
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

func NewEffects(d downloader.Downloader, wg *sync.WaitGroup, mu *sync.Mutex) *Effect {
	return &Effect{
		downloader: d,
		wg:         wg,
		mu:         mu,
	}
}

func (e *Effect) GetAllEffectLib() []byte {
	d := e.downloader

	d.SetDomain("com")
	d.SetGordon()
	d.SetProduction(d.GetProduction())
	d.SetPath("/")
	d.SetFileName("effectmap.xml")

	byte, _ := d.Fetch()

	return byte
}
