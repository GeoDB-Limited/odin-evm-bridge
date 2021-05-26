package config

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
)

// Config defines an interface of global service configurations.
type Config interface {
	Logger() *logrus.Logger
	EthereumClient() *ethclient.Client
	EthereumConfig() EthereumConfig
	EthereumSigner() (common.Address, *ecdsa.PrivateKey)
}

// Config defines global service configurations.
type config struct {
	Log      string         `yaml:"log"`
	Ethereum EthereumConfig `yaml:"ethereum"`
}

type EthereumConfig struct {
	Endpoint   string   `yaml:"endpoint"`
	GasLimit   *big.Int `yaml:"gas_limit"`
	GasPrice   *big.Int `yaml:"gas_price"`
	PrivateKey string   `yaml:"private_key"`
}

// NewConfig returns global service configurations.
func NewConfig(path string) Config {
	cfg := config{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to read config: %s", path)))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("failed to unmarshal config: %s", path)))
	}

	return &cfg
}

// Logger returns new configured logger.
func (c *config) Logger() *logrus.Logger {
	level, err := logrus.ParseLevel(c.Log)
	if err != nil {
		panic(errors.Wrapf(err, "failed to parse logging level %s", c.Log))
	}

	logger := logrus.New()
	logger.SetLevel(level)

	return logger
}

// EthereumClient returns ethereum client.
func (c *config) EthereumClient() *ethclient.Client {
	ethClient, err := ethclient.Dial(c.Ethereum.Endpoint)
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", c.Ethereum.Endpoint))
	}

	return ethClient
}

// EthereumSigner returns address and private key of ethereum signer.
func (c *config) EthereumSigner() (common.Address, *ecdsa.PrivateKey) {
	pk, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "error casting private key to ECDSA"))
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(errors.Wrap(err, "error casting public key to ECDSA"))
	}

	accAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return accAddress, pk
}

// EthereumConfig returns ethereum config.
func (c *config) EthereumConfig() EthereumConfig {
	return c.Ethereum
}
