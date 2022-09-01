package output

import (
	"github.com/imayberoot/ggshare/pkg/goshare"
)

func NewOutputProviderByName(name string, conf *goshare.Config) goshare.OutputProvider {
	//We have only one outputprovider at the moment
	return NewStdoutput(conf)
}
