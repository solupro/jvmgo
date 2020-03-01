package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath: absPath}
}


func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(this.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {

		if f.Name != className {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return nil, nil, err
		}
		defer rc.Close()

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, nil, err
		}

		return data, this, nil
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (this *ZipEntry) String() string {
	return this.absPath
}



