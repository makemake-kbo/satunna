package main;

import (
	"os"
	"fmt"
	"encoding/base64"
	"testing"

	"golang.org/x/crypto/blake2b"	
	"github.com/makemake-kbo/satunna/vm"
)

func TestHelloName(t *testing.T) {
	//want := "cZQw4YPVvuEKljZJ7cT3La6fxQs9oqwoR1-haKj6BNofit-plK-AgXGJMFsPlPEFTnkUkSQrngvwOjVVWyQkgA===="

	file, err := os.Open(".gitignore");
	hash, err := blake2b.New512(nil);
	if err != nil {
		panic(err);
	}

	hash.Write(vm.RunVM(file, deriveSeedFromFile(file)));
	fmt.Println("Hash: ", base64.URLEncoding.EncodeToString(hash.Sum(nil)));
	//outputHash := base64.URLEncoding.EncodeToString(hash.Sum(nil));
	// if outputHash != want || err != nil {
	// 	t.Fatalf(`vm.RunVM(file, deriveSeedFromFile(file)) = %q, %v, want match for %#q, nil`, outputHash, err, want)
	// }
}
