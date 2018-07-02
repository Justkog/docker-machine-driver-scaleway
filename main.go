package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/justkog/docker-machine-driver-scaleway/driver"
)

func main() {
	plugin.RegisterDriver(scaleway.NewDriver("", ""))
}
