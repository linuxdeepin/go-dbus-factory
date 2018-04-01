package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"strconv"

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

	dirName := filepath.Base(dir)
	lastDotIdx := strings.LastIndexByte(dirName, '.')
	pkg := dirName[lastDotIdx+1:]

	sf := NewSourceFile(pkg)

	sf.AddGoImport("errors")
	sf.AddGoImport("fmt")
	sf.AddGoImport("unsafe")
	sf.AddGoImport("pkg.deepin.io/lib/dbus1")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil")
	sf.AddGoImport("pkg.deepin.io/lib/dbusutil/proxy")

	sf.GoBody.Pn("/* prevent compile error */")
	sf.GoBody.Pn("var _ = errors.New")
	sf.GoBody.Pn("var _ dbusutil.SignalHandlerId")
	sf.GoBody.Pn("var _ = fmt.Sprintf")
	sf.GoBody.Pn("")

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
			if ifc.Name == "org.freedesktop.DBus.ObjectManager" {
				ifcCfg.TypeDefined = true
				sf.AddGoImport("github.com/linuxdeepin/go-dbus-factory/object_manager")
				sf.GoBody.Pn("object_manager.ObjectManager // interface %s", ifc.Name)
			} else {
				sf.GoBody.Pn("%s // interface %s", ifcCfg.Type, ifc.Name)
			}
		}

		sf.GoBody.Pn("proxy.Object")
		sf.GoBody.Pn("}\n")

		writeNewObject(sf.GoBody, srvCfg.Service, objCfg)

		for _, ifc := range node.Interfaces {
			ifcCfg := objCfg.getInterface(ifc.Name)
			if ifcCfg == nil || ifcCfg.TypeDefined {
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
				var methodSameName bool
				for _, method := range ifc.Methods {
					if method.Name == prop.Name {
						methodSameName = true
						break
					}
				}
				writeProperty(sf.GoBody, prop, ifcCfg, methodSameName)
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
		sb.Pn("    return obj")
	} else {
		sb.Pn("func New%s(conn *dbus.Conn, path dbus.ObjectPath) (*%s,error) {",
			cfg.Type, cfg.Type)
		sb.Pn("    if !path.IsValid() {")
		sb.Pn("        return nil, errors.New(\"path is invalid\")")
		sb.Pn("    }")
		sb.Pn("    obj := new(%s)", cfg.Type)
		sb.Pn("    obj.Object.Init_(conn, %q,path)", serviceName)
		sb.Pn("    return obj, nil")
	}
	sb.Pn("}\n")
}

func writeObjectAccessMethod(sb *SourceBody, ifc introspect.Interface, objCfg *ObjectConfig) {
	ifcCfg := objCfg.getInterface(ifc.Name)
	if ifcCfg.Accessor == "" {
		return
	}

	sb.Pn("func (obj *%s) %s() *%s {", objCfg.Type, ifcCfg.Accessor,
		ifcCfg.Type)
	sb.Pn("    return &obj.%s", ifcCfg.Type)
	sb.Pn("}\n")
}

func writeImplementerMethods(sb *SourceBody, ifc introspect.Interface, ifcCfg *InterfaceConfig) {
	sb.Pn("type %s struct{}", ifcCfg.Type)

	sb.Pn("func (v *%s) GetObject_() *proxy.Object {", ifcCfg.Type)
	sb.Pn("    return (*proxy.Object)(unsafe.Pointer(v))")
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

	argFixes := ifcCfg.getArgFixes("m/" + method.Name)
	// check and warn
	for _, arg := range args {
		argType := getArgType(arg, argFixes)
		if strings.Contains(argType, "[]interface {}") {
			log.Printf("Warning: found []interface{} in %s.%s arg %s\n",
				ifcCfg.Name, method.Name, arg.Name)
		}
	}

	// GoXXX

	sb.Pn("func (v *%s) Go%s(flags dbus.Flags, ch chan *dbus.Call %s) *dbus.Call {",
		ifcCfg.Type, method.Name, getArgsProto(inArgs, false, argFixes))
	sb.Pn("    return v.GetObject_().Go_(v.GetInterfaceName_()+\".%s\", flags, ch %s)",
		method.Name, getArgsName(inArgs, false))
	sb.Pn("}\n")

	// StoreXXX
	if len(outArgs) > 0 {
		sb.Pn("func (*%s) Store%s(call *dbus.Call) (%s,err error) {", ifcCfg.Type,
			method.Name, getArgsProto(outArgs, true, argFixes))
		sb.Pn("    err = call.Store(%s)", getArgsRefName(outArgs, true))
		sb.Pn("    return")
		sb.Pn("}\n")
	}

	// Call
	if len(outArgs) == 0 {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) error {",
			ifcCfg.Type, method.Name, getArgsProto(inArgs, false, argFixes))
		sb.Pn("return (<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done).Err",
			method.Name, getArgsName(inArgs, false))
		sb.Pn("}\n")
	} else {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) (%s, err error) {",
			ifcCfg.Type, method.Name, getArgsProto(inArgs, false, argFixes),
			getArgsProto(outArgs, true, argFixes))
		sb.Pn("return v.Store%s(", method.Name)
		sb.Pn("<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done)",
			method.Name, getArgsName(inArgs, false))
		sb.Pn("}\n")
	}
}

func writeSignal(sb *SourceBody, signal introspect.Signal, ifcCfg *InterfaceConfig) {
	sb.Pn("// signal %s\n", signal.Name)

	args := getArgs(signal.Args)
	argFixes := ifcCfg.getArgFixes("s/" + signal.Name)

	sb.Pn("func (v *%s) Connect%s(cb func(%s)) (dbusutil.SignalHandlerId, error) {",
		ifcCfg.Type, signal.Name, getArgsProto(args, true, argFixes))
	sb.Pn("if cb == nil {")
	sb.Pn("   return 0, errors.New(\"nil callback\")")
	sb.Pn("}")
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
			argType := getArgType(arg, argFixes)
			sb.Pn("var %s %s", arg.Name, argType)
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
		return "proxy.Prop" + val
	} else if len(ty) == 2 && ty[0] == 'a' {
		val := propBaseTypeMap[ty[1:]]
		if val == "" {
			return ""
		}
		return "proxy.Prop" + val + "Array"
	}
	return ""
}

func writeProperty(sb *SourceBody, prop introspect.Property, ifcCfg *InterfaceConfig,
	methodSameName bool) {
	sb.Pn("// property %s %s\n", prop.Name, prop.Type)

	var funcName string
	if methodSameName {
		funcName = "Prop" + prop.Name
	} else {
		funcName = prop.Name
	}

	propType := getPropType(prop.Type)
	if propType != "" {
		sb.Pn("func (v *%s) %s() %s {", ifcCfg.Type, funcName, propType)
		sb.Pn("    return %s{", propType)
		sb.Pn("        Impl: v,")
		sb.Pn("        Name: %q,", prop.Name)
		sb.Pn("    }")

		sb.Pn("}\n")
	} else {
		propFix := ifcCfg.getPropertyFix(prop.Name)
		if propFix == nil {
			panic(fmt.Errorf("failed to get property fix for %s.%s", ifcCfg.Name, prop.Name))
		}

		sb.Pn("func (v *%s) %s() %s {", ifcCfg.Type, prop.Name, propFix.Type)
		sb.Pn("    return %s{", propFix.Type)
		sb.Pn("        Impl: v,")
		sb.Pn("    }")
		sb.Pn("}\n")

		sb.Pn("type %s struct {", propFix.Type)
		sb.Pn("Impl proxy.Implementer")
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
		sb.Pn("if cb == nil {")
		sb.Pn("    return errors.New(\"nil callback\")")
		sb.Pn("}")
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

func initNameMap() map[string]int {
	goKeywords := []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}
	nameMap := make(map[string]int, len(goKeywords))
	for _, keyword := range goKeywords {
		nameMap[keyword] = 1
	}
	nameMap["flags"] = 1
	nameMap["ch"] = 1

	return nameMap
}

func getName(nameMap map[string]int, name string) string {
	count := nameMap[name]
	defer func() {
		nameMap[name] = count + 1
	}()

	if count == 0 {
		return name
	}
	return name + strconv.Itoa(count-1)
}

func getArgs(args []introspect.Arg) []introspect.Arg {
	nameMap := initNameMap()
	var ret []introspect.Arg
	fixName := false
	for idx, arg := range args {

		if idx == 0 {
			if arg.Name == "" {
				fixName = true
				nameMap["arg"] = 1
			}
		}

		arg0 := arg
		if fixName {
			arg0.Name = getName(nameMap, "arg")
		} else {
			arg0.Name = getName(nameMap, arg0.Name)
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

func getArgType(arg introspect.Arg, argFixes ArgFixes) string {
	argType := argFixes.get(arg.Name).Type
	if argType != "" {
		return argType
	}
	return TypeFor(arg.Type).String()
}

func getArgsProto(args []introspect.Arg, begin bool, argFixes ArgFixes) string {
	var buf bytes.Buffer

	for _, arg := range args {
		argType := getArgType(arg, argFixes)
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
