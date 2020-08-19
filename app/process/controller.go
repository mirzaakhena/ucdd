package process

import (
	"fmt"
	"os"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
)

func controllerPart(tp model.ThePackage, usecase model.TheUsecase) {
	dir := fmt.Sprintf("../../../../%s/backend/controller/restapi", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
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

	{
		templateFile := "../templates/backend/controller/consumer/consumer._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/controller/consumer/%s.go", tp.PackagePath, strings.ToLower(usecase.Name))
		basic(&tp, templateFile, outputFile, usecase, 0664)
	}
}
