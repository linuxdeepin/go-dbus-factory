package secrets

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Service struct {
	service // interface org.freedesktop.Secret.Service
	proxy.Object
}

func NewService(conn *dbus.Conn) *Service {
	obj := new(Service)
	obj.Object.Init_(conn, "org.freedesktop.secrets", "/org/freedesktop/secrets")
	return obj
}

type service struct{}

func (v *service) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*service) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Service"
}

// method OpenSession

func (v *service) GoOpenSession(flags dbus.Flags, ch chan *dbus.Call, algorithm string, input dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenSession", flags, ch, algorithm, input)
}

func (*service) StoreOpenSession(call *dbus.Call) (output dbus.Variant, result dbus.ObjectPath, err error) {
	err = call.Store(&output, &result)
	return
}

func (v *service) OpenSession(flags dbus.Flags, algorithm string, input dbus.Variant) (output dbus.Variant, result dbus.ObjectPath, err error) {
	return v.StoreOpenSession(
		<-v.GoOpenSession(flags, make(chan *dbus.Call, 1), algorithm, input).Done)
}

// method CreateCollection

func (v *service) GoCreateCollection(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, alias string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateCollection", flags, ch, properties, alias)
}

func (*service) StoreCreateCollection(call *dbus.Call) (collection dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&collection, &prompt)
	return
}

func (v *service) CreateCollection(flags dbus.Flags, properties map[string]dbus.Variant, alias string) (collection dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	return v.StoreCreateCollection(
		<-v.GoCreateCollection(flags, make(chan *dbus.Call, 1), properties, alias).Done)
}

// method SearchItems

func (v *service) GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchItems", flags, ch, attributes)
}

func (*service) StoreSearchItems(call *dbus.Call) (unlocked []dbus.ObjectPath, locked []dbus.ObjectPath, err error) {
	err = call.Store(&unlocked, &locked)
	return
}

func (v *service) SearchItems(flags dbus.Flags, attributes map[string]string) (unlocked []dbus.ObjectPath, locked []dbus.ObjectPath, err error) {
	return v.StoreSearchItems(
		<-v.GoSearchItems(flags, make(chan *dbus.Call, 1), attributes).Done)
}

// method Unlock

func (v *service) GoUnlock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unlock", flags, ch, objects)
}

func (*service) StoreUnlock(call *dbus.Call) (unlocked []dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&unlocked, &prompt)
	return
}

func (v *service) Unlock(flags dbus.Flags, objects []dbus.ObjectPath) (unlocked []dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	return v.StoreUnlock(
		<-v.GoUnlock(flags, make(chan *dbus.Call, 1), objects).Done)
}

// method Lock

func (v *service) GoLock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Lock", flags, ch, objects)
}

func (*service) StoreLock(call *dbus.Call) (locked []dbus.ObjectPath, Prompt dbus.ObjectPath, err error) {
	err = call.Store(&locked, &Prompt)
	return
}

func (v *service) Lock(flags dbus.Flags, objects []dbus.ObjectPath) (locked []dbus.ObjectPath, Prompt dbus.ObjectPath, err error) {
	return v.StoreLock(
		<-v.GoLock(flags, make(chan *dbus.Call, 1), objects).Done)
}

// method LockService

func (v *service) GoLockService(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockService", flags, ch)
}

func (v *service) LockService(flags dbus.Flags) error {
	return (<-v.GoLockService(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ChangeLock

func (v *service) GoChangeLock(flags dbus.Flags, ch chan *dbus.Call, collection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeLock", flags, ch, collection)
}

func (*service) StoreChangeLock(call *dbus.Call) (prompt dbus.ObjectPath, err error) {
	err = call.Store(&prompt)
	return
}

func (v *service) ChangeLock(flags dbus.Flags, collection dbus.ObjectPath) (prompt dbus.ObjectPath, err error) {
	return v.StoreChangeLock(
		<-v.GoChangeLock(flags, make(chan *dbus.Call, 1), collection).Done)
}

// method GetSecrets

func (v *service) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, items []dbus.ObjectPath, session dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, items, session)
}

func (*service) StoreGetSecrets(call *dbus.Call) (secrets map[dbus.ObjectPath]Secret, err error) {
	err = call.Store(&secrets)
	return
}

func (v *service) GetSecrets(flags dbus.Flags, items []dbus.ObjectPath, session dbus.ObjectPath) (secrets map[dbus.ObjectPath]Secret, err error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), items, session).Done)
}

// method ReadAlias

func (v *service) GoReadAlias(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReadAlias", flags, ch, name)
}

func (*service) StoreReadAlias(call *dbus.Call) (collection dbus.ObjectPath, err error) {
	err = call.Store(&collection)
	return
}

func (v *service) ReadAlias(flags dbus.Flags, name string) (collection dbus.ObjectPath, err error) {
	return v.StoreReadAlias(
		<-v.GoReadAlias(flags, make(chan *dbus.Call, 1), name).Done)
}

// method SetAlias

func (v *service) GoSetAlias(flags dbus.Flags, ch chan *dbus.Call, name string, collection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAlias", flags, ch, name, collection)
}

func (v *service) SetAlias(flags dbus.Flags, name string, collection dbus.ObjectPath) error {
	return (<-v.GoSetAlias(flags, make(chan *dbus.Call, 1), name, collection).Done).Err
}

// signal CollectionCreated

func (v *service) ConnectCollectionCreated(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionCreated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionCreated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CollectionDeleted

func (v *service) ConnectCollectionDeleted(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CollectionChanged

func (v *service) ConnectCollectionChanged(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Collections ao

func (v *service) Collections() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Collections",
	}
}

type Collection struct {
	collection // interface org.freedesktop.Secret.Collection
	proxy.Object
}

func NewCollection(conn *dbus.Conn, path dbus.ObjectPath) (*Collection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Collection)
	obj.Object.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type collection struct{}

func (v *collection) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*collection) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Collection"
}

// method Delete

func (v *collection) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (*collection) StoreDelete(call *dbus.Call) (prompt dbus.ObjectPath, err error) {
	err = call.Store(&prompt)
	return
}

func (v *collection) Delete(flags dbus.Flags) (prompt dbus.ObjectPath, err error) {
	return v.StoreDelete(
		<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done)
}

// method SearchItems

func (v *collection) GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchItems", flags, ch, attributes)
}

func (*collection) StoreSearchItems(call *dbus.Call) (results []dbus.ObjectPath, err error) {
	err = call.Store(&results)
	return
}

func (v *collection) SearchItems(flags dbus.Flags, attributes map[string]string) (results []dbus.ObjectPath, err error) {
	return v.StoreSearchItems(
		<-v.GoSearchItems(flags, make(chan *dbus.Call, 1), attributes).Done)
}

// method CreateItem

func (v *collection) GoCreateItem(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, secret Secret, replace bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateItem", flags, ch, properties, secret, replace)
}

func (*collection) StoreCreateItem(call *dbus.Call) (item dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&item, &prompt)
	return
}

func (v *collection) CreateItem(flags dbus.Flags, properties map[string]dbus.Variant, secret Secret, replace bool) (item dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	return v.StoreCreateItem(
		<-v.GoCreateItem(flags, make(chan *dbus.Call, 1), properties, secret, replace).Done)
}

// signal ItemCreated

func (v *collection) ConnectItemCreated(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemCreated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemCreated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemDeleted

func (v *collection) ConnectItemDeleted(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemChanged

func (v *collection) ConnectItemChanged(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Items ao

func (v *collection) Items() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Items",
	}
}

// property Label s

func (v *collection) Label() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Label",
	}
}

// property Locked b

func (v *collection) Locked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property Created t

func (v *collection) Created() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Created",
	}
}

// property Modified t

func (v *collection) Modified() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Modified",
	}
}

type Item struct {
	item // interface org.freedesktop.Secret.Item
	proxy.Object
}

func NewItem(conn *dbus.Conn, path dbus.ObjectPath) (*Item, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Item)
	obj.Object.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type item struct{}

func (v *item) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*item) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Item"
}

// method Delete

func (v *item) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (*item) StoreDelete(call *dbus.Call) (Prompt dbus.ObjectPath, err error) {
	err = call.Store(&Prompt)
	return
}

func (v *item) Delete(flags dbus.Flags) (Prompt dbus.ObjectPath, err error) {
	return v.StoreDelete(
		<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSecret

func (v *item) GoGetSecret(flags dbus.Flags, ch chan *dbus.Call, session dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecret", flags, ch, session)
}

func (*item) StoreGetSecret(call *dbus.Call) (secret Secret, err error) {
	err = call.Store(&secret)
	return
}

func (v *item) GetSecret(flags dbus.Flags, session dbus.ObjectPath) (secret Secret, err error) {
	return v.StoreGetSecret(
		<-v.GoGetSecret(flags, make(chan *dbus.Call, 1), session).Done)
}

// method SetSecret

func (v *item) GoSetSecret(flags dbus.Flags, ch chan *dbus.Call, secret Secret) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSecret", flags, ch, secret)
}

func (v *item) SetSecret(flags dbus.Flags, secret Secret) error {
	return (<-v.GoSetSecret(flags, make(chan *dbus.Call, 1), secret).Done).Err
}

// property Locked b

func (v *item) Locked() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property Attributes a{ss}

func (v *item) Attributes() PropItemAttributes {
	return PropItemAttributes{
		Impl: v,
	}
}

type PropItemAttributes struct {
	Impl proxy.Implementer
}

func (p PropItemAttributes) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Attributes", &value)
	return
}

func (p PropItemAttributes) Set(flags dbus.Flags, value map[string]string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), "Attributes", value)
}

func (p PropItemAttributes) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Attributes", cb0)
}

// property Label s

func (v *item) Label() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Label",
	}
}

// property Type s

func (v *item) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Created t

func (v *item) Created() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Created",
	}
}

// property Modified t

func (v *item) Modified() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Modified",
	}
}

type Session struct {
	session // interface org.freedesktop.Secret.Session
	proxy.Object
}

func NewSession(conn *dbus.Conn, path dbus.ObjectPath) (*Session, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Session)
	obj.Object.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type session struct{}

func (v *session) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*session) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Session"
}

// method Close

func (v *session) GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Close", flags, ch)
}

func (v *session) Close(flags dbus.Flags) error {
	return (<-v.GoClose(flags, make(chan *dbus.Call, 1)).Done).Err
}

type Prompt struct {
	prompt // interface org.freedesktop.Secret.Prompt
	proxy.Object
}

func NewPrompt(conn *dbus.Conn, path dbus.ObjectPath) (*Prompt, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Prompt)
	obj.Object.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type prompt struct{}

func (v *prompt) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*prompt) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Prompt"
}

// method Prompt

func (v *prompt) GoPrompt(flags dbus.Flags, ch chan *dbus.Call, window_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Prompt", flags, ch, window_id)
}

func (v *prompt) Prompt(flags dbus.Flags, window_id string) error {
	return (<-v.GoPrompt(flags, make(chan *dbus.Call, 1), window_id).Done).Err
}

// method Dismiss

func (v *prompt) GoDismiss(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Dismiss", flags, ch)
}

func (v *prompt) Dismiss(flags dbus.Flags) error {
	return (<-v.GoDismiss(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Completed

func (v *prompt) ConnectCompleted(cb func(dismissed bool, result dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Completed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Completed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var dismissed bool
		var result dbus.Variant
		err := dbus.Store(sig.Body, &dismissed, &result)
		if err == nil {
			cb(dismissed, result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
