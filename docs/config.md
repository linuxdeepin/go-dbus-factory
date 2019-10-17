

```go
type ServiceConfig struct {
	Service       string // optional
	Objects       []*ObjectConfig
	PropertyTypes []PropertyTypeConfig
}

type ObjectConfig struct {
	Type       string
	Path       string   // optional
	XMLFile    string   // optional
	XMLFiles   []string // optional
	Interfaces []*InterfaceConfig
}

type InterfaceConfig struct {
	Name        string
	Type        string
	Accessor    string
	Fixes       map[string]json.RawMessage
	TypeDefined bool
}

type PropertyTypeConfig struct {
	Type       string
	ValueType  string
	EmptyValue string
}
```

ObjectConfig 的 XMLFiles 字段可以用 glob 通配符。

InterfaceConfig 中的Fixes 字段 map 的键值规则：

键根据不同对象选择不同的前缀，方法的为 "m/"，属性的为 "p/"，信号的为 "s/"。

如果以 "m/" 或 "s/" 开头，值部分为 []ArgFix。

```go
type ArgFix struct {
	Name string
	Type string
}

```

如果以 "p/" 开关，值部分为 PropertyFix。
```go
type PropertyFix struct {
	// PropertyTypeConfig 已展开
	Type       string
	ValueType  string
	EmptyValue string

	RenameTo string
	RefType  string
}
```
Type 为属性的 Go 类型，一般为 "Prop" + interface 基本名 + 属性名。
 ValueType 为属性值的 Go 类型。
 EmptyValue 为属性值为空时的 Go 表示，如果 ValueType 以 "[]" 或 "map[" 开头，可以省略。可以使用 $T 指代 ValueType 的值。
RenameTo 重命名属性。
