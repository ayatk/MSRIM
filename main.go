package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
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
		}
	}
	// jarとjsonがない場合の処理
	if jar == "" {
		log.Fatal("Not Found: minecraft jar file")
	}

	// jarを解凍してtmpディレクトリに展開
	UnJar(jar, "tmp")

	// 出力先のディレクトリ、なければ作る
	out := "output"
	if !Exists(out) {
		if err := os.Mkdir(out, 0777); err != nil {
			log.Fatal(err)
		}
	}

	// 使い終わったtmpディレクトリを削除
	if err := os.RemoveAll("tmp"); err != nil {
		log.Fatal(err)
	}
}
