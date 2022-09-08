package downloader

import (
	"encoding/json"
	"testing"

	"github.com/Izzxt/hat/client"

	"github.com/stretchr/testify/assert"
)

var pathUrl string = "/furnidata_json/0"
var domain string = "com"

func TestFetch(t *testing.T) {
	c := client.NewClient()

	d := NewDownloader(c)
	d.SetPath(pathUrl)
	d.SetDomain(domain)

	bodyBytes, _ := d.Fetch()

	var furni Furni

	if err := json.Unmarshal(bodyBytes, &furni); err != nil {
		t.Errorf("Error while trying to unmarshal json: %s", err)
	}

	if furni.RoomItemTypes.FurniType == nil {
		t.Error("expected furnitype object, got nil")
	}
}

func TestSetPath(t *testing.T) {
	c := client.NewClient()

	assert := assert.New(t)

	d := NewDownloader(c)
	d.SetPath(pathUrl)

	expected := "/furnidata_json/0"

	assert.Equal(d.pathUrl, expected)
}
