package process

import (
	"fmt"
	"os"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
)

func interactorPart(tp model.ThePackage, usecase model.TheUsecase) {
	if !(!tp.IgnoreAll || (tp.IgnoreAll && usecase.PartialUpdateInteractor)) {
		return
	}
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
