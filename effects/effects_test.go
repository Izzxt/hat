package effects

import (
	"strings"
	"testing"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/xml"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEffectLib(t *testing.T) {
	assert := assert.New(t)
	c := client.NewClient()
	d := downloader.NewDownloader(c)

	d.SetProduction("PRODUCTION-202209021210-342804575")

	e := NewEffects(*d, nil, nil)

	eBtye := e.GetAllEffectLib()

	var effect xml.EffectMap

	xml.Parse(&effect, strings.NewReader(string(eBtye)))

	assert.Equal("Dance1", effect.Effect[0].Lib)
	assert.Equal("Dance2", effect.Effect[1].Lib)
}
