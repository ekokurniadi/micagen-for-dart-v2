package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ekokurniadi/micagen-for-dart-v2/schemas"
)

func GenerateCoreFailures(project schemas.Project) {
	createFileCoreFailures(project)
	writeFileCoreFailures(project)
}

func writeFileCoreFailures(project schemas.Project) (string, error) {
	filepath, err := filepath.Abs("./core/error/failures.dart")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//Write some text line-by-line to file.
	_, err = file.WriteString("import 'package:equatable/equatable.dart';\n\nabstract class Failures extends Equatable {\n\tString get errorMessage;\n\t@override\n\tList<Object> get props => [errorMessage];\n}\n\nclass ServerFailure extends Failures {\n\t@override\n\t\n\tfinal String errorMessage;\n\n\tServerFailure({required this.errorMessage});\n\n\t@override\n\tList<Object> get props => [errorMessage];\n}\n\nclass DatabaseFailure extends Failures {\n\t@override\n\tfinal String errorMessage;\n\n\tDatabaseFailure({required this.errorMessage});\n\n\t@override\n\tList<Object> get props => [errorMessage];\n\n}")

	if isError(err) {
		return "", err
	}

	return "", nil

}

func createFileCoreFailures(project schemas.Project) (string, error) {
	filepath, err := filepath.Abs("./core/error/failures.dart")

	if err != nil {
		log.Fatal("error")
		return filepath, err
	}
	_, err = os.Stat(filepath)
	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return "", err
	}

	filename, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err.Error())
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}
