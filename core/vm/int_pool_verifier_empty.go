// Copyright 2017 The promethium Authors
// This file is part of the promethium library.
//
// The promethium library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The promethium library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the promethium library. If not, see <http://www.gnu.org/licenses/>.

// +build !VERIFY_EVM_INTEGER_POOL

package vm

const verifyPool = false

func verifyIntegerPool(ip *intPool) {}