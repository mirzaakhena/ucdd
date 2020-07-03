package process

import (
	"fmt"
	"strings"

	"github.com/mirzaakhena/ucdd/app/model"
)

func createDynamicBackendFile(tp *model.ThePackage, et *model.TheDomain) {

	{
		templateFile := fmt.Sprintf("../templates/backend/repository/repository._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/repository/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/controller/restapi/controller._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/model/model._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/model/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/service/crud/service._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/service/crud/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

}
