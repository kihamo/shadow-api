package api // import "github.com/kihamo/shadow-api"

//go:generate goimports -w ./
//go:generate /bin/bash -c "cd components/api/internal && go-bindata-assetfs -pkg=internal templates/... assets/..."
