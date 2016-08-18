package az_Cyrl_AZ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type az_Cyrl_AZ struct {
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

// New returns a new instance of translator for the 'az_Cyrl_AZ' locale
func New() locales.Translator {
	return &az_Cyrl_AZ{
		locale:                 "az_Cyrl_AZ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 4, 5, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e}, {0x66, 0x65, 0x76}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x79}, {0x69, 0x79, 0x6e}, {0x69, 0x79, 0x6c}, {0x61, 0x76, 0x71}, {0x73, 0x65, 0x6e}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x79}, {0x64, 0x65, 0x6b}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}, {0x37}, {0x38}, {0x39}, {0x31, 0x30}, {0x31, 0x31}, {0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e, 0x76, 0x61, 0x72}, {0x66, 0x65, 0x76, 0x72, 0x61, 0x6c}, {0x6d, 0x61, 0x72, 0x74}, {0x61, 0x70, 0x72, 0x65, 0x6c}, {0x6d, 0x61, 0x79}, {0x69, 0x79, 0x75, 0x6e}, {0x69, 0x79, 0x75, 0x6c}, {0x61, 0x76, 0x71, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x6e, 0x74, 0x79, 0x61, 0x62, 0x72}, {0x6f, 0x6b, 0x74, 0x79, 0x61, 0x62, 0x72}, {0x6e, 0x6f, 0x79, 0x61, 0x62, 0x72}, {0x64, 0x65, 0x6b, 0x61, 0x62, 0x72}},
		daysAbbreviated:        [][]uint8{{0x42, 0x2e}, {0x42, 0x2e, 0x45, 0x2e}, {0xc3, 0x87, 0x2e, 0x41, 0x2e}, {0xc3, 0x87, 0x2e}, {0x43, 0x2e, 0x41, 0x2e}, {0x43, 0x2e}, {0xc5, 0x9e, 0x2e}},
		daysNarrow:             [][]uint8{{0x37}, {0x31}, {0x32}, {0x33}, {0x34}, {0x35}, {0x36}},
		daysShort:              [][]uint8{{0x42, 0x2e}, {0x42, 0x2e, 0x45, 0x2e}, {0xc3, 0x87, 0x2e, 0x41, 0x2e}, {0xc3, 0x87, 0x2e}, {0x43, 0x2e, 0x41, 0x2e}, {0x43, 0x2e}, {0xc5, 0x9e, 0x2e}},
		daysWide:               [][]uint8{{0x62, 0x61, 0x7a, 0x61, 0x72}, {0x62, 0x61, 0x7a, 0x61, 0x72, 0x20, 0x65, 0x72, 0x74, 0xc9, 0x99, 0x73, 0x69}, {0xc3, 0xa7, 0xc9, 0x99, 0x72, 0xc5, 0x9f, 0xc9, 0x99, 0x6e, 0x62, 0xc9, 0x99, 0x20, 0x61, 0x78, 0xc5, 0x9f, 0x61, 0x6d, 0xc4, 0xb1}, {0xc3, 0xa7, 0xc9, 0x99, 0x72, 0xc5, 0x9f, 0xc9, 0x99, 0x6e, 0x62, 0xc9, 0x99}, {0x63, 0xc3, 0xbc, 0x6d, 0xc9, 0x99, 0x20, 0x61, 0x78, 0xc5, 0x9f, 0x61, 0x6d, 0xc4, 0xb1}, {0x63, 0xc3, 0xbc, 0x6d, 0xc9, 0x99}, {0xc5, 0x9f, 0xc9, 0x99, 0x6e, 0x62, 0xc9, 0x99}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x65, 0x2e, 0xc9, 0x99, 0x2e}, {0x62, 0x2e, 0x65, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x65, 0x72, 0x61, 0x6d, 0xc4, 0xb1, 0x7a, 0x64, 0x61, 0x6e, 0x20, 0xc9, 0x99, 0x76, 0x76, 0xc9, 0x99, 0x6c}, {0x65, 0x72, 0x61, 0x6d, 0xc4, 0xb1, 0x7a}},
		timezones:              map[string][]uint8{"OEZ": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ACDT": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AEST": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "SGT": {0x53, 0x69, 0x6e, 0x71, 0x61, 0x70, 0x75, 0x72, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HAST": {0x48, 0x61, 0x76, 0x61, 0x79, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "OESZ": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HKT": {0x48, 0x6f, 0x6e, 0x71, 0x20, 0x4b, 0x6f, 0x6e, 0x71, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "NZDT": {0x59, 0x65, 0x6e, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x73, 0xc4, 0xb1, 0x7a, 0x20, 0x51, 0x76, 0x69, 0x61, 0x6e, 0x61, 0x73, 0xc4, 0xb1, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WAT": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HNT": {0x4e, 0x79, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HAT": {0x4e, 0x79, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WAST": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "PDT": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x61, 0x6b, 0x69, 0x74, 0x20, 0x4f, 0x6b, 0x65, 0x61, 0x6e, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ACWDT": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "UYT": {0x55, 0x72, 0x75, 0x71, 0x76, 0x61, 0x79, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "EST": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AWDT": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WIT": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0xc4, 0xb0, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WART": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WARST": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "EAT": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "JDT": {0x59, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CLT": {0xc3, 0x87, 0x69, 0x6c, 0x69, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CDT": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CAT": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x61, 0x75, 0x20, 0x59, 0x61, 0x79, 0x20, 0x76, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "MDT": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x44, 0x61, 0xc4, 0x9f, 0x6c, 0xc4, 0xb1, 0x71, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "SAST": {0x43, 0xc9, 0x99, 0x6e, 0x75, 0x62, 0x69, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HKST": {0x48, 0x6f, 0x6e, 0x71, 0x20, 0x4b, 0x6f, 0x6e, 0x71, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AEDT": {0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ACWST": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CLST": {0xc3, 0x87, 0x69, 0x6c, 0x69, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WITA": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0xc4, 0xb0, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AKDT": {0x41, 0x6c, 0x79, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WIB": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0xc4, 0xb0, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AWST": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WEZ": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "JST": {0x59, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "BT": {0x42, 0x75, 0x74, 0x61, 0x6e, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "IST": {0x48, 0x69, 0x6e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ACST": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x61, 0x75, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74}, "MST": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x44, 0x61, 0xc4, 0x9f, 0x6c, 0xc4, 0xb1, 0x71, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "TMT": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0xc9, 0x99, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "MEZ": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CHAST": {0xc3, 0x87, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "HADT": {0x48, 0x61, 0x76, 0x61, 0x79, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "WESZ": {0x51, 0xc9, 0x99, 0x72, 0x62, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "∅∅∅": {0x42, 0x72, 0x61, 0x7a, 0x69, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ChST": {0xc3, 0x87, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "GYT": {0x51, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x79, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "EDT": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0xc5, 0x9e, 0xc9, 0x99, 0x72, 0x71, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "GMT": {0x51, 0x72, 0x69, 0x6e, 0x76, 0x69, 0xc3, 0xa7, 0x20, 0x4f, 0x72, 0x74, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "UYST": {0x55, 0x72, 0x75, 0x71, 0x76, 0x61, 0x79, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "MESZ": {0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "PST": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x61, 0x6b, 0x69, 0x74, 0x20, 0x4f, 0x6b, 0x65, 0x61, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "AKST": {0x41, 0x6c, 0x79, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "TMST": {0x54, 0xc3, 0xbc, 0x72, 0x6b, 0x6d, 0xc9, 0x99, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CST": {0xc5, 0x9e, 0x69, 0x6d, 0x61, 0x6c, 0x69, 0x20, 0x4d, 0xc9, 0x99, 0x72, 0x6b, 0xc9, 0x99, 0x7a, 0x69, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "NZST": {0x59, 0x65, 0x6e, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x79, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}, "CHADT": {0xc3, 0x87, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x20, 0x59, 0x61, 0x79, 0x20, 0x56, 0x61, 0x78, 0x74, 0xc4, 0xb1}},
	}
}

// Locale returns the current translators string locale
func (az *az_Cyrl_AZ) Locale() string {
	return az.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'az_Cyrl_AZ'
func (az *az_Cyrl_AZ) PluralsCardinal() []locales.PluralRule {
	return az.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'az_Cyrl_AZ'
func (az *az_Cyrl_AZ) PluralsOrdinal() []locales.PluralRule {
	return az.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'az_Cyrl_AZ'
func (az *az_Cyrl_AZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'az_Cyrl_AZ'
func (az *az_Cyrl_AZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100
	iMod1000 := i % 1000

	if (iMod10 == 1 || iMod10 == 2 || iMod10 == 5 || iMod10 == 7 || iMod10 == 8) || (iMod100 == 20 || iMod100 == 50 || iMod100 == 70 || iMod100 == 80) {
		return locales.PluralRuleOne
	} else if (iMod10 == 3 || iMod10 == 4) || (iMod1000 == 100 || iMod1000 == 200 || iMod1000 == 300 || iMod1000 == 400 || iMod1000 == 500 || iMod1000 == 600 || iMod1000 == 700 || iMod1000 == 800 || iMod1000 == 900) {
		return locales.PluralRuleFew
	} else if (i == 0) || (iMod10 == 6) || (iMod100 == 40 || iMod100 == 60 || iMod100 == 90) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'az_Cyrl_AZ'
func (az *az_Cyrl_AZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := az.CardinalPluralRule(num1, v1)
	end := az.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'az_Cyrl_AZ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(az.decimal) + len(az.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'az_Cyrl_AZ' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (az *az_Cyrl_AZ) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(az.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, az.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(az.decimal) + len(az.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(az.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, az.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'az_Cyrl_AZ'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(az.decimal) + len(az.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(az.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, az.currencyNegativePrefix[j])
		}

		b = append(b, az.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(az.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, az.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, az.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := az.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateTimeShort(t time.Time) []byte {

	b := az.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, az.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateTimeMedium(t time.Time) []byte {

	b := az.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, az.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateTimeLong(t time.Time) []byte {

	b := az.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, az.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'az_Cyrl_AZ'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (az *az_Cyrl_AZ) FmtDateTimeFull(t time.Time) []byte {

	b := az.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, az.FmtTimeFull(t)...)

	return b
}
