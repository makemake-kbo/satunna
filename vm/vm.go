package vm;

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	//"io"
	//"strconv"
	"bytes"
	"fmt"
	"unsafe"
)

/*read contents of the file as a uint64 from memory. what could go wrong?*/
func readUInt64Unsafe(file *os.File) uint64 {
    return *(*uint64)(unsafe.Pointer(&file[0]))
}


/*we run the VM on the file here*/
func runVM (file *os.File, seed uint32) uint64 {
	var rawFileAsUInt64 uint64 = ReadUInt64Unsafe(file);

	fmt.Println(seed);

}
