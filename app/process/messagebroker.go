package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func messagebrokerPart(tp model.ThePackage) {
	dir := fmt.Sprintf("../../../../%s/backend/shared/messagebroker", tp.PackagePath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	{
		templateFile := "../templates/backend/shared/messagebroker/contract._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/contract.go", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/messagebroker/public._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/public.go", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/messagebroker/implementation-consumer._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-consumer.go", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/messagebroker/implementation-producer._go"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-producer.go", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := "../templates/backend/shared/messagebroker/docker-compose._yml"
		outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/docker-compose.yml", tp.PackagePath)
		basic(&tp, templateFile, outputFile, tp, 0664)
	}
}
