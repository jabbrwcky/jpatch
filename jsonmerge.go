package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

var force bool
var pretty bool

func main() {
	outFile := flag.String("o", "-", "Specify the outputfile")
	flag.BoolVar(&force, "f", false, "Force to continue even on type mismatches")
	flag.BoolVar(&pretty, "p", false, "Pretty-print the json output")

	flag.Parse()
	files := flag.Args()

	if len(files) < 2 {
		fmt.Println("Need at least an input file and a patch file")
		os.Exit(1)
	}

	var err error
	out := os.Stdout

	if *outFile != "-" {
		var err error
		out, err = os.OpenFile(*outFile, os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Unable to open outputfile!")
			os.Exit(8)
		}
	}

	var source map[string]interface{}

	if source, err = readJson(files[0]); err != nil {
		fmt.Println("Error unmarshalling source", err)
	}

	var patch map[string]interface{}

	if patch, err = readJson(files[1]); err != nil {
		fmt.Println("Error unmarshalling patch", err)
	}

	var marshalled []byte
	if marshalled, err = json.Marshal(merge(source, patch)); err != nil {
		log.Fatalf("Failed to marshal result!", err)
	}

	if pretty {
		var pretty bytes.Buffer
		json.Indent(&pretty, marshalled, "", "  ")
		pretty.WriteTo(out)
	} else {
		out.Write(marshalled)
	}
}

func merge(src map[string]interface{}, mod map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k := range src {
		replVal := mod[k]
		if replVal == nil {
			result[k] = src[k]
		} else {
			replace(k, src, mod[k], &result)
			delete(mod, k)
		}
	}

	/* process new keys introduced by replacement map */
	for k, v := range mod {
		result[k] = v
	}

	return result
}

func replace(key string, src map[string]interface{}, replacement interface{}, result *map[string]interface{}) error {
	switch rt := replacement.(type) {
	case map[string]interface{}:
		if reflect.TypeOf(src[key]) != reflect.TypeOf(rt) {
			logTypeMismatch(key, src[key], rt)
			(*result)[key] = rt
		} else {
			(*result)[key] = merge(src[key].(map[string]interface{}), rt)
		}
	case bool:
		if reflect.TypeOf(src[key]).Kind() != reflect.Bool {
			logTypeMismatch(key, src[key], rt)
		}
		(*result)[key] = rt
	case float64:
		if reflect.TypeOf(src[key]).Kind() != reflect.Float64 {
			logTypeMismatch(key, src[key], rt)
		}
		(*result)[key] = rt
	case string:
		if reflect.TypeOf(src[key]).Kind() != reflect.String {
			logTypeMismatch(key, src[key], rt)
		}
		(*result)[key] = rt
	case nil:
		/* delete key */
	case []interface{}:
		if reflect.TypeOf(src[key]) != reflect.TypeOf(rt) {
			logTypeMismatch(key, src[key], rt)
		}
		(*result)[key] = rt
	}
	return nil
}

func logTypeMismatch(key string, src interface{}, mod interface{}) {
	if force {
		log.Printf("Type mismatch for key %s. Type of source is %s, replacement is of type %s. \n", key, reflect.TypeOf(src), reflect.TypeOf(mod))
	} else {
		log.Fatalf("Type mismatch for key %s. Type of source is %s, replacement is of type %s. \n", key, reflect.TypeOf(src), reflect.TypeOf(mod))
	}
}

func readJson(file string) (map[string]interface{}, error) {
	var buffer []byte
	var err error
	if buffer, err = ioutil.ReadFile(file); err != nil {
		fmt.Println(err)
		return nil, err
	}

	parsed := make(map[string]interface{})

	if err := json.Unmarshal(buffer, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}
