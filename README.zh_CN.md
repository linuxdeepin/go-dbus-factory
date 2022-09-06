# go-dbus-factory

用于自动化生成 DBus 服务的方便的 go 绑定代码。

命令如果没有特殊说明，都在项目根目录下执行。

构建生成器，执行命令：
```
make bin
```
将产生 generator 二进制。

对单个文件夹进行自动化代码生成, 如执行命令：
```
./generator org.freedesktop.notifications
```

对所有文件夹执行自动化代码生成，执行命令
```
./gen.sh
```

这个脚本自动查找包含特征文件 config.json 的文件夹进行生成。

## 文件夹

将生成一个 DBus 服务的代码所需的文件集中放到一个文件夹，文件夹名一般是 DBus 服务名全小写，如标准的通知服务
org.freedesktop.Notifications 文件夹为 org.freedesktop.notifications, 自动生成的代码包的包名为点分割的最后一个组件，比如 notifications。


文件夹内有内省xml, 配置文件 config.json 和 go 代码。

### 内省 xml
内省 xml 可以用 gdbus introspect 命令获取，比如获取 org.freedesktop.Notifications 服务的 /org/freedesktop/Notifications Object 的内省 xml，执行命令：
```
gdbus introspect -e -d org.freedesktop.Notifications -o /org/freedesktop/Notifications -x > Notifications.xml
```

内省 xml 的根元素的 tag 为 node，文件名为 config.json 配置文件中 Object 的 Type 加后缀 `.xml`。

### config.json

简单的例子，如 org.freedesktop.Notifications 服务的配置文件
```json
{
  "Service": "org.freedesktop.Notifications",
  "Objects": [
    {
      "Type": "Notifications",
      "Path": "/org/freedesktop/Notifications",
      "Interfaces": [
        {
          "Name": "org.freedesktop.Notifications",
          "Type": "notifications"
        }
      ]
    }
  ]
}
```

有关配置文件的详细查看生成器的代码 `_tool/generator/config.go`

其中 Object 如果 Path 是固定的，则指定Path，如果 Path 不固定，则不指定 Path。

如果 Object 具有多个接口，则需要指定 InterfaceConfig 的 Accessor 字段， 如果只具有一个接口，则不需要。
参考 `com.deepin.daemon.apps/config.json`。

如果 Object 具有 org.freedestkop.DBus.ObjectManager 接口，则 InterfaceConfig 只需写 Name 字段就行，生成器对这个 interface 进行了特殊处理。

Type 字段一般对应自动生成代码里的 struct 类型的名。
比如 ObjectConfig 的 Type 与 
InterfaceConfig 的 Type。

#### 对接口的修正

 用 InterfaceConfig 的 Fixes 字段指定， 其类型为 `map[string]json.RawMessage`。

键是属性、方法、信号的名加上特定的前缀：

* 属性前缀 `p/`
* 方法前缀 `m/`
* 信号前缀 `s/`

如要修正属性 Name，则 Fixes 中的键为 "p/Name"。

值可被解析为 ArgFixes 或 PropertyFix 类型。
ArgFixes 用于对方法或信号的某个参数的的类型修正，ArgFix 的 Name 指定了参数名，Type 为这个参数指定了新的类型。

PropertyFix 用与对属性的类型进行修正。


```go
type ArgFixes []ArgFix

type ArgFix struct {
	Name string
	Type string
}

type PropertyFix struct {
	Type       string
	ValueType  string
	EmptyValue string
}
```

对属性的修正是必须的，如果不写则生成器会painc，对方法和信号的修正是可选的，如果不写则会有警告。

如警告信息：
```
Warning: found []interface{} in org.freedesktop.login1.Manager.CreateSession arg properties
```

### go代码
自动生成的代码命名为 auto.go, 手写代码命名任意，一般为 manual.go。
手写代码一般作为自动代码的补充，比如定义struct类型。

## 帮助
任何使用问题都可以通过以下方式寻求帮助：

* [Gitter](https://gitter.im/orgs/linuxdeepin/rooms)
* [IRC channel](https://webchat.freenode.net/?channels=deepin)
* [WiKi](https://wiki.deepin.org)
* [官方论坛](https://bbs.deepin.org)
* [开发者中心](https://github.com/linuxdeepin/go-dbus-factory) 

## 贡献指南

我们鼓励您报告问题并做出更改

- [开发者代码贡献指南](https://github.com/linuxdeepin/developer-center/wiki/Contribution-Guidelines-for-Developers) 

## 开源许可证
go-dbus-factory 在 [GPL-3.0-or-later](LICENSE) 下发布。