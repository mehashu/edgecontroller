// Copyright 2019 Smart-Edge.com, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gorilla

import (
	"context"
	"log"

	cce "github.com/smartedgemec/controller-ce"
)

func handleCreateNodesApps(
	ctx context.Context,
	ps cce.PersistenceService,
	e cce.Persistable,
) error {
	app, err := ps.Read(ctx, e.(*cce.NodeApp).AppID, &cce.App{})
	if err != nil {
		return err
	}
	log.Printf("Loaded app %s", app.GetID())
	log.Println(app)

	nodeCC, err := connectNode(ctx, ps, e.(*cce.NodeApp))
	if err != nil {
		return err
	}

	log.Println("Connection to node established:", nodeCC.Node)

	if err := nodeCC.AppDeploySvcCli.Deploy(ctx, app.(*cce.App)); err != nil {
		return err
	}

	log.Printf("App %s deployed to node:", app.GetID())

	return nil
}

func handleCreateNodesDNSConfigs(
	ctx context.Context,
	ps cce.PersistenceService,
	e cce.Persistable,
) error {
	nodeCC, err := connectNode(ctx, ps, e.(*cce.NodeDNSConfig))
	if err != nil {
		return err
	}

	// TODO add gRPC calls
	log.Println("Connection to node established:", nodeCC.Node)

	return nil
}
