// Package skeleton makes skeletons to be filled out with solutions.
package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"flag"
	"path/filepath"
	"text/template"

)

//go:embed tmpls/*.go
var fs embed.FS

func main(){
	var year int
	var day int

	flag.IntVar(&year, "year", 2015, "Year as YYYY")
	flag.IntVar(&day, "day", 01, "Day as DD")
	flag.Parse()
	fmt.Printf("Creating %d/%d\n",year,day)

	Run(day, year)

}
// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fs, "tmpls/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	mainFilename := filepath.Join("./", fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join("./", fmt.Sprintf("%d/day%02d/main_test.go", year, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", nil)
	fmt.Printf("templates made for %d-day%d\n", year, day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
