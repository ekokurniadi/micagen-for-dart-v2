package generator_v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFeatureName(featureName string) {
	writeFolder(featureName)
}

func writeFolder(featureName string) error {
	var realFeatureName []string
	splitText := strings.Split(featureName, " ")

	for a := 0; a < len(splitText); a++ {
		realFeatureName = append(realFeatureName, strings.ToLower(splitText[a]))
	}
	realName := strings.Join(realFeatureName, "_")
	root := filepath.Join("./", "features")
	err := os.MkdirAll(root, 0775)

	if err != nil {
		fmt.Println("your directory is already exist but it's ok")
	}
	path := filepath.Join(root, "/", realName)

	err = os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("your directory is already exist but it's ok")
	}

	var childCoreFolder = []string{"error", "usecases", "helpers", "constant", "extensions", "module"}
	var baseFolder = []string{"data", "domain", "presentations"}
	var dataChildFolder = []string{"datasources", "models", "repositories"}
	var dataSourceChildFolder = []string{"local", "remote"}
	var domainChildFolder = []string{"entities", "repositories", "usecases"}
	var presentationChildFolder = []string{"bloc", "pages", "widgets"}
	var pagesChildFolder = []string{"tablet", "mobile"}

	core := filepath.Join("./", "core")
	os.MkdirAll(core, 0755)

	for x := 0; x < len(childCoreFolder); x++ {
		pathCore := filepath.Join("./core/", childCoreFolder[x])
		os.MkdirAll(pathCore, 0755)
	}

	for a := 0; a < len(baseFolder); a++ {
		basePath := filepath.Join(path, baseFolder[a])
		os.MkdirAll(basePath, 0755)
		for b := 0; b < len(dataChildFolder); b++ {
			dataChildFolderPath := filepath.Join(path+"/"+baseFolder[0], dataChildFolder[b])
			os.MkdirAll(dataChildFolderPath, 0755)

		}

		for ds := 0; ds < len(dataSourceChildFolder); ds++ {
			dataSourceChildFolderPath := filepath.Join(path+"/"+baseFolder[0], dataChildFolder[0], dataSourceChildFolder[ds])
			os.MkdirAll(dataSourceChildFolderPath, 0755)
		}

		for c := 0; c < len(domainChildFolder); c++ {
			domainChildFolderPath := filepath.Join(path+"/"+baseFolder[1], domainChildFolder[c])
			os.MkdirAll(domainChildFolderPath, 0755)

		}
		for d := 0; d < len(presentationChildFolder); d++ {
			presentationChildFolderPath := filepath.Join(path+"/"+baseFolder[2], presentationChildFolder[d])
			os.MkdirAll(presentationChildFolderPath, 0755)

		}

	}
	for x := 0; x < len(pagesChildFolder); x++ {
		pagesChildFolderPath := filepath.Join(path+"/"+presentationChildFolder[1], pagesChildFolder[x])
		os.MkdirAll(pagesChildFolderPath, 0755)
	}

	return nil
}
