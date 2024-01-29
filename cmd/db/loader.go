package main

import (
	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"github.com/sirrobot01/oauth-sso/api/models"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("sqlite").Load(models.Models...)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		return
	}
}
