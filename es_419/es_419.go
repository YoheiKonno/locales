package es_419

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type es_419 struct {
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

// New returns a new instance of translator for the 'es_419' locale
func New() locales.Translator {
	return &es_419{
		locale:             "es_419",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     []locales.PluralRule{6},
		decimal:            []byte{0x2e},
		group:              []byte{0x2c},
		minus:              []byte{0x2d},
		percent:            []byte{0x25},
		perMille:           []byte{0xe2, 0x80, 0xb0},
		timeSeparator:      []byte{0x3a},
		inifinity:          []byte{0xe2, 0x88, 0x9e},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0xc2, 0xa3}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x42, 0x73, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:      []byte{0xc2, 0xa0},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0x65, 0x6e, 0x65, 0x2e}, {0x66, 0x65, 0x62, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x61, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x79, 0x2e}, {0x6a, 0x75, 0x6e, 0x2e}, {0x6a, 0x75, 0x6c, 0x2e}, {0x61, 0x67, 0x6f, 0x2e}, {0x73, 0x65, 0x70, 0x2e}, {0x6f, 0x63, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x69, 0x63, 0x2e}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0x65}, {0x66}, {0x6d}, {0x61}, {0x6d}, {0x6a}, {0x6a}, {0x61}, {0x73}, {0x6f}, {0x6e}, {0x64}},
		monthsWide:         [][]uint8{[]uint8(nil), {0x65, 0x6e, 0x65, 0x72, 0x6f}, {0x66, 0x65, 0x62, 0x72, 0x65, 0x72, 0x6f}, {0x6d, 0x61, 0x72, 0x7a, 0x6f}, {0x61, 0x62, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x79, 0x6f}, {0x6a, 0x75, 0x6e, 0x69, 0x6f}, {0x6a, 0x75, 0x6c, 0x69, 0x6f}, {0x61, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x73, 0x65, 0x70, 0x74, 0x69, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x6f, 0x63, 0x74, 0x75, 0x62, 0x72, 0x65}, {0x6e, 0x6f, 0x76, 0x69, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0x69, 0x63, 0x69, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:    [][]uint8{{0x64, 0x6f, 0x6d, 0x2e}, {0x6c, 0x75, 0x6e, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x6d, 0x69, 0xc3, 0xa9, 0x2e}, {0x6a, 0x75, 0x65, 0x2e}, {0x76, 0x69, 0x65, 0x2e}, {0x73, 0xc3, 0xa1, 0x62, 0x2e}},
		daysNarrow:         [][]uint8{{0x64}, {0x6c}, {0x6d}, {0x6d}, {0x6a}, {0x76}, {0x73}},
		daysShort:          [][]uint8{{0x44, 0x4f}, {0x4c, 0x55}, {0x4d, 0x41}, {0x4d, 0x49}, {0x4a, 0x55}, {0x56, 0x49}, {0x53, 0x41}},
		daysWide:           [][]uint8{{0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x6f}, {0x6c, 0x75, 0x6e, 0x65, 0x73}, {0x6d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x6d, 0x69, 0xc3, 0xa9, 0x72, 0x63, 0x6f, 0x6c, 0x65, 0x73}, {0x6a, 0x75, 0x65, 0x76, 0x65, 0x73}, {0x76, 0x69, 0x65, 0x72, 0x6e, 0x65, 0x73}, {0x73, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated: [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsNarrow:      [][]uint8{[]uint8(nil), []uint8(nil)},
		periodsWide:        [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		erasAbbreviated:    [][]uint8{{0x61, 0x2e, 0x20, 0x43, 0x2e}, {0x64, 0x2e, 0x20, 0x43, 0x2e}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0x61, 0x6e, 0x74, 0x65, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x64, 0x65, 0x73, 0x70, 0x75, 0xc3, 0xa9, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:          map[string][]uint8{"SAST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x64, 0xc3, 0xa1, 0x66, 0x72, 0x69, 0x63, 0x61}, "ECT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "AWDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "UYT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "JDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "ART": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "WART": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "COST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "HADT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0xc3, 0xa1, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x61, 0x73}, "AWST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "UYST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79}, "ADT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "MST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x73, 0x20, 0x4d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x61, 0x73}, "COT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x63, 0x72, 0x65}, "SRT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "MDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x73, 0x20, 0x4d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x61, 0x73}, "ACWST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ACDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "AEDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "LHST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WAT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HNT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "MESZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "GYT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "CAT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "TMT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "WIT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ACST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "PST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "GFT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "WAST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "PDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "EST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "BT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa1, 0x6e}, "WARST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "MYT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x73, 0x69, 0x61}, "AST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "EDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "OESZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AEST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "SGT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "EAT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ARST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "NZDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x75, 0x65, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "AKDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "WIB": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HAT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "ACWDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "MEZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "HKST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "GMT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "CHADT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "TMST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "OEZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WEZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CLT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "CLST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "NZST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x75, 0x65, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "WESZ": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "BOT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "JST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "IST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "WITA": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ChST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "VET": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "HKT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CHAST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "LHDT": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HAST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0xc3, 0xa1, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x61, 0x73}, "AKST": {0x68, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}},
	}
}

// Locale returns the current translators string locale
func (es *es_419) Locale() string {
	return es.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'es_419'
func (es *es_419) PluralsCardinal() []locales.PluralRule {
	return es.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'es_419'
func (es *es_419) PluralsOrdinal() []locales.PluralRule {
	return es.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'es_419'
func (es *es_419) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'es_419'
func (es *es_419) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'es_419'
func (es *es_419) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'es_419' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(es.decimal) + len(es.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'es_419' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (es *es_419) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(es.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, es.percentSuffix...)

	b = append(b, es.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(es.decimal) + len(es.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
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

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'es_419'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(es.decimal) + len(es.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
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

		b = append(b, es.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, es.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := es.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateTimeShort(t time.Time) []byte {

	b := es.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, es.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateTimeMedium(t time.Time) []byte {

	b := es.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, es.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateTimeLong(t time.Time) []byte {

	b := es.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, es.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'es_419'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (es *es_419) FmtDateTimeFull(t time.Time) []byte {

	b := es.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, es.FmtTimeFull(t)...)

	return b
}
