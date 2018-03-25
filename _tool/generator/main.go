package main

import (
	"encoding/xml"
	"os"

	"log"

	"github.com/davecgh/go-spew/spew"
	"pkg.deepin.io/lib/dbus1/introspect"
)

type ObjectConfig struct {
	ServiceName string
	Path        string // optional
	TypeName    string

	Interfaces map[string]*InterfaceConfig
	//                   ^interfaceName
}

type InterfaceConfig struct {
	TypeName           string
	ObjectAccessMethod string
}

func main() {
	node, err := ParseNode("data/com.deepin.dde.daemon.Dock/Dock.xml")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(node)

	// serviceName := "com.deepin.dde.daemon.Dock"
	// path := "/com/deepin/dde/daemon/Dock"
	// typeName := "DockObject"
	objCfg := &ObjectConfig{
		ServiceName: "com.deepin.dde.daemon.Dock",
		Path:        "/com/deepin/dde/daemon/Dock",
		TypeName:    "DockObject",
		Interfaces: map[string]*InterfaceConfig{
			"com.deepin.dde.daemon.Dock": {
				TypeName:           "dock",
				ObjectAccessMethod: "Dock",
			},
		},
	}
	// objCfg.Path = ""

	// DockObject

	// dock interface

	sf := NewSourceFile("libdock")

	sf.AddGoImport("pkg.deepin.io/lib/dbus1")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil/client")

	sf.GoBody.Pn("type %s struct {", objCfg.TypeName)

	for _, ifc := range node.Interfaces {
		ifcCfg := objCfg.Interfaces[ifc.Name]
		if ifcCfg == nil {
			continue
		}
		sf.GoBody.Pn("%s // interface %s", ifcCfg.TypeName, ifc.Name)
	}

	sf.GoBody.Pn("client.Object")
	sf.GoBody.Pn("}\n")

	writeNewObject(sf.GoBody, objCfg)

	for _, ifc := range node.Interfaces {
		ifcCfg := objCfg.Interfaces[ifc.Name]
		if ifcCfg == nil {
			continue
		}

		writeObjectAccessMethod(sf.GoBody, ifc, objCfg)
	}
	sf.Print()
}

func writeNewObject(sb *SourceBody, cfg *ObjectConfig) {
	if cfg.Path != "" {
		sb.Pn("func New%s(conn *dbus.Conn) *%s {", cfg.TypeName, cfg.TypeName)
		sb.Pn("    obj := new(%s)", cfg.TypeName)
		sb.Pn("    obj.Object.Init_(conn, %q,%q)", cfg.ServiceName, cfg.Path)
	} else {
		sb.Pn("func New%s(conn *dbus.Conn, path dbus.ObjectPath) *%s {",
			cfg.TypeName, cfg.TypeName)
		sb.Pn("    obj := new(%s)", cfg.TypeName)
		sb.Pn("    obj.Object.Init_(conn, %q,path)", cfg.ServiceName)
	}
	sb.Pn("    return obj")
	sb.Pn("}\n")
}

func writeObjectAccessMethod(sb *SourceBody, ifc introspect.Interface, objCfg *ObjectConfig) {
	ifcCfg := objCfg.Interfaces[ifc.Name]
	sb.Pn("func (obj *%s) %s() *%s {", objCfg.TypeName, ifcCfg.ObjectAccessMethod,
		ifcCfg.TypeName)
	sb.Pn("    return &obj.%s", ifcCfg.TypeName)
	sb.Pn("}\n")
}

func ParseNode(file string) (*introspect.Node, error) {
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
	return &node, nil
}
