package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hoklet/BDAS/lab1/pkg"
)

var (
	pathToXml = flag.String("xml", "/tmp/file.xml", "Path to XML file to obfuscate/deobfuscate")
	outputPath = flag.String("output", "/tmp/output-file.xml", "Path to output file")
	deobfuscate = flag.Bool("deobfuscate", false, "Deobfuscate file")
)

func main() {
	flag.Parse()

	if *deobfuscate {
		fmt.Println("Deobfuscation started")
	} else {
		fmt.Println("Obfuscation started")
	}

	var err error
	err = pkg.ObfuscateDeobfuscateXml(*pathToXml, *outputPath, *deobfuscate)

	if err != nil {
		fmt.Printf("Obfuscation/Deobfuscation is failed with %v.\nTry to cleanup corrupted outputfile.\n", err)

		err = os.Remove(*outputPath)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Fatalf("Failed to cleanup corrupted file %v", err)
			}
		}
	} else {
		fmt.Println("Program executed successfully")
	}
}
