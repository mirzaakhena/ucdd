package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/ucdd/app/model"
)

func makeBackendDirectory(tp *model.ThePackage) {

	var dir string

	dir = fmt.Sprintf("../../../../%s/backend/app", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/controller/restapi", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/datasource", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/domain/domain-name/usecase/contract/inport", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/domain/domain-name/usecase/contract/outport", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/domain/domain-name/usecase/interactor", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/domain/domain-name/usecase/model", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	// ====================================================================================================

	dir = fmt.Sprintf("../../../../%s/backend/shared/utils", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/token", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/config", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/log", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/transaction", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/messagebroker", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/httpclient", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/constant", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/extractor", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/backend/shared/error", tp.PackagePath)
	os.MkdirAll(dir, 0777)

}
