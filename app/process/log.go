package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func logPart(tp model.ThePackage) {

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
