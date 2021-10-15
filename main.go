package main;

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	"fmt"
	
	"github.com/makemake-kbo/satunna/vm"
)

func main() {

	inputData := os.Args[1];

	file, err := os.Open(inputData);

	if checkIfFile(err) {
		fmt.Println("File cant be read, hashing as string");
		fmt.Println("string", deriveSeedFromString(inputData));
	} else {
		mifamilia, _ := os.Open(inputData); // file gets garbage collected but not when its passed to deriveSeedFromFile?
		vm.RunVM(mifamilia, deriveSeedFromFile(file));
	}

}
