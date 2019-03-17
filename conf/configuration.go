package conf

type Configuration struct {
	GenesisBlockPath string
	BlockchainPath   string
	UTXOSetPath      string
}

// It is accessible as cli or filePath.
// Cli overwrite configuration file.
func NewConfiguration(confFilePath string) Configuration {
	return Configuration{}
}
