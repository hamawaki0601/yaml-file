# yaml-file
[YAML](https://github.com/go-yaml/yaml)を使って、ファイルから読み込めるものです。  
また、デフォルト値の設定も可能にしました。

## Installation

```
$ go get github.com/hamawaki0601/yaml-file
```

## Usage

```
type server struct {
	Host string
	Port int `default:"80"`
}

func main() {
	s := server{}
	err := Load("server.yml", &s)
}
```