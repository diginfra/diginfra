package diginfra

import _ "embed"

//go:embed diginfra-usage-example.yml
var referenceUsageFileContents []byte

func GetReferenceUsageFileContents() *[]byte {
	return &referenceUsageFileContents
}
