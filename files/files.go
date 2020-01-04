package files

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func GetFiles(dir string, showHidden bool, search string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// filter by search
	if search != "" {
		filtered := files[:0]
		for _, file := range files {
			if strings.Index(strings.ToLower(file.Name()), strings.ToLower(search)) != -1 {
				filtered = append(filtered, file)
			}
		}
		files = filtered
	}

	// filter by hidden
	if !showHidden {
		notHidden := files[:0]
		for _, file := range files {
			if strings.Index(file.Name(), ".") != 0 {
				notHidden = append(notHidden, file)
			}
		}
		files = notHidden
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir() && !files[j].IsDir() {
			return true
		}
		if !files[i].IsDir() && files[j].IsDir() {
			return false
		}
		return strings.ToLower(files[i].Name()) < strings.ToLower(files[j].Name())
	})

	return files
}

func RemoveFile(path string, file os.FileInfo) error {
	return os.RemoveAll(path + "/" + file.Name())
}
