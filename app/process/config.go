package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func configPart(tp model.ThePackage) {

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
