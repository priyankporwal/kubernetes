/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_node

import (
	"flag"
)

var kubeletAddress = flag.String("kubelet-address", "http://127.0.0.1:10255", "Host and port of the kubelet")

var disableKubenet = flag.Bool("disable-kubenet", false, "If true, start kubelet without kubenet")
var buildServices = flag.Bool("build-services", true, "If true, build local executables")
var startServices = flag.Bool("start-services", true, "If true, start local node services")
var stopServices = flag.Bool("stop-services", true, "If true, stop local node services after running tests")

type SharedContext struct {
	PodConfigPath string
}
