package process

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
	"unicode"

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

func writeFile(pkg *model.ThePackage, templateFile, outputFilePath string, object interface{}, perm os.FileMode) {

	fmt.Println(outputFilePath)

	var buffer bytes.Buffer

	// this process can be refactor later
	{
		// open template file
		file, err := os.Open(templateFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			row := scanner.Text()
			buffer.WriteString(row)
			buffer.WriteString("\n")
		}
	}

	// function that used in templates
	funcMap := template.FuncMap{
		"CamelCase":   CamelCase,
		"PascalCase":  PascalCase,
		"SnakeCase":   SnakeCase,
		"UpperCase":   UpperCase,
		"LowerCase":   LowerCase,
		"AppName":     func() string { return pkg.ApplicationName },
		"PackagePath": func() string { return pkg.PackagePath },
	}

	// bind func and template
	t := template.Must(template.New("codes").Funcs(funcMap).Parse(buffer.String()))

	// prepare output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	defer outputFile.Close()

	// write template into the file
	if err := t.Execute(outputFile, object); err != nil {
		panic(err)
	}

}

// CamelCase is
func CamelCase(name string) string {

	// force it!
	// this is bad. But we can figure out later
	{
		if name == "IPAddress" {
			return "ipAddress"
		}

		if name == "ID" {
			return "id"
		}
	}

	out := []rune(name)
	out[0] = unicode.ToLower([]rune(name)[0])
	return string(out)
}

// UpperCase is
func UpperCase(name string) string {
	return strings.ToUpper(name)
}

// LowerCase is
func LowerCase(name string) string {
	return strings.ToLower(name)
}

// PascalCase is
func PascalCase(name string) string {
	return name
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// SnakeCase is
func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
