package config

import "github.com/BurntSushi/toml"

type RecipeConfig struct {
	Type []RecipeType `toml:"recipes"`
}

type RecipeType struct {
	// 使用するGUIタイプ
	Type    string   `toml:"type"`
	Recipes []Recipe `toml:"recipe"`
}

// 生成される画像の名前,配置,使用する画像の配列
type Recipe struct {
	// 生成される画像に付けられる名前
	Name   string `toml:"name"`
	Shape  []int  `toml:"shape"`
	Images []Img  `toml:"images"`
}

// Shapeに対応する画像と個数
type Img struct {
	Slot   uint   `toml:"slot"`
	Image  string `toml:"image"`
	number uint   `toml:"num"`
}

// レシピが記述されたtomlを読み込んでRecipeConfigに変換
func LoadRecipe(fp string) RecipeConfig {
	var conf RecipeConfig

	_, err := toml.DecodeFile(fp, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
