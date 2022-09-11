package xml

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFigureMapParse(t *testing.T) {
	assert := assert.New(t)
	figureData := `
	<map>
		<lib id="hh_people_pool" revision="50491">
			<part id="s03" type="ss"/>
			<part id="s04" type="ss"/>
		</lib>
		<lib id="hh_human_acc_eye" revision="24465">
			<part id="1" type="ea"/>
			<part id="2" type="ea"/>
		</lib>
	</map>`

	var figureMap FigureMap

	Parse(&figureMap, strings.NewReader(figureData))

	assert.Equal("s04", figureMap.Lib[0].Part[1].Id)
	assert.Equal("hh_human_acc_eye", figureMap.Lib[1].Id)
}

func TestFurnidataRoomItemType(t *testing.T) {
	assert := assert.New(t)
	data := `
	<furnidata>
		<roomitemtypes>
			<furnitype id="13" classname="shelves_norja">
				<revision>61856</revision>
				<category>shelf</category>
			</furnitype>
			<furnitype id="14" classname="shelves_polyfon">
				<revision>48082</revision>
				<category>shelf</category>
			</furnitype>
		</roomitemtypes>
	</furnidata>
	`

	var furniData FurniData

	Parse(&furniData, strings.NewReader(data))

	assert.Equal("13", furniData.RoomItemType.FurniType[0].Id)
	assert.Equal("shelves_polyfon", furniData.RoomItemType.FurniType[1].ClassName)
}

func TestFurnidataWallItemType(t *testing.T) {
	assert := assert.New(t)
	data := `
	<furnidata>
		<wallitemtypes>
			<furnitype id="4748" classname="wall_ph">
				<revision>68275</revision>
				<category>other</category>
				<name>wall_ph name</name>
				<description>wall_ph desc</description>
			</furnitype>
			<furnitype id="4749" classname="nft_h22_sharkaquarium">
				<revision>68283</revision>
				<category>wall_decoration</category>
				<name>nft_h22_sharkaquarium name</name>
				<description>nft_h22_sharkaquarium desc</description>
			</furnitype>
		</wallitemtypes>
	</furnidata>
	`

	var furniData FurniData

	Parse(&furniData, strings.NewReader(data))

	assert.Equal("wall_ph", furniData.WallItemType.FurniType[0].ClassName)
	assert.Equal("4749", furniData.WallItemType.FurniType[1].Id)
}

func TestEffectMap(t *testing.T) {
	assert := assert.New(t)
	data := `
	<map>
		<effect id="211" lib="CyberKongz" type="fx" revision="67951"/>
		<effect id="212" lib="Metakey" type="fx" revision="67970"/>
	</map>
	`

	var effectMap EffectMap

	Parse(&effectMap, strings.NewReader(data))

	assert.Equal("CyberKongz", effectMap.Effect[0].Lib)
	assert.Equal("212", effectMap.Effect[1].Id)
}
