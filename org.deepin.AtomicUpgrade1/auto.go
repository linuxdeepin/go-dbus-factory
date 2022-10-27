// Code generated by "./generator org.deepin.AtomicUpgrade1"; DO NOT EDIT.

// Copyright: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// License-Identifier: GPL-3.0-or-later
package AtomicUpgrade1

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type AtomicUpgrade1 interface {
	atomicUpgrade1 // interface org.deepin.AtomicUpgrade1
	proxy.Object
}

type objectAtomicUpgrade1 struct {
	interfaceAtomicUpgrade1 // interface org.deepin.AtomicUpgrade1
	proxy.ImplObject
}

func NewAtomicUpgrade1(conn *dbus.Conn) AtomicUpgrade1 {
	obj := new(objectAtomicUpgrade1)
	obj.ImplObject.Init_(conn, "org.deepin.AtomicUpgrade1", "/org/deepin/AtomicUpgrade1")
	return obj
}

type atomicUpgrade1 interface {
	GoCancelRollback(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CancelRollback(flags dbus.Flags) error
	GoCommit(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	Commit(flags dbus.Flags, arg0 string) error
	GoDelete(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	Delete(flags dbus.Flags, arg0 string) error
	GoGetGrubTitle(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetGrubTitle(flags dbus.Flags, arg0 string) (string, error)
	GoListVersion(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListVersion(flags dbus.Flags) ([]string, error)
	GoQuerySubject(flags dbus.Flags, ch chan *dbus.Call, arg0 []string) *dbus.Call
	QuerySubject(flags dbus.Flags, arg0 []string) ([]string, error)
	GoRollback(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	Rollback(flags dbus.Flags, arg0 string) error
	GoSetDefaultConfig(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	SetDefaultConfig(flags dbus.Flags, arg0 string) error
	ConnectStateChanged(cb func(op int32, state int32, target string, desc string)) (dbusutil.SignalHandlerId, error)
	ActiveVersion() proxy.PropString
	RepoUUID() proxy.PropString
	DefaultConfig() proxy.PropString
	Running() proxy.PropBool
}

type interfaceAtomicUpgrade1 struct{}

func (v *interfaceAtomicUpgrade1) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceAtomicUpgrade1) GetInterfaceName_() string {
	return "org.deepin.AtomicUpgrade1"
}

// method CancelRollback

func (v *interfaceAtomicUpgrade1) GoCancelRollback(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelRollback", flags, ch)
}

func (v *interfaceAtomicUpgrade1) CancelRollback(flags dbus.Flags) error {
	return (<-v.GoCancelRollback(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Commit

func (v *interfaceAtomicUpgrade1) GoCommit(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Commit", flags, ch, arg0)
}

func (v *interfaceAtomicUpgrade1) Commit(flags dbus.Flags, arg0 string) error {
	return (<-v.GoCommit(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method Delete

func (v *interfaceAtomicUpgrade1) GoDelete(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, arg0)
}

func (v *interfaceAtomicUpgrade1) Delete(flags dbus.Flags, arg0 string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetGrubTitle

func (v *interfaceAtomicUpgrade1) GoGetGrubTitle(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetGrubTitle", flags, ch, arg0)
}

func (*interfaceAtomicUpgrade1) StoreGetGrubTitle(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceAtomicUpgrade1) GetGrubTitle(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetGrubTitle(
		<-v.GoGetGrubTitle(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListVersion

func (v *interfaceAtomicUpgrade1) GoListVersion(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListVersion", flags, ch)
}

func (*interfaceAtomicUpgrade1) StoreListVersion(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceAtomicUpgrade1) ListVersion(flags dbus.Flags) ([]string, error) {
	return v.StoreListVersion(
		<-v.GoListVersion(flags, make(chan *dbus.Call, 1)).Done)
}

// method QuerySubject

func (v *interfaceAtomicUpgrade1) GoQuerySubject(flags dbus.Flags, ch chan *dbus.Call, arg0 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".QuerySubject", flags, ch, arg0)
}

func (*interfaceAtomicUpgrade1) StoreQuerySubject(call *dbus.Call) (arg1 []string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceAtomicUpgrade1) QuerySubject(flags dbus.Flags, arg0 []string) ([]string, error) {
	return v.StoreQuerySubject(
		<-v.GoQuerySubject(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method Rollback

func (v *interfaceAtomicUpgrade1) GoRollback(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Rollback", flags, ch, arg0)
}

func (v *interfaceAtomicUpgrade1) Rollback(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRollback(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetDefaultConfig

func (v *interfaceAtomicUpgrade1) GoSetDefaultConfig(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultConfig", flags, ch, arg0)
}

func (v *interfaceAtomicUpgrade1) SetDefaultConfig(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetDefaultConfig(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// signal StateChanged

func (v *interfaceAtomicUpgrade1) ConnectStateChanged(cb func(op int32, state int32, target string, desc string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "StateChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".StateChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var op int32
		var state int32
		var target string
		var desc string
		err := dbus.Store(sig.Body, &op, &state, &target, &desc)
		if err == nil {
			cb(op, state, target, desc)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ActiveVersion s

func (v *interfaceAtomicUpgrade1) ActiveVersion() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ActiveVersion",
	}
}

// property RepoUUID s

func (v *interfaceAtomicUpgrade1) RepoUUID() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "RepoUUID",
	}
}

// property DefaultConfig s

func (v *interfaceAtomicUpgrade1) DefaultConfig() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DefaultConfig",
	}
}

// property Running b

func (v *interfaceAtomicUpgrade1) Running() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Running",
	}
}
