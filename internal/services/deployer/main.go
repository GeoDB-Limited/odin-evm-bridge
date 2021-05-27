package deployer

import (
	"context"
	"github.com/GeoDB-Limited/odin-evm-bridge/internal/config"
	"github.com/GeoDB-Limited/odin-evm-bridge/system-contracts/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
)

// Service defines a service that deploys a bridge contract.
type Service struct {
	config   config.Config
	context  context.Context
	logger   *logrus.Logger
	ethereum *ethclient.Client
}

// New creates a service that deploys a bridge contract.
func New(ctx context.Context, cfg config.Config) *Service {
	return &Service{
		config:   cfg,
		context:  ctx,
		logger:   cfg.Logger(),
		ethereum: cfg.EthereumClient(),
	}
}

// Run performs deploying a bridge smart contract.
func (s *Service) Run() (err error) {
	bridgeAddress, tx, err := s.deployBridgeContract()
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	s.logger.WithFields(logrus.Fields{
		"tx_hash":          tx.Hash(),
		"contract_address": bridgeAddress.Hex(),
	}).Info("Bridge contract deployed")

	if err := s.saveBridgeAddress(*bridgeAddress); err != nil {
		return errors.Wrap(err, "failed to save bridge contract address")
	}

	datasetAddress, tx, err := s.deployDatasetContract(*bridgeAddress, 2)
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	s.logger.WithFields(logrus.Fields{
		"tx_hash":          tx.Hash(),
		"contract_address": datasetAddress.Hex(),
	}).Info("Dataset contract deployed")

	return nil
}

// deployBridgeContract deploys a bridge contract.
func (s *Service) deployBridgeContract() (*common.Address, *types.Transaction, error) {
	txOpts, err := s.getTxOpts()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get tx opts")
	}

	contractAddress, tx, _, err := generated.DeployCacheBridge(txOpts, s.ethereum, s.config.Validators())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to deploy contract")
	}

	return &contractAddress, tx, nil
}

// deployDatasetContract deploys a dataset contract.
func (s *Service) deployDatasetContract(bridgeAddr common.Address, oracleScriptId uint64) (*common.Address, *types.Transaction, error) {
	txOpts, err := s.getTxOpts()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get tx opts")
	}

	contractAddress, tx, _, err := generated.DeployOdinDataset(txOpts, s.ethereum, bridgeAddr, oracleScriptId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to deploy contract")
	}

	return &contractAddress, tx, nil
}

// getTxOpts returns tx opts
func (s *Service) getTxOpts() (*bind.TransactOpts, error) {
	chainId, err := s.ethereum.NetworkID(s.context)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	address, pk := s.config.EthereumSigner()
	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction options")
	}

	nonce, err := s.ethereum.PendingNonceAt(s.context, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a nonce")
	}

	txOpts.Nonce = new(big.Int).SetUint64(nonce)

	ethConfig := s.config.EthereumConfig()
	txOpts.GasLimit = ethConfig.GasLimit.Uint64()
	txOpts.GasPrice = ethConfig.GasPrice

	return txOpts, nil
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (s *Service) saveBridgeAddress(address common.Address) error {
	f, err := os.OpenFile(s.config.BridgeAddressStorage(), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}

	if _, err := f.WriteString(address.String()); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}
