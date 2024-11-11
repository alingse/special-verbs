package create

import (
	"reflect"
	"testing"

	"github.com/alingse/special-verbs/schema"
)

func TestAddSpecialVerb(t *testing.T) {
	testCases := map[string]struct {
		verb     string
		resource schema.GroupResource
	}{
		"existing verb": {
			verb:     "use",
			resource: schema.GroupResource{Group: "my.custom.io", Resource: "one"},
		},
		"new verb": {
			verb:     "new",
			resource: schema.GroupResource{Group: "my.custom.io", Resource: "two"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			AddSpecialVerb(tc.verb, tc.resource)
			resources, ok := specialVerbs[tc.verb]
			t.Logf("run for `%s` and verb `%+v` got ok %+v resources %+v \n", name, tc.verb, ok, resources)

			if !ok {
				t.Errorf("missing expected verb: %s", tc.verb)
			}

			for _, res := range resources {
				if reflect.DeepEqual(tc.resource, res) {
					t.Logf("check deepEqual for res %+v with tc.resource %+v got true", res, tc.resource)
					return
				}
				t.Logf("check deepEqual for res %+v with tc.resource %+v got false", res, tc.resource)
			}
			t.Errorf("missing expected resource:%#v", tc.resource)
		})
	}
}
