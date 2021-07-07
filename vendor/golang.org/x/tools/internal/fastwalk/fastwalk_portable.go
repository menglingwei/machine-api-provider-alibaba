// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build appengine || (!linux && !darwin && !freebsd && !openbsd && !netbsd)
=======
>>>>>>> 79bfea2d (update vendor)
=======
//go:build appengine || (!linux && !darwin && !freebsd && !openbsd && !netbsd)
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
//go:build appengine || (!linux && !darwin && !freebsd && !openbsd && !netbsd)
>>>>>>> 03397665 (update api)
// +build appengine !linux,!darwin,!freebsd,!openbsd,!netbsd

package fastwalk

import (
	"io/ioutil"
	"os"
)

// readDir calls fn for each directory entry in dirName.
// It does not descend into directories or follow symlinks.
// If fn returns a non-nil error, readDir returns with that error
// immediately.
func readDir(dirName string, fn func(dirName, entName string, typ os.FileMode) error) error {
	fis, err := ioutil.ReadDir(dirName)
	if err != nil {
		return err
	}
	skipFiles := false
	for _, fi := range fis {
		if fi.Mode().IsRegular() && skipFiles {
			continue
		}
		if err := fn(dirName, fi.Name(), fi.Mode()&os.ModeType); err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			if err == ErrSkipFiles {
=======
			if err == SkipFiles {
>>>>>>> 79bfea2d (update vendor)
=======
			if err == ErrSkipFiles {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
			if err == ErrSkipFiles {
>>>>>>> 03397665 (update api)
				skipFiles = true
				continue
			}
			return err
		}
	}
	return nil
}
