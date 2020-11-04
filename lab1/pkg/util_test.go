package pkg

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testInputPath = "/tmp/file.xml"
	testOutputPath = "/tmp/outputFile.xml"
)

func TestObfuscateXmlString(t *testing.T) {
	testStrings := []string{"<employee id=\"111\">", "<lastName>Gupta</lastName>", "<location>USA</location>"}
	assertStrings := []string{"<KrueygKK 9J=\"kkk\">", "<eOp0TOrK>Dau0O</eOp0TOrK>", "<ey7O09yt>VHQ</ey7O09yt>"}

	for i, _ := range testStrings {
		res := obfuscateXmlString(testStrings[i])
		if res != assertStrings[i] {
			t.Errorf("Obfuscate string is failed. The result should be %s but it is %s", res, assertStrings[i])
		}
	}
}

func TestDeobfuscateXmlString(t *testing.T) {
	testStrings := []string{"<KrueygKK 9J=\"kkk\">", "<eOp0TOrK>Dau0O</eOp0TOrK>", "<ey7O09yt>VHQ</ey7O09yt>"}
	assertStrings := []string{"<employee id=\"111\">", "<lastName>Gupta</lastName>", "<location>USA</location>"}

	for i, _ := range testStrings {
		res := deobfuscateXmlString(testStrings[i])
		if res != assertStrings[i] {
			t.Errorf("Deobfuscate string is failed. The result should be %s but it is %s", res, assertStrings[i])
		}
	}
}

func TestObfuscateDeobfuscateXml(t *testing.T) {
	testXml := []byte(`
            <employees>
               <employee id="111">
                   <firstName>Lokesh</firstName>
                   <lastName>Gupta</lastName>
                   <location>India</location>
               </employee>
            </employees>
            `)

	if err := ioutil.WriteFile(testInputPath, testXml, 0777); err != nil {
		t.Errorf("Failed to create test file %s", testInputPath)
	}
	defer func() {
		if err := os.Remove(testInputPath); err != nil {
			t.Errorf("Failed to remove temp file %s", testInputPath)
		}
	}()

	if err := ObfuscateDeobfuscateXml(testInputPath, testOutputPath, false); err != nil {
		t.Errorf("Failed to obfuscate file %v", err)
	}
	defer func() {
		if err := os.Remove(testOutputPath); err != nil {
			t.Errorf("Failed to remove temp file %s", testOutputPath)
		}
	}()

	if err := ObfuscateDeobfuscateXml(testInputPath, testOutputPath, true); err != nil {
		t.Errorf("Failed to deobfuscate file %v", err)
	}
}
