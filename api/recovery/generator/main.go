// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	// The file name to be generated as output.
	blobFileName string = "../internal/blob.go"
)

// Define vars for build template
var (
	conv = map[string]interface{}{"conv": fmtByteSlice}
	tmpl = template.Must(template.New("").Funcs(conv).Parse(`// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

// Code generated by go generate; DO NOT EDIT.

func init() {
    {{- range $name, $file := . }}
        stackRegistry.Add("{{ $name }}", []byte{ {{ conv $file }} })
    {{- end }}
}`))
)

// fmtByteSlice
//
// Format's a byte slice for execution of the template.
func fmtByteSlice(s []byte) string {
	builder := strings.Builder{}

	for _, v := range s {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}

	return builder.String()
}

// main
//
// Reads all .go files in the api/directory and encodes
// them as a map in blob.go as a map of filenames and
// the file contents for stack traces.
func main() {

	// Remove file thats already been created
	os.Remove(blobFileName)

	// Create map for filenames
	stack := make(map[string][]byte)

	err := filepath.Walk("../../../api", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") && !strings.Contains(path, "_test.go") {
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			file := strings.Replace(path, "../../../", "/verbis/", -1)
			fmt.Println(file)
			stack[file] = contents
		}

		return nil
	})

	if err != nil {
		die("Error walking file directory", err)
	}

	// Create buffer
	builder := &bytes.Buffer{}

	// Execute template
	if err = tmpl.Execute(builder, stack); err != nil {
		die("Error executing template", err)
	}

	// Formatting generated code
	data, err := format.Source(builder.Bytes())
	if err != nil {
		die("Error formatting generated code", err)
	}

	// Writing blob file
	if err = ioutil.WriteFile(blobFileName, data, os.ModePerm); err != nil {
		die("Error writing blob file", err)
	}

	return
}

// die
//
// Exit out of the generation
func die(msg string, err error) {
	if err != nil {
		log.Fatal(msg + ": " + err.Error())
	}
}