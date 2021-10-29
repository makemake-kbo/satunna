package main;

import (
	"os"
	"fmt"
	"unsafe"
	"encoding/base64"

	"golang.org/x/crypto/blake2b"	
	"github.com/makemake-kbo/satunna/vm"
)

func main() {

	inputData := os.Args[1];

	file, err := os.Open(inputData);

	if checkIfFile(err) {
		fmt.Println("File cant be read, hashing as string");
		fmt.Println("string", deriveSeedFromString(inputData));
	} else {
		mifamilia, _ := os.Open(inputData); /*file gets garbage collected but not when its passed to deriveSeedFromFile?*/
		
		/*run VM on the file*/

		hash, err := blake2b.New512(nil);
		if err != nil {
			panic(err);
		}

		hash.Write(vm.RunVM(mifamilia, deriveSeedFromFile(file)));
		fmt.Println("Hash: ", base64.URLEncoding.EncodeToString(hash.Sum(nil)));
	}

}
