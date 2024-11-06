//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package net

import (
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/spiffe/spike/internal/config"
	"github.com/spiffe/spike/internal/entity/v1/reqres"
	"github.com/spiffe/spike/internal/net"
	"strconv"
)

// DeleteSecret deletes a secret from SPIKE Nexus.
func DeleteSecret(source *workloadapi.X509Source,
	path string, versions []string) error {
	var vv = []int{}

	if len(versions) == 0 {
		vv = []int{0}
	}

	for _, version := range versions {
		v, e := strconv.Atoi(version)
		if e != nil {
			continue
		}
		vv = append(vv, v)
	}

	r := reqres.SecretDeleteRequest{
		Path:     path,
		Versions: vv,
	}

	mr, err := json.Marshal(r)
	if err != nil {
		return errors.Join(
			errors.New("deleteSecret: I am having problem generating the payload"),
			err,
		)
	}

	client, err := net.CreateMtlsClient(source, config.IsNexus)
	if err != nil {
		return err
	}

	_, err = net.Post(client, urlSecretDelete, mr)

	return err
}
