# Run tests in check section
# disable for bootstrapping
%bcond_with check


%global pkgname github.com/linuxdeepin/go-dbus-factory

#2020-04-26增加%global with_debug 1
%global with_debug 1

%if 0%{?with_debug}
%global debug_package   %{nil}
%endif

%global sname golang-github-linuxdeepin-go-dbus-factory

Name:           golang-github-linuxdeepin-dbus-factory
Version:        1.9.2
Release:        1
Summary:        GO DBus factory for Deepin Desktop Environment
License:        GPLv3
URL:            %{gourl}
Source0:        %{sname}-%{version}.tar.gz
BuildRequires:  compiler(go-compiler)
BuildRequires:  go-srpm-macros

%description
%{summary}.

%package -n %{sname}-devel
Summary:        %{summary}
BuildArch:      noarch


%description -n %{sname}-devel
%{summary}.



%prep
%forgeautosetup


%install
install -d -p %{buildroot}%{gopath}/src/%{pkgname}/
for file in $(find . -iname "*.go" | grep -v "_tool") ; do
    install -d -p %{buildroot}%{gopath}/src/%{pkgname}/$(dirname $file)
    cp -pav $file %{buildroot}%{gopath}/src/%{pkgname}/$file
    echo "%%{gopath}/src/%%{pkgname}/$file" >> devel.file-list
done

cp -pav README.md %{buildroot}%{gopath}/src/%{pkgname}/README.md
cp -pav CHANGELOG.md %{buildroot}%{gopath}/src/%{pkgname}/CHANGELOG.md
echo "%%{gopath}/src/%%{pkgname}/README.md" >> devel.file-list
echo "%%{gopath}/src/%%{pkgname}/CHANGELOG.md" >> devel.file-list

%if %{with check}
%check
%gochecks
%endif

%files -n %{sname}-devel -f devel.file-list
%doc README.md CHANGELOG.md

%changelog
* Thu Mar 18 2021 uoser <uoser@uniontech.com> - 1.9.2-1
- Update to 1.9.2
