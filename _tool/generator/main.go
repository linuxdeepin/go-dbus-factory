// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/godbus/dbus/introspect"
	"github.com/linuxdeepin/go-lib/strv"
	"github.com/linuxdeepin/go-lib/utils"
)

var reservedSignals = strv.Strv{
	"PropertiesChanged",
	"InterfaceAdded",
	"InterfaceRemoved",
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	dir := os.Args[1]
	log.Println("dir:", dir)

	configFile := filepath.Join(dir, "config.json")
	srvCfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	dirName := filepath.Base(dir)
	lastDotIdx := strings.LastIndexByte(dirName, '.')
	pkg := dirName[lastDotIdx+1:]

	sf := NewSourceFile(pkg)
	msf := NewSourceFile(pkg)

	sf.AddGoImport("errors")
	sf.AddGoImport("fmt")
	sf.AddGoImport("unsafe")
	sf.AddGoImport("github.com/godbus/dbus")
	sf.AddGoImport("github.com/linuxdeepin/go-lib/dbusutil")
	sf.AddGoImport("github.com/linuxdeepin/go-lib/dbusutil/proxy")
	sf.AddGoImport("github.com/linuxdeepin/go-dbus-factory/object_manager")

	msf.AddGoImport("fmt")
	msf.AddGoImport("github.com/godbus/dbus")
	msf.AddGoImport("github.com/linuxdeepin/go-lib/dbusutil")
	msf.AddGoImport("github.com/linuxdeepin/go-lib/dbusutil/proxy")
	msf.AddGoImport("github.com/stretchr/testify/mock")

	srvCfg.autoFill()

	for _, objCfg := range srvCfg.Objects {
		log.Println("Object", objCfg.Type)
		interfaces, err := objCfg.loadXml(dir)
		if err != nil {
			log.Fatal(err)
		}

		sf.GoBody.Pn("type %s interface {", objCfg.Type)

		for _, ifcCfg := range objCfg.Interfaces {
			if len(objCfg.Interfaces) > 1 {
				if ifcCfg.Name == "org.freedesktop.DBus.ObjectManager" {
					sf.GoBody.Pn("%s() object_manager.ObjectManager // interface %s", ifcCfg.Accessor, ifcCfg.Name)
				} else {
					sf.GoBody.Pn("%s() %s // interface %s", ifcCfg.Accessor, ifcCfg.Type, ifcCfg.Name)
				}
			} else {
				if ifcCfg.Name == "org.freedesktop.DBus.ObjectManager" {
					ifcCfg.TypeDefined = true
					sf.GoBody.Pn("object_manager.ObjectManager // interface %s", ifcCfg.Name)
				} else {
					sf.GoBody.Pn("%s // interface %s", ifcCfg.Type, ifcCfg.Name)
				}
			}
		}

		sf.GoBody.Pn("proxy.Object")
		sf.GoBody.Pn("}\n")

		sf.GoBody.Pn("type %s struct {", toObjectImplName(objCfg.Type))

		for _, ifcCfg := range objCfg.Interfaces {
			if ifcCfg.Name == "org.freedesktop.DBus.ObjectManager" {
				ifcCfg.TypeDefined = true
				sf.GoBody.Pn("object_manager.InterfaceObjectManager // interface %s", ifcCfg.Name)
			} else {
				sf.GoBody.Pn("%s // interface %s", toInterfaceImplName(ifcCfg.Type), ifcCfg.Name)
			}
		}

		sf.GoBody.Pn("proxy.ImplObject")
		sf.GoBody.Pn("}\n")

		msf.GoBody.Pn("type %s struct {", toObjectMockName(objCfg.Type))

		for _, ifcCfg := range objCfg.Interfaces {
			if ifcCfg.Name == "org.freedesktop.DBus.ObjectManager" {
				ifcCfg.TypeDefined = true
				msf.GoBody.Pn("object_manager.MockInterfaceObjectManager // interface %s", ifcCfg.Name)
			} else {
				msf.GoBody.Pn("%s // interface %s", toInterfaceMockName(ifcCfg.Type), ifcCfg.Name)
			}
		}

		msf.GoBody.Pn("proxy.MockObject")
		msf.GoBody.Pn("}\n")

		writeNewObject(sf.GoBody, srvCfg.Service, objCfg)

		for _, ifcCfg := range objCfg.Interfaces {
			ifc := getInterfaceByName(interfaces, ifcCfg.Name)

			if len(objCfg.Interfaces) > 1 {
				writeImplementerAccessorMethod(sf.GoBody, ifcCfg, objCfg)
			}

			if ifcCfg.TypeDefined {
				continue
			}

			if ifc == nil {
				log.Fatalf("not found interface %q for object %s",
					ifcCfg.Name, objCfg.Type)
			}

			writeInterfaceInterface(sf.GoBody, ifc, ifcCfg)
			writeInterfaceImpl(sf.GoBody, ifc, ifcCfg)
			writeInterfaceMock(msf.GoBody, ifc, ifcCfg)

			for _, method := range ifc.Methods {
				writeMethodImpl(sf.GoBody, method, ifcCfg)
				writeMethodMock(msf.GoBody, method, ifcCfg)
			}

			for _, signal := range ifc.Signals {
				writeSignalImpl(sf.GoBody, signal, ifcCfg)
				writeSignalMock(msf.GoBody, signal, ifcCfg)
			}

			for _, prop := range ifc.Properties {
				var methodSameName bool
				for _, method := range ifc.Methods {
					if strings.Title(method.Name) == strings.Title(prop.Name) {
						methodSameName = true
						break
					}
				}
				writeProperty(sf.GoBody, msf.GoBody, prop, ifcCfg, methodSameName)
			}
		}
	}

	for _, propTypeCfg := range srvCfg.PropertyTypes {
		writePropTypeInterface(sf.GoBody, propTypeCfg)
		writePropTypeImpl(sf.GoBody, propTypeCfg)
		writePropTypeMock(msf.GoBody, propTypeCfg)
	}

	autoGoFile := filepath.Join(dir, "auto.go")
	md5sum, ok := utils.SumFileMd5(autoGoFile)
	sf.Save(autoGoFile)
	if ok {
		md5sum1, ok1 := utils.SumFileMd5(autoGoFile)
		if ok1 && md5sum == md5sum1 {
			log.Println("auto.go not changed")
		}
	}

	autoMockGoFile := filepath.Join(dir, "auto_mock.go")
	md5sum, ok = utils.SumFileMd5(autoMockGoFile)
	msf.Save(autoMockGoFile)
	if ok {
		md5sum1, ok1 := utils.SumFileMd5(autoMockGoFile)
		if ok1 && md5sum == md5sum1 {
			log.Println("auto_mock.go not changed")
		}
	}
}

func writeNewObject(sb *SourceBody, serviceName string, cfg *ObjectConfig) {
	var inArgs []string
	outTypes := []string{cfg.Type}
	if serviceName == "" {
		inArgs = append(inArgs, "serviceName string")
	}

	if cfg.Path == "" {
		inArgs = append(inArgs, "path dbus.ObjectPath")
		outTypes = append(outTypes, "error")
	}

	outTypesStr := strings.Join(outTypes, ",")
	if len(outTypes) > 1 {
		outTypesStr = "(" + outTypesStr + ")"
	}

	sb.Pn("func New%s(conn *dbus.Conn, %s) %s {", cfg.Type,
		strings.Join(inArgs, ","), outTypesStr)

	// check path
	if cfg.Path == "" {
		sb.Pn("if !path.IsValid() {")
		sb.Pn("    return nil, errors.New(\"path is invalid\")")
		sb.Pn("}")
	}

	sb.Pn("obj := new(%s)", toObjectImplName(cfg.Type))
	var inArgServiceName string
	var inArgPath string

	if serviceName != "" {
		inArgServiceName = strconv.Quote(serviceName)
	} else {
		inArgServiceName = "serviceName"
	}

	if cfg.Path != "" {
		inArgPath = strconv.Quote(cfg.Path)
	} else {
		inArgPath = "path"
	}

	sb.Pn("obj.ImplObject.Init_(conn, %s, %s)", inArgServiceName, inArgPath)

	// return
	var returnExpr string
	if cfg.Path == "" {
		returnExpr = "obj, nil"
	} else {
		returnExpr = "obj"
	}
	sb.Pn("return %s", returnExpr)
	sb.Pn("}\n")
}

func writeImplementerAccessorMethod(sb *SourceBody, ifcCfg *InterfaceConfig, objCfg *ObjectConfig) {
	if ifcCfg.Name == "org.freedesktop.DBus.ObjectManager" {
		sb.Pn("func (obj *%s) ObjectManager() object_manager.ObjectManager {", toObjectImplName(objCfg.Type))
		sb.Pn("    return &obj.InterfaceObjectManager")
		sb.Pn("}\n")
	} else {
		sb.Pn("func (obj *%s) %s() %s {", toObjectImplName(objCfg.Type), ifcCfg.Accessor,
			ifcCfg.Type)
		sb.Pn("    return &obj.%s", toInterfaceImplName(ifcCfg.Type))
		sb.Pn("}\n")
	}
}

func writeInterfaceInterface(sb *SourceBody, ifc *introspect.Interface, ifcCfg *InterfaceConfig) {
	sb.Pn("type %s interface{", ifcCfg.Type)

	if ifcCfg.CustomInterfaceName {
		sb.Pn("SetInterfaceName_(name string)")
	}

	for _, method := range ifc.Methods {
		methodName := strings.Title(method.Name)
		args := getArgs(method.Args)
		inArgs := getInArgs(args)
		outArgs := getOutArgs(args)
		argFixes := ifcCfg.getMethodFix(method.Name)

		sb.Pn("Go%s(flags dbus.Flags, ch chan *dbus.Call %s) *dbus.Call",
			methodName, getArgsProto(inArgs, false, argFixes))

		if len(outArgs) == 0 {
			sb.Pn("%s(flags dbus.Flags %s) error",
				methodName, getArgsProto(inArgs, false, argFixes))
		} else {
			sb.Pn("%s(flags dbus.Flags %s) (%s, error)",
				methodName, getArgsProto(inArgs, false, argFixes),
				getArgsType(outArgs, true, argFixes))
		}
	}

	for _, signal := range ifc.Signals {
		args := getArgs(signal.Args)
		argFixes := ifcCfg.getSignalFix(signal.Name)
		methodName := strings.Title(signal.Name)

		if reservedSignals.Contains(methodName) {
			methodName = "Signal" + methodName
		}

		sb.Pn("Connect%s(cb func(%s)) (dbusutil.SignalHandlerId, error)",
			methodName, getArgsProto(args, true, argFixes))
	}

	for _, prop := range ifc.Properties {
		var methodSameName bool
		for _, method := range ifc.Methods {
			if strings.Title(method.Name) == strings.Title(prop.Name) {
				methodSameName = true
				break
			}
		}

		propFix := ifcCfg.getPropertyFix(prop.Name)

		var renameTo string
		if propFix != nil {
			renameTo = propFix.RenameTo
		}

		// get funcName
		funcName := strings.Title(prop.Name)
		if renameTo != "" {
			funcName = renameTo
		} else if methodSameName {
			funcName = "Prop" + funcName
		}

		propIface, _, _ := getPropType(prop.Type)
		if propIface == "" {
			if propFix == nil {
				panic(fmt.Errorf("failed to get property fix for %s.%s",
					ifcCfg.Name, prop.Name))
			}

			if propFix.RefType != "" {
				propIface = propFix.RefType
			} else {
				propIface = propFix.Type
			}
		}

		sb.Pn("%s() %s", funcName, propIface)
	}
	sb.Pn("}\n")
}

func writeInterfaceMock(sb *SourceBody, ifc *introspect.Interface, ifcCfg *InterfaceConfig) {
	t := toInterfaceMockName(ifcCfg.Type)
	sb.Pn("type %s struct{", t)
	sb.Pn("    mock.Mock")
	sb.Pn("}\n")

	if ifcCfg.CustomInterfaceName {
		sb.Pn("func (v *%s) SetInterfaceName_(string) {", t)
		sb.Pn("}\n")
	}

}

func writeInterfaceImpl(sb *SourceBody, ifc *introspect.Interface, ifcCfg *InterfaceConfig) {
	t := toInterfaceImplName(ifcCfg.Type)

	sb.Pn("type %s struct{}", t)

	sb.Pn("func (v *%s) GetObject_() *proxy.ImplObject {", t)
	sb.Pn("    return (*proxy.ImplObject)(unsafe.Pointer(v))")
	sb.Pn("}\n")

	if ifcCfg.CustomInterfaceName {
		sb.Pn("func (v *%s) SetInterfaceName_(name string) {", t)
		sb.Pn(`    v.GetObject_().SetExtra("customIfc", name)`)
		sb.Pn("}\n")
	}

	if !ifcCfg.NoGetInterfaceName {
		if ifcCfg.CustomInterfaceName {
			sb.Pn("func (v *%s) GetInterfaceName_() string {", t)
			sb.Pn(`    ifcName, _ := v.GetObject_().GetExtra("customIfc")`)
			sb.Pn("    ifcNameStr, _ := ifcName.(string)")
			sb.Pn("    return ifcNameStr")
			sb.Pn("}\n")
		} else {
			sb.Pn("func (*%s) GetInterfaceName_() string {", t)
			sb.Pn("    return %q", ifc.Name)
			sb.Pn("}\n")
		}
	}
}

func writeMethodImpl(sb *SourceBody, method introspect.Method, ifcCfg *InterfaceConfig) {
	t := toInterfaceImplName(ifcCfg.Type)

	sb.Pn("// method %s\n", method.Name)
	methodName := strings.Title(method.Name)
	args := getArgs(method.Args)
	inArgs := getInArgs(args)
	outArgs := getOutArgs(args)
	// sb.Pn("// in %#v", inArgs)
	// sb.Pn("// out %#v", outArgs)

	argFixes := ifcCfg.getMethodFix(method.Name)
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
		t, methodName, getArgsProto(inArgs, false, argFixes))
	sb.Pn("    return v.GetObject_().Go_(v.GetInterfaceName_()+\".%s\", flags, ch %s)",
		method.Name, getArgsName(inArgs, false))
	sb.Pn("}\n")

	// StoreXXX
	if len(outArgs) > 0 {
		sb.Pn("func (*%s) Store%s(call *dbus.Call) (%s,err error) {", t,
			methodName, getArgsProto(outArgs, true, argFixes))
		sb.Pn("    err = call.Store(%s)", getArgsRefName(outArgs, true))
		sb.Pn("    return")
		sb.Pn("}\n")
	}

	// Call
	if len(outArgs) == 0 {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) error {",
			t, methodName, getArgsProto(inArgs, false, argFixes))
		sb.Pn("return (<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done).Err",
			methodName, getArgsName(inArgs, false))
		sb.Pn("}\n")
	} else {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) (%s, error) {",
			t, methodName, getArgsProto(inArgs, false, argFixes),
			getArgsType(outArgs, true, argFixes))
		sb.Pn("return v.Store%s(", methodName)
		sb.Pn("<-v.Go%s(flags, make(chan *dbus.Call, 1) %s).Done)",
			methodName, getArgsName(inArgs, false))
		sb.Pn("}\n")
	}
}

func writeMethodMock(sb *SourceBody, method introspect.Method, ifcCfg *InterfaceConfig) {
	t := toInterfaceMockName(ifcCfg.Type)

	sb.Pn("// method %s\n", method.Name)

	argFixes := ifcCfg.getMethodFix(method.Name)

	methodName := strings.Title(method.Name)
	args := getArgs(method.Args)
	inArgs := getInArgs(args)
	outArgs := getOutArgs(args)

	// GoXXX
	sb.Pn("func (v *%s) Go%s(flags dbus.Flags, ch chan *dbus.Call %s) *dbus.Call {",
		t, methodName, getArgsProto(inArgs, false, argFixes))
	sb.Pn(" mockArgs := v.Called(flags, ch %s)\n", getArgsName(inArgs, false))
	sb.Pn(` ret, ok := mockArgs.Get(0).(*dbus.Call)
			if !ok {
				panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %%v", mockArgs.Get(0)))
			}
			`)
	sb.Pn("return ret")
	sb.Pn("}\n")

	// Call
	if len(outArgs) == 0 {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) error {",
			t, methodName, getArgsProto(inArgs, false, argFixes))
	} else {
		sb.Pn("func (v *%s) %s(flags dbus.Flags %s) (%s, error) {",
			t, methodName, getArgsProto(inArgs, false, argFixes),
			getArgsType(outArgs, true, argFixes))
	}

	sb.Pn("mockArgs := v.Called(flags %s)\n", getArgsName(inArgs, false))

	var rets []string
	for i, a := range outArgs {
		argType := getArgType(a, argFixes)
		switch argType {
		case "string", "int", "bool":
			ret := fmt.Sprintf("mockArgs.%s(%d)", strings.Title(argType), i)
			rets = append(rets, ret)
		default:
			retName := fmt.Sprintf("ret%d", i)

			sb.Pn("%s, ok := mockArgs.Get(%d).(%s)", retName, i, argType)
			sb.Pn(`if !ok {
					panic(fmt.Sprintf("assert: arguments: %%d failed because object wasn't correct type: %%v", %[1]d, mockArgs.Get(%[1]d)))
				}
				`, i)

			rets = append(rets, retName)
		}

	}

	rets = append(rets, fmt.Sprintf("mockArgs.Error(%d)", len(outArgs)))

	sb.Pn("return %s", strings.Join(rets, ","))
	sb.Pn("}\n")
}

func writeSignalImpl(sb *SourceBody, signal introspect.Signal, ifcCfg *InterfaceConfig) {
	t := toInterfaceImplName(ifcCfg.Type)

	sb.Pn("// signal %s\n", signal.Name)

	args := getArgs(signal.Args)
	argFixes := ifcCfg.getSignalFix(signal.Name)
	methodName := strings.Title(signal.Name)

	if reservedSignals.Contains(methodName) {
		methodName = "Signal" + methodName
	}

	sb.Pn("func (v *%s) Connect%s(cb func(%s)) (dbusutil.SignalHandlerId, error) {",
		t, methodName, getArgsProto(args, true, argFixes))
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

func writeSignalMock(sb *SourceBody, signal introspect.Signal, ifcCfg *InterfaceConfig) {
	t := toInterfaceMockName(ifcCfg.Type)

	args := getArgs(signal.Args)
	argFixes := ifcCfg.getSignalFix(signal.Name)
	methodName := strings.Title(signal.Name)

	if reservedSignals.Contains(methodName) {
		methodName = "Signal" + methodName
	}

	sb.Pn("// signal %s\n", signal.Name)

	sb.Pn("func (v *%s) Connect%s(cb func(%s)) (dbusutil.SignalHandlerId, error) {",
		t, methodName, getArgsProto(args, true, argFixes))

	sb.Pn("mockArgs := v.Called(cb)\n")

	var rets []string
	retName := "ret0"
	sb.Pn("%s, ok := mockArgs.Get(%d).(%s)", retName, 0, "dbusutil.SignalHandlerId")
	sb.Pn(`if !ok {
			panic(fmt.Sprintf("assert: arguments: %%d failed because object wasn't correct type: %%v", %[1]d, mockArgs.Get(%[1]d)))
		}
		`, 0)

	rets = append(rets, retName)
	rets = append(rets, "mockArgs.Error(1)")

	sb.Pn("return %s", strings.Join(rets, ","))
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

func getPropType(ty string) (iface string, stct string, mock string) {
	if len(ty) == 1 {
		val := propBaseTypeMap[ty]
		if val == "" {
			return
		}
		iface = "proxy.Prop" + val
		stct = "proxy.ImplProp" + val
		mock = "proxy.MockProp" + val
	} else if len(ty) == 2 && ty[0] == 'a' {
		val := propBaseTypeMap[ty[1:]]
		if val == "" {
			return
		}
		iface = "proxy.Prop" + val + "Array"
		stct = "proxy.ImplProp" + val + "Array"
		mock = "proxy.MockProp" + val + "Array"
	}
	return
}

func writePropTypeInterface(sb *SourceBody, propTypeCfg *PropertyTypeConfig) {
	const propName = "p.Name"

	sb.Pn("type %s interface {", propTypeCfg.Type)
	sb.Pn("    Get(flags dbus.Flags) (value %s, err error)", propTypeCfg.ValueType)
	sb.Pn("    Set(flags dbus.Flags, value %s) error", propTypeCfg.ValueType)
	sb.Pn("    ConnectChanged(cb func(hasValue bool, value %s)) error", propTypeCfg.ValueType)
	sb.Pn("}\n")
}

func writePropTypeImpl(sb *SourceBody, propTypeCfg *PropertyTypeConfig) {
	const propName = "p.Name"

	sb.Pn("type impl%s struct {", propTypeCfg.Type)
	sb.Pn("    Impl proxy.Implementer")
	sb.Pn("    Name string")
	sb.Pn("}\n")

	writePropGetImpl(sb, propTypeCfg, propName)
	writePropSetImpl(sb, propTypeCfg, propName)
	writePropConnectChangedImpl(sb, propTypeCfg, propName)
}

func writePropGetImpl(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p impl%s) Get(flags dbus.Flags) (value %s, err error) {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn("err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),")
	sb.Pn("%s, &value)", propName)
	sb.Pn("    return")
	sb.Pn("}\n")
}

func writePropSetImpl(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p impl%s) Set(flags dbus.Flags, value %s) error {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn("return p.Impl.GetObject_().SetProperty_(flags,"+
		" p.Impl.GetInterfaceName_(), %s, value)", propName)
	sb.Pn("}\n")
}

func writePropConnectChangedImpl(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p impl%s) ConnectChanged(cb func(hasValue bool, value %s)) error {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn("if cb == nil {")
	sb.Pn("    return errors.New(\"nil callback\")")
	sb.Pn("}")
	sb.Pn("cb0 := func(hasValue bool, value interface{}) {")

	sb.Pn("if hasValue {")
	sb.Pn("    var v %s", propTypeCfg.ValueType)
	sb.Pn("    err := dbus.Store([]interface{}{value}, &v)")
	sb.Pn("    if err != nil {")
	sb.Pn("        return")
	sb.Pn("    }")
	sb.Pn("    cb(true, v)")
	sb.Pn("} else {")
	sb.Pn("    cb(false, %s)", propTypeCfg.EmptyValue)
	sb.Pn("}")

	sb.Pn("}") // end cb0

	sb.Pn("return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),")
	sb.Pn("%s, cb0)", propName)
	sb.Pn("}\n")
}

func writePropTypeMock(sb *SourceBody, propTypeCfg *PropertyTypeConfig) {
	const propName = "p.Name"

	sb.Pn("type Mock%s struct {", propTypeCfg.Type)
	sb.Pn("    mock.Mock")
	sb.Pn("}\n")

	writePropGetMock(sb, propTypeCfg, propName)
	writePropSetMock(sb, propTypeCfg, propName)
	writePropConnectChangedMock(sb, propTypeCfg, propName)
}

func writePropGetMock(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p Mock%s) Get(flags dbus.Flags) (value %s, err error) {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn(`	args := p.Called(flags)

			var ok bool
			value, ok = args.Get(0).(%s)
			if !ok {
				panic(fmt.Sprintf("assert: arguments: %%d failed because object wasn't correct type: %%v", 0, args.Get(0)))
			}

			err = args.Error(1)

			return`, propTypeCfg.ValueType)
	sb.Pn("}\n")
}

func writePropSetMock(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p Mock%s) Set(flags dbus.Flags, value %s) error {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn(`	args := p.Called(flags, value)

			return args.Error(0)`)
	sb.Pn("}\n")
}

func writePropConnectChangedMock(sb *SourceBody, propTypeCfg *PropertyTypeConfig, propName string) {
	sb.Pn("func (p Mock%s) ConnectChanged(cb func(hasValue bool, value %s)) error {",
		propTypeCfg.Type, propTypeCfg.ValueType)
	sb.Pn(`	args := p.Called(cb)

			return args.Error(0)`)
	sb.Pn("}\n")
}

func writeProperty(sb *SourceBody, msb *SourceBody, prop introspect.Property, ifcCfg *InterfaceConfig,
	methodSameName bool) {
	t := toInterfaceImplName(ifcCfg.Type)
	mt := toInterfaceMockName(ifcCfg.Type)

	propFix := ifcCfg.getPropertyFix(prop.Name)

	var renameTo string
	if propFix != nil {
		renameTo = propFix.RenameTo
	}

	// get funcName
	funcName := strings.Title(prop.Name)
	if renameTo != "" {
		funcName = renameTo
	} else if methodSameName {
		funcName = "Prop" + funcName
	}

	propIface, propType, propMockType := getPropType(prop.Type)
	if propIface == "" {
		if propFix == nil {
			panic(fmt.Errorf("failed to get property fix for %s.%s",
				ifcCfg.Name, prop.Name))
		}

		if propFix.RefType != "" {
			propIface = propFix.RefType
			propType = "impl" + propFix.RefType
			propMockType = "Mock" + propFix.RefType
		} else {
			propIface = propFix.Type
			propType = "impl" + propFix.Type
			propMockType = "Mock" + propFix.Type

			writePropTypeInterface(sb, &propFix.PropertyTypeConfig)
			writePropTypeImpl(sb, &propFix.PropertyTypeConfig)
			writePropTypeMock(msb, &propFix.PropertyTypeConfig)
		}
	}

	sb.Pn("// property %s %s\n", prop.Name, prop.Type)

	sb.Pn("func (v *%s) %s() %s {", t, funcName, propIface)
	sb.Pn("    return &%s{", propType)
	sb.Pn("        Impl: v,")
	sb.Pn("        Name: %q,", prop.Name)
	sb.Pn("    }")
	sb.Pn("}\n")

	msb.Pn("// property %s %s\n", prop.Name, prop.Type)

	msb.Pn("func (v *%s) %s() %s {", mt, funcName, propIface)
	msb.Pn("    mockArgs := v.Called()\n")
	msb.Pn(`    ret0, ok := mockArgs.Get(0).(*%s)`, propMockType)
	msb.Pn(`    if !ok {`)
	msb.Pn(`        panic(fmt.Sprintf("assert: arguments: %%d failed because object wasn't correct type: %%v", 0, mockArgs.Get(0)))`)
	msb.Pn("    }\n")
	msb.Pn("    return ret0")
	msb.Pn("}\n")
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

func getArgsType(args []introspect.Arg, begin bool, argFixes ArgFixes) string {
	var buf bytes.Buffer

	for _, arg := range args {
		argType := getArgType(arg, argFixes)
		buf.WriteString(fmt.Sprintf(",%s", argType))
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

func toObjectImplName(t string) string {
	return "object" + strings.Title(t)
}

func toObjectMockName(t string) string {
	return "Mock" + strings.Title(t)
}

func toInterfaceImplName(t string) string {
	return "interface" + strings.Title(t)
}

func toInterfaceMockName(t string) string {
	return "MockInterface" + strings.Title(t)
}

func fixList(str string, begin bool) string {
	if str == "" || !begin {
		return str
	}

	return str[1:]
}

func getInterfaceByName(interfaces []introspect.Interface, name string) *introspect.Interface {
	for _, ifc := range interfaces {
		if ifc.Name == name {
			return &ifc
		}
	}
	return nil
}
