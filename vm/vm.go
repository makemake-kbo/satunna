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
func RunVM(file *os.File, seed uint32) []uint64 {
	var rawFileAsUInt64 []uint64 = readFileAsUInt64Unsafe(file);
	fmt.Println("mi familia ", rawFileAsUInt64);

	ss := string(seed);

	var processedFile []uint64
	for i := 0; i < len(ss); i++ {
		switch string(ss[i]) {
		case "0":
			processedFile = IADD_M(rawFileAsUInt64);
		case "1":
			processedFile = ISUB_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
		case "2":
			processedFile = IMUL_M(rawFileAsUInt64);
		case "3":
			processedFile = IMULH_M(rawFileAsUInt64);
		case "4":
			processedFile = IMULHS_M(rawFileAsUInt64, seed);
		case "5":
			processedFile = IXOR_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
		case "6":
			processedFile = IOR_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
		case "7":
			processedFile = ISQR_R(rawFileAsUInt64);
		case "8":
			processedFile = ISIN_R(rawFileAsUInt64);
		case "9":
			processedFile = IADD_M(rawFileAsUInt64); // i ran out of ideas for opcodes
			processedFile = IADD_M(rawFileAsUInt64);
		default:
			panic("UNKNOWN OPCODE: EXECUTION ABORTED")
		}
	}

	return processedFile;

}
