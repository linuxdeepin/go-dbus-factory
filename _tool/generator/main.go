package main

import (
	"encoding/xml"
	"os"

	"log"

	"fmt"

	"bytes"

	"strings"

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

	sf.AddGoImport("fmt")
	sf.AddGoImport("unsafe")
	sf.AddGoImport("pkg.deepin.io/lib/dbus1")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil")
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

		for _, signal := range ifc.Signals {
			writeSignal(sf.GoBody, signal, ifcCfg)
		}

		for _, prop := range ifc.Properties {
			writeProperty(sf.GoBody, prop, ifcCfg)
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
	sb.Pn("// method %s\n", method.Name)
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

func writeSignal(sb *SourceBody, signal introspect.Signal, ifcCfg *InterfaceConfig) {
	sb.Pn("// signal %s\n", signal.Name)

	args := getArgs(signal.Args)

	sb.Pn("func (v *%s) Connect%s(cb func(%s)) (dbusutil.SignalHandlerId, error) {",
		ifcCfg.TypeName, signal.Name, getArgsProto(args, true))
	sb.Pn("obj := v.GetObject_()")
	sb.Pn("rule := fmt.Sprintf(")
	sb.writeStr(`"type='signal',interface='%s',member='%s',path='%s',sender='%s'",` + "\n")
	sb.Pn("v.GetInterfaceName_(), %q, obj.Path_(), obj.ServiceName_())\n", signal.Name)

	sb.Pn("sigRule := &dbusutil.SignalRule{")
	sb.Pn("Path: obj.Path_(),")
	sb.Pn("Name: v.GetInterfaceName_() + \".%s\",", signal.Name)
	sb.Pn("}")
	sb.Pn("handlerFunc := func(sig *dbus.Signal) {")

	if len(args) > 0 {
		for _, arg := range args {
			sb.Pn("var %s %s", arg.Name, TypeFor(arg.Type).String())
		}
		sb.Pn("err := dbus.Store(sig.Body %s)", getArgsRefName(args, false))
		sb.Pn("if err == nil {")
		sb.Pn("    cb(%s)", getArgsName(args, true))
		sb.Pn("}")
	} else {
		sb.Pn("cb()")
	}

	sb.Pn("}\n") // end handlerFunc

	sb.Pn("return obj.ConnectSignal_(rule, sigRule, handlerFunc)")

	sb.Pn("}\n")
}

var propBaseTypeMap = map[string]string{
	"y": "Byte",
	"b": "Bool",
	"n": "Int16",
	"q": "Uint16",
	"i": "Int32",
	"u": "Uint32",
	"x": "Int64",
	"t": "Uint64",
	"d": "Double",
	"s": "String",
	"o": "ObjectPath",
}

func getPropType(ty string) string {
	if len(ty) == 1 {
		val := propBaseTypeMap[ty]
		if val == "" {
			return ""
		}
		return "client.Prop" + val
	} else if len(ty) == 2 && ty[0] == 'a' {
		val := propBaseTypeMap[ty[1:]]
		if val == "" {
			return ""
		}
		return "client.Prop" + val + "Array"
	}
	return ""
}

type PropConfig struct {
	Type         string
	GoType       string
	GoEmptyValue string
}

func writeProperty(sb *SourceBody, prop introspect.Property, ifcCfg *InterfaceConfig) {
	sb.Pn("// property %s %s\n", prop.Name, prop.Type)

	propType := getPropType(prop.Type)
	if propType != "" {
		sb.Pn("func (v *%s) %s() %s {", ifcCfg.TypeName, prop.Name, propType)
		sb.Pn("    return %s{", propType)
		sb.Pn("        Impl: v,")
		sb.Pn("        Name: %q,", prop.Name)
		sb.Pn("    }")

		sb.Pn("}\n")
	} else {
		//
		pc := PropConfig{
			Type:         "DockFrontendWindowRect",
			GoType:       "FrontendWindowRect",
			GoEmptyValue: "FrontendWindowRect{}",
		}
		sb.Pn("func (v *%s) %s() %s {", ifcCfg.TypeName, prop.Name, pc.Type)
		sb.Pn("    return %s{", pc.Type)
		sb.Pn("        Impl: v,")
		sb.Pn("    }")
		sb.Pn("}\n")

		sb.Pn("type %s struct {", pc.Type)
		sb.Pn("Impl client.Implementer")
		sb.Pn("}\n")

		// Get
		if strings.Contains(prop.Access, "read") {
			sb.Pn("func (p %s) Get(flags dbus.Flags) (value %s, err error) {",
				pc.Type, pc.GoType)
			sb.Pn("err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),")
			sb.Pn("%q, &value)", prop.Name)
			sb.Pn("    return")
			sb.Pn("}\n")
		}

		// Set
		if strings.Contains(prop.Access, "write") {
			sb.Pn("func (p %s) Set(flags dbus.Flags, value %s) error {",
				pc.Type, pc.GoType)
			sb.Pn("return p.Impl.GetObject_().SetProperty_(flags,"+
				" p.Impl.GetInterfaceName_(), %q, value)", prop.Name)
			sb.Pn("}\n")
		}

		// ConnectChanged
		sb.Pn("func (p %s) ConnectChanged(cb func(hasValue bool, value %s)) error {",
			pc.Type, pc.GoType)
		sb.Pn("cb0 := func(hasValue bool, value interface{}) {")

		sb.Pn("if hasValue {")
		sb.Pn("    var v %s", pc.GoType)
		sb.Pn("    err := dbus.Store([]interface{}{value}, &v)")
		sb.Pn("    if err != nil {")
		sb.Pn("        return")
		sb.Pn("    }")
		sb.Pn("    cb(true, v)")
		sb.Pn("} else {")
		sb.Pn("    cb(false, %s)", pc.GoEmptyValue)
		sb.Pn("}")

		sb.Pn("}") // end cb0

		sb.Pn("return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),")
		sb.Pn("%q, cb0)", prop.Name)
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
