package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"pkg.deepin.io/lib/dbus1/introspect"
)

type ServiceConfig struct {
	Service string // optional
	Objects []*ObjectConfig
}

type ObjectConfig struct {
	Type       string
	Path       string   // optional
	XMLFile    string   // optional
	XMLFiles   []string // optional
	Interfaces []*InterfaceConfig
}

func (oc *ObjectConfig) getXmlFiles(dir string) []string {
	var result []string
	if len(oc.XMLFiles) > 0 {
		for _, pat := range oc.XMLFiles {
			matches, _ := filepath.Glob(filepath.Join(dir, pat+".xml"))
			result = append(result, matches...)
		}

	} else if oc.XMLFile != "" {
		result = append(result, filepath.Join(dir, oc.XMLFile+".xml"))
	} else {
		result = append(result, filepath.Join(dir, oc.Type+".xml"))
	}
	return result
}

func (oc *ObjectConfig) loadXml(dir string) ([]introspect.Interface, error) {
	var result []introspect.Interface
	xmlFiles := oc.getXmlFiles(dir)
	for _, file := range xmlFiles {
		interfaces, err := getInterfacesFromXmlFile(file)
		log.Println("load xml file:", file)
		if err != nil {
			return nil, err
		}
		result = append(result, interfaces...)
	}
	return result, nil
}

func getInterfacesFromXmlFile(file string) ([]introspect.Interface, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	var node introspect.Node
	err = dec.Decode(&node)
	if err != nil {
		return nil, err
	}
	return node.Interfaces, nil
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
	Name        string
	Type        string
	Accessor    string
	Fixes       map[string]json.RawMessage
	TypeDefined bool
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
	RenameTo   string
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
