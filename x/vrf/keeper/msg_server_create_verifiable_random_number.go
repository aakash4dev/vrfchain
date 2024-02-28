package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"

	"github.com/aakash4dev/vrfchain/x/vrf/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVerifiableRandomNumber(goCtx context.Context, msg *types.MsgCreateVerifiableRandomNumber) (*types.MsgCreateVerifiableRandomNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	shaseedByte := msg.Shaseed
	publickeyByte := msg.Publickey
	rByte := msg.R
	sByte := msg.S
	maxRange := msg.MaxRange

	// shaseedByte: byte to byte32
	var shaSeed [32]byte
	copy(shaSeed[:], shaseedByte)

	// unmarshal publickeyByte
	var publicKey *ecdsa.PublicKey
	pubKeyError := json.Unmarshal(publickeyByte, publicKey)
	if pubKeyError != nil {
		return nil, status.Error(codes.InvalidArgument, pubKeyError.Error())
	}

	// unmarshal rByte to r *big.Int
	var r *big.Int
	rError := json.Unmarshal(rByte, r)
	if rError != nil {
		return nil, status.Error(codes.InvalidArgument, rError.Error())
	}

	// unmarshal sByte to s *big.Int
	var s *big.Int
	sError := json.Unmarshal(sByte, s)
	if sError != nil {
		return nil, status.Error(codes.InvalidArgument, sError.Error())
	}

	// generate psudo random number
	_, psudoRandomNumber, err := generateSudoRandomNumber(shaSeed, publicKey.X, publicKey.Y, r, s, maxRange)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// verify signature
	isValid := ecdsa.Verify(publicKey, shaseedByte, r, s)
	if !isValid {
		return nil, status.Error(codes.Internal, "Signature verification failed.")
	}

	// store the number
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix("/psudorandomnumber")) //types.PostKey))
	// psudoRandomNumber to bytes
	psudoRandomNumberBytes := psudoRandomNumber.Bytes()
	store.Set([]byte("podID"), psudoRandomNumberBytes)

	return &types.MsgCreateVerifiableRandomNumberResponse{}, nil
}

func generateSudoRandomNumber(shaSeedByte [32]byte, x, y, r, s *big.Int, maxRange int64) (hashInt [32]byte, psudoRandomNumber *big.Int, err error) {
	xBytes := x.Bytes()
	yBytes := y.Bytes()
	rBytes := r.Bytes()
	sBytes := s.Bytes()

	// Concatenate all byte slices into one
	allBytes := append(shaSeedByte[:], xBytes...)
	allBytes = append(allBytes, yBytes...)
	allBytes = append(allBytes, rBytes...)
	allBytes = append(allBytes, sBytes...)

	// Hash the combined byte slice using SHA-256
	hashInt = sha256.Sum256(allBytes)
	hash := new(big.Int).SetBytes(hashInt[:])
	psudoRandomNumber = new(big.Int).Mod(hash, big.NewInt(maxRange+1))
	return hashInt, psudoRandomNumber, nil
}
