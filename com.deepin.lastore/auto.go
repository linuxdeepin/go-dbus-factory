package lastore

import "errors"
import "fmt"
import "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Lastore struct {
	manager // interface com.deepin.lastore.Manager
	updater // interface com.deepin.lastore.Updater
	proxy.Object
}

func NewLastore(conn *dbus.Conn) *Lastore {
	obj := new(Lastore)
	obj.Object.Init_(conn, "com.deepin.lastore", "/com/deepin/lastore")
	return obj
}

func (obj *Lastore) Manager() *manager {
	return &obj.manager
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "com.deepin.lastore.Manager"
}

// method CleanArchives

func (v *manager) GoCleanArchives(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanArchives", flags, ch)
}

func (*manager) StoreCleanArchives(call *dbus.Call) (arg0 dbus.ObjectPath, err error) {
	err = call.Store(&arg0)
	return
}

func (v *manager) CleanArchives(flags dbus.Flags) (arg0 dbus.ObjectPath, err error) {
	return v.StoreCleanArchives(
		<-v.GoCleanArchives(flags, make(chan *dbus.Call, 1)).Done)
}

// method CleanJob

func (v *manager) GoCleanJob(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanJob", flags, ch, arg0)
}

func (v *manager) CleanJob(flags dbus.Flags, arg0 string) error {
	return (<-v.GoCleanJob(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method DistUpgrade

func (v *manager) GoDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DistUpgrade", flags, ch)
}

func (*manager) StoreDistUpgrade(call *dbus.Call) (arg0 dbus.ObjectPath, err error) {
	err = call.Store(&arg0)
	return
}

func (v *manager) DistUpgrade(flags dbus.Flags) (arg0 dbus.ObjectPath, err error) {
	return v.StoreDistUpgrade(
		<-v.GoDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

// method InstallPackage

func (v *manager) GoInstallPackage(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InstallPackage", flags, ch, arg0, arg1)
}

func (*manager) StoreInstallPackage(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *manager) InstallPackage(flags dbus.Flags, arg0 string, arg1 string) (arg2 dbus.ObjectPath, err error) {
	return v.StoreInstallPackage(
		<-v.GoInstallPackage(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method PackageDesktopPath

func (v *manager) GoPackageDesktopPath(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageDesktopPath", flags, ch, arg0)
}

func (*manager) StorePackageDesktopPath(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *manager) PackageDesktopPath(flags dbus.Flags, arg0 string) (arg1 string, err error) {
	return v.StorePackageDesktopPath(
		<-v.GoPackageDesktopPath(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method PackageExists

func (v *manager) GoPackageExists(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageExists", flags, ch, arg0)
}

func (*manager) StorePackageExists(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *manager) PackageExists(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StorePackageExists(
		<-v.GoPackageExists(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method PackageInstallable

func (v *manager) GoPackageInstallable(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageInstallable", flags, ch, arg0)
}

func (*manager) StorePackageInstallable(call *dbus.Call) (arg1 bool, err error) {
	err = call.Store(&arg1)
	return
}

func (v *manager) PackageInstallable(flags dbus.Flags, arg0 string) (arg1 bool, err error) {
	return v.StorePackageInstallable(
		<-v.GoPackageInstallable(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method PackagesDownloadSize

func (v *manager) GoPackagesDownloadSize(flags dbus.Flags, ch chan *dbus.Call, arg0 []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackagesDownloadSize", flags, ch, arg0)
}

func (*manager) StorePackagesDownloadSize(call *dbus.Call) (arg1 int64, err error) {
	err = call.Store(&arg1)
	return
}

func (v *manager) PackagesDownloadSize(flags dbus.Flags, arg0 []string) (arg1 int64, err error) {
	return v.StorePackagesDownloadSize(
		<-v.GoPackagesDownloadSize(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method PauseJob

func (v *manager) GoPauseJob(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PauseJob", flags, ch, arg0)
}

func (v *manager) PauseJob(flags dbus.Flags, arg0 string) error {
	return (<-v.GoPauseJob(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method PrepareDistUpgrade

func (v *manager) GoPrepareDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PrepareDistUpgrade", flags, ch)
}

func (*manager) StorePrepareDistUpgrade(call *dbus.Call) (arg0 dbus.ObjectPath, err error) {
	err = call.Store(&arg0)
	return
}

func (v *manager) PrepareDistUpgrade(flags dbus.Flags) (arg0 dbus.ObjectPath, err error) {
	return v.StorePrepareDistUpgrade(
		<-v.GoPrepareDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

// method RecordLocaleInfo

func (v *manager) GoRecordLocaleInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RecordLocaleInfo", flags, ch, arg0)
}

func (v *manager) RecordLocaleInfo(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRecordLocaleInfo(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method RemovePackage

func (v *manager) GoRemovePackage(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemovePackage", flags, ch, arg0, arg1)
}

func (*manager) StoreRemovePackage(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *manager) RemovePackage(flags dbus.Flags, arg0 string, arg1 string) (arg2 dbus.ObjectPath, err error) {
	return v.StoreRemovePackage(
		<-v.GoRemovePackage(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method SetAutoClean

func (v *manager) GoSetAutoClean(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoClean", flags, ch, arg0)
}

func (v *manager) SetAutoClean(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoSetAutoClean(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetCurrentX11Id

func (v *manager) GoSetCurrentX11Id(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentX11Id", flags, ch, arg0, arg1)
}

func (v *manager) SetCurrentX11Id(flags dbus.Flags, arg0 string, arg1 string) error {
	return (<-v.GoSetCurrentX11Id(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method SetLogger

func (v *manager) GoSetLogger(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogger", flags, ch, arg0, arg1, arg2)
}

func (v *manager) SetLogger(flags dbus.Flags, arg0 string, arg1 string, arg2 string) error {
	return (<-v.GoSetLogger(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetRegion

func (v *manager) GoSetRegion(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRegion", flags, ch, arg0)
}

func (v *manager) SetRegion(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetRegion(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method StartJob

func (v *manager) GoStartJob(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartJob", flags, ch, arg0)
}

func (v *manager) StartJob(flags dbus.Flags, arg0 string) error {
	return (<-v.GoStartJob(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method UpdatePackage

func (v *manager) GoUpdatePackage(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdatePackage", flags, ch, arg0, arg1)
}

func (*manager) StoreUpdatePackage(call *dbus.Call) (arg2 dbus.ObjectPath, err error) {
	err = call.Store(&arg2)
	return
}

func (v *manager) UpdatePackage(flags dbus.Flags, arg0 string, arg1 string) (arg2 dbus.ObjectPath, err error) {
	return v.StoreUpdatePackage(
		<-v.GoUpdatePackage(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method UpdateSource

func (v *manager) GoUpdateSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateSource", flags, ch)
}

func (*manager) StoreUpdateSource(call *dbus.Call) (arg0 dbus.ObjectPath, err error) {
	err = call.Store(&arg0)
	return
}

func (v *manager) UpdateSource(flags dbus.Flags) (arg0 dbus.ObjectPath, err error) {
	return v.StoreUpdateSource(
		<-v.GoUpdateSource(flags, make(chan *dbus.Call, 1)).Done)
}

// property JobList ao

func (v *manager) JobList() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "JobList",
	}
}

// property SystemArchitectures as

func (v *manager) SystemArchitectures() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "SystemArchitectures",
	}
}

// property UpgradableApps as

func (v *manager) UpgradableApps() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpgradableApps",
	}
}

// property SystemOnChanging b

func (v *manager) SystemOnChanging() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SystemOnChanging",
	}
}

// property AutoClean b

func (v *manager) AutoClean() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoClean",
	}
}

func (obj *Lastore) Updater() *updater {
	return &obj.updater
}

type updater struct{}

func (v *updater) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*updater) GetInterfaceName_() string {
	return "com.deepin.lastore.Updater"
}

// method ApplicationUpdateInfos

func (v *updater) GoApplicationUpdateInfos(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationUpdateInfos", flags, ch, arg0)
}

func (*updater) StoreApplicationUpdateInfos(call *dbus.Call) (arg1 []ApplicationUpdateInfo, err error) {
	err = call.Store(&arg1)
	return
}

func (v *updater) ApplicationUpdateInfos(flags dbus.Flags, arg0 string) (arg1 []ApplicationUpdateInfo, err error) {
	return v.StoreApplicationUpdateInfos(
		<-v.GoApplicationUpdateInfos(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method ListMirrorSources

func (v *updater) GoListMirrorSources(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListMirrorSources", flags, ch, arg0)
}

func (*updater) StoreListMirrorSources(call *dbus.Call) (arg1 []MirrorSource, err error) {
	err = call.Store(&arg1)
	return
}

func (v *updater) ListMirrorSources(flags dbus.Flags, arg0 string) (arg1 []MirrorSource, err error) {
	return v.StoreListMirrorSources(
		<-v.GoListMirrorSources(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method RestoreSystemSource

func (v *updater) GoRestoreSystemSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestoreSystemSource", flags, ch)
}

func (v *updater) RestoreSystemSource(flags dbus.Flags) error {
	return (<-v.GoRestoreSystemSource(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetAutoCheckUpdates

func (v *updater) GoSetAutoCheckUpdates(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoCheckUpdates", flags, ch, arg0)
}

func (v *updater) SetAutoCheckUpdates(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoSetAutoCheckUpdates(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetAutoDownloadUpdates

func (v *updater) GoSetAutoDownloadUpdates(flags dbus.Flags, ch chan *dbus.Call, arg0 bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoDownloadUpdates", flags, ch, arg0)
}

func (v *updater) SetAutoDownloadUpdates(flags dbus.Flags, arg0 bool) error {
	return (<-v.GoSetAutoDownloadUpdates(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method SetMirrorSource

func (v *updater) GoSetMirrorSource(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMirrorSource", flags, ch, arg0)
}

func (v *updater) SetMirrorSource(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetMirrorSource(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// property AutoCheckUpdates b

func (v *updater) AutoCheckUpdates() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoCheckUpdates",
	}
}

// property AutoDownloadUpdates b

func (v *updater) AutoDownloadUpdates() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoDownloadUpdates",
	}
}

// property MirrorSource s

func (v *updater) MirrorSource() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MirrorSource",
	}
}

// property UpdatableApps as

func (v *updater) UpdatableApps() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpdatableApps",
	}
}

// property UpdatablePackages as

func (v *updater) UpdatablePackages() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpdatablePackages",
	}
}

type Job struct {
	job // interface com.deepin.lastore.Job
	proxy.Object
}

func NewJob(conn *dbus.Conn, path dbus.ObjectPath) (*Job, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Job)
	obj.Object.Init_(conn, "com.deepin.lastore", path)
	return obj, nil
}

type job struct{}

func (v *job) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*job) GetInterfaceName_() string {
	return "com.deepin.lastore.Job"
}

// property Id s

func (v *job) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Name s

func (v *job) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Packages as

func (v *job) Packages() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Packages",
	}
}

// property Type s

func (v *job) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Status s

func (v *job) Status() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Status",
	}
}

// property Progress d

func (v *job) Progress() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Progress",
	}
}

// property Speed d

func (v *job) Speed() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Speed",
	}
}

// property Description s

func (v *job) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property Cancelable b

func (v *job) Cancelable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Cancelable",
	}
}
