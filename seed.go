package main;

import (
	"golang.org/x/crypto/blake2b"
	"os"
	"io"
	"strconv"
	"bytes"
)

func deriveSeedFromFile (file *os.File) uint32 {
	buf := bytes.NewBuffer(nil);

	/*copies the file 3 times into buffer, memory and i/o hardening happens here*/
	for i := 0; i < 3; i++ {
		io.Copy(buf, file);
	}

	hash, err := blake2b.New(8, nil);

	hash.Write([]byte(buf.Bytes())); /*giving data from buffer here*/

	bs := string(hash.Sum(nil)); /*TODO: clean this up*/

	n, err := strconv.ParseUint(bs, 16, 32); /*jeet*/
	if err != nil {
		panic(err);
	}

	return uint32(n);
}
