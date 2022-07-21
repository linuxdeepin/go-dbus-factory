package lastore

type ApplicationUpdateInfo struct {
	Id             string
	Name           string
	Icon           string
	CurrentVersion string
	LastVersion    string
}

type MirrorSource struct {
	Id   string
	Url  string
	Name string
}
