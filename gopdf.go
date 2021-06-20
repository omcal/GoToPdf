package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	file := os.Args[1]
	content, err := ioutil.ReadFile(file)
	stringSlience := strings.Split(file, ".")
	fileName := stringSlience[0]
	fmt.Println(stringSlience[0])
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MultiCell(190, 5, tr(string(content)), "0", "0", false)
	_ = pdf.OutputFileAndClose(fileName + ".pdf")
	fmt.Println("PDF Created")

}
