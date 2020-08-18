package model

// ThePackage is root of application definition
type ThePackage struct {
	ApplicationName string       ``                   // name of application
	PackagePath     string       `yaml:"packagePath"` // golang path of application
	Usecases        []TheUsecase `yaml:"usecases"`    // list of usecases
	RestAPI         bool         `yaml:"restapi"`
	Consumer        bool         `yaml:"consumer"`
}

// TheUsecase is
type TheUsecase struct {
	Name             string       `yaml:"name"`             //
	Inport           TheInport    `yaml:"inport"`           //
	Outports         []TheOutport `yaml:"outports"`         //
	Ignore           bool         `yaml:"ignore"`           //
	IgnoreBinder     bool         `yaml:"ignoreBinder"`     //
	IgnoreInteractor bool         `yaml:"ignoreInteractor"` //
	IgnoreDatasource bool         `yaml:"ignoreDatasource"` //
	IgnoreController bool         `yaml:"ignoreController"` //
}

// TheInport is
type TheInport struct {
	Ignore         bool     `yaml:"ignore"`         //
	RequestFields  []string `yaml:"requestFields"`  //
	ResponseFields []string `yaml:"responseFields"` //
}

// TheOutport is
type TheOutport struct {
	Name           string   `yaml:"name"`           //
	Ignore         bool     `yaml:"ignore"`         //
	RequestFields  []string `yaml:"requestFields"`  //
	ResponseFields []string `yaml:"responseFields"` //
}
