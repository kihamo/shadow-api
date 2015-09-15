package api // import "github.com/kihamo/shadow-api"

//go:generate goimports -w ./
//go:generate sh -c "cd service && go-bindata-assetfs -pkg=service templates/... public/..."
