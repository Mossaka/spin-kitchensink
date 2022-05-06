package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

const pathInfoEnv = "PATH_INFO"

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		path := os.Getenv(pathInfoEnv)
		buf, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read file %v: %v", path, err)
		}

		if _, err = io.Copy(w, buf); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file %s: %s\n", path, err)
		}
	})
}

func main() {}
