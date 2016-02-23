// JSONStructs project JSONStructs.go
package JSONStructs

/*
	{
		"couchdb":"Welcome",
		"uuid":"d3ab89c42df6962d4ffb2c12d24d6576",
		"version":"1.6.1",
		"vendor": {
			"version":"1.6.1",
			"name":"The Apache Software Foundation"
		}
	}
*/

type SlashResponse struct {
	CouchDB string            `json:"couchdb"`
	UUID    string            `json:"uuid"`
	Version string            `json:"version"`
	Vendor  map[string]string `json:"vendor"`
}

type UUIDSResponse struct {
	UUIDS []string `json:"uuids"`
}

type OKResponse struct{
	OK bool	`json:"ok"`	
}

type DocOKResponse struct{
	ID string	`json:"id"`
	OK bool		`json:"ok"`
	REV string	`json:"rev"`
}

type DBErrorResponse struct{
	ErrorMsg string	`json:"error"`
	Reason string	`json:"reason"`
}
