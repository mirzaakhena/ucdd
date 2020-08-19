package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func mainPart(tp model.ThePackage) {
	dir := fmt.Sprintf("../../../../%s/backend/app", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
		templateFile := "../templates/backend/app/main._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/app/main.go", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}
}
