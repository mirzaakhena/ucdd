package model

// ThePackage is root of application definition
type ThePackage struct {
	ApplicationName string       ``                   // name of application
	PackagePath     string       `yaml:"packagePath"` // golang path of application
	Usecases        []TheUsecase `yaml:"usecases"`    // list of usecases
	RestAPI         bool         `yaml:"restapi"`
	MessageBroker   bool         `yaml:"messagebroker"`
}

// TheUsecase is
type TheUsecase struct {
	Name     string       `yaml:"name"`     // name of usecase
	Outports []TheOutport `yaml:"outports"` // name of usecase
	Inport   TheInport    `yaml:"inport"`   // name of usecase
	Ignore   bool         `yaml:"ignore"`
}

// TheInport is
type TheInport struct {
	RequestFields  []string `yaml:"requestFields"`  //
	ResponseFields []string `yaml:"responseFields"` //
}

// TheOutport is
type TheOutport struct {
	Name           string   `yaml:"name"`           // name of outport
	RequestFields  []string `yaml:"requestFields"`  //
	ResponseFields []string `yaml:"responseFields"` //
}
