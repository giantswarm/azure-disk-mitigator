package key

import (
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
)

func ToEvent(v interface{}) (corev1.Event, error) {
	if v == nil {
		return corev1.Event{}, microerror.Maskf(wrongTypeError, "expected non-nil, got %#v'", v)
	}

	e, ok := v.(*corev1.Event)
	if !ok {
		return corev1.Event{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", e, v)
	}

	c := e.DeepCopy()

	return *c, nil
}

func EventIsWarning(event corev1.Event) bool {
	return event.Type == corev1.EventTypeWarning
}

func EventReason(event corev1.Event) string {
	return event.Reason
}
