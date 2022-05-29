package gdrive

var MimeTypes = newMimeTypes()

type MimeType struct {
	name           string
	googleMimeType string
}

type MimeTypeRegistry struct {
	CSV MimeType
}

var mimeTypeMap = map[string]MimeType{
	"csv": {name: "text/csv", googleMimeType: "application/vnd.google-apps.spreadsheet"},
}

// newMimeTypes based on https://stackoverflow.com/a/49551778/4154982
func newMimeTypes() MimeTypeRegistry {
	return MimeTypeRegistry{
		CSV: mimeTypeMap["csv"]}
}

func (mt MimeType) Name() string {
	return mt.name
}

func (mt MimeType) GoogleMimeType() string {
	return mt.googleMimeType
}
