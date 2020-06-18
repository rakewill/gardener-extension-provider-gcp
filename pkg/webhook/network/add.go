// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"github.com/gardener/gardener-extension-provider-gcp/pkg/gcp"
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/extensions/pkg/webhook/network"

	"github.com/gardener/gardener-extension-networking-calico/pkg/calico"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var logger = log.Log.WithName("networking-calico-webhook")

// AddToManager creates a webhook and adds it to the manager.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")
	return network.New(mgr, network.Args{
		CloudProvider:   gcp.Type,
		NetworkProvider: calico.Type,
		Types:           []runtime.Object{&extensionsv1alpha1.Network{}},
		Mutator:         network.NewMutator(logger, mutateNetworkConfig),
	})
}
