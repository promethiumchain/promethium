// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://ba72d977865eb7f4860b03630440dac4e9bd6d8323a12ac86b7f50e5891701b66ac1bc43f92802b5c1264b685253022397bb26b3c65b789383df19a09efac567@78.141.222.49:30303",
	"enode://b45a19a5d7dbd4a18a5241d128a8b06cd44d2387a28738ecdf06baf12408af989dca4aab46b2b2a3dedc35181ce46b4afbeeed4a968eaef5bb9206ca7b340831@209.250.245.131:30303",
	"enode://e3183145894240b731a744007e797010e83082d59db35d2f8ea80dd312000eeda2de278193d742f3bbaf8b7938f904368e35cd1c8f8659110f253ef85df19ebd@209.250.230.114:30303",
	"enode://8f305295090ec2783ce6b2d6410772b8d0bd19ed9907cf2d33bd9f0a70cea2e8aea9daa77eea7026ca2709efae0e0a5c98a72eb2a0d26312d19c03ddd40aa9b5@104.238.132.71:30303",
	"enode://e8726131a8b4c211f1d93db017d83328b08361e318dfe4d938107e7b414b123d622b201c74ac380fb08c469464ce44956aec8d65185949f6974eb4984e2faedd@95.179.214.242:30303",
	"enode://3a50f5a3b941cb934286ad783218d60409bea962cc8b2b6d52d1402fc550692801648dbeb08637acfed7b37a39ff93fc192bd1f39f5e2c336f92c805e9d49c1e@173.199.70.154:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://123d132d4da2a85db7d3cdf91b517eb5e8815008d64dc488a88b4fc57532c53f5d0efe0d812159cac668b0f805a2454f5803d59db5cabc3a9ed7c23baf06aea3@199.247.30.54:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://ba72d977865eb7f4860b03630440dac4e9bd6d8323a12ac86b7f50e5891701b66ac1bc43f92802b5c1264b685253022397bb26b3c65b789383df19a09efac567@78.141.222.49:30303",
	"enode://b45a19a5d7dbd4a18a5241d128a8b06cd44d2387a28738ecdf06baf12408af989dca4aab46b2b2a3dedc35181ce46b4afbeeed4a968eaef5bb9206ca7b340831@209.250.245.131:30303",
	"enode://e3183145894240b731a744007e797010e83082d59db35d2f8ea80dd312000eeda2de278193d742f3bbaf8b7938f904368e35cd1c8f8659110f253ef85df19ebd@209.250.230.114:30303",
	"enode://8f305295090ec2783ce6b2d6410772b8d0bd19ed9907cf2d33bd9f0a70cea2e8aea9daa77eea7026ca2709efae0e0a5c98a72eb2a0d26312d19c03ddd40aa9b5@104.238.132.71:30303",
	"enode://e8726131a8b4c211f1d93db017d83328b08361e318dfe4d938107e7b414b123d622b201c74ac380fb08c469464ce44956aec8d65185949f6974eb4984e2faedd@95.179.214.242:30303",
	"enode://3a50f5a3b941cb934286ad783218d60409bea962cc8b2b6d52d1402fc550692801648dbeb08637acfed7b37a39ff93fc192bd1f39f5e2c336f92c805e9d49c1e@173.199.70.154:30303",
}
