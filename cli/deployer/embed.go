package deployer

import "embed"

//go:embed templates/*.tmpl
var templatesFS embed.FS
