package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
    obsSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    deobsSource = "QWE1ASD2ZXC3RTY4FGH5VBN6UIOP7JKL8M9qwertyuiop0asdfghjklzxcvbnm"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close the file %v\n", err)
	}
}

func ObfuscateDeobfuscateXml(inputPath, outputPath string, isDeobfuscationMode bool) error {
	sourceFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer closeFile(sourceFile)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer closeFile(outputFile)

	scanner := bufio.NewScanner(sourceFile)
	writer := bufio.NewWriter(outputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() != nil {
			return fmt.Errorf("failed to read file during obfuscation: %v", scanner.Err())
		}

		var processedLine string
		if isDeobfuscationMode {
			processedLine = deobfuscateXmlString(line)
		} else {
			processedLine = obfuscateXmlString(line)
		}

		if _, err = writer.WriteString(processedLine + "\n"); err != nil {
			return fmt.Errorf("failed to write file during obfuscation: %v", err)
		}
	}

	return writer.Flush()
}

func obfuscateXmlString(xmlStr string) string {
	var result strings.Builder
	for _, symbol := range xmlStr {
		// Index returns -1 if symbol is not in source
		if i := strings.IndexRune(obsSource, symbol); i != -1 {
			result.WriteByte(deobsSource[i])
			continue
		}
		result.WriteRune(symbol)
	}

	return result.String()
}

func deobfuscateXmlString(xmlStr string) string {
	var result strings.Builder
	for _, symbol := range xmlStr {
		if i := strings.IndexRune(deobsSource, symbol); i != -1 {
			result.WriteByte(obsSource[i])
			continue
		}
		result.WriteRune(symbol)
	}

	return result.String()
}
