package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// YAMLToJSON handles "yaml2json [INPUT] [OUTPUT]"
func YAMLToJSON(args []string) {
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
	wantsPretty := false // TODO -p --pretty
	var bytes []byte
	var err error
	for _, d := range rows {
		if wantsPretty {
			bytes, err = json.MarshalIndent(d, "", "  ")
		} else {
			bytes, err = json.Marshal(d)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to encode JSON: %s\n", err)
			os.Exit(1)
			return
		}
		fmt.Fprintf(w, "%s\n", string(bytes))
	}
}

func convertii(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convertii(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertii(v)
		}
	}
	return i
}
