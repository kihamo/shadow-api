package instance

import (
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow-api/components/api/internal"
)

func NewComponent() shadow.Component {
	return &internal.Component{}
}
