package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/usr/share/windows-binaries/nc.exe")
	check(err)
	pefile, err := pe.NewFile(f)
	check(err)
	defer f.Close()
	defer pefile.Close()

	dosHeader := make([]byte, 96)
	sizeOffset := make([]byte, 4)

	// Dec to ASCII (searching for MZ)
	_, err = f.Read(dosHeader)
	check(err)
	fmt.Println("[[-----DOS Header / Stub-----]")
	fmt.Printf("[+] Magic Value: %s%s\n", string(dosHeader[0]), string(dosHeader[1]))

	// Validate PE+0+0 (Valid PE format)
	pe_sig_offset := int64(binary.LittleEndian.Uint32(dosHeader[0x3c:]))
	f.ReadAt(sizeOffset[:], pe_sig_offset)
	fmt.Println("[-----Signature Header-----]")
	fmt.Printf("[+] LFANEW Value: %s\n", string(sizeOffset))

	// Create the reader and read COFF Header
	sr := io.NewSectionReader(f, 0, 1<<63-1)
	_, err = sr.Seek(pe_sig_offset+4, os.SEEK_SET)
	check(err)
	binary.Read(sr, binary.LittleEndian, &pefile.FileHeader)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
