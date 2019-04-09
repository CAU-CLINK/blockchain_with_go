package conf

import "os"

var confFilePath = os.Getenv("GOPATH") + "/src/github.com/CAU-CLINK/blockchain_with_go/conf/config.yaml"

type Configuration struct {
	GenesisBlockPath string
	BlockchainPath   string
	UTXOSetPath      string
}

// GOPATH 설정유무 확인, conf package 호출 시 최초 실행되는 func
func init() {
	if os.Getenv("GOPATH") == "" {
		panic("Need to set GOPATH")
	}
}

// TODO : implements me w/ test case
// It is accessible as cli or filePath.
// Cli overwrite configuration file.
func NewConfiguration(confFilePath string) Configuration {
	return Configuration{}
}

func ConfigPath(path string) {
	confFilePath = path
}
