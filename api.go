package api // import "github.com/kihamo/shadow-api"

//go:generate goimports -w ./
//go:generate sh -c "cd components/api && go-bindata-assetfs -pkg=api templates/... public/..."
