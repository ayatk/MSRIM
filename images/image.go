package images

import (
	"image"
	"image/png"
	"os"
)

// パス指定されたpng画像をimage.Image型にデコード
func Decode(fp string) (image.Image, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}

// image.Image型を指定された先にpng形式で書き出し
// ファイル名がかぶっていた時はエラーを返す
// NOTE: 今後jpegやgifに対応させるかも
func Encode(fp string, i image.Image) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, i)
}

