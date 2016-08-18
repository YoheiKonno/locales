package vi_VN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type vi_VN struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsShort            [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'vi_VN' locale
func New() locales.Translator {
	return &vi_VN{
		locale:                 "vi_VN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x74, 0x68, 0x67, 0x20, 0x31}, {0x74, 0x68, 0x67, 0x20, 0x32}, {0x74, 0x68, 0x67, 0x20, 0x33}, {0x74, 0x68, 0x67, 0x20, 0x34}, {0x74, 0x68, 0x67, 0x20, 0x35}, {0x74, 0x68, 0x67, 0x20, 0x36}, {0x74, 0x68, 0x67, 0x20, 0x37}, {0x74, 0x68, 0x67, 0x20, 0x38}, {0x74, 0x68, 0x67, 0x20, 0x39}, {0x74, 0x68, 0x67, 0x20, 0x31, 0x30}, {0x74, 0x68, 0x67, 0x20, 0x31, 0x31}, {0x74, 0x68, 0x67, 0x20, 0x31, 0x32}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}, {0x37}, {0x38}, {0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x31}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x32}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x33}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x34}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x35}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x36}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x37}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x38}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x39}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x31, 0x30}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x31, 0x31}, {0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x20, 0x31, 0x32}},
		daysAbbreviated:        [][]uint8{{0x43, 0x4e}, {0x54, 0x68, 0x20, 0x32}, {0x54, 0x68, 0x20, 0x33}, {0x54, 0x68, 0x20, 0x34}, {0x54, 0x68, 0x20, 0x35}, {0x54, 0x68, 0x20, 0x36}, {0x54, 0x68, 0x20, 0x37}},
		daysNarrow:             [][]uint8{{0x43, 0x4e}, {0x54, 0x32}, {0x54, 0x33}, {0x54, 0x34}, {0x54, 0x35}, {0x54, 0x36}, {0x54, 0x37}},
		daysShort:              [][]uint8{{0x43, 0x4e}, {0x54, 0x32}, {0x54, 0x33}, {0x54, 0x34}, {0x54, 0x35}, {0x54, 0x36}, {0x54, 0x37}},
		daysWide:               [][]uint8{{0x43, 0x68, 0xe1, 0xbb, 0xa7, 0x20, 0x4e, 0x68, 0xe1, 0xba, 0xad, 0x74}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x48, 0x61, 0x69}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x42, 0x61}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x54, 0xc6, 0xb0}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x4e, 0xc4, 0x83, 0x6d}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x53, 0xc3, 0xa1, 0x75}, {0x54, 0x68, 0xe1, 0xbb, 0xa9, 0x20, 0x42, 0xe1, 0xba, 0xa3, 0x79}},
		periodsAbbreviated:     [][]uint8{{0x53, 0x41}, {0x43, 0x48}},
		periodsNarrow:          [][]uint8{{0x73}, {0x63}},
		periodsWide:            [][]uint8{{0x53, 0x41}, {0x43, 0x48}},
		erasAbbreviated:        [][]uint8{{0x74, 0x72, 0x2e, 0x20, 0x43, 0x4e}, {0x73, 0x61, 0x75, 0x20, 0x43, 0x4e}},
		erasNarrow:             [][]uint8{{0x74, 0x72, 0x2e, 0x20, 0x43, 0x4e}, {0x73, 0x61, 0x75, 0x20, 0x43, 0x4e}},
		erasWide:               [][]uint8{{0x74, 0x72, 0x2e, 0x20, 0x43, 0x4e}, {0x73, 0x61, 0x75, 0x20, 0x43, 0x4e}},
		timezones:              map[string][]uint8{"AWDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "MESZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0xc3, 0x82, 0x75}, "HKST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x48, 0xe1, 0xbb, 0x93, 0x6e, 0x67, 0x20, 0x4b, 0xc3, 0xb4, 0x6e, 0x67}, "ACDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "AEST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "AKST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "AKDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "WAT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x50, 0x68, 0x69}, "SRT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "WARST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x6d, 0xc3, 0xb9, 0x61, 0x20, 0x68, 0xc3, 0xa8, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x74, 0xc3, 0xa2, 0x79, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "TMST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "IST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0xe1, 0xba, 0xa4, 0x6e, 0x20, 0xc4, 0x90, 0xe1, 0xbb, 0x99}, "CST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x63, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x74, 0x72, 0x75, 0x6e, 0x67}, "CDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x62, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0xc3, 0xa0, 0x79, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x74, 0x72, 0x75, 0x6e, 0x67}, "ADT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x68, 0xc3, 0xa8, 0x20, 0xc4, 0x90, 0xe1, 0xba, 0xa1, 0x69, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x44, 0xc6, 0xb0, 0xc6, 0xa1, 0x6e, 0x67}, "WITA": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "NZDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "ACST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "COST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "UYST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "ACWST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "ACWDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "LHST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "BOT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "ECT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "MEZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0xc3, 0x82, 0x75}, "BT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}, "CHAST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "CAT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x50, 0x68, 0x69}, "NZST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "ChST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "OESZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0xc3, 0x82, 0x75}, "GFT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x74, 0x68, 0x75, 0xe1, 0xbb, 0x99, 0x63, 0x20, 0x50, 0x68, 0xc3, 0xa1, 0x70}, "EDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x62, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0xc3, 0xa0, 0x79, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0xc4, 0x91, 0xc3, 0xb4, 0x6e, 0x67}, "WART": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x63, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x6d, 0xc3, 0xa2, 0x79, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "OEZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0xc3, 0x82, 0x75}, "WIB": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "MYT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "PDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x62, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0xc3, 0xa0, 0x79, 0x20, 0x54, 0x68, 0xc3, 0xa1, 0x69, 0x20, 0x42, 0xc3, 0xac, 0x6e, 0x68, 0x20, 0x44, 0xc6, 0xb0, 0xc6, 0xa1, 0x6e, 0x67}, "UYT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "HNT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "TMT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ART": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "COT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "SGT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65}, "AST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0xc4, 0x90, 0xe1, 0xba, 0xa1, 0x69, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x44, 0xc6, 0xb0, 0xc6, 0xa1, 0x6e, 0x67}, "JST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4e, 0x68, 0xe1, 0xba, 0xad, 0x74, 0x20, 0x42, 0xe1, 0xba, 0xa3, 0x6e}, "SAST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4e, 0x61, 0x6d, 0x20, 0x50, 0x68, 0x69}, "CLT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "CLST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "EST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x63, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0xc4, 0x91, 0xc3, 0xb4, 0x6e, 0x67}, "GYT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "LHDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WESZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x68, 0xc3, 0xa8, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0xc3, 0x82, 0x75}, "HAT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "GMT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x54, 0x72, 0x75, 0x6e, 0x67, 0x20, 0x62, 0xc3, 0xac, 0x6e, 0x68, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "HAST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "WAST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x50, 0x68, 0x69}, "MST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x63, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x6e, 0xc3, 0xba, 0x69}, "VET": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "WIT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "CHADT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "HADT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e}, "ARST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "∅∅∅": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x41, 0x63, 0x72, 0x65}, "WEZ": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0xc3, 0x82, 0x75}, "AWST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x54, 0xc3, 0xa2, 0x79, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "MDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x62, 0x61, 0x6e, 0x20, 0x6e, 0x67, 0xc3, 0xa0, 0x79, 0x20, 0x6d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0x6e, 0xc3, 0xba, 0x69}, "HKT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x43, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x48, 0xe1, 0xbb, 0x93, 0x6e, 0x67, 0x20, 0x4b, 0xc3, 0xb4, 0x6e, 0x67}, "AEDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4d, 0x69, 0xe1, 0xbb, 0x81, 0x6e, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0x4e, 0xc6, 0xb0, 0xe1, 0xbb, 0x9b, 0x63, 0x20, 0xc3, 0x9a, 0x63}, "EAT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0xc4, 0x90, 0xc3, 0xb4, 0x6e, 0x67, 0x20, 0x50, 0x68, 0x69}, "PST": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x63, 0x68, 0x75, 0xe1, 0xba, 0xa9, 0x6e, 0x20, 0x54, 0x68, 0xc3, 0xa1, 0x69, 0x20, 0x42, 0xc3, 0xac, 0x6e, 0x68, 0x20, 0x44, 0xc6, 0xb0, 0xc6, 0xa1, 0x6e, 0x67}, "JDT": {0x47, 0x69, 0xe1, 0xbb, 0x9d, 0x20, 0x4d, 0xc3, 0xb9, 0x61, 0x20, 0x48, 0xc3, 0xa8, 0x20, 0x4e, 0x68, 0xe1, 0xba, 0xad, 0x74, 0x20, 0x42, 0xe1, 0xba, 0xa3, 0x6e}},
	}
}

// Locale returns the current translators string locale
func (vi *vi_VN) Locale() string {
	return vi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vi_VN'
func (vi *vi_VN) PluralsCardinal() []locales.PluralRule {
	return vi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vi_VN'
func (vi *vi_VN) PluralsOrdinal() []locales.PluralRule {
	return vi.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vi_VN'
func (vi *vi_VN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vi_VN'
func (vi *vi_VN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vi_VN'
func (vi *vi_VN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vi_VN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(vi.decimal) + len(vi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vi_VN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vi *vi_VN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(vi.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, vi.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(vi.decimal) + len(vi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, vi.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vi_VN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(vi.decimal) + len(vi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, vi.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, vi.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, vi.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x27, 0x4e, 0x67, 0xc3, 0xa0, 0x79, 0x27, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x74, 0x68, 0xc3, 0xa1, 0x6e, 0x67, 0x27, 0x20}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6e, 0xc4, 0x83, 0x6d, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, vi.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c}...)
	b = append(b, []byte{0x27, 0x6e, 0x67, 0xc3, 0xa0, 0x79, 0x27, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6e, 0xc4, 0x83, 0x6d, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateTimeShort(t time.Time) []byte {

	b := vi.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, vi.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateTimeMedium(t time.Time) []byte {

	b := vi.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, vi.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateTimeLong(t time.Time) []byte {

	b := vi.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, vi.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'vi_VN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vi *vi_VN) FmtDateTimeFull(t time.Time) []byte {

	b := vi.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, vi.FmtTimeFull(t)...)

	return b
}
