package vm

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	"io"
	//"strconv"
	"bytes"
	"fmt"
)

/*read contents of the file as a uint64 from memory. what could go wrong?*/
func readFileAsUInt64Unsafe(file *os.File) []uint64 {
	buf := bytes.NewBuffer(nil);
	io.Copy(buf, file);

	baf := []byte(buf.Bytes());

	var n []uint64;

	for i := 0; i < len(baf); i++ {
		n = append(n, uint64(baf[i]));
	}

	return n;
}

/*RunVM is used to get an intenger representation of the file after it has been processed according to the seed*/
func RunVM(file *os.File, seed uint32) uint64 {
	var rawFileAsUInt64 []uint64 = readFileAsUInt64Unsafe(file);
	fmt.Println("mi familia ", rawFileAsUInt64);


	return rawFileAsUInt64[0];

}

/*	buf := bytes.NewBuffer(nil);
	io.Copy(buf, file);
	fmt.Println("mi filebuf ", buf);

	var mifamilia []byte = buf.Bytes();
	var n []uint64;

	for i := 0; i < len(mifamilia); i++ {
		n = append(n, uint64(mifamilia[i]));
	}

	return n;*/