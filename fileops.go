/*
fileops contains functions used specifically for file managment
*/
package main;

import (
	"fmt"
)

/*panics if it fails to read file*/
func panicOnFail(e error) {
	if e != nil {
		panic(e)
	}
}

/*returns true if file can be read, else false*/
func checkIfFile(e error) bool {
	if e != nil {
		return true;
	}
	fmt.Println(e);
	return false;
}
