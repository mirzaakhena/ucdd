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

	// fmt.Printf("%v\n", utils.GetJSON(tp))

	for _, usecase := range tp.Usecases {

		// INPORT
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/inport", tp.PackagePath, strings.ToLower(usecase.Name))
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/usecase/usecasesname/inport/inport._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/inport/inport.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// OUTPORT
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/outport", tp.PackagePath, strings.ToLower(usecase.Name))
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/usecase/usecasesname/outport/outport._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/outport/outport.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// INTERACTOR
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/interactor", tp.PackagePath, strings.ToLower(usecase.Name))
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/usecase/usecasesname/interactor/interactor._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/interactor/%s.go", tp.PackagePath, strings.ToLower(usecase.Name), strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// CONTROLLER
		{
			dir := fmt.Sprintf("../../../../%s/backend/controller/restapi", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/controller/restapi/restapi._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/%s.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/controller/restapi/request._http")
				outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/request.http", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// DOMAIN
		{
			dir := fmt.Sprintf("../../../../%s/backend/domain", tp.PackagePath)
			os.MkdirAll(dir, 0777)
		}

		// REAL
		{
			dir := fmt.Sprintf("../../../../%s/backend/datasource/real", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/datasource/real/datasource._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/datasource/real/%s_outport.go", tp.PackagePath, SnakeCase(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

		}

		// MOCKS
		{
			dir := fmt.Sprintf("../../../../%s/backend/datasource/mocks", tp.PackagePath)
			os.MkdirAll(dir, 0777)
		}

		// TRANSACTION
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/transaction", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/transaction/contract._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/transaction/implementation._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/transaction/public._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}
		}

		// LOG
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/log", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/log/contract._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/log/implementation._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/log/public._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// CONFIG
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/config", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/config/contract._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/config/implementation._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/shared/config/public._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/config._toml")
				outputFile := fmt.Sprintf("../../../../%s/backend/config.toml", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// MAIN
		{
			dir := fmt.Sprintf("../../../../%s/backend/app", tp.PackagePath)
			os.MkdirAll(dir, 0777)

			{
				templateFile := fmt.Sprintf("../templates/backend/app/main._go")
				outputFile := fmt.Sprintf("../../../../%s/backend/app/main.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

	}

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
