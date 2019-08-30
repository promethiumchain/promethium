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
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

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
