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

		if usecase.Ignore {
			continue
		}

		// INPORT
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/inport", tp.PackagePath, strings.ToLower(usecase.Name))
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/usecase/usecasesname/inport/inport._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/inport/inport.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// OUTPORT
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/outport", tp.PackagePath, strings.ToLower(usecase.Name))
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/usecase/usecasesname/outport/outport._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/outport/outport.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// INTERACTOR
		{
			dir := fmt.Sprintf("../../../../%s/backend/usecase/%s/interactor", tp.PackagePath, strings.ToLower(usecase.Name))
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/usecase/usecasesname/interactor/interactor._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/interactor/interactor.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

			{
				templateFile := "../templates/backend/usecase/usecasesname/interactor/interactor_test._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/usecase/%s/interactor/normal_test.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}
		}

		// BINDER
		{
			dir := fmt.Sprintf("../../../../%s/backend/binder", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/binder/runner._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/binder/runner.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/binder/setup_component._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/binder/setup_component.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/binder/wiring_component._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/binder/wiring_component.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// CONTROLLER
		{
			dir := fmt.Sprintf("../../../../%s/backend/controller/restapi", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			if tp.Consumer {
				dir = fmt.Sprintf("../../../../%s/backend/controller/consumer", tp.PackagePath)
				if err := os.MkdirAll(dir, 0777); err != nil {
					panic(err)
				}
			}

			{
				templateFile := "../templates/backend/controller/restapi/restapi._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/%s.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

			// {
			// 	templateFile := "../templates/backend/controller/restapi/request._http"
			// 	outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/request.http", tp.PackagePath)
			// 	basic(&tp, templateFile, outputFile, tp, 0664)
			// }

			if tp.Consumer {
				templateFile := "../templates/backend/controller/consumer/consumer._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/controller/consumer/%s.go", tp.PackagePath, strings.ToLower(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

		}

		// DOMAIN
		{
			dir := fmt.Sprintf("../../../../%s/backend/domain", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}
		}

		// REAL
		{
			dir := fmt.Sprintf("../../../../%s/backend/datasource/production", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/datasource/production/datasource._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/datasource/production/%s_outport.go", tp.PackagePath, SnakeCase(usecase.Name))
				basic(&tp, templateFile, outputFile, usecase, 0664)
			}

		}

		// MOCKS
		{
			dir := fmt.Sprintf("../../../../%s/backend/datasource/mocks", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}
		}

		// DATABASE
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/database", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/shared/database/contract._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/database/implementation._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/database/public._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}
		}

		// LOG
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/log", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/shared/log/contract._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/log/implementation._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/log/public._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// MESSAGE_BROKER
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/messagebroker", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/shared/messagebroker/contract._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/messagebroker/public._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/messagebroker/implementation-consumer._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-consumer.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/messagebroker/implementation-producer._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-producer.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/messagebroker/docker-compose._yml"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/docker-compose.yml", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// CONFIG
		{
			dir := fmt.Sprintf("../../../../%s/backend/shared/config", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/shared/config/contract._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/contract.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/config/implementation._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/implementation.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/shared/config/public._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/public.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := "../templates/backend/config._toml"
				outputFile := fmt.Sprintf("../../../../%s/backend/config.toml", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// MAIN
		{
			dir := fmt.Sprintf("../../../../%s/backend/app", tp.PackagePath)
			if err := os.MkdirAll(dir, 0777); err != nil {
				panic(err)
			}

			{
				templateFile := "../templates/backend/app/main._go"
				outputFile := fmt.Sprintf("../../../../%s/backend/app/main.go", tp.PackagePath)
				basic(&tp, templateFile, outputFile, tp, 0664)
			}

		}

		// generateMock(tp.PackagePath, usecase.Name)

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

	goFormat(tp.PackagePath)

	goGet()

	gitCommit()

}

func gitCommit() {
	x := "git --git-dir $GOPATH/src/github.com/mirzaakhena/accounting/.git log"
	runcmd(x, false)
}

func runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
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
