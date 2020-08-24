package main

import (
	"reflect"
	"strings"

	"pkg.deepin.io/lib/dbus1"
)

// copy from pkg.deepin.io/lib/dbus1/sig.go

var (
	byteType        = reflect.TypeOf(byte(0))
	boolType        = reflect.TypeOf(false)
	uint8Type       = reflect.TypeOf(uint8(0))
	int16Type       = reflect.TypeOf(int16(0))
	uint16Type      = reflect.TypeOf(uint16(0))
	intType         = reflect.TypeOf(int(0))
	uintType        = reflect.TypeOf(uint(0))
	int32Type       = reflect.TypeOf(int32(0))
	uint32Type      = reflect.TypeOf(uint32(0))
	int64Type       = reflect.TypeOf(int64(0))
	uint64Type      = reflect.TypeOf(uint64(0))
	float64Type     = reflect.TypeOf(float64(0))
	stringType      = reflect.TypeOf("")
	signatureType   = reflect.TypeOf(dbus.Signature{})
	objectPathType  = reflect.TypeOf(dbus.ObjectPath(""))
	variantType     = reflect.TypeOf(dbus.Variant{})
	interfacesType  = reflect.TypeOf([]interface{}{})
	interfaceType   = reflect.TypeOf((*interface{})(nil)).Elem()
	unixFDType      = reflect.TypeOf(dbus.UnixFD(0))
	unixFDIndexType = reflect.TypeOf(dbus.UnixFDIndex(0))
)

var sigToType = map[byte]reflect.Type{
	'y': byteType,
	'b': boolType,
	'n': int16Type,
	'q': uint16Type,
	'i': int32Type,
	'u': uint32Type,
	'x': int64Type,
	't': uint64Type,
	'd': float64Type,
	's': stringType,
	'g': signatureType,
	'o': objectPathType,
	'v': variantType,
	'h': unixFDType,
}

// Try to read a single type from this string. If it was successful, err is nil
// and rem is the remaining unparsed part. Otherwise, err is a non-nil
// SignatureError and rem is "". depth is the current recursion depth which may
// not be greater than 64 and should be given as 0 on the first call.
func validSingle(s string, depth int) (err error, rem string) {
	if s == "" {
		return dbus.SignatureError{Sig: s, Reason: "empty signature"}, ""
	}
	if depth > 64 {
		return dbus.SignatureError{Sig: s, Reason: "container nesting too deep"}, ""
	}
	switch s[0] {
	case 'y', 'b', 'n', 'q', 'i', 'u', 'x', 't', 'd', 's', 'g', 'o', 'v', 'h':
		return nil, s[1:]
	case 'a':
		if len(s) > 1 && s[1] == '{' {
			i := findMatching(s[1:], '{', '}')
			if i == -1 {
				return dbus.SignatureError{Sig: s, Reason: "unmatched '{'"}, ""
			}
			i++
			rem = s[i+1:]
			s = s[2:i]
			if err, _ = validSingle(s[:1], depth+1); err != nil {
				return err, ""
			}
			err, nr := validSingle(s[1:], depth+1)
			if err != nil {
				return err, ""
			}
			if nr != "" {
				return dbus.SignatureError{Sig: s, Reason: "too many types in dict"}, ""
			}
			return nil, rem
		}
		return validSingle(s[1:], depth+1)
	case '(':
		i := findMatching(s, '(', ')')
		if i == -1 {
			return dbus.SignatureError{Sig: s, Reason: "unmatched ')'"}, ""
		}
		rem = s[i+1:]
		s = s[1:i]
		for err == nil && s != "" {
			err, s = validSingle(s, depth+1)
		}
		if err != nil {
			rem = ""
		}
		return
	}
	return dbus.SignatureError{Sig: s, Reason: "invalid type character"}, ""
}

func findMatching(s string, left, right rune) int {
	n := 0
	for i, v := range s {
		if v == left {
			n++
		} else if v == right {
			n--
		}
		if n == 0 {
			return i
		}
	}
	return -1
}

// typeFor returns the type of the given signature. It ignores any left over
// characters and panics if s doesn't start with a valid type signature.
func typeFor(s string) (t reflect.Type) {
	err, _ := validSingle(s, 0)
	if err != nil {
		panic(err)
	}

	if t, ok := sigToType[s[0]]; ok {
		return t
	}
	switch s[0] {
	case 'a':
		if s[1] == '{' {
			i := strings.LastIndex(s, "}")
			t = reflect.MapOf(sigToType[s[2]], typeFor(s[3:i]))
		} else {
			t = reflect.SliceOf(typeFor(s[1:]))
		}
	case '(':
		t = interfacesType
	}
	return
}

func TypeFor(s string) reflect.Type {
	return typeFor(s)
}
