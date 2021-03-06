package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func datasourcePart(tp model.ThePackage, usecase model.TheUsecase) {

	dir := fmt.Sprintf("../../../../%s/backend/datasource/production", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
		templateFile := "../templates/backend/datasource/production/datasource._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/datasource/production/%s_outport.go", tp.PackagePath, SnakeCase(usecase.Name))
		writeFile(&tp, templateFile, outputFile, usecase, 0664)
	}

	{
		dir := fmt.Sprintf("../../../../%s/backend/datasource/mocks", tp.PackagePath)
		if err := os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}
	}

	{
		// mocks file will generated by mockery
	}

}
