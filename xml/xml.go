package xml

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
)

type FigureMap struct {
	XMLName xml.Name    `xml:"map"`
	Lib     []FigureLib `xml:"lib"`
}

type FigureLib struct {
	XMLName  xml.Name     `xml:"lib"`
	Part     []FigurePart `xml:"part"`
	Id       string       `xml:"id,attr"`
	Revision string       `xml:"revision,attr"`
}

type FigurePart struct {
	XMLName xml.Name `xml:"part"`
	Id      string   `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
}

type FurniData struct {
	XMLName      xml.Name     `xml:"furnidata"`
	RoomItemType RoomItemType `xml:"roomitemtypes"`
	WallItemType WallItemType `xml:"wallitemtypes"`
}

type RoomItemType struct {
	XMLName   xml.Name    `xml:"roomitemtypes"`
	FurniType []FurniType `xml:"furnitype"`
}

type WallItemType struct {
	XMLName   xml.Name    `xml:"wallitemtypes"`
	FurniType []FurniType `xml:"furnitype"`
}

type FurniType struct {
	XMLName   xml.Name `xml:"furnitype"`
	Id        string   `xml:"id,attr"`
	ClassName string   `xml:"classname,attr"`
}

type EffectMap struct {
	XMLName xml.Name     `xml:"map"`
	Effect  []EffectAttr `xml:"effect"`
}

type EffectAttr struct {
	XMLName  xml.Name `xml:"effect"`
	Id       string   `xml:"id,attr"`
	Lib      string   `xml:"lib,attr"`
	Type     string   `xml:"type,attr"`
	Revision string   `xml:"revision,attr"`
}

func Parse(obj interface{}, xmlReader io.Reader) {

	byteValue, err := ioutil.ReadAll(xmlReader)

	if err != nil {
		log.Fatalf("Error while trying to read xml : %s", err)
	}

	xml.Unmarshal(byteValue, &obj)
}
