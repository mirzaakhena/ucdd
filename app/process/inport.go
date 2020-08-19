package process

import (
	"fmt"
	"os"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
)

func inportPart(tp model.ThePackage, usecase model.TheUsecase) {
	if !(!tp.IgnoreAll || (tp.IgnoreAll && usecase.Inport.PartialUpdateThis)) {
		return
	}

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
