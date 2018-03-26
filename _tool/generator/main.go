package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"pkg.deepin.io/lib/dbus1/introspect"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	dir := os.Args[1]
	log.Println("dir:", dir)

	configFile := filepath.Join(dir, "config.json")
	srvCfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(srvCfg)

	sf := NewSourceFile("libdock")

	sf.AddGoImport("fmt")
	sf.AddGoImport("unsafe")
	sf.AddGoImport("pkg.deepin.io/lib/dbus1")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil/client")

	for _, objCfg := range srvCfg.Objects {
		node, err := ParseNode(filepath.Join(dir, objCfg.Type+".xml"))
		if err != nil {
			log.Fatal(err)
		}

		sf.GoBody.Pn("type %s struct {", objCfg.Type)

		for _, ifc := range node.Interfaces {
			ifcCfg := objCfg.getInterface(ifc.Name)
			if ifcCfg == nil {
				continue
			}
			sf.GoBody.Pn("%s // interface %s", ifcCfg.Type, ifc.Name)
		}

		sf.GoBody.Pn("client.Object")
		sf.GoBody.Pn("}\n")

		writeNewObject(sf.GoBody, srvCfg.Service, objCfg)

		for _, ifc := range node.Interfaces {
			ifcCfg := objCfg.getInterface(ifc.Name)
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
	}

	// sf.Print()
	sf.Save(filepath.Join(dir, "auto.go"))
}

func writeNewObject(sb *SourceBody, serviceName string, cfg *ObjectConfig) {
	if cfg.Path != "" {
		sb.Pn("func New%s(conn *dbus.Conn) *%s {", cfg.Type, cfg.Type)
		sb.Pn("    obj := new(%s)", cfg.Type)
		sb.Pn("    obj.Object.Init_(conn, %q,%q)", serviceName, cfg.Path)
	} else {
		sb.Pn("func New%s(conn *dbus.Conn, path dbus.ObjectPath) *%s {",
			cfg.Type, cfg.Type)
		sb.Pn("    obj := new(%s)", cfg.Type)
		sb.Pn("    obj.Object.Init_(conn, %q,path)", serviceName)
	}
	sb.Pn("    return obj")
	sb.Pn("}\n")
}

func writeObjectAccessMethod(sb *SourceBody, ifc introspect.Interface, objCfg *ObjectConfig) {
	ifcCfg := objCfg.getInterface(ifc.Name)
	sb.Pn("func (obj *%s) %s() *%s {", objCfg.Type, ifcCfg.Accessor,
		ifcCfg.Type)
	sb.Pn("    return &obj.%s", ifcCfg.Type)
	sb.Pn("}\n")
}

func writeImplementerMethods(sb *SourceBody, ifc introspect.Interface, ifcCfg *InterfaceConfig) {
	sb.Pn("type %s struct{}", ifcCfg.Type)

	sb.Pn("func (v *%s) GetObject_() *client.Object {", ifcCfg.Type)
	sb.Pn("    return (*client.Object)(unsafe.Pointer(v))")
	sb.Pn("}\n")

	sb.Pn("func (*%s) GetInterfaceName_() string {", ifcCfg.Type)
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
		ifcCfg.Type, method.Name, getArgsProto(inArgs, false))
	sb.Pn("    return v.GetObject_().Go_(v.GetInterfaceName_()+\".%s\", flags, ch %s)",
		method.Name, getArgsName(inArgs, false))
	sb.Pn("}\n")

	// StoreXXX
	if len(outArgs) > 0 {
		sb.Pn("func (*%s) Store%s(call *dbus.Call) (%s,err error) {", ifcCfg.Type,
			method.Name, getArgsProto(outArgs, true))
		sb.Pn("    err = call.Store(%s)", getArgsRefName(outArgs, true))
		sb.Pn("    return")
		sb.Pn("}\n")
	}

	// Call
	if len(outArgs) == 0 {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) error {",
			ifcCfg.Type, method.Name, getArgsProto(inArgs, false))
		sb.Pn("return (<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done).Err",
			method.Name, getArgsName(inArgs, false))
		sb.Pn("}\n")
	} else {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) (%s, err error) {",
			ifcCfg.Type, method.Name, getArgsProto(inArgs, false),
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
		ifcCfg.Type, signal.Name, getArgsProto(args, true))
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

func writeProperty(sb *SourceBody, prop introspect.Property, ifcCfg *InterfaceConfig) {
	sb.Pn("// property %s %s\n", prop.Name, prop.Type)

	propType := getPropType(prop.Type)
	if propType != "" {
		sb.Pn("func (v *%s) %s() %s {", ifcCfg.Type, prop.Name, propType)
		sb.Pn("    return %s{", propType)
		sb.Pn("        Impl: v,")
		sb.Pn("        Name: %q,", prop.Name)
		sb.Pn("    }")

		sb.Pn("}\n")
	} else {
		propFix := ifcCfg.getPropertyFix(prop.Name)

		sb.Pn("func (v *%s) %s() %s {", ifcCfg.Type, prop.Name, propFix.Type)
		sb.Pn("    return %s{", propFix.Type)
		sb.Pn("        Impl: v,")
		sb.Pn("    }")
		sb.Pn("}\n")

		sb.Pn("type %s struct {", propFix.Type)
		sb.Pn("Impl client.Implementer")
		sb.Pn("}\n")

		// Get
		if strings.Contains(prop.Access, "read") {
			sb.Pn("func (p %s) Get(flags dbus.Flags) (value %s, err error) {",
				propFix.Type, propFix.ValueType)
			sb.Pn("err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),")
			sb.Pn("%q, &value)", prop.Name)
			sb.Pn("    return")
			sb.Pn("}\n")
		}

		// Set
		if strings.Contains(prop.Access, "write") {
			sb.Pn("func (p %s) Set(flags dbus.Flags, value %s) error {",
				propFix.Type, propFix.ValueType)
			sb.Pn("return p.Impl.GetObject_().SetProperty_(flags,"+
				" p.Impl.GetInterfaceName_(), %q, value)", prop.Name)
			sb.Pn("}\n")
		}

		// ConnectChanged
		sb.Pn("func (p %s) ConnectChanged(cb func(hasValue bool, value %s)) error {",
			propFix.Type, propFix.ValueType)
		sb.Pn("cb0 := func(hasValue bool, value interface{}) {")

		sb.Pn("if hasValue {")
		sb.Pn("    var v %s", propFix.ValueType)
		sb.Pn("    err := dbus.Store([]interface{}{value}, &v)")
		sb.Pn("    if err != nil {")
		sb.Pn("        return")
		sb.Pn("    }")
		sb.Pn("    cb(true, v)")
		sb.Pn("} else {")
		sb.Pn("    cb(false, %s)", propFix.EmptyValue)
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
