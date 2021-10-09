package main;

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	"fmt"
	
	"github.com/makemake-kbo/satunna/vm"
)

func main() {

	/* currently used to get input data to derive the seed.
	 * change later to accomodate user input and any other args.
	 */
	inputData := os.Args[1];

	file, err := os.Open(inputData);

	if checkIfFile(err) {
		fmt.Println("File cant be read, hashing as string");
		fmt.Println("string", deriveSeedFromString(inputData));
	} else {
		fmt.Println("xd", deriveSeedFromFile(file));
		vm.runVM(file, seed);
	}

}
