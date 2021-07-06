package resources

import (
	"embed"
)

var (
	//go:embed config/*
	confFS embed.FS

	//go:embed testdata/*
	testFS embed.FS
)

func read(fs embed.FS, path string) ([]byte, bool) {
	if b, err := fs.ReadFile(path); err != nil {
		return nil, false
	} else {
		return b, true
	}
}

func ReadTestFile(path string) ([]byte, bool) {
	return read(testFS, path)
}

func ReadConfFile(path string) ([]byte, bool) {
	return read(confFS, path)
}
