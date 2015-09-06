package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

type RSA struct {
	publicKey *big.Int
	privateKey *big.Int
	modulus *big.Int
}

func (rsa *RSA) initialize(n int) {
	p := genRandomPrime(n/2)
	q := genRandomPrime(n/2)
	phi := big.NewInt(1)
	phi.Sub(p, big.NewInt(1))
	mul := big.NewInt(1)
	mul.Sub(q, big.NewInt(1))
	phi.Mul(phi, mul)
	rsa.modulus = big.NewInt(1)
	rsa.modulus.Mul(p,q)
	rsa.publicKey = big.NewInt(65537); // common public key
	rsa.privateKey = big.NewInt(1)
	rsa.privateKey.ModInverse(rsa.publicKey, phi)
}

func (rsa *RSA) encrypt(msg *big.Int) *big.Int{
	return msg.Exp(msg,rsa.publicKey,rsa.modulus)
}

func (rsa *RSA) decrypt(enc *big.Int) *big.Int{
	return enc.Exp(enc,rsa.privateKey,rsa.modulus)
}
	
func genRandom(n int) *big.Int {
	max_num := big.NewInt(0)
	max_num.SetBit(max_num,n,1)
	num, err := rand.Int(rand.Reader, max_num)
	if err != nil {
	    panic(err.Error())
	}
	return num	
}

/** gen Random Prime numbers with atleast n bits **/
func genRandomPrime(n int) *big.Int {
	num := big.NewInt(2)
	for true {
		num := genRandom(n)
		num.SetBit(num,n,1)
		if num.ProbablyPrime(100) {
			return num	
		}
	}
	return num
}


func main() {
	n := 64;
	message := genRandom(n-1)
	fmt.Println("Message:",message)
	obj := new(RSA)
	obj.initialize(n)
	enc := obj.encrypt(message);
	fmt.Println("Encrypted Message:",enc)
	dec := obj.decrypt(enc);
	fmt.Println("Decrypted Message:",dec)
}
