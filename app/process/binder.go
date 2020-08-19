package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func binderPart(tp model.ThePackage) {
	dir := fmt.Sprintf("../../../../%s/backend/binder", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
		templateFile := "../templates/backend/binder/runner._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/binder/runner.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/binder/setup_component._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/binder/setup_component.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/binder/wiring_component._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/binder/wiring_component.go", tp.PackagePath)
		writeFile(&tp, templateFile, outputFile, tp, 0664)
	}
}
