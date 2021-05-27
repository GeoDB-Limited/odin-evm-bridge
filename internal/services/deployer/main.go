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
	contractAddress, tx, err := s.deployContract()
	if err != nil {
		return errors.Wrap(err, "failed to deploy contract")
	}

	s.logger.WithFields(logrus.Fields{
		"tx_hash":          tx.Hash(),
		"contract_address": contractAddress.Hex(),
	}).Info("Contract deployed")

	return nil
}

// deployContract deploys a bridge contract.
func (s *Service) deployContract() (*common.Address, *types.Transaction, error) {
	chainId, err := s.ethereum.NetworkID(s.context)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get chain id")
	}

	address, pk := s.config.EthereumSigner()

	txOpts, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create transaction options")
	}

	nonce, err := s.ethereum.PendingNonceAt(s.context, address)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get a nonce")
	}

	txOpts.Nonce = new(big.Int).SetUint64(nonce)

	ethConfig := s.config.EthereumConfig()
	txOpts.GasLimit = ethConfig.GasLimit.Uint64()
	txOpts.GasPrice = ethConfig.GasPrice

	contractAddress, tx, _, err := generated.DeployCacheBridge(txOpts, s.ethereum, s.config.Validators())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to deploy contract")
	}

	return &contractAddress, tx, nil
}
