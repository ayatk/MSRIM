package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ayatk/MinecraftRecipeMaker/archive"
	"github.com/ayatk/MinecraftRecipeMaker/config"
	"github.com/ayatk/MinecraftRecipeMaker/utils"
)

// Minecraftのjarファイルとレシピのファイル名
var jar, cfg string

func main() {
	// カレントディレクトリのファイルを読み込み
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		switch path.Ext(file.Name()) {
		case ".jar":
			jar = file.Name()
		case ".toml":
			cfg = file.Name()
		}
	}
	// jarとレシピファイルがない場合の処理
	if jar == "" {
		log.Fatal("Not Found: minecraft jar file")
	}
	if cfg == "" {
		log.Fatal("Not Found: Recipe file")
	}

	// レシピ読み込み
	recipe := config.LoadRecipe(cfg)
	// jarを解凍してtmpディレクトリに展開
	archive.UnJar(jar, "tmp/"+jar)

	// 出力先のディレクトリ、なければ作る
	out := "output"
	if !utils.Exists(out) {
		if err := os.Mkdir(out, 0777); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(recipe)

	// 使い終わったtmpディレクトリを削除
	if err := os.RemoveAll("tmp"); err != nil {
		log.Fatal(err)
	}
}
