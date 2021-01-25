package statik

import (
	"github.com/rakyll/statik/fs"
	"io/ioutil"
)

func ReadFile(filename string) ([]byte, error) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	openApiFile, err := statikFS.Open(filename)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(openApiFile)
	if err != nil {
		panic(err)
	}
	return data, err
}
