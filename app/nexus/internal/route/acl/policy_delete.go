//    \\ SPIKE: Secure your secrets with SPIFFE.
//  \\\\\ Copyright 2024-present SPIKE contributors.
// \\\\\\\ SPDX-License-Identifier: Apache-2.0

package acl

import (
	"github.com/spiffe/spike/internal/log"
	"net/http"
)

func RouteDeletePolicy(
	w http.ResponseWriter, r *http.Request, audit *log.AuditEntry,
) error {
	return nil
}
