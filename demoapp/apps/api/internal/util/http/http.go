package http

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	errMsgRequiredMissing    = "required parameter is missing"
	errMsgMinValueConstraint = "provided parameter is not respecting minimum value constraint"
	errMsgMaxValueConstraint = "provided parameter is not respecting maximum value constraint"
)

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	wHeader := w.Header()

	f, ok := i.(*os.File)
	if ok {
		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		wHeader.Set("Content-Type", http.DetectContentType(data))
		wHeader.Set("Content-Disposition", "attachment; filename="+f.Name())
		if status != nil {
			w.WriteHeader(*status)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		_, err = w.Write(data)
		return err
	}
	wHeader.Set("Content-Type", "application/json; charset=UTF-8")

	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if i != nil {
		return json.NewEncoder(w).Encode(i)
	}

	return nil
}

// ParseNumericParameter parses a numeric parameter to its respective type.
func ParseNumericParameter[T Number](param string, fn Operation[T], checks ...Constraint[T]) (T, error) {
	v, ok, err := fn(param)
	if err != nil {
		return 0, err
	}

	if !ok {
		for _, check := range checks {
			if err := check(v); err != nil {
				return 0, err
			}
		}
	}

	return v, nil
}

// ParseBoolParameter parses a string parameter to a bool
func ParseBoolParameter(param string, fn Operation[bool]) (bool, error) {
	v, _, err := fn(param)
	return v, err
}

// ParseNumericArrayParameter parses a string parameter containing array of values to its respective type.
func ParseNumericArrayParameter[T Number](param, delim string, required bool, fn Operation[T], checks ...Constraint[T]) ([]T, error) {
	if param == "" {
		if required {
			return nil, errors.New(errMsgRequiredMissing)
		}

		return nil, nil
	}

	str := strings.Split(param, delim)
	values := make([]T, len(str))

	for i, s := range str {
		v, ok, err := fn(s)
		if err != nil {
			return nil, err
		}

		if !ok {
			for _, check := range checks {
				if err := check(v); err != nil {
					return nil, err
				}
			}
		}

		values[i] = v
	}

	return values, nil
}

// parseQuery parses query parameters and returns an error if any malformed value pairs are encountered.
func ParseQuery(rawQuery string) (url.Values, error) {
	return url.ParseQuery(rawQuery)
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	_, fileHeader, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	return readFileHeaderToTempFile(fileHeader)
}

// ReadFormFilesToTempFiles reads files array data from a request form and writes it to a temporary files
func ReadFormFilesToTempFiles(r *http.Request, key string) ([]*os.File, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	files := make([]*os.File, 0, len(r.MultipartForm.File[key]))

	for _, fileHeader := range r.MultipartForm.File[key] {
		file, err := readFileHeaderToTempFile(fileHeader)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

// readFileHeaderToTempFile reads multipart.FileHeader and writes it to a temporary file
func readFileHeaderToTempFile(fileHeader *multipart.FileHeader) (*os.File, error) {
	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer formFile.Close()

	// Use .* as suffix, because the asterisk is a placeholder for the random value,
	// and the period allows consumers of this file to remove the suffix to obtain the original file name
	file, err := os.CreateTemp("", fileHeader.Filename+".*")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	_, err = io.Copy(file, formFile)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func parseTimes(param string) ([]time.Time, error) {
	splits := strings.Split(param, ",")
	times := make([]time.Time, 0, len(splits))
	for _, v := range splits {
		t, err := parseTime(v)
		if err != nil {
			return nil, err
		}
		times = append(times, t)
	}
	return times, nil
}

// parseTime will parses a string parameter into a time.Time using the RFC3339 format
func parseTime(param string) (time.Time, error) {
	if param == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, param)
}

type Number interface {
	~int32 | ~int64 | ~float32 | ~float64
}

type ParseString[T Number | string | bool] func(v string) (T, error)

// parseFloat64 parses a string parameter to an float64.
func parseFloat64(param string) (float64, error) {
	if param == "" {
		return 0, nil
	}

	return strconv.ParseFloat(param, 64)
}

// ParseFloat32 parses a string parameter to an float32.
func ParseFloat32(param string) (float32, error) {
	if param == "" {
		return 0, nil
	}

	v, err := strconv.ParseFloat(param, 32)
	return float32(v), err
}

// ParseInt64 parses a string parameter to an int64.
func ParseInt64(param string) (int64, error) {
	if param == "" {
		return 0, nil
	}

	return strconv.ParseInt(param, 10, 64)
}

// ParseInt32 parses a string parameter to an int32.
func ParseInt32(param string) (int32, error) {
	if param == "" {
		return 0, nil
	}

	val, err := strconv.ParseInt(param, 10, 32)
	return int32(val), err
}

// ParseBool parses a string parameter to an bool.
func ParseBool(param string) (bool, error) {
	if param == "" {
		return false, nil
	}

	return strconv.ParseBool(param)
}

type Operation[T Number | string | bool] func(actual string) (T, bool, error)

func WithRequire[T Number | string | bool](parse ParseString[T]) Operation[T] {
	var empty T
	return func(actual string) (T, bool, error) {
		if actual == "" {
			return empty, false, errors.New(errMsgRequiredMissing)
		}

		v, err := parse(actual)
		return v, false, err
	}
}

func WithDefaultOrParse[T Number | string | bool](def T, parse ParseString[T]) Operation[T] {
	return func(actual string) (T, bool, error) {
		if actual == "" {
			return def, true, nil
		}

		v, err := parse(actual)
		return v, false, err
	}
}

func WithParse[T Number | string | bool](parse ParseString[T]) Operation[T] {
	return func(actual string) (T, bool, error) {
		v, err := parse(actual)
		return v, false, err
	}
}

type Constraint[T Number | string | bool] func(actual T) error

func WithMinimum[T Number](expected T) Constraint[T] {
	return func(actual T) error {
		if actual < expected {
			return errors.New(errMsgMinValueConstraint)
		}

		return nil
	}
}

func WithMaximum[T Number](expected T) Constraint[T] {
	return func(actual T) error {
		if actual > expected {
			return errors.New(errMsgMaxValueConstraint)
		}

		return nil
	}
}

func IsZeroValue(val interface{}) bool {
	return val == nil || reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
}

// AssertRecurseInterfaceRequired recursively checks each struct in a slice against the callback.
// This method traverse nested slices in a preorder fashion.
func AssertRecurseInterfaceRequired[T any](obj interface{}, callback func(T) error) error {
	return AssertRecurseValueRequired(reflect.ValueOf(obj), callback)
}

// AssertRecurseValueRequired checks each struct in the nested slice against the callback.
// This method traverse nested slices in a preorder fashion. ErrTypeAssertionError is thrown if
// the underlying struct does not match type T.
func AssertRecurseValueRequired[T any](value reflect.Value, callback func(T) error) error {
	switch value.Kind() {
	// If it is a struct we check using callback
	case reflect.Struct:
		obj, ok := value.Interface().(T)
		if !ok {
			return errors.New("assertion failed")
		}

		if err := callback(obj); err != nil {
			return err
		}

	// If it is a slice we continue recursion
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := AssertRecurseValueRequired(value.Index(i), callback); err != nil {
				return err
			}
		}
	default:
		return errors.New("unknown type")
	}
	return nil
}
