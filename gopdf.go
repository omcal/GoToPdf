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

	file := os.Args[1]                        //Go takes file as a argv and it start from 1 because Args[0] is go
	content, err := ioutil.ReadFile(file)     //reading file  and error
	stringSlience := strings.Split(file, ".") // Spliting filename and its extantion
	fileName := stringSlience[0]              //to string
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")        // here I create a pdf file
	tr := pdf.UnicodeTranslatorFromDescriptor("") //to convert utf-8 might be we come upon non english words like é ş etc.
	pdf.AddPage()                                 //adding page
	pdf.SetFont("Arial", "B", 18)                 //setting font we use Arial as a font
	pdf.MultiCell(190, 5, tr(string(content)), "0", "0", false)
	_ = pdf.OutputFileAndClose(fileName + ".pdf")
	fmt.Printf("%s File Converted to PDF...", fileName) //Message

}
