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
	"enode://3a50f5a3b941cb934286ad783218d60409bea962cc8b2b6d52d1402fc550692801648dbeb08637acfed7b37a39ff93fc192bd1f39f5e2c336f92c805e9d49c1e@173.199.70.154:30303",
	"enode://bb9e4b9e3bec143ebd802d37561b568d560912a4faa6c10e6e714c83af827829bd77de25510b590b812096a0408ed427d3f9e098502578b94c72455b3c13eb3d@45.76.45.30:30303",
	"enode://85470f54eb5ac86d367dfa1377a148f87f9c4a2e6743cf569e65fc6b31ee1ab6b4c8b68e10d58acb36ecd2b987a9613d87866852b08a237d58869028deed19d8@209.250.251.109:30303",
	"enode://ac79ab5972ff0b7cad9e5f81944eff68947b51c6ea99f0e7847c8487234cf34bd46118164e22b1cafa269640ec1f04238abc02aab0df95352adaa1142278c3e3@80.240.23.152:30303",
	"enode://458605844b62ef5ba0820228476291747b9623185c6ec4742b2a8372aa61673fde4e33a27a90cd4c88c92906c8b5c134dcf976b277d64dcb502f0e9c12c927c0@217.69.3.196:30303",
	"enode://d3d4f7fbdf1cb8337a994912e578e80476082e16401abc7a60d88f51654954d5949eed6d74eedb3dbf866b66ceb53d6273dcf82de18239e8db149cbf1167b80d@209.250.240.205:30303",
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
	"enode://3a50f5a3b941cb934286ad783218d60409bea962cc8b2b6d52d1402fc550692801648dbeb08637acfed7b37a39ff93fc192bd1f39f5e2c336f92c805e9d49c1e@173.199.70.154:30303",
	"enode://bb9e4b9e3bec143ebd802d37561b568d560912a4faa6c10e6e714c83af827829bd77de25510b590b812096a0408ed427d3f9e098502578b94c72455b3c13eb3d@45.76.45.30:30303",
	"enode://85470f54eb5ac86d367dfa1377a148f87f9c4a2e6743cf569e65fc6b31ee1ab6b4c8b68e10d58acb36ecd2b987a9613d87866852b08a237d58869028deed19d8@209.250.251.109:30303",
	"enode://ac79ab5972ff0b7cad9e5f81944eff68947b51c6ea99f0e7847c8487234cf34bd46118164e22b1cafa269640ec1f04238abc02aab0df95352adaa1142278c3e3@80.240.23.152:30303",
	"enode://458605844b62ef5ba0820228476291747b9623185c6ec4742b2a8372aa61673fde4e33a27a90cd4c88c92906c8b5c134dcf976b277d64dcb502f0e9c12c927c0@217.69.3.196:30303",
	"enode://d3d4f7fbdf1cb8337a994912e578e80476082e16401abc7a60d88f51654954d5949eed6d74eedb3dbf866b66ceb53d6273dcf82de18239e8db149cbf1167b80d@209.250.240.205:30303",
}
