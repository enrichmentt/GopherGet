package models

type InformationResponse struct {
	Data struct {
		SourceIdentifier       string   `json:"SourceIdentifier"`
		ServerSupportedVersion []string `json:"ServerSupportedVersion"`
	} `json:"Data"`
}

type SearchRequest struct {
	Query string `json:"Query"`
}

type SearchResponse struct {
	Data []PackageMatch `json:"Data"`
}

type PackageMatch struct {
	Package struct {
		Id        string `json:"Id"`
		Name      string `json:"Name"`
		Version   string `json:"Version"`
		Publisher string `json:"Publisher"`
	} `json:"Package"`
}

type ManifestResponse struct {
	Data Manifest `json:"Manifest"`
}

type Manifest struct {
	Id         string      `json:"Id"`
	Name       string      `json:"Name"`
	Version    string      `json:"Version"`
	Publisher  string      `json:"Publisher"`
	Installers []Installer `json:"Installers"`
}

type Installer struct {
	InstallerUrl       string `json:"InstallerUrl"`
	InstallerSha256    string `json:"InstallerSha256"`
	InstallerType      string `json:"InstallerType"`
	Silent             string `json:"Silent"`
	SilentWithProgress string `json:"SilentWuthProgress"`
}
