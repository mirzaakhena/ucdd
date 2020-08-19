package model

// ThePackage is root of application definition
type ThePackage struct {
	ApplicationName string       ``                   // name of application
	PackagePath     string       `yaml:"packagePath"` // golang path of application
	Usecases        []TheUsecase `yaml:"usecases"`    // list of usecases
	IgnoreAll       bool         `yaml:"ignoreAll"`   //
}

// TheUsecase is
type TheUsecase struct {
	Name                    string       `yaml:"name"`                    //
	Inport                  TheInport    `yaml:"inport"`                  //
	Outports                []TheOutport `yaml:"outports"`                //
	DefaultDatasource       string       `yaml:"defaultDatasource"`       //
	Datasources             []Datasource `yaml:"datasources"`             //
	Controllers             []Controller `yaml:"controllers"`             //
	PartialUpdateThis       bool         `yaml:"partialUpdateThis"`       //
	PartialUpdateBinder     bool         `yaml:"partialUpdateBinder"`     //
	PartialUpdateInteractor bool         `yaml:"partialUpdateInteractor"` //
	PartialUpdateDatasource bool         `yaml:"partialUpdateDatasource"` //
	PartialUpdateController bool         `yaml:"partialUpdateController"` //
}

// TheInport is
type TheInport struct {
	PartialUpdateThis bool     `yaml:"partialUpdateThis"` //
	RequestFields     []string `yaml:"requestFields"`     //
	ResponseFields    []string `yaml:"responseFields"`    //
}

// TheOutport is
type TheOutport struct {
	Name              string   `yaml:"name"`              //
	PartialUpdateThis bool     `yaml:"partialUpdateThis"` //
	RequestFields     []string `yaml:"requestFields"`     //
	ResponseFields    []string `yaml:"responseFields"`    //
}

// Datasource is
type Datasource struct {
	Name              string `yaml:"name"`              //
	PartialUpdateThis bool   `yaml:"partialUpdateThis"` //
}

// Controller is
type Controller struct {
	Type              string `yaml:"name"`              //
	PartialUpdateThis bool   `yaml:"partialUpdateThis"` //
	Framework         string `yaml:"framework"`         //
}
