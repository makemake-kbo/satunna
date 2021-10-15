package main;

import (
	"golang.org/x/crypto/blake2b"
	"os"
	"io"
	"bytes"
	"fmt"
)


func deriveSeedFromFile (file *os.File) uint32 {
	var n uint32;
	buf := bytes.NewBuffer(nil);
	
	/*primitive memory hardening*/
	for i := 0; i < 16; i++ {
		io.Copy(buf, file);
	}

	/*creates a new hash with 8 byte lenght*/
	hash, err := blake2b.New(8, nil);
	if err != nil {
		panic(err);
	}

	hash.Write([]byte(buf.Bytes())); /*giving data from buffer here*/

	bs := hash.Sum(nil);/*TODO: clean this up*/
	fmt.Println("seed hash value: ", bs);
	for i := 0; i < len(bs); i++ {
		n += uint32(bs[i]);
	}

	/* math.Pow doesnt work so we do this
	 */
	var nj uint32; /*n gets GC'd if we use it in the for loop*/
	for i := 0; i < int(n); i++ {
		nj = n*n*n;
	}
	return nj;
}


func deriveSeedFromString (string string) uint32 {
	var n uint32;
	buf := []byte(string);

	/*creates a new hash with 8 byte lenght*/
	hash, err := blake2b.New(8, nil);
	if err != nil {
		panic(err);
	}

	hash.Write([]byte(buf)); /*giving data from buffer here*/

	bs := hash.Sum(nil);
	/*TODO: Debbbuging prints*/
	fmt.Println("seed hash value: ", bs);
	fmt.Printf("seed hash type: %T\n", bs);

	for i := 0; i < len(bs); i++ {
		n += uint32(bs[i]);
	}

	var nj uint32; /*n gets GC'd if we use it in the for loop*/
	for i := 0; i < int(n); i++ {
		nj = n*n*n;
	}
	return nj;
}
