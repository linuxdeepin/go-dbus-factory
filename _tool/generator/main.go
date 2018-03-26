package main

import (
	"encoding/xml"
	"os"

	"log"

	"fmt"

	"bytes"

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

	sf.AddGoImport("unsafe")
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
		writeImplementerMethods(sf.GoBody, ifc, ifcCfg)

		for _, method := range ifc.Methods {
			writeMethod(sf.GoBody, method, ifcCfg)
		}
	}
	sf.Print()
	sf.Save("./auto.go")
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

func writeImplementerMethods(sb *SourceBody, ifc introspect.Interface, ifcCfg *InterfaceConfig) {
	sb.Pn("type %s struct{}", ifcCfg.TypeName)

	sb.Pn("func (v *%s) GetObject_() *client.Object {", ifcCfg.TypeName)
	sb.Pn("    return (*client.Object)(unsafe.Pointer(v))")
	sb.Pn("}\n")

	sb.Pn("func (*%s) GetInterfaceName_() string {", ifcCfg.TypeName)
	sb.Pn("    return %q", ifc.Name)
	sb.Pn("}\n")
}

func writeMethod(sb *SourceBody, method introspect.Method, ifcCfg *InterfaceConfig) {
	args := getArgs(method.Args)
	inArgs := getInArgs(args)
	outArgs := getOutArgs(args)
	// sb.Pn("// in %#v", inArgs)
	// sb.Pn("// out %#v", outArgs)

	// GoXXX

	sb.Pn("func (v *%s) Go%s(flags dbus.Flags, ch chan *dbus.Call %s) *dbus.Call {",
		ifcCfg.TypeName, method.Name, getArgsProto(inArgs, false))
	sb.Pn("    return v.GetObject_().Go_(v.GetInterfaceName_()+\".%s\", flags, ch %s)",
		method.Name, getArgsName(inArgs, false))
	sb.Pn("}\n")

	// StoreXXX
	if len(outArgs) > 0 {
		sb.Pn("func (*%s) Store%s(call *dbus.Call) (%s,err error) {", ifcCfg.TypeName,
			method.Name, getArgsProto(outArgs, true))
		sb.Pn("    err = call.Store(%s)", getArgsRefName(outArgs, true))
		sb.Pn("    return")
		sb.Pn("}\n")
	}

	// Call
	if len(outArgs) == 0 {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) error {",
			ifcCfg.TypeName, method.Name, getArgsProto(inArgs, false))
		sb.Pn("return (<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done).Err",
			method.Name, getArgsName(inArgs, false))
		sb.Pn("}\n")
	} else {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) (%s, err error) {",
			ifcCfg.TypeName, method.Name, getArgsProto(inArgs, false),
			getArgsProto(outArgs, true))
		sb.Pn("return v.Store%s(", method.Name)
		sb.Pn("<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done)",
			method.Name, getArgsName(inArgs, false))
		sb.Pn("}\n")
	}
}

func getArgs(args []introspect.Arg) []introspect.Arg {
	argNameIdx := 0
	var ret []introspect.Arg
	fixName := false
	for idx, arg := range args {

		if idx == 0 {
			if arg.Name == "" {
				fixName = true
			}
		}

		arg0 := arg
		// fix arg name
		if fixName {
			arg0.Name = fmt.Sprintf("arg%d", argNameIdx)
			argNameIdx++
		}
		ret = append(ret, arg0)
	}
	return ret
}

func getInArgs(args []introspect.Arg) []introspect.Arg {
	var ret []introspect.Arg
	for _, arg := range args {
		if arg.Direction != "out" {
			ret = append(ret, arg)
		}
	}
	return ret
}

func getOutArgs(args []introspect.Arg) []introspect.Arg {
	var ret []introspect.Arg
	for _, arg := range args {
		if arg.Direction == "out" {
			ret = append(ret, arg)
		}
	}
	return ret
}

func getArgsProto(args []introspect.Arg, begin bool) string {
	var buf bytes.Buffer

	for _, arg := range args {
		argType := TypeFor(arg.Type).String()
		buf.WriteString(fmt.Sprintf(",%s %s", arg.Name, argType))
	}

	return fixList(buf.String(), begin)
}

func getArgsName(args []introspect.Arg, begin bool) string {
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(fmt.Sprintf(",%s", arg.Name))
	}
	return fixList(buf.String(), begin)
}

func getArgsRefName(args []introspect.Arg, begin bool) string {
	var buf bytes.Buffer
	for _, arg := range args {
		buf.WriteString(fmt.Sprintf(",&%s", arg.Name))
	}
	return fixList(buf.String(), begin)
}

func fixList(str string, begin bool) string {
	if str == "" || !begin {
		return str
	}

	return str[1:]
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
