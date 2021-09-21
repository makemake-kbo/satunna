package satunna

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	"fmt"
)

func main() {

	/* currently used to get input data to derive the seed.
	 * change later to accomodate user input and any other args.
	 */
	inputData := os.Args[1];

	fmt.Println(inputData);
}