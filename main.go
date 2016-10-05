package main

import "os"

//	when we run `go generate` from the cli it will run the
//	`go run` command oulined below
//	**Important: sure to include the comment below for the generator to see**
//go:generate go run generators/generator.go

//This func takes the name that you want the generated []byte to be as a binary
func genFile(p string) (int, error) {
	//Create the binary file
	file, err := os.Create(p)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	//Write the data to the the created file
	return file.Write(data)
}

func main() {
	p := "./final-bin"
	genFile(p)
	//Finally make the file executable
	//I know... 777 is bad
	os.Chmod(p, 0777)
}
