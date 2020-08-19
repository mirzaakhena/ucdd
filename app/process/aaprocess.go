package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
	"gopkg.in/yaml.v2"
)

// RunProcess is
func RunProcess(file string) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// prepare root object
	tp := model.ThePackage{}

	// read yaml file
	if err = yaml.Unmarshal(content, &tp); err != nil {
		log.Fatalf("error: %+v", err)
	}

	// set the application name from package path
	{
		s := strings.Split(tp.PackagePath, "/")
		tp.ApplicationName = s[len(s)-1]
	}

	for _, usecase := range tp.Usecases {

		if tp.IgnoreAll && !usecase.PartialUpdateThis {
			continue
		}

		// usecase
		{
			inportPart(tp, usecase)
			outportPart(tp, usecase)
			interactorPart(tp, usecase)
		}

		// outport implementation and inport invoker
		{
			controllerPart(tp, usecase)
			datasourcePart(tp, usecase)
		}

		// shared
		{
			databasePart(tp)
			logPart(tp)
			messagebrokerPart(tp)
			configPart(tp)
		}

		// binder component
		{
			binderPart(tp)
		}

		// entry point
		{
			mainPart(tp)
		}

		generateMock(tp.PackagePath, usecase.Name)

	}

	goFormat(tp.PackagePath)

	goGet()

}

func generateMock(packagePath, usecaseName string) {
	fmt.Printf("mockery %s", usecaseName)
	cmd := exec.Command("mockery", "-all", "-case", "snake", "-output", fmt.Sprintf("../../../../%s/backend/datasource/mocks/", packagePath), "-dir", fmt.Sprintf("../../../../%s/backend/usecase/%s/outport/", packagePath, strings.ToLower(usecaseName)))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// do go format
func goFormat(path string) {
	fmt.Println("go fmt")
	cmd := exec.Command("go", "fmt", fmt.Sprintf("%s/...", path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// call go get
func goGet() {
	fmt.Println("go get")
	cmd := exec.Command("go", "get", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
