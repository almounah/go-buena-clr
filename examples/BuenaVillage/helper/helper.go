// Used to do a quick and dirty xor on the file

package main

import (
	"flag"
	"fmt"
	"os"
)

func main()  {

    key := byte(133)

    pShellcodePath := flag.String("file", "", "Path Of the Shellcode")
    flag.Parse()

    shellcodePath := *pShellcodePath

    clearShellcodeByte, err := os.ReadFile(shellcodePath)
    if err != nil {
        fmt.Println("Error Opening file")
        fmt.Println(err.Error())
    }

    var encryptedShellcode []byte

    for i := 0; i < len(clearShellcodeByte); i++ {
        encryptedShellcode = append(encryptedShellcode, clearShellcodeByte[i] ^ key )
    }

    filename := "file.enc"

	err = os.WriteFile(filename, encryptedShellcode, 0644)
	if err != nil {
		fmt.Println("Failed to write file:", err)
		return
	}
}
