package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mirzaakhena/common/utils"
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

	fmt.Printf("%v\n", utils.GetJSON(tp))

	// single directory generated
	{
		// makeBackendDirectory(&tp)
		// makeFrontendDirectory(&tp)
	}

	// // single static file generated
	// {
	// 	createStaticBackendFile(&tp)
	// 	createStaticFrontendFile(&tp)
	// }

	// // multiple generated file for entity
	// for _, et := range tp.Entities {
	// 	createDynamicBackendFile(&tp, &et)
	// 	createDynamicFrontendFile(&tp, &et)
	// }

	// // multiple generated file for enum
	// for _, en := range tp.Enums {
	// 	createEnumFile(&tp, &en)
	// }

	// fmt.Printf(">>>>> done Run\n")

	// goFormat(tp.PackagePath)

	// goGet()

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
	cmd := exec.Command("go", "get", fmt.Sprintf("./..."))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
