package de_DE

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type de_DE struct {
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

// New returns a new instance of translator for the 'de_DE' locale
func New() locales.Translator {
	return &de_DE{
		locale:                 "de_DE",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x2e}, {0x46, 0x65, 0x62, 0x2e}, {0x4d, 0xc3, 0xa4, 0x72, 0x7a}, {0x41, 0x70, 0x72, 0x2e}, {0x4d, 0x61, 0x69}, {0x4a, 0x75, 0x6e, 0x69}, {0x4a, 0x75, 0x6c, 0x69}, {0x41, 0x75, 0x67, 0x2e}, {0x53, 0x65, 0x70, 0x2e}, {0x4f, 0x6b, 0x74, 0x2e}, {0x4e, 0x6f, 0x76, 0x2e}, {0x44, 0x65, 0x7a, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72}, {0x4d, 0xc3, 0xa4, 0x72, 0x7a}, {0x41, 0x70, 0x72, 0x69, 0x6c}, {0x4d, 0x61, 0x69}, {0x4a, 0x75, 0x6e, 0x69}, {0x4a, 0x75, 0x6c, 0x69}, {0x41, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x53, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x4f, 0x6b, 0x74, 0x6f, 0x62, 0x65, 0x72}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x44, 0x65, 0x7a, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x53, 0x6f, 0x2e}, {0x4d, 0x6f, 0x2e}, {0x44, 0x69, 0x2e}, {0x4d, 0x69, 0x2e}, {0x44, 0x6f, 0x2e}, {0x46, 0x72, 0x2e}, {0x53, 0x61, 0x2e}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x44}, {0x4d}, {0x44}, {0x46}, {0x53}},
		daysShort:              [][]uint8{{0x53, 0x6f, 0x2e}, {0x4d, 0x6f, 0x2e}, {0x44, 0x69, 0x2e}, {0x4d, 0x69, 0x2e}, {0x44, 0x6f, 0x2e}, {0x46, 0x72, 0x2e}, {0x53, 0x61, 0x2e}},
		daysWide:               [][]uint8{{0x53, 0x6f, 0x6e, 0x6e, 0x74, 0x61, 0x67}, {0x4d, 0x6f, 0x6e, 0x74, 0x61, 0x67}, {0x44, 0x69, 0x65, 0x6e, 0x73, 0x74, 0x61, 0x67}, {0x4d, 0x69, 0x74, 0x74, 0x77, 0x6f, 0x63, 0x68}, {0x44, 0x6f, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x74, 0x61, 0x67}, {0x46, 0x72, 0x65, 0x69, 0x74, 0x61, 0x67}, {0x53, 0x61, 0x6d, 0x73, 0x74, 0x61, 0x67}},
		periodsAbbreviated:     [][]uint8{{0x76, 0x6f, 0x72, 0x6d, 0x2e}, {0x6e, 0x61, 0x63, 0x68, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x76, 0x6d, 0x2e}, {0x6e, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x76, 0x6f, 0x72, 0x6d, 0x2e}, {0x6e, 0x61, 0x63, 0x68, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		erasWide:               [][]uint8{{0x76, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}, {0x6e, 0x2e, 0x20, 0x43, 0x68, 0x72, 0x2e}},
		timezones:              map[string][]uint8{"CST": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x49, 0x6e, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "WITA": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "ACST": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "SAST": {0x53, 0xc3, 0xbc, 0x64, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "AEDT": {0x4f, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "WESZ": {0x57, 0x65, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "ECT": {0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "CAT": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "WAST": {0x57, 0x65, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "OEZ": {0x4f, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "WEZ": {0x57, 0x65, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "MDT": {0x4d, 0x61, 0x63, 0x61, 0x75, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "WIT": {0x4f, 0x73, 0x74, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "WART": {0x57, 0x65, 0x73, 0x74, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "NZST": {0x4e, 0x65, 0x75, 0x73, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "JST": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ACWST": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2d, 0x2f, 0x57, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "EDT": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4f, 0x73, 0x74, 0x6b, 0xc3, 0xbc, 0x73, 0x74, 0x65, 0x6e, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "OESZ": {0x4f, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "ACDT": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "MST": {0x4d, 0x61, 0x63, 0x61, 0x75, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "NZDT": {0x4e, 0x65, 0x75, 0x73, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x65, 0x6e, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "AWDT": {0x57, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "CDT": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x49, 0x6e, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "PDT": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x57, 0x65, 0x73, 0x74, 0x6b, 0xc3, 0xbc, 0x73, 0x74, 0x65, 0x6e, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "CLST": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "WARST": {0x57, 0x65, 0x73, 0x74, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "GMT": {0x4d, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "WIB": {0x57, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "AWST": {0x57, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "PST": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x57, 0x65, 0x73, 0x74, 0x6b, 0xc3, 0xbc, 0x73, 0x74, 0x65, 0x6e, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "HNT": {0x4e, 0x65, 0x75, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "ACWDT": {0x5a, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2d, 0x2f, 0x57, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "EST": {0x4e, 0x6f, 0x72, 0x64, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4f, 0x73, 0x74, 0x6b, 0xc3, 0xbc, 0x73, 0x74, 0x65, 0x6e, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "MESZ": {0x4d, 0x69, 0x74, 0x74, 0x65, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "∅∅∅": {0x41, 0x63, 0x72, 0x65, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "HAT": {0x4e, 0x65, 0x75, 0x66, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x79, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x7a, 0xc3, 0xb6, 0x73, 0x69, 0x73, 0x63, 0x68, 0x2d, 0x47, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x2d, 0x5a, 0x65, 0x69, 0x74}, "WAT": {0x57, 0x65, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "AEST": {0x4f, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "EAT": {0x4f, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x5a, 0x65, 0x69, 0x74}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x65, 0x6e, 0x2d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "CLT": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "MEZ": {0x4d, 0x69, 0x74, 0x74, 0x65, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa4, 0x69, 0x73, 0x63, 0x68, 0x65, 0x20, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x7a, 0x65, 0x69, 0x74}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x2d, 0x53, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x7a, 0x65, 0x69, 0x74}},
	}
}

// Locale returns the current translators string locale
func (de *de_DE) Locale() string {
	return de.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'de_DE'
func (de *de_DE) PluralsCardinal() []locales.PluralRule {
	return de.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'de_DE'
func (de *de_DE) PluralsOrdinal() []locales.PluralRule {
	return de.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'de_DE'
func (de *de_DE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'de_DE'
func (de *de_DE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'de_DE'
func (de *de_DE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := de.CardinalPluralRule(num1, v1)
	end := de.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'de_DE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(de.decimal) + len(de.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, de.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, de.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, de.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'de_DE' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (de *de_DE) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(de.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, de.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, de.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, de.percentSuffix...)

	b = append(b, de.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := de.currencies[currency]
	l := len(s) + len(de.decimal) + len(de.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, de.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, de.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, de.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, de.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, de.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'de_DE'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := de.currencies[currency]
	l := len(s) + len(de.decimal) + len(de.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, de.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, de.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, de.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, de.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, de.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, de.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateMedium(t time.Time) []byte {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, de.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, de.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, de.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, de.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, de.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, de.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, de.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, de.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, de.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, de.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := de.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateTimeShort(t time.Time) []byte {

	b := de.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, de.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateTimeMedium(t time.Time) []byte {

	b := de.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, de.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateTimeLong(t time.Time) []byte {

	b := de.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, de.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'de_DE'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (de *de_DE) FmtDateTimeFull(t time.Time) []byte {

	b := de.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, de.FmtTimeFull(t)...)

	return b
}
