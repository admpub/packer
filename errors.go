package packer

import "errors"

var ErrNotFound = errors.New(`no package manager found`)
var ErrUnsuppored = errors.New(`not supported`)
