package input

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Nekroze/tuireframe/pkg/ir"
)

func LoadScreen(filename string) (ir.Screen, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadScreen(file), nil
}

func LoadMeta(filename string) (out MetaFile, err error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return out, err
	}
	return out, json.Unmarshal(raw, &out)
}
