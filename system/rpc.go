package system

import (
	"bytes"
	
	"github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/scale"
)

func AccountNonce(client substrate.Client, accountPubKey []byte) (uint64, error) {
	m, err := client.MetaData(true)
	if err != nil {
		return 0, err
	}
	key, err := substrate.NewStorageKey(*m,"System", "AccountNonce", accountPubKey)
	if err != nil {
		return 0, err
	}

	s := substrate.NewStateRPC(client)
	data, err := s.Storage(key, nil)
	if err != nil {
		return 0, err
	}

	buf := bytes.NewBuffer(data)
	tempDec := scale.NewDecoder(buf)
	var nonce uint64
	err = tempDec.Decode(&nonce)
	if err != nil {
		return 0, err
	}
	
	return nonce, nil
}

func BlockHash(client substrate.Client, blockNumber uint64) (substrate.Hash, error) {
	// TODO
	return nil, nil
}