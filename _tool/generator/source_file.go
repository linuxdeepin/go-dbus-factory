package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type SourceFile struct {
	Pkg       string
	GoImports []string
	GoBody    *SourceBody
}

func NewSourceFile(pkg string) *SourceFile {
	sf := &SourceFile{
		Pkg:    pkg,
		GoBody: &SourceBody{},
	}
	return sf
}

func (v *SourceFile) Print() {
	v.WriteTo(os.Stdout)
}

func (v *SourceFile) Save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("fail to create file:", err)
	}
	defer f.Close()

	v.WriteTo(f)

	err = f.Sync()
	if err != nil {
		log.Fatal("fail to sync file:", err)
	}

	err = exec.Command("go", "fmt", filename).Run()
	if err != nil {
		log.Fatal("failed to format file:", filename)
	}
}

func (v *SourceFile) WriteTo(w io.Writer) {
	io.WriteString(w, "package "+v.Pkg+"\n")

	sort.Strings(v.GoImports)
	for _, imp := range v.GoImports {
		io.WriteString(w, "import "+imp+"\n")
	}

	w.Write(v.GoBody.buf.Bytes())
}

// unsafe => "unsafe"
// or x,github.com/path/ => x "path"
func (s *SourceFile) AddGoImport(imp string) {
	var importStr string
	if strings.Contains(imp, ",") {
		parts := strings.SplitN(imp, ",", 2)
		importStr = fmt.Sprintf("%s %q", parts[0], parts[1])
	} else {
		importStr = `"` + imp + `"`
	}

	for _, imp0 := range s.GoImports {
		if imp0 == importStr {
			return
		}
	}
	s.GoImports = append(s.GoImports, importStr)
}

type SourceBody struct {
	buf bytes.Buffer
}

func (v *SourceBody) writeStr(str string) {
	v.buf.WriteString(str)
}

func (v *SourceBody) Pn(format string, a ...interface{}) {
	v.P(format, a...)
	v.buf.WriteByte('\n')
}

func (v *SourceBody) P(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	v.writeStr(str)
}
