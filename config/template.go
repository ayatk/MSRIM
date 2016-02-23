package config

import (
	"image"

	"github.com/BurntSushi/toml"
)

//
type TomlConfig struct {
	Type     string         `toml:"type"`
	Image    string         `toml:"image"`
	Crop     [][]int        `toml:"crop"`
	Slot     [][]int        `toml:"slot"`
	Progress []TomlProgress `toml:"progress"`
}

type TomlProgress struct {
	Crop  [][]int `toml:"crop"`
	Paste []int   `toml:"paste"`
}

type Config struct {
	// GUIタイプを設定します
	Type string
	// 使用する画像を指定します
	Image string
	// GUIに使われる画像の切り取り範囲を指定します
	Crop image.Rectangle
	// どのスロットがどの位置にあるのかを指定します
	Slot []image.Point
	// 進捗バーを切り貼りする座標を指定します
	Progress []Progress
}

type Progress struct {
	// 進捗バーを切り取り範囲を指定します
	Crop image.Rectangle
	// 切り取られた進捗バーを貼り付ける座標です
	// 切り取られた画像の左上の角がくる座標を指定してください
	Paste image.Point
}

func LoadConfig(fp string) (*Config, error) {
	var tml TomlConfig
	_, err := toml.DecodeFile(fp, &tml)
	if err != nil {
		return nil, err
	}

	var slot []image.Point
	for i := range tml.Slot {
		slot = append(slot, itop(tml.Slot[i]))
	}

	var prgs []Progress
	var prg Progress
	for i := range tml.Progress {
		prg = Progress{
			Crop:  itor(tml.Progress[i].Crop),
			Paste: itop(tml.Progress[i].Paste),
		}
		prgs = append(prgs, prg)
	}

	// Configを初期化
	conf := Config{
		Type:     tml.Type,
		Image:    tml.Image,
		Crop:     itor(tml.Crop),
		Slot:     slot,
		Progress: prgs,
	}

	return &conf, nil
}

// Int型配列からimage.Point型に変換
func itop(p []int) image.Point {
	return image.Point{p[0], p[1]}
}

// Int型2次元配列からimage.Rectangle型に変換
func itor(p [][]int) image.Rectangle {
	return image.Rect(p[0][0], p[0][1], p[1][0], p[1][1])
}
