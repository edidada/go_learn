package main

// Version identifies the release documented by this directory.
const Version = "1.6"

// LanguageChangeCount is zero because Go 1.6 introduced no language-specification
// changes. It is intentionally executable so this conclusion is covered by CI.
func LanguageChangeCount() int {
	return 0
}
