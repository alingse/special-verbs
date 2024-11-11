package create

import (
	"github.com/alingse/special-verbs/schema"
)

var (
	specialVerbs = map[string][]schema.GroupResource{
		"use": {
			{
				Group:    "policy",
				Resource: "podsecuritypolicies",
			},
			{
				Group:    "extensions",
				Resource: "podsecuritypolicies",
			},
		},
		"bind": {
			{
				Group:    "rbac.authorization.k8s.io",
				Resource: "roles",
			},
			{
				Group:    "rbac.authorization.k8s.io",
				Resource: "clusterroles",
			},
		},
		"escalate": {
			{
				Group:    "rbac.authorization.k8s.io",
				Resource: "roles",
			},
			{
				Group:    "rbac.authorization.k8s.io",
				Resource: "clusterroles",
			},
		},
		"impersonate": {
			{
				Group:    "",
				Resource: "users",
			},
			{
				Group:    "",
				Resource: "serviceaccounts",
			},
			{
				Group:    "",
				Resource: "groups",
			},
			{
				Group:    "authentication.k8s.io",
				Resource: "userextras",
			},
		},
	}
)

// AddSpecialVerb allows the addition of items to the `specialVerbs` map for non-k8s native resources.
func AddSpecialVerb(verb string, gr schema.GroupResource) {
	resources, ok := specialVerbs[verb]
	if !ok {
		resources = make([]schema.GroupResource, 1)
	}
	resources = append(resources, gr)
	specialVerbs[verb] = resources
}
