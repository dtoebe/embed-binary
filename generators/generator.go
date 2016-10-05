package main

//Much of this file is from http://stackoverflow.com/questions/17796043/golang-embedding-text-file-into-compiled-executable

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//Empty string slice to place the byte lierals in
	var dataSlice []string

	//This will be the generated go file
	//Note: the path is ./ because we are using the go generate command
	// inside the main.go file
	outfile, err := os.Create("./generateddata.go")
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()

	//This is the file we will turn to []bytes
	infile, err := ioutil.ReadFile("./sample-data/generated-bin")
	if err != nil {
		panic(err.Error())
	}

	//Write the initial data to the generated go file
	//package main
	//var data = []byte{ EOF... so far
	outfile.Write([]byte("package main\n\nvar (\n\tdata = []byte{"))

	//Here we loop over each byte in the []byte from the sample
	//data file.
	//Take the byte literal and format it as a string
	//Then append it to the empty []string dataSlice we created at
	//the top of the func
	//Depending on the size of infile this could take a bit
	for _, b := range infile {
		bString := fmt.Sprintf("%v", b)
		dataSlice = append(dataSlice, bString)
	}

	//We join the []string together seperating it with commas
	//Remember we are writing a go src file so everything has to be a string
	dataString := strings.Join(dataSlice, ", ")

	//Write all the data to the generated file
	outfile.WriteString(dataString)

	//Finish off the data
	outfile.Write([]byte("}\n)"))
}
