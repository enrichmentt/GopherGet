package storage

import "github.com/enrichmentt/GopherGet/internal/models"

var packages = map[string][]models.PackageMatch{
	"test": {{
		Package: struct {
			Id        string "json:\"Id\""
			Name      string "json:\"Name\""
			Version   string "json:\"Version\""
			Publisher string "json:\"Publisher\""
		}{Id: "test", Name: "Test App", Version: "1.0.0", Publisher: "GopherGet"},
	}},
}

var manifests = map[string]*models.Manifest{
	"test": {
		Id:        "Test",
		Name:      "Test App",
		Version:   "1.0.0",
		Publisher: "GopherGet",
		Installers: []models.Installer{{
			InstallerUrl:  "https://example.com/test.txt",
			InstallerType: "exe",
			Silent:        "/s",
		}},
	},
}

func SearchPackages(query string) []models.PackageMatch {
	if pkg, ok := packages[query]; ok {
		return pkg
	}
	return []models.PackageMatch{}
}
