package pkgs

import (
	"errors"
	"github.com/XuHandsome/stocks/pkgs/config"
	"github.com/XuHandsome/stocks/pkgs/dashboard"
)

func Run(mainConfig config.MainConfig) error {
	d := dashboard.New(mainConfig)
	if err := d.Start(); err != nil {
		d.Stop()
		return err
	}

	return errors.New("exit")
}
