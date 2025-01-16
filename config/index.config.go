package config

import (
	"path/filepath"
	"runtime"
)

var (
	//Get Current file full path runtime

	_, b, _, _ = runtime.Caller(0)

	//Root Folder of this project

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)
