package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	lowerPath := strings.ToLower(path)
	if strings.HasSuffix(lowerPath, ".jar") || strings.HasSuffix(lowerPath, ".zip") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}