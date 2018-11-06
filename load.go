package file

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const loadError = "%s is not YAML or incorrect format"

// Load は 指定したファイル(filename)を読み込み、 outに値を展開します.
// ファイルがない場合や、yamlではない場合にエラーを返します.
func Load(filename string, out interface{}) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	SetDefaultValue(out)

	err = yaml.Unmarshal(buf, out)
	if err != nil {
		return fmt.Errorf(loadError, filename)
	}
	return nil
}
