// Copyright (C) 2017, Beijing Bochen Technology Co.,Ltd.  All rights reserved.
//
// This file is part of L0
//
// The L0 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The L0 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/bocheninc/L0/components/utils"
)

const (
	testPrivateKey = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
)

func TestSign(t *testing.T) {
	priv, _ := HexToECDSA(testPrivateKey)
	msg := Sha256([]byte("hello"))

	sig, err := priv.Sign(msg[:])
	if err != nil {
		t.Errorf("Sign Error %s", err)
	}

	if sig.Validate() {
		t.Errorf("Signature Validate error %s", err)
	}

	s := Signature{}
	s.SetBytes(sig[:], false)

	if !s.Validate() {
		t.Errorf("Signature Validate error %v", s.Validate())
	}

	pub, err := sig.RecoverPublicKey(msg[:])
	if err != nil {
		t.Errorf("SigToPub Error %v - %v - %s", sig, pub, err)
	}

	pub2 := priv.Public()
	if !bytes.Equal(pub.Bytes(), pub2.Bytes()) {
		t.Errorf("public key not match! %0x - %0x ", pub.Bytes(), pub2.Bytes())
	}
}

func TestHash(t *testing.T) {
	var (
		testSha256HashStr    = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
		testRipemd160HashStr = "8eb208f7e05d987a9b044a8e98c6b087f15a0bfc"
	)
	// Ripemd160 hash
	h := Ripemd160([]byte("abc"))
	if hex.EncodeToString(h) != testRipemd160HashStr {
		t.Errorf("Rimped160(%s) = %s, except %s !", "abc", testRipemd160HashStr, hex.EncodeToString(h))
	}

	// Sha256 hash
	h = Sha256([]byte("hello")).Bytes()
	hash := NewHash(h)
	if hash.String() != testSha256HashStr {
		t.Errorf("Sha256(%s) = %s, except %s !", "hello", testSha256HashStr, hash.String())
	}

	// Setbytes
	h2 := Sha256([]byte("hello"))
	hash2 := new(Hash)
	hash2.SetBytes(h2[:])
	if hash2.String() != testSha256HashStr {
		t.Errorf("Sha256(%s) = %s, except %s !", "hello", testSha256HashStr, hash.String())
	}
	fmt.Println(DoubleSha256(utils.HexToBytes("925a82a0970b3fe1e5831d307a54458ba1b935768ee789c11a55c55b97f360b1")))
}

func TestLoadAndSaveECDSA(t *testing.T) {
	priv, _ := HexToECDSA(testPrivateKey)
	priv.SaveECDSA("nodekey")
	priv2, _ := LoadECDSA("nodekey")
	if !bytes.Equal(priv.SecretBytes(), priv2.SecretBytes()) {
		t.Errorf("load and save private key error %s != %s", priv, priv2)
	}
}
