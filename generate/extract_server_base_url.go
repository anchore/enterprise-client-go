package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/getkin/kin-openapi/openapi3"
)

// This program generates server_base_url.go for each client package
const (
	source = "server_base_url.go"
)

var (
	apis = []string{
		"enterprise",
	}
	tmp = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT. This file was generated by robots at {{ .Timestamp }}
package {{ .Package }}

const ServerBaseURL = {{ printf "%q" .URL }}
`))
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	for _, api := range apis {
		path := filepath.Join("pkg", api, source)

		fmt.Printf("generating %q...\n", path)

		openapiDocPath := filepath.Join("pkg", api, "api", "openapi.yaml")
		openapiF, err := os.Open(openapiDocPath)
		if err != nil {
			return fmt.Errorf("unable to open openapi doc %q: %w", openapiDocPath, err)
		}

		openapiDoc, err := ioutil.ReadAll(openapiF)
		if err != nil {
			return fmt.Errorf("unable to read openapi doc %q: %w", openapiDocPath, err)
		}
		if err = openapiF.Close(); err != nil {
			return fmt.Errorf("unable to close openapi doc %q: %w", openapiDocPath, err)
		}

		url, err := extractServerBaseURL(string(openapiDoc))
		if err != nil {
			return fmt.Errorf("unable to extract server base url from openapi=%q doc: %w", openapiDocPath, err)
		}

		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("unable to create %q: %+v", path, err)
		}

		err = tmp.Execute(f, struct {
			Timestamp time.Time
			URL       string
			Package   string
		}{
			Timestamp: time.Now(),
			URL:       url,
			Package:   api,
		})

		if err := f.Close(); err != nil {
			return fmt.Errorf("unable to close %q: %+v", path, err)
		}

		if err != nil {
			return fmt.Errorf("unable to generate template: %+v", err)
		}
	}
	return nil
}
func extractServerBaseURL(openapiDoc string) (string, error) {
	var doc openapi3.T

	if err := yaml.Unmarshal([]byte(openapiDoc), &doc); err != nil {
		return "", err
	}

	var baseURL string
	for _, s := range doc.Servers {
		baseURL = s.URL
		break
	}

	if baseURL == "" {
		return "", fmt.Errorf("no base url for the enterprise API found")
	}
	return baseURL, nil
}
