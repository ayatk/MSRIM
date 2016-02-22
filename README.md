# MinecraftRecipeMaker
[![GitHub release](http://img.shields.io/github/release/S--Minecraft/MinecraftSourceRecipeImageMaker.svg?style=flat-square)][release]

[release]: https://github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/releases

Minecraftのテクスチャからレシピの画像を生成します

## 使い方
1. パッケージファイルを[ダウンロード][release]して解凍する
1. 解凍してできたフォルダに`minecraft.jar`もしくは`バージョン番号.jar`をコピペする
1. 名前はなんでもいいので設定ファイルを書いてバイナリファイルと同じ階層に保存する
1. ターミナルかコマンドプロンプトでバイナリを実行する
1. `output`フォルダに画像ができてる!!✌('ω'✌ )三✌('ω')✌三( ✌'ω')✌

MODのテクスチャを使う場合は[こちら]()

## 設定について
設定ファイルはtomlの形式で記述し、エンコーディングはUTF-8にしてください
``` toml
# レシピのサンプル

[[recipes]]
  # crafting_table, furnace, brewing_standが標準で使えます
  type = "crafting_table"
  [[recipes.recipe]]
    # 生成される画像の名前です
    name = "test"

    # どのアイテムをどのように置くかを指定します
    shape = [
      # 作業台の場合、shapeの1つ目から9つ目は各入力スロットで10こ目は出力スロットです
      # すべてint型で記述してください
      # 各スロットに0を入力すれば空のスロットになります
      1, 0, 1,
      0, 2, 0,
      1, 0, 1,
      3
    ]

    # shapeで指定したスロットにどの画像が入るのかを記述します
    [[recipes.recipe.images]]
      # slotに指定した数字の位置にimageで指定した画像が入ります
      slot = 1
      image = "gold_nugget"
      # アイテム数を指定します
      # アイテムが1個の場合は省略できます
      num = 5

    # shapeで指定した数だけ画像を入力してください
    [[recipes.recipe.images]]
      slot = 2
      # block/dartのように書くとブロック形式に変形して描画されます
      image = "block/dirt"

    [[recipes.recipe.images]]
      slot = 3
      image = "apple_golden"
      num = 2
```

### かまど
``` toml
shape = [
  # shapeの1つ目が材料、2つ目が燃料、3つ目が出力です
  1, 2, 3
]
```

### 醸造台
``` toml
shape = [
  # 1つ目が材料、右から順に2,3,4つ目が出力です
  1, 2, 3, 4
]
```

## Contribution
1. Fork (https://github.com/S--Minecraft/MinecraftSourceRecipeImageMaker/fork)
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `go fmt`
1. Create a new Pull Request

## License
copyright (c) 2016 S--Minecraft
