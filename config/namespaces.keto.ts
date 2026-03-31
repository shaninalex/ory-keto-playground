// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { Namespace, SubjectSet, Context } from "@ory/keto-namespace-types"

class User implements Namespace {}

class Role implements Namespace {
  related: {
    member: User[]
  }
}

class PurchaseOrder implements Namespace {
    related: {
        creator: SubjectSet<Role, "member">[]
        viewer: SubjectSet<Role, "member">[]
    }
}
