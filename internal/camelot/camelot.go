package camelot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type DirectoryWithTables struct {
	splitPDFs []SplitPDF
}

type SplitPDF struct {
	AbsolutePath string
	RawOutput    string
}
type PDFPageImage struct{}

// java -jar bin/tabula.jar -p all -g -d -l ~/GOMOD/go-pardocs/raw/BukanLisan/Jawapan\ Bertulis\ dikemaskini\ pada\ 18\ Sept\ 2019.pdf

func ExtractTablesFromSplitDirectory(splitDir string) DirectoryWithTables {

	// For each PDF file in the split; get the basename and run above command on it
	// If nothing; mark as nothing .. else create the corresponding csv file to be cleaned
	fInfo, err := ioutil.ReadDir(splitDir)
	if err != nil {
		panic(err)
	}
	// In case it is relative ..
	absolutePathDir, aerr := filepath.Abs(splitDir)
	if aerr != nil {
		panic(aerr)
	}
	splitPDFs := make([]SplitPDF, 0)
	// If want subfolders, we can use https://flaviocopes.com/go-list-files/
	for _, f := range fInfo {
		// Ignore if dir ..
		if !f.IsDir() {
			absoluePathPDF := absolutePathDir + "/" + f.Name()
			// DEBUG
			//fmt.Println("Checking ... ", absoluePathPDF)
			output := detectTablesInPDF(absoluePathPDF)
			if output != "" {
				//  If has output; is file type .. and can go on; else skip ..
				splitPDF := SplitPDF{
					AbsolutePath: absoluePathPDF,
					RawOutput:    output,
				}
				// append to  the type if has output
				splitPDFs = append(splitPDFs, splitPDF)
			} else {
				fmt.Println("Ignoring  ...", f.Name())
			}
		}
	}
	// DEBUG
	//spew.Dump(splitPDFs)
	return DirectoryWithTables{splitPDFs: splitPDFs}
}

func detectTablesInPDF(absolutePDFPath string) string {
	// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
	// java -jar bin/tabula.jar -p all -g -d -l ~/GOMOD/go-pardocs/raw/BukanLisan/Jawapan\ Bertulis\ dikemaskini\ pada\ 18\ Sept\ 2019.pdf
	cmd := exec.Command("java", "-jar", "/Users/leow/GOMOD/go-camelot/bin/tabula.jar", "-p", "all", "-g", "-d", "-l", absolutePDFPath)
	// DEBUG
	//fmt.Println("RUN: ", cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// Stdout, StdErr
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	// DEBUG
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if errStr != "" {
		fmt.Println("ERR: ", errStr)
	}
	return outStr
}
func (sp SplitPDF) SaveCSV() {

}

func (sp SplitPDF) SaveYAML() {

}
