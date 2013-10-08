// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gokyle/gokyle.talks/pkg/present"
	"github.com/gokyle/gokyle.talks/pkg/socket"
)

const basePkg = "github.com/gokyle/gokyle.talks/present"

var basePath string

func main() {
        var port string
        if port = os.Getenv("PORT"); port == "" {
                port = "8080"
        }
	httpListen := flag.String("http", "127.0.0.1:" + port, "host:port to listen on")
	flag.StringVar(&basePath, "base", "", "base path for slide template and static resources")
	flag.BoolVar(&present.PlayEnabled, "play", true, "enable playground (permit execution of arbitrary user code)")
	flag.Parse()

	if basePath == "" {
		p, err := build.Default.Import(basePkg, "", build.FindOnly)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't find gopresent files: %v\n", err)
			fmt.Fprintf(os.Stderr, basePathMessage, basePkg)
			os.Exit(1)
		}
		basePath = p.Dir
	}

	if present.PlayEnabled {
		playScript(basePath, "socket.js")
		http.Handle("/socket", socket.Handler)
	}
	http.Handle("/static/", http.FileServer(http.Dir(basePath)))

	if !strings.HasPrefix(*httpListen, "127.0.0.1") &&
		!strings.HasPrefix(*httpListen, "localhost") &&
		present.PlayEnabled {
		log.Print(localhostWarning)
	}

	log.Printf("Open your web browser and visit http://%s/", *httpListen)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}

const basePathMessage = `
By default, gopresent locates the slide template files and associated
static content by looking for a %q package
in your Go workspaces (GOPATH).

You may use the -base flag to specify an alternate location.
`

const localhostWarning = `
WARNING!  WARNING!  WARNING!

The present server appears to be listening on an address that is not localhost.
Anyone with access to this address and port will have access to this machine as
the user running present.

To avoid this message, listen on localhost or run with -play=false.

If you don't understand this message, hit Control-C to terminate this process.

WARNING!  WARNING!  WARNING!
`
