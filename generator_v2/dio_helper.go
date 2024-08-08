package generator_v2

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ekokurniadi/micagen-for-dart-v2/schemas"
)

func GenerateDioHelper(project schemas.Project) {
	createFileDioHelper(project)
	// writeFileBuildYaml()
}

func createFileDioHelper(project schemas.Project) (string, error) {

	url := ""
	if project.Config.UseInjectable {
		url = "https://raw.githubusercontent.com/ekokurniadi/micagen-for-dart-v2/main/generator_v2/dio_helper_injectable.txt"
	} else {
		url = "https://raw.githubusercontent.com/ekokurniadi/micagen-for-dart-v2/main/generator_v2/dio_helper.txt"
	}

	filenames := "./core/helpers/dio_helper.dart"

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	file, err := os.Create(filenames)
	if err != nil {
		log.Fatal(err.Error())
		return file.Name(), err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Create %s is successfully \n", file.Name())
	return file.Name(), nil
}
