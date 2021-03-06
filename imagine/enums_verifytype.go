// Code generated by "enumer -type=verifyType -trimprefix=verifyType -text -transform=kebab -output enums_verifytype.go"; DO NOT EDIT.

package main

import (
	"fmt"
)

const _verifyTypeName = "errornonepurgeupdatecreate"

var _verifyTypeIndex = [...]uint8{0, 5, 9, 14, 20, 26}

func (i verifyType) String() string {
	if i < 0 || i >= verifyType(len(_verifyTypeIndex)-1) {
		return fmt.Sprintf("verifyType(%d)", i)
	}
	return _verifyTypeName[_verifyTypeIndex[i]:_verifyTypeIndex[i+1]]
}

var _verifyTypeValues = []verifyType{0, 1, 2, 3, 4}

var _verifyTypeNameToValueMap = map[string]verifyType{
	_verifyTypeName[0:5]:   0,
	_verifyTypeName[5:9]:   1,
	_verifyTypeName[9:14]:  2,
	_verifyTypeName[14:20]: 3,
	_verifyTypeName[20:26]: 4,
}

// verifyTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func verifyTypeString(s string) (verifyType, error) {
	if val, ok := _verifyTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to verifyType values", s)
}

// verifyTypeValues returns all values of the enum
func verifyTypeValues() []verifyType {
	return _verifyTypeValues
}

// IsAverifyType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i verifyType) IsAverifyType() bool {
	for _, v := range _verifyTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for verifyType
func (i verifyType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for verifyType
func (i *verifyType) UnmarshalText(text []byte) error {
	var err error
	*i, err = verifyTypeString(string(text))
	return err
}
