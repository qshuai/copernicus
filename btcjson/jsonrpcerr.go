// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

// Standard JSON-RPC 2.0 errors.
var (
	ErrRPCInvalidRequest = &RPCError{
		Code:    -32600,
		Message: "Invalid request",
	}
	ErrRPCMethodNotFound = &RPCError{
		Code:    -32601,
		Message: "Method not found",
	}
	ErrRPCInvalidParams = &RPCError{
		Code:    -32602,
		Message: "Invalid parameters",
	}
	ErrRPCInternal = &RPCError{
		Code:    -32603,
		Message: "Internal error",
	}
	ErrRPCParse = &RPCError{
		Code:    -32700,
		Message: "Parse error",
	}
)

// General application defined JSON errors.
const (
	ErrRPCMisc                RPCErrorCode = -1
	ErrRPCForbiddenBySafeMode RPCErrorCode = -2
	ErrRPCType                RPCErrorCode = -3
	ErrRPCInvalidAddressOrKey RPCErrorCode = -5
	ErrRPCOutOfMemory         RPCErrorCode = -7
	ErrRPCInvalidParameter    RPCErrorCode = -8
	ErrRPCDatabase            RPCErrorCode = -20
	ErrRPCDeserialization     RPCErrorCode = -22
	ErrRPCVerify              RPCErrorCode = -25
)

// Peer-to-peer client errors.
const (
	ErrRPCClientNotConnected      RPCErrorCode = -9
	ErrRPCClientInInitialDownload RPCErrorCode = -10
	ErrRPCClientNodeNotAdded      RPCErrorCode = -24
)

// Wallet JSON errors
const (
	ErrRPCWallet                    RPCErrorCode = -4
	ErrRPCWalletInsufficientFunds   RPCErrorCode = -6
	ErrRPCWalletInvalidAccountName  RPCErrorCode = -11
	ErrRPCWalletKeypoolRanOut       RPCErrorCode = -12
	ErrRPCWalletUnlockNeeded        RPCErrorCode = -13
	ErrRPCWalletPassphraseIncorrect RPCErrorCode = -14
	ErrRPCWalletWrongEncState       RPCErrorCode = -15
	ErrRPCWalletEncryptionFailed    RPCErrorCode = -16
	ErrRPCWalletAlreadyUnlocked     RPCErrorCode = -17
)

// Specific Errors related to commands.  These are the ones a user of the RPC
// server are most likely to see.  Generally, the codes should match one of the
// more general errors above.
const (
	ErrRPCBlockNotFound     RPCErrorCode = -5
	ErrRPCBlockCount        RPCErrorCode = -5
	ErrRPCBestBlockHash     RPCErrorCode = -5
	ErrRPCDifficulty        RPCErrorCode = -5
	ErrRPCOutOfRange        RPCErrorCode = -1
	ErrRPCNoTxInfo          RPCErrorCode = -5
	ErrRPCNoNewestBlockInfo RPCErrorCode = -5
	ErrRPCInvalidTxVout     RPCErrorCode = -5
	ErrRPCRawTxString       RPCErrorCode = -32602
	ErrRPCDecodeHexString   RPCErrorCode = -22
)

// Errors that are specific to btcd.
const (
	ErrRPCNoWallet      RPCErrorCode = -1
	ErrRPCUnimplemented RPCErrorCode = -1
)

const (
	// custom
	ErrUnDefined        RPCErrorCode = 404
	ErrInvalidParameter RPCErrorCode = -30

	// BCH v0.16.0
	// Standard JSON-RPC 2.0 errors
	// RPC_INVALID_REQUEST is internally mapped to HTTP_BAD_REQUEST (400).
	// It should not be used for application-layer errors.
	RpcInvalidRequest = -32600
	// RPC_METHOD_NOT_FOUND is internally mapped to HTTP_NOT_FOUND (404).
	// It should not be used for application-layer errors.
	RpcMethodNotFound  = -32601
	RPC_INVALID_PARAMS = -32602
	// RPC_INTERNAL_ERROR should only be used for genuine errors in bitcoind
	// (for exampled datadir corruption).
	RpcInternalError = -32603
	RpcParseError    = -32700

	// General application defined errors
	// std::exception thrown in command handling
	RpcMiscError = -1
	// Server is in safe mode, and command is not allowed in safe mode
	RpcForbiddenBySafeMode = -2
	// Unexpected type was passed as parameter
	RpcTypeError = -3
	// Invalid address or key
	RpcInvalidAddressOrKey = -5
	// Ran out of memory during operation
	RpcOutOfMemory = -7
	// Invalid, missing or duplicate parameter
	RpcInvalidParameter = -8
	// Database error
	RpcDatabaseError = -20
	// Error parsing or validating structure in raw format
	RpcDeserializationError = -22
	// General error during transaction or block submission
	RpcVerifyError = -25
	// Transaction or block was rejected by network rules
	RpcVerifyRejected = -26
	// Transaction already in chain
	RpcVerifyAlreadyInChain = -27
	// Client still warming up
	RpcInWarmup = -28

	// Aliases for backward compatibility
	RpcTransactionError          = RpcVerifyError
	RpcTransactionRejected       = RpcVerifyRejected
	RpcTransactionAlreadyInChain = RpcVerifyAlreadyInChain

	// P2P client errors
	// Bitcoin is not connected
	RpcClientNotConnected = -9
	// Still downloading initial blocks
	RpcClientInInitialDownload = -10
	// Node is already added
	RpcClientNodeAlreadyAdded = -23
	// Node has not been added before
	RpcClientNodeNotAdded = -24
	// Node to disconnect not found in connected nodes
	RpcClientNodeNotConnected = -29
	// Invalid IP/Subnet
	RpcClientInvalidIpOrSubnet = -30
	// No valid connection manager instance found
	RpcClientP2pDisabled = -31

	// Wallet errors
	// Unspecified problem with wallet (key not found etc.)
	RpcWalletError = -4
	// Not enough funds in wallet or account
	RpcWalletInsufficientFunds = -6
	// Invalid account name
	RpcWalletInvalidAccountName = -11
	// Keypool ran out, call keypoolrefill first
	RpcWalletKeypoolRanOut = -12
	// Enter the wallet passphrase with walletpassphrase first
	RpcWalletUnlockNeeded = -13
	// The wallet passphrase entered was incorrect
	RpcWalletPassphraseIncorrect = -14
	// Command given in wrong wallet encryption state (encrypting an encrypted wallet etc.)
	RpcWalletWrongEncState = -15
	// Failed to encrypt the wallet
	RpcWalletEncryptionFailed = -16
	// Wallet is already unlocked
	RpcWalletAlreadyUnlocked = -17
)
