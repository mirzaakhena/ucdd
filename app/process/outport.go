package process

import (
	"fmt"
	"os"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
)

func outportPart(tp model.ThePackage, usecase model.TheUsecase) {
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
