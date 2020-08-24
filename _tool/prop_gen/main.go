package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const propCode = `
type Prop{{.Type}} struct {
	Impl Implementer
	Name string
}

func (p Prop{{.Type}}) Get(flags dbus.Flags) (value {{.GoType}}, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p Prop{{.Type}}) Set(flags dbus.Flags, value {{.GoType}}) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p Prop{{.Type}}) ConnectChanged(cb func(hasValue bool, value {{.GoType}})) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.({{.GoType}})
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, {{.GoEmptyValue}})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}
`

type config struct {
	Type         string
	GoType       string
	GoEmptyValue string
}

func main() {
	configs := []config{
		{
			Type:         "Bool",
			GoType:       "bool",
			GoEmptyValue: "false",
		},

		{
			Type:         "String",
			GoType:       "string",
			GoEmptyValue: `""`,
		},

		{
			Type:         "ObjectPath",
			GoType:       "dbus.ObjectPath",
			GoEmptyValue: `""`,
		},

		{
			Type:         "Double",
			GoType:       "float64",
			GoEmptyValue: "0",
		},
	}

	addNumType := func(name string) {
		configs = append(configs, config{
			Type:         name,
			GoType:       strings.ToLower(name),
			GoEmptyValue: "0",
		})
	}

	for _, name := range []string{"Byte", "Int16", "Uint16", "Int32", "Uint32", "Int64",
		"Uint64"} {
		addNumType(name)
	}

	t := template.Must(template.New("propCode").Parse(propCode))

	fmt.Println("package proxy")
	fmt.Println("import \"errors\"")
	fmt.Println("import \"pkg.deepin.io/lib/dbus1\"")
	fmt.Println("\nvar errNilCallback = errors.New(\"nil callback\")")

	for _, cfg := range configs {
		err := t.Execute(os.Stdout, cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, cfg := range configs {
		cfg.Type = cfg.Type + "Array"
		cfg.GoType = "[]" + cfg.GoType
		cfg.GoEmptyValue = "nil"
		err := t.Execute(os.Stdout, cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

}
