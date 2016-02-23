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
//
// NOTE: 今後jpegやgifに対応させるかも
func Encode(fp string, i image.Image) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, i)
}

// (x0, y0)から(x1, y1)まで切り取り
func Crop(i *image.Image, x0, y0, x1, y1 int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, x1-x0, y1-y0))
	size := (*img).Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			(*img).Set(x, y, (*i).At(x1+x, y1+y))
		}
	}
	return img
}

// Cropの座標指定をintのかわりにimage.Rectangleを用いる
func CropRect(i *image.Image, r image.Rectangle) image.Image {
	return Crop(i, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
}
