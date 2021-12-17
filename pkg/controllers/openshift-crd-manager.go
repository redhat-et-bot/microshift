/*
Copyright © 2021 Microshift Contributors

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
package controllers

import (
	"context"

	"github.com/openshift/microshift/pkg/assets"
	"github.com/openshift/microshift/pkg/config"

	"github.com/sirupsen/logrus"
)

type OpenShiftCRDManager struct {
	cfg *config.MicroshiftConfig
}

func NewOpenShiftCRDManager(cfg *config.MicroshiftConfig) *OpenShiftCRDManager {
	s := &OpenShiftCRDManager{}
	s.cfg = cfg
	return s
}

func (s *OpenShiftCRDManager) Name() string { return "openshift-crd-manager" }
func (s *OpenShiftCRDManager) Dependencies() []string {
	return []string{"kube-apiserver"}
}

func (s *OpenShiftCRDManager) Run(ctx context.Context, ready chan<- struct{}, stopped chan<- struct{}) error {
	defer close(ready)
	// To-DO add readiness check
	if err := assets.ApplyCRDs(s.cfg); err != nil {
		logrus.Errorf("%s unable to apply default CRDs: %v", s.Name(), err)
	}
	logrus.Infof("%s applied default CRDs", s.Name())

	return ctx.Err()
}