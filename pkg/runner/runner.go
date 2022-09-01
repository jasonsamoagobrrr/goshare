package runner

import (
	"github.com/imayberoot/ggshare/pkg/goshare"
)

func NewRunnerByName(name string, conf *goshare.Config, replay bool) goshare.RunnerProvider {
	// We have only one Runner at the moment
	return NewBasicRunner(conf, replay)
}
