// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
	"github.com/edgexfoundry/device-service"
	"github.com/edgexfoundry/device-service/driver"
)

const (
	serviceName string = "device-service"
)

func main() {
	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
