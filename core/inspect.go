package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

// unmarshal response to query into empty interface
func Inspect(flag string, cmd string, action string, options ...string) error {
	var respBody interface{}
	// get the default region from the configfile
	// annoying when this function doesn't support default region, even if it messies the code a bit
	region, ok := HasRegion(options...)
	if !ok {
		region = DefaultRegion()
		options = append(options, "Region="+region)
	}
	resp, err := DoAction(cmd, action, options...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &respBody)
	filter := buildFilter(flag)

	output := executeTemplate(filter, respBody)
	return output
}

// builds filter to execute against template
func buildFilter(flag string) string {
	if strings.Contains(flag, "{{") == true && strings.Contains(flag, "}}") == true {
		filter := fmt.Sprintf("%s\n", flag)
		return filter
	} else if _, err := os.Stat(flag); !os.IsNotExist(err) {
		buf := bytes.NewBuffer(nil)
		file, err := os.Open(flag)
		// os.Open always returns err. freaks compiler out, but no issues in runtime
		// leaving as log.Fatal(err). elegant sltn later
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(buf, file)
		file.Close()
		filter := string(buf.Bytes())
		return filter
	} else {
		fmt.Println("Please enter an input file or string. See help for details: tcpicli -h")
		os.Exit(1)
	}

	return ""
}

// executes template using filter
func executeTemplate(filter string, respBody interface{}) error {
	tmpl, err := template.New("template").Parse(filter)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, respBody)
	if err != nil {
		return err
	}

	return nil
}
