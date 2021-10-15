package main;

import (
	"golang.org/x/crypto/blake2b"
	"os"
	"io"
	//"strconv"
	"bytes"
	"fmt"
)

func deriveSeedFromFile (file *os.File) uint32 {
	var n uint32;
	buf := bytes.NewBuffer(nil);
	io.Copy(buf, file);

	/*creates a new hash with 8 byte lenght*/
	hash, err := blake2b.New(16, nil);
	if err != nil {
		panic(err);
	}

	hash.Write([]byte(buf.Bytes())); /*giving data from buffer here*/

	bs := hash.Sum(nil);/*TODO: clean this up*/
	fmt.Println("seed hash value: ", bs);
	for i := 0; i < len(bs); i++ {
		n += uint32(bs[i]);
	}

	return n;
}


func deriveSeedFromString (string string) uint32 {
	var n uint32;
	buf := []byte(string);

	/*creates a new hash with 8 byte lenght*/
	hash, err := blake2b.New(16, nil);
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

	return n;
}
