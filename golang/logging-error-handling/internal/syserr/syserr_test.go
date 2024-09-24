package syserr_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"loggingerrorhandling/internal/syserr"
)

func TestGetStackFormatted(t *testing.T) {
	type setup struct {
		makeError func() error
	}
	type verify struct {
		err error
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return call stack": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						nestedFn := func() error {
							return syserr.NewInternal("hello")
						}

						return nestedFn()
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				var sysErr *syserr.Error
				assert.True(t, errors.As(verify.err, &sysErr))

				stack := sysErr.GetStackFormatted()

				assert.True(t, len(stack) > 0)
			},
		},
		"Should wrap a generic error and return a message": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return syserr.Wrap(errors.New("bar"), syserr.InternalCode, "foo")
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				msg := verify.err.Error()
				assert.Equal(t, "foo: bar", msg)
			},
		},
		"Should be wrap-able by a generic error and return a message": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(syserr.New(syserr.InternalCode, "bar"), "foo")
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				msg := verify.err.Error()
				assert.Equal(t, "foo: bar", msg)
			},
		},
		"Should support long chain of messages when wrapped": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								errors.Wrap(
									syserr.NewInternal("wheel"),
									"baz",
								),
								syserr.InternalCode,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				msg := verify.err.Error()
				assert.Equal(t, "foo: bar: baz: wheel", msg)
			},
		},
		// todo: should support errors.Is
		// todo: should support errors.As
		// todo: should support code extraction
		// todo: should support stack extraction
		// todo: support causer?
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			err := setup.makeError()

			testCase.verifyFunc(t, setup, &verify{
				err: err,
			})
		})
	}
}
