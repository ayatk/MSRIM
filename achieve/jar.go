package achieve

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// jarの解凍、内部の処理はzipの解凍と全く同じ
func UnJar(src, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	var rc io.ReadCloser
	for _, f := range reader.File {
		rc, err = f.Open()
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		if err != nil {
			return err
		}

		s := dest + "/" + f.Name
		d, _ := filepath.Split(s)

		if _, err = os.Stat(d); err != nil {
			os.MkdirAll(d, 0755)
		}

		if err = ioutil.WriteFile(s, buf.Bytes(), 0755); err != nil {
			return err
		}
		rc.Close()
	}

	return nil
}
