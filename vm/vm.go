package vm

import (
	//"golang.org/x/crypto/blake2b"
	"os"
	"io"
	"strconv"
	"bytes"
	"encoding/binary"
	"fmt"
)


/*IntToSlice converts int to a slice so we can use it in the VM*/
func IntToSlice(n int64, sequence []int64) []int64 {
    if n != 0 {
        i := n % 10
        // sequence = append(sequence, i) // reverse order output
        sequence = append([]int64{i}, sequence...)
        return IntToSlice(n/10, sequence)
    }
    return sequence
}


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
func RunVM(file *os.File, seed uint32) []byte {
	var rawFileAsUInt64 []uint64 = readFileAsUInt64Unsafe(file);

	ss := strconv.Itoa(int(seed));

	var processedFile []uint64

	fmt.Println("Starting VM...");
	for i := 0; i < len(ss); i++ {
		switch string(ss[i]) {
		case "0":
			processedFile = IADD_M(rawFileAsUInt64);
			fmt.Println("IADD_M");
		case "1":
			processedFile = ISUB_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
			fmt.Println("ISUB_R");
		case "2":
			processedFile = IMUL_M(rawFileAsUInt64);
			fmt.Println("IADD_M");
		case "3":
			processedFile = IMULH_M(rawFileAsUInt64);
			fmt.Println("IMUL_M");
		case "4":
			processedFile = IMULS_M(rawFileAsUInt64, seed);
			fmt.Println("IMULS_M");
		case "5":
			processedFile = IXOR_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
			fmt.Println("IXOR_R");
		case "6":
			processedFile = IOR_R(rawFileAsUInt64, IMUL_M(rawFileAsUInt64));
			fmt.Println("IOR_R");
		case "7":
			processedFile = ISQR_R(rawFileAsUInt64);
			fmt.Println("ISQR_R");
		case "8":
			processedFile = ISIN_R(rawFileAsUInt64);
			fmt.Println("ISIN_R");
		case "9":
			processedFile = IADD_M(rawFileAsUInt64); // i ran out of ideas for opcodes
			processedFile = IADD_M(rawFileAsUInt64);
			fmt.Println("IADD_M * 2");
		default:
			panic("UNKNOWN OPCODE: EXECUTION ABORTED")
		}
	}

	var result uint64;

	for i := 0; i < len(processedFile); i++ {
		result += processedFile[i];
	}

	b := make([]byte, 8);
	binary.LittleEndian.PutUint64(b, result);

	fmt.Println("Final output: ", b);

	return b;

}
