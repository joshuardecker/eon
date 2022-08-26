package version

const VERSION string = "0.1B"

// Get and return the version as a string
func GetVersion() string {

	return VERSION
}

// GEt abd return the version as a byte array.
func GetVersionBytes() []byte {

	return []byte(VERSION)
}
