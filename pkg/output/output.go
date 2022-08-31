package output



package output

import (
	//"github.com/imayberoot/goshare/pkg/goshare"
	"https://github.com/imayberoot/goshare/tree/devboi/pkg"

)

func NewOutputProviderByName(name string, conf *goshare.Config) goshare.OutputProvider {
	//We have only one outputprovider at the moment
	return NewStdoutput(conf)
}
