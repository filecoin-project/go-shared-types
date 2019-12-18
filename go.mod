module github.com/filecoin-project/go-shared-types

go 1.13

require (
	github.com/filecoin-project/filecoin-ffi v0.0.0-20191210104338-2383ce072e95
	github.com/gopherjs/gopherjs v0.0.0-20190812055157-5d271430af9f // indirect
	github.com/ipfs/go-cid v0.0.4-0.20191112011718-79e75dffeb10 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.3
	github.com/ipfs/go-ipld-format v0.0.2 // indirect
	github.com/ipsn/go-secp256k1 v0.0.0-20180726113642-9d62b9f0bc52
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/minio/sha256-simd v0.1.1 // indirect
	github.com/multiformats/go-varint v0.0.1
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/warpfork/go-wish v0.0.0-20190328234359-8b3e70f8e830 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20191216205031-b047b6acb3c0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
