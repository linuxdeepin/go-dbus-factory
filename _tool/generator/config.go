package main

import (
	"encoding/json"
	"io/ioutil"
)

type ServiceConfig struct {
	Service string
	Objects []*ObjectConfig
}

type ObjectConfig struct {
	Type string
	Path string // optional

	Interfaces []*InterfaceConfig
}

func (oc *ObjectConfig) getInterface(name string) *InterfaceConfig {
	for _, ifc := range oc.Interfaces {
		if ifc.Name == name {
			return ifc
		}
	}
	return nil
}

type InterfaceConfig struct {
	Name     string
	Type     string
	Accessor string
	Fixes    map[string]json.RawMessage
}

func (ic *InterfaceConfig) getPropertyFix(name string) *PropertyFix {
	rawMsg, ok := ic.Fixes["p/"+name]
	if !ok {
		return nil
	}
	var propFix PropertyFix
	err := json.Unmarshal(rawMsg, &propFix)
	if err != nil {
		return nil
	}
	return &propFix
}

func (ic *InterfaceConfig) getArgFixes(name string) ArgFixes {
	rawMsg, ok := ic.Fixes[name]
	if !ok {
		return nil
	}
	var argFixes ArgFixes
	err := json.Unmarshal(rawMsg, &argFixes)
	if err != nil {
		return nil
	}
	return argFixes
}

type ArgFixes []ArgFix

func (fixes ArgFixes) get(name string) ArgFix {
	for _, fix := range fixes {
		if fix.Name == name {
			return fix
		}
	}
	return ArgFix{}
}

type ArgFix struct {
	Name string
	Type string
}

type PropertyFix struct {
	Type       string
	ValueType  string
	EmptyValue string
}

func loadConfig(file string) (*ServiceConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var s ServiceConfig
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
