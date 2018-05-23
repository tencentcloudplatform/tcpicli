package autogen

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/tencentcloudplatform/tcpicli/core"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
	"time"
)

type Pkg interface {
	DoAction(action string, query ...string) ([]byte, error)
}

type Gen struct {
	DocRoot      string
	Seq          []string
	SliceOptions map[string]string
	FuncMap      map[string][]string
	FixStruct    map[string][]string
	Version      map[string][]string
	PkgName      string
	Pkg          Pkg
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (g *Gen) Run() {
	if g.Pkg == nil {
		log.Println(`Gen.Pkg cannot be nil
Consider use template:

// +build ignore

package main

import (
	. "."
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	gen := &autogen.Gen{
		Pkg : new(Pkg),
		...
		...
	}

}
`)
	}
	for _, action := range g.Seq {
		sp := strings.Split(action, " ")
		if len(sp) > 1 {
			switch sp[0] {
			case "DO":
				cmd := action[3:]
				fmt.Println("DO", cmd)
				c := exec.Command("/bin/sh", "-c", cmd)
				c.Env = os.Environ()
				c.Stderr = os.Stderr
				c.Stdout = os.Stdout
				c.Run()
			case "SET":
				key := strings.Split(action[4:], "=")[0]
				val := action[5+len(key):]
				fmt.Println("SET", key, "=", val)
				c := exec.Command("/bin/sh", "-c", val)
				c.Env = os.Environ()
				var buf bytes.Buffer
				c.Stdout = &buf
				c.Run()
				os.Setenv(key, strings.TrimRight(buf.String(), "\n"))
			}
			continue
		}
		q := g.FuncMap[action][1:]
		var query []string
		for _, v := range q {
			v = os.Expand(v, os.Getenv)
			query = append(query, v)
		}
		document := fmt.Sprintf("%s%s", g.DocRoot, g.FuncMap[action][0])
		fmt.Println("start generate", action)
		b, err := g.Pkg.DoAction(action, query...)
		if err != nil {
			log.Println(action, err.Error())
			continue
		}
		str := fmt.Sprintf("%sResp", action)
		res, err := gojson.Generate(bytes.NewBuffer(b), gojson.ParseJson, str, g.PkgName, []string{"json"}, false, true)
		if err != nil {
			log.Println(action, err.Error())
			continue
		}
		s := strings.Join(strings.Split(string(res), "\n")[1:], "\n")

		if val, ok := g.FixStruct[action]; ok {
			for i := 0; i < len(val); i += 2 {
				name := val[i]
				typ := val[i+1]
				// 				re, err := regexp.Compile(`([^\s]+)\s+[^\s]+\s+(.json:"` + name + `".)`)
				re, err := regexp.Compile(`(` + gojson.FmtFieldName(name) + `)(.|\n)*(.json:"` + name + `".)`)
				if err != nil {
					log.Println(action, name, typ, err.Error())
					continue
				}
				// 				s = re.ReplaceAllString(s, "${1} "+typ+" ${2}")
				s = re.ReplaceAllString(s, "${1} "+typ+" ${3}")
			}
		}

		if version, ok := g.Version[action]; ok {
			responseField := version[0]
			var versionErr string
			if len(responseField) == 0 {
				responseField = "RequestId"
				versionErr = "`json:\"" + responseField + "\"`\nError interface{} `json:\"Error,omitempty\"`"
			} else if len(responseField) > 0 {
				versionErr = "`json:\"" + responseField + ",omitempty\"`\nError interface{} `json:\"Error,omitempty\"`"
			}
			re, err := regexp.Compile("(`json:\"" + responseField + "\"`)")
			if err != nil {
				log.Println(responseField, err.Error())
			}
			s = re.ReplaceAllString(s, versionErr)
		}

		hasSliceOptions := func(action string) string {
			if val, ok := g.SliceOptions[action]; ok {
				return val
			}
			return ""
		}

		var buf bytes.Buffer
		packageTemplate.Execute(&buf, struct {
			Timestamp    time.Time
			PkgName      string
			Action       string
			Struct       string
			Document     string
			CorePwd      string
			SliceOptions string
		}{
			Timestamp:    time.Now(),
			PkgName:      g.PkgName,
			Action:       action,
			Struct:       s,
			Document:     document,
			CorePwd:      core.PWD,
			SliceOptions: hasSliceOptions(action),
		})
		var overwrite bool

		// if current file does not exist then need overwrite
		fname := action + ".go"
		current, err := ioutil.ReadFile(fname)
		if err != nil {
			overwrite = true
		}

		// exclude first line timestamp

		currentS := strings.Join(strings.Split(string(current), "\n")[1:], "\n")
		bufS := strings.Join(strings.Split(buf.String(), "\n")[1:], "\n")

		bufS = gofmt(bufS)

		// if file exist and generate file not same as oldfile then overwrite
		if !overwrite && currentS != bufS {
			overwrite = true
		}
		if overwrite {
			err = ioutil.WriteFile(fname, []byte(gofmt(buf.String())), 0644)
			if err != nil {
				log.Println(action, err.Error())
				continue
			}
			fmt.Println("finish generate", action)
		} else {
			fmt.Println("finish nochange", action)
		}
	}
}

func gofmt(str string) string {
	b, err := format.Source([]byte(str))
	if err != nil {
		panic(err)
	}
	return string(b)
}

var funcMap = template.FuncMap{
	"capFirst": func(s string) string {
		if len(s) > 0 {
			return strings.ToUpper(s[:1]) + s[1:]
		}
		return s
	},
}
var packageTemplate = template.Must(template.New("").Funcs(funcMap).Parse(`// {{ .Timestamp }}
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package {{.PkgName}}
import (
	"encoding/json"{{if .SliceOptions}}
	"fmt"{{end}}
	"{{.CorePwd}}"
)

{{.Struct}}

// Implement {{.Document}}
func (c *{{.PkgName|capFirst}}Client) {{.Action}}({{if .SliceOptions}}{{.SliceOptions}}{{else}}options{{end}} ...string) (*{{.Action}}Resp, error) { {{if .SliceOptions}}
	var options []string
	for k, v := range {{.SliceOptions}} {
		options = append(options, fmt.Sprintf("{{.SliceOptions}}.%v=%v", k, v))
	}{{end}}
	resp, err := c.DoAction("{{.Action}}", options...)
	if err != nil {
		return nil, err
	}
	var s {{.Action}}Resp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func {{.Action}}({{if .SliceOptions}}{{.SliceOptions}}{{else}}options{{end}} ...string) (*{{.Action}}Resp, error) {
	return DefaultClient.{{.Action}}({{if .SliceOptions}}{{.SliceOptions}}{{else}}options{{end}}...)
}

func (r *{{.Action}}Resp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
`))
