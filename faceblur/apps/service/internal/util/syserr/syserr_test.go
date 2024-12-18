package syserr_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"service/internal/util/syserr"
)

var (
	ErrNotFound = errors.New("resource not found")
	SysInternal = syserr.New(syserr.InternalCode, "foo")
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
				stack := syserr.GetStackFormatted(verify.err)

				assert.True(t, len(stack) > 0)
			},
		},
		"Should return call stack of a wrapped error": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						nestedFn := func() error {
							return errors.Wrap(errors.Wrap(syserr.NewInternal("hello"), "foo"), "bar")
						}

						return nestedFn()
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				stack := syserr.GetStackFormatted(verify.err)

				assert.True(t, len(stack) > 0)
			},
		},
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

func TestAs(t *testing.T) {
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
		"Should support wrapped errors": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								ErrNotFound,
								syserr.InternalCode,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				var customErr *syserr.Error
				ok := errors.As(verify.err, &customErr)
				assert.True(t, ok)
				assert.Equal(t, "bar: resource not found", customErr.Error())
				assert.Equal(t, "foo: bar: resource not found", verify.err.Error())
			},
		},
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

func TestIs(t *testing.T) {
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
		"Should detect generic error": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								ErrNotFound,
								syserr.InternalCode,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				isNotFound := errors.Is(verify.err, ErrNotFound)
				assert.True(t, isNotFound)
			},
		},
		"Should detect custom error": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								SysInternal,
								syserr.BadInputCode,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				isInternal := errors.Is(verify.err, SysInternal)
				assert.True(t, isInternal)
			},
		},
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

func TestError(t *testing.T) {
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

func TestGetCode(t *testing.T) {
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
		"Should get the code from wrapped error": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								ErrNotFound,
								syserr.BadInputCode,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				code := syserr.GetCode(verify.err)
				assert.Equal(t, syserr.BadInputCode, code)
			},
		},
		"Should accept nil": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return nil
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				code := syserr.GetCode(verify.err)
				assert.Equal(t, syserr.InternalCode, code)
			},
		},
		"Should return the first code met": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return syserr.Wrap(
							syserr.Wrap(
								ErrNotFound,
								syserr.BadInputCode,
								"bar",
							),
							syserr.NotFoundCode,
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				code := syserr.GetCode(verify.err)
				assert.Equal(t, syserr.NotFoundCode, code)
			},
		},
		"Should accept generic errors": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							errors.Wrap(
								ErrNotFound,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				code := syserr.GetCode(verify.err)
				assert.Equal(t, syserr.InternalCode, code)
			},
		},
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

func TestGetFields(t *testing.T) {
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
		"Should get the fields from wrapped error": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								ErrNotFound,
								syserr.BadInputCode,
								"bar",
								syserr.F("one", "two"),
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				fields := syserr.GetFields(verify.err)
				assert.Len(t, fields, 1)
				assert.Equal(t, "one", fields[0].Key)
				assert.Equal(t, "two", fields[0].Value)
			},
		},
		"Should get the union of fields from multiple wrapped errors": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							syserr.Wrap(
								syserr.Wrap(
									errors.Wrap(
										syserr.NewInternal("internal", syserr.F("five", "six")),
										"baz",
									),
									syserr.NotFoundCode,
									"bar",
									syserr.F("three", "four"),
								),
								syserr.BadInputCode,
								"bar",
								syserr.F("one", "two"),
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				fields := syserr.GetFields(verify.err)
				assert.Len(t, fields, 3)
				assert.Equal(t, "one", fields[0].Key)
				assert.Equal(t, "two", fields[0].Value)
				assert.Equal(t, "three", fields[1].Key)
				assert.Equal(t, "four", fields[1].Value)
				assert.Equal(t, "five", fields[2].Key)
				assert.Equal(t, "six", fields[2].Value)
			},
		},
		"Should accept nil": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return nil
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				fields := syserr.GetFields(verify.err)
				assert.Len(t, fields, 0)
			},
		},
		"Should accept generic errors": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{
					makeError: func() error {
						return errors.Wrap(
							errors.Wrap(
								ErrNotFound,
								"bar",
							),
							"foo",
						)
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				fields := syserr.GetFields(verify.err)
				assert.Len(t, fields, 0)
			},
		},
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
