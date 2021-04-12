package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/godbus/dbus/introspect"
)

type ServiceConfig struct {
	Service       string // optional
	Objects       []*ObjectConfig
	PropertyTypes []*PropertyTypeConfig
}

func (sc *ServiceConfig) autoFill() {
	for _, objCfg := range sc.Objects {
		objCfg.autoFill()
	}
	for _, propTypeCfg := range sc.PropertyTypes {
		propTypeCfg.autoFill("", nil)
	}
}

type PropertyTypeConfig struct {
	Type       string
	ValueType  string
	EmptyValue string
}

func (ptc *PropertyTypeConfig) autoFill(propName string, ifcCfg *InterfaceConfig) {
	if ptc.Type == "" &&
		ifcCfg != nil &&
		ifcCfg.Accessor != "" &&
		propName != "" {
		ptc.Type = "Prop" + ifcCfg.Accessor + propName
	}

	if ptc.EmptyValue == "" {
		if strings.HasPrefix(ptc.ValueType, "map[") ||
			strings.HasPrefix(ptc.ValueType, "[]") {
			ptc.EmptyValue = "nil"
		} else {
			ptc.EmptyValue = ptc.ValueType + "{}"
		}
	} else {
		ptc.EmptyValue = strings.Replace(ptc.EmptyValue, "$T",
			ptc.ValueType, 1)
	}
}

type ObjectConfig struct {
	Type       string
	Path       string   // optional
	XMLFile    string   // optional
	XMLFiles   []string // optional
	Interfaces []*InterfaceConfig
}

func (oc *ObjectConfig) autoFill() {
	for _, ifcCfg := range oc.Interfaces {
		ifcCfg.autoFill()
	}
}

func (oc *ObjectConfig) getXmlFiles(dir string) []string {
	var result []string
	if len(oc.XMLFiles) > 0 {
		for _, pat := range oc.XMLFiles {
			matches, _ := filepath.Glob(filepath.Join(dir, pat+".xml"))
			result = append(result, matches...)
		}

	} else if oc.XMLFile != "" {
		if oc.XMLFile == "-" {
			return nil
		}
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
	Name                string
	Type                string
	Accessor            string
	Fixes               map[string]json.RawMessage
	TypeDefined         bool
	CustomInterfaceName bool
	NoGetInterfaceName  bool

	MethodFixes map[string]ArgFixes
	PropFixes   map[string]*PropertyFix
	SignalFixes map[string]ArgFixes
}

func (ic *InterfaceConfig) autoFill() {
	dotRIdx := strings.LastIndex(ic.Name, ".")
	if dotRIdx == -1 {
		panic(fmt.Sprintf("not found dot in interface name %q", ic.Name))
	}
	if ic.Accessor == "" {
		ic.Accessor = strings.Title(ic.Name[dotRIdx+1:])
	}

	if ic.Type == "" && ic.Accessor != "" {
		firstLetter := strings.ToLower(string(ic.Accessor[0]))
		ic.Type = firstLetter + ic.Accessor[1:]
	}

	for key, fix := range ic.Fixes {
		if strings.HasPrefix(key, "p/") {
			propName := strings.TrimPrefix(key, "p/")
			var propFix PropertyFix
			err := json.Unmarshal(fix, &propFix)
			if err != nil {
				panic(err)
			}
			if ic.PropFixes == nil {
				ic.PropFixes = make(map[string]*PropertyFix)
			}
			ic.PropFixes[propName] = &propFix

		} else if strings.HasPrefix(key, "s/") {
			signalName := strings.TrimPrefix(key, "s/")
			var argFixes ArgFixes
			err := json.Unmarshal(fix, &argFixes)
			if err != nil {
				panic(err)
			}
			if ic.SignalFixes == nil {
				ic.SignalFixes = make(map[string]ArgFixes)
			}
			ic.SignalFixes[signalName] = argFixes

		} else if strings.HasPrefix(key, "m/") {
			methodName := strings.TrimPrefix(key, "m/")
			var argFixes ArgFixes
			err := json.Unmarshal(fix, &argFixes)
			if err != nil {
				panic(err)
			}
			if ic.MethodFixes == nil {
				ic.MethodFixes = make(map[string]ArgFixes)
			}
			ic.MethodFixes[methodName] = argFixes
		}
	}
	for propName, propFix := range ic.PropFixes {
		propFix.autoFill(propName, ic)
	}
}

func (ic *InterfaceConfig) getPropertyFix(name string) *PropertyFix {
	v := ic.PropFixes[name]
	return v
}

func (ic *InterfaceConfig) getMethodFix(name string) ArgFixes {
	v := ic.MethodFixes[name]
	return v
}

func (ic *InterfaceConfig) getSignalFix(name string) ArgFixes {
	v := ic.SignalFixes[name]
	return v
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
	PropertyTypeConfig
	RenameTo string
	RefType  string
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
