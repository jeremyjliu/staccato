package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jeremyliu/staccato/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
