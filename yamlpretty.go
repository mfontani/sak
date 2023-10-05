package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// YAMLPretty handles "yamlpretty [INPUT] [OUTPUT]"
func YAMLPretty(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	d := yaml.NewDecoder(r)
	var data interface{}
	var rows []interface{}
	for {
		err := d.Decode(&data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Failed to decode YAML: %s\n", err)
			os.Exit(1)
			return
		}
		data = convertii(data)
		rows = append(rows, data)
	}
	var bytes []byte
	var err error
	for _, d := range rows {
		bytes, err = yaml.Marshal(d)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to encode YAML: %s\n", err)
			os.Exit(1)
			return
		}
		fmt.Fprintf(w, "---\n%s", string(bytes))
	}
}
