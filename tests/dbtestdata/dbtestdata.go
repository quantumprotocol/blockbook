package dbtestdata

import (
	"blockbook/bchain"
	"encoding/hex"
	"math/big"

	"github.com/golang/glog"
)

const (
	TxidB1T1 = "00b2c06055e5e90e9c82bd4181fde310104391a7fa4f289b1704e5d90caa3840"
	TxidB1T2 = "effd9ef509383d536b1c8af5bf434c8efbf521a4f2befd4022bbd68694b4ac75"
	TxidB2T1 = "7c3be24063f268aaa1ed81b64776798f56088757641a34fb156c4f51ed2e9d25"
	TxidB2T2 = "3d90d15ed026dc45e19ffb52875ed18fa9e8012ad123d7f7212176e2b0ebdb71"
	TxidB2T3 = "05e2e48aeabdd9b75def7b48d756ba304713c2aba7b522bf9dbc893fc4231b07"
	TxidB2T4 = "fdd824a780cbb718eeb766eb05d83fdefc793a27082cd5e67f856d69798cf7db"

	Addr1 = "mfcWp7DB6NuaZsExybTTXpVgWz559Np4Ti"  // 76a914010d39800f86122416e28f485029acf77507169288ac
	Addr2 = "mtGXQvBowMkBpnhLckhxhbwYK44Gs9eEtz"  // 76a9148bdf0aa3c567aa5975c2e61321b8bebbe7293df688ac
	Addr3 = "mv9uLThosiEnGRbVPS7Vhyw6VssbVRsiAw"  // 76a914a08eae93007f22668ab5e4a9c83c8cd1c325e3e088ac
	Addr4 = "2Mz1CYoppGGsLNUGF2YDhTif6J661JitALS" // a9144a21db08fb6882cb152e1ff06780a430740f770487
	Addr5 = "2NEVv9LJmAnY99W1pFoc5UJjVdypBqdnvu1" // a914e921fc4912a315078f370d959f2c4f7b6d2a683c87
	Addr6 = "mzB8cYrfRwFRFAGTDzV8LkUQy5BQicxGhX"  // 76a914ccaaaf374e1b06cb83118453d102587b4273d09588ac
	Addr7 = "mtR97eM2HPWVM6c8FGLGcukgaHHQv7THoL"  // 76a9148d802c045445df49613f6a70ddd2e48526f3701f88ac
	Addr8 = "mwwoKQE5Lb1G4picHSHDQKg8jw424PF9SC"  // 76a914b434eb0c1a3b7a02e8a29cc616e791ef1e0bf51f88ac
	Addr9 = "mmJx9Y8ayz9h14yd9fgCW1bUKoEpkBAquP"  // 76a9143f8ba3fda3ba7b69f5818086e12223c6dd25e3c888ac
	AddrA = "mzVznVsCHkVHX9UN8WPFASWUUHtxnNn4Jj"  // 76a914d03c0d863d189b23b061a95ad32940b65837609f88ac
)

var (
	SatZero   = big.NewInt(0)
	SatB1T1A1 = big.NewInt(100000000)
	SatB1T1A2 = big.NewInt(12345)
	SatB1T2A3 = big.NewInt(1234567890123)
	SatB1T2A4 = big.NewInt(1)
	SatB1T2A5 = big.NewInt(9876)
	SatB2T1A6 = big.NewInt(317283951061)
	SatB2T1A7 = big.NewInt(917283951061)
	SatB2T2A8 = big.NewInt(118641975500)
	SatB2T2A9 = big.NewInt(198641975500)
	SatB2T3A5 = big.NewInt(9000)
	SatB2T4AA = big.NewInt(1360030331)
)

func AddressToPubKeyHex(addr string, parser bchain.BlockChainParser) string {
	if addr == "" {
		return ""
	}
	b, err := parser.GetAddrDescFromAddress(addr)
	if err != nil {
		glog.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func GetTestBitcoinTypeBlock1(parser bchain.BlockChainParser) *bchain.Block {
	return &bchain.Block{
		BlockHeader: bchain.BlockHeader{
			Height:        225493,
			Hash:          "0000000076fbbed90fd75b0e18856aa35baa984e9c9d444cf746ad85e94e2997",
			Size:          1234567,
			Time:          1534858021,
			Confirmations: 2,
		},
		Txs: []bchain.Tx{
			bchain.Tx{
				Txid: TxidB1T1,
				Vin:  []bchain.Vin{},
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr1, parser),
						},
						ValueSat: *SatB1T1A1,
					},
					bchain.Vout{
						N: 1,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr2, parser),
						},
						ValueSat: *SatB1T1A2,
					},
				},
				Blocktime:     22549300000,
				Time:          22549300000,
				Confirmations: 2,
			},
			bchain.Tx{
				Txid: TxidB1T2,
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr3, parser),
						},
						ValueSat: *SatB1T2A3,
					},
					bchain.Vout{
						N: 1,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr4, parser),
						},
						ValueSat: *SatB1T2A4,
					},
					bchain.Vout{
						N: 2,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr5, parser),
						},
						ValueSat: *SatB1T2A5,
					},
				},
				Blocktime:     22549300001,
				Time:          22549300001,
				Confirmations: 2,
			},
		},
	}
}

func GetTestBitcoinTypeBlock2(parser bchain.BlockChainParser) *bchain.Block {
	return &bchain.Block{
		BlockHeader: bchain.BlockHeader{
			Height:        225494,
			Hash:          "00000000eb0443fd7dc4a1ed5c686a8e995057805f9a161d9a5a77a95e72b7b6",
			Size:          2345678,
			Time:          1534859123,
			Confirmations: 1,
		},
		Txs: []bchain.Tx{
			bchain.Tx{
				Txid: TxidB2T1,
				Vin: []bchain.Vin{
					// addr3
					bchain.Vin{
						Txid: TxidB1T2,
						Vout: 0,
					},
					// addr2
					bchain.Vin{
						Txid: TxidB1T1,
						Vout: 1,
					},
				},
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr6, parser),
						},
						ValueSat: *SatB2T1A6,
					},
					bchain.Vout{
						N: 1,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr7, parser),
						},
						ValueSat: *SatB2T1A7,
					},
				},
				Blocktime:     22549400000,
				Time:          22549400000,
				Confirmations: 1,
			},
			bchain.Tx{
				Txid: TxidB2T2,
				Vin: []bchain.Vin{
					// spending an output in the same block - addr6
					bchain.Vin{
						Txid: TxidB2T1,
						Vout: 0,
					},
					// spending an output in the previous block - addr4
					bchain.Vin{
						Txid: TxidB1T2,
						Vout: 1,
					},
				},
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr8, parser),
						},
						ValueSat: *SatB2T2A8,
					},
					bchain.Vout{
						N: 1,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr9, parser),
						},
						ValueSat: *SatB2T2A9,
					},
				},
				Blocktime:     22549400001,
				Time:          22549400001,
				Confirmations: 1,
			},
			// transaction from the same address in the previous block
			bchain.Tx{
				Txid: TxidB2T3,
				Vin: []bchain.Vin{
					// addr5
					bchain.Vin{
						Txid: TxidB1T2,
						Vout: 2,
					},
				},
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(Addr5, parser),
						},
						ValueSat: *SatB2T3A5,
					},
				},
				Blocktime:     22549400002,
				Time:          22549400002,
				Confirmations: 1,
			},
			// mining transaction
			bchain.Tx{
				Txid: TxidB2T4,
				Vin: []bchain.Vin{
					bchain.Vin{
						Coinbase: "03bf1e1504aede765b726567696f6e312f50726f6a65637420425443506f6f6c2f01000001bf7e000000000000",
					},
				},
				Vout: []bchain.Vout{
					bchain.Vout{
						N: 0,
						ScriptPubKey: bchain.ScriptPubKey{
							Hex: AddressToPubKeyHex(AddrA, parser),
						},
						ValueSat: *SatB2T4AA,
					},
					bchain.Vout{
						N:            1,
						ScriptPubKey: bchain.ScriptPubKey{},
						ValueSat:     *SatZero,
					},
				},
				Blocktime:     22549400003,
				Time:          22549400003,
				Confirmations: 1,
			},
		},
	}
}
