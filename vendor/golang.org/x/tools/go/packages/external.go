// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file enables an external tool to intercept package requests.
// If the tool is present then its results are used in preference to
// the go list command.

package packages

import (
	"bytes"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 03397665 (update api)
	exec "golang.org/x/sys/execabs"
	"os"
	"strings"
)

// The Driver Protocol
//
// The driver, given the inputs to a call to Load, returns metadata about the packages specified.
// This allows for different build systems to support go/packages by telling go/packages how the
// packages' source is organized.
// The driver is a binary, either specified by the GOPACKAGESDRIVER environment variable or in
// the path as gopackagesdriver. It's given the inputs to load in its argv. See the package
// documentation in doc.go for the full description of the patterns that need to be supported.
// A driver receives as a JSON-serialized driverRequest struct in standard input and will
// produce a JSON-serialized driverResponse (see definition in packages.go) in its standard output.

// driverRequest is used to provide the portion of Load's Config that is needed by a driver.
type driverRequest struct {
	Mode LoadMode `json:"mode"`
	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`
	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`
	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`
	// Overlay maps file paths (relative to the driver's working directory) to the byte contents
	// of overlay files.
	Overlay map[string][]byte `json:"overlay"`
}

// findExternalDriver returns the file path of a tool that supplies
<<<<<<< HEAD
=======
	"os/exec"
	"strings"
)

// findExternalTool returns the file path of a tool that supplies
>>>>>>> 79bfea2d (update vendor)
=======
	exec "golang.org/x/sys/execabs"
	"os"
	"strings"
)

// The Driver Protocol
//
// The driver, given the inputs to a call to Load, returns metadata about the packages specified.
// This allows for different build systems to support go/packages by telling go/packages how the
// packages' source is organized.
// The driver is a binary, either specified by the GOPACKAGESDRIVER environment variable or in
// the path as gopackagesdriver. It's given the inputs to load in its argv. See the package
// documentation in doc.go for the full description of the patterns that need to be supported.
// A driver receives as a JSON-serialized driverRequest struct in standard input and will
// produce a JSON-serialized driverResponse (see definition in packages.go) in its standard output.

// driverRequest is used to provide the portion of Load's Config that is needed by a driver.
type driverRequest struct {
	Mode LoadMode `json:"mode"`
	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`
	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`
	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`
	// Overlay maps file paths (relative to the driver's working directory) to the byte contents
	// of overlay files.
	Overlay map[string][]byte `json:"overlay"`
}

// findExternalDriver returns the file path of a tool that supplies
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
// the build system package structure, or "" if not found."
// If GOPACKAGESDRIVER is set in the environment findExternalTool returns its
// value, otherwise it searches for a binary named gopackagesdriver on the PATH.
func findExternalDriver(cfg *Config) driver {
	const toolPrefix = "GOPACKAGESDRIVER="
	tool := ""
	for _, env := range cfg.Env {
		if val := strings.TrimPrefix(env, toolPrefix); val != env {
			tool = val
		}
	}
	if tool != "" && tool == "off" {
		return nil
	}
	if tool == "" {
		var err error
		tool, err = exec.LookPath("gopackagesdriver")
		if err != nil {
			return nil
		}
	}
	return func(cfg *Config, words ...string) (*driverResponse, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
		req, err := json.Marshal(driverRequest{
			Mode:       cfg.Mode,
			Env:        cfg.Env,
			BuildFlags: cfg.BuildFlags,
			Tests:      cfg.Tests,
			Overlay:    cfg.Overlay,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to encode message to driver tool: %v", err)
<<<<<<< HEAD
<<<<<<< HEAD
		}

		buf := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd := exec.CommandContext(cfg.Context, tool, words...)
		cmd.Dir = cfg.Dir
		cmd.Env = cfg.Env
		cmd.Stdin = bytes.NewReader(req)
		cmd.Stdout = buf
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%v: %v: %s", tool, err, cmd.Stderr)
		}
		if len(stderr.Bytes()) != 0 && os.Getenv("GOPACKAGESPRINTDRIVERERRORS") != "" {
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd), stderr)
		}

=======
		buf := new(bytes.Buffer)
		fullargs := []string{
			"list",
			fmt.Sprintf("-test=%t", cfg.Tests),
			fmt.Sprintf("-export=%t", usesExportData(cfg)),
			fmt.Sprintf("-deps=%t", cfg.Mode >= LoadImports),
		}
		for _, f := range cfg.BuildFlags {
			fullargs = append(fullargs, fmt.Sprintf("-buildflag=%v", f))
=======
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
		}

		buf := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd := exec.CommandContext(cfg.Context, tool, words...)
		cmd.Dir = cfg.Dir
		cmd.Env = cfg.Env
		cmd.Stdin = bytes.NewReader(req)
		cmd.Stdout = buf
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%v: %v: %s", tool, err, cmd.Stderr)
		}
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 79bfea2d (update vendor)
=======
=======
>>>>>>> 03397665 (update api)
		if len(stderr.Bytes()) != 0 && os.Getenv("GOPACKAGESPRINTDRIVERERRORS") != "" {
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd), stderr)
		}

<<<<<<< HEAD
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
>>>>>>> 03397665 (update api)
		var response driverResponse
		if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
			return nil, err
		}
		return &response, nil
	}
}
