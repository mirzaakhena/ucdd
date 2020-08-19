package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func databasePart(tp model.ThePackage) {

	dir := fmt.Sprintf("../../../../%s/backend/shared/database", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
		templateFile := "../templates/backend/shared/database/contract._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/contract.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/database/implementation._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/implementation.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/database/public._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/database/public.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}

}
