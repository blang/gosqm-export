package main

import (
	"flag"
	"fmt"
	"github.com/blang/gosqm"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

const LF = "\r\n"

var (
	output = flag.String("output", "", "file to write to, empty to print to stdout")
	input  = flag.String("input", "mission.sqm", "mission.sqm to read from")
)

var tmpl *template.Template

func init() {
	flag.Parse()
}

func main() {
	tmpl, err := readTemplate()
	if err != nil {
		fmt.Printf("Could not read template: %s"+LF, err)
		return
	}

	if *input == "" {
		fmt.Print("Specify path to mission.sqm" + LF)
		return
	}
	f, err := os.Open(*input)
	if err != nil {
		fmt.Printf("Can't open mission.sqm: %s"+LF, err)
		return
	}
	defer f.Close()
	dec := gosqm.NewDecoder(f)
	missionFile, err := dec.Decode()
	if err != nil {
		fmt.Printf("Error while reading mission.sqm: %s"+LF, err)
		return
	}
	var out io.Writer

	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Printf("Can't write to output: %s"+LF, err)
			return
		}
		defer f.Close()
		out = io.Writer(f)
	} else {
		out = os.Stdout
	}

	if missionFile.Mission != nil {
		err = tmpl.Execute(out, missionFile.Mission)
		if err != nil {
			fmt.Printf("Error while executing template %s"+LF, err)
			return
		}

	} else {
		fmt.Println("No Mission found")
		return
	}
}

func readTemplate() (*template.Template, error) {
	b, err := ioutil.ReadFile("export.tmpl")
	if err != nil {
		fmt.Printf("Error reading template %s", err)
		return nil, err
	}
	funcMap := template.FuncMap{
		"exportGroup": exportGroup,
	}
	return template.New("export").Funcs(funcMap).Parse(string(b))
}
