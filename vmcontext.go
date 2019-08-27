package types

import (
	"github.com/filecoin-project/go-amt-ipld"
	"github.com/filecoin-project/go-lotus/chain/actors/aerrors"
	"github.com/filecoin-project/go-lotus/chain/address"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-hamt-ipld"
)

type Storage interface {
	Put(interface{}) (cid.Cid, aerrors.ActorError)
	Get(cid.Cid, interface{}) aerrors.ActorError

	GetHead() cid.Cid

	// Commit sets the new head of the actors state as long as the current
	// state matches 'oldh'
	Commit(oldh cid.Cid, newh cid.Cid) aerrors.ActorError
}

type StateTree interface {
	SetActor(addr address.Address, act *Actor) error
	GetActor(addr address.Address) (*Actor, error)
}

type VMContext interface {
	Message() *Message
	Origin() address.Address
	Ipld() *hamt.CborIpldStore
	Send(to address.Address, method uint64, value BigInt, params []byte) ([]byte, aerrors.ActorError)
	BlockHeight() uint64
	GasUsed() BigInt
	Storage() Storage
	StateTree() (StateTree, aerrors.ActorError)
	VerifySignature(sig *Signature, from address.Address, data []byte) aerrors.ActorError
	ChargeGas(uint64) aerrors.ActorError
}

type storageWrapper struct {
	s Storage
}

func (sw *storageWrapper) Put(i interface{}) (cid.Cid, error) {
	c, err := sw.s.Put(i)
	if err != nil {
		return cid.Undef, err
	}

	return c, nil
}

func (sw *storageWrapper) Get(c cid.Cid, out interface{}) error {
	if err := sw.s.Get(c, out); err != nil {
		return err
	}

	return nil
}

func WrapStorage(s Storage) amt.Blocks {
	return &storageWrapper{s}
}