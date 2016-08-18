package zu

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type zu struct {
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

// New returns a new instance of translator for the 'zu' locale
func New() locales.Translator {
	return &zu{
		locale:                 "zu",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e}, {0x46, 0x65, 0x62}, {0x4d, 0x61, 0x73}, {0x45, 0x70, 0x68}, {0x4d, 0x65, 0x79}, {0x4a, 0x75, 0x6e}, {0x4a, 0x75, 0x6c}, {0x41, 0x67, 0x61}, {0x53, 0x65, 0x70}, {0x4f, 0x6b, 0x74}, {0x4e, 0x6f, 0x76}, {0x44, 0x69, 0x73}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x45}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4a, 0x61, 0x6e, 0x75, 0x77, 0x61, 0x72, 0x69}, {0x46, 0x65, 0x62, 0x72, 0x75, 0x77, 0x61, 0x72, 0x69}, {0x4d, 0x61, 0x73, 0x68, 0x69}, {0x45, 0x70, 0x68, 0x72, 0x65, 0x6c, 0x69}, {0x4d, 0x65, 0x79, 0x69}, {0x4a, 0x75, 0x6e, 0x69}, {0x4a, 0x75, 0x6c, 0x61, 0x79, 0x69}, {0x41, 0x67, 0x61, 0x73, 0x74, 0x69}, {0x53, 0x65, 0x70, 0x74, 0x68, 0x65, 0x6d, 0x62, 0x61}, {0x4f, 0x6b, 0x74, 0x68, 0x6f, 0x62, 0x61}, {0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x61}, {0x44, 0x69, 0x73, 0x65, 0x6d, 0x62, 0x61}},
		daysAbbreviated:        [][]uint8{{0x53, 0x6f, 0x6e}, {0x4d, 0x73, 0x6f}, {0x42, 0x69, 0x6c}, {0x54, 0x68, 0x61}, {0x53, 0x69, 0x6e}, {0x48, 0x6c, 0x61}, {0x4d, 0x67, 0x71}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x42}, {0x54}, {0x53}, {0x48}, {0x4d}},
		daysShort:              [][]uint8{{0x53, 0x6f, 0x6e}, {0x4d, 0x73, 0x6f}, {0x42, 0x69, 0x6c}, {0x54, 0x68, 0x61}, {0x53, 0x69, 0x6e}, {0x48, 0x6c, 0x61}, {0x4d, 0x67, 0x71}},
		daysWide:               [][]uint8{{0x49, 0x53, 0x6f, 0x6e, 0x74, 0x6f}, {0x55, 0x4d, 0x73, 0x6f, 0x6d, 0x62, 0x75, 0x6c, 0x75, 0x6b, 0x6f}, {0x55, 0x4c, 0x77, 0x65, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69}, {0x55, 0x4c, 0x77, 0x65, 0x73, 0x69, 0x74, 0x68, 0x61, 0x74, 0x68, 0x75}, {0x55, 0x4c, 0x77, 0x65, 0x73, 0x69, 0x6e, 0x65}, {0x55, 0x4c, 0x77, 0x65, 0x73, 0x69, 0x68, 0x6c, 0x61, 0x6e, 0x75}, {0x55, 0x4d, 0x67, 0x71, 0x69, 0x62, 0x65, 0x6c, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		timezones:              map[string][]uint8{"PST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "ART": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "HNT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "OEZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "SRT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x65}, "WEZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "WAT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4e, 0x74, 0x73, 0x68, 0x6f, 0x6e, 0x61, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "CLST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x68, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f}, "NZST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "WIT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "CHAST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AWDT": {0x49, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "UYST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f}, "GFT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x46, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x61}, "EDT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "WARST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4e, 0x79, 0x61, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x20, 0x6e, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f}, "MESZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "ACWDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x57, 0x65, 0x73, 0x74}, "ARST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f}, "TMT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "TMST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ACST": {0x65, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x69, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "AST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "ECT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "MDT": {0x4d, 0x44, 0x54}, "NZDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x65, 0x77, 0x20, 0x5a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "MST": {0x4d, 0x53, 0x54}, "COT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "HKT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "HADT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "ADT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "HAT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "MEZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "CLT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x68, 0x69, 0x6c, 0x65, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "CDT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "WITA": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "CAT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x61}, "EAT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4d, 0x70, 0x75, 0x6d, 0x61, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "PDT": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "COST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f}, "GMT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x4d, 0x65, 0x61, 0x6e}, "CHADT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ChST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "BT": {0x65, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65}, "GYT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "SGT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65}, "SAST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4e, 0x69, 0x6e, 0x67, 0x69, 0x7a, 0x69, 0x6d, 0x75, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "VET": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "JST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "EST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "LHDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "AEST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74}, "AEDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x45, 0x61, 0x73, 0x74}, "HAST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "AWST": {0x49, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "OESZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "HKST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "AKDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x73, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69}, "WAST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4e, 0x74, 0x73, 0x68, 0x6f, 0x6e, 0x61, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61}, "JDT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4a, 0x61, 0x70, 0x61, 0x6e}, "WART": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x4e, 0x79, 0x61, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x20, 0x6e, 0x65, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "IST": {0x65, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x49, 0x6e, 0x64, 0x69, 0x61, 0x20, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x54, 0x69, 0x6d, 0x65}, "LHST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65}, "BOT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "ACWST": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x57, 0x65, 0x73, 0x74, 0x20, 0x65, 0x73, 0x69, 0x76, 0x61, 0x6d, 0x69, 0x6c, 0x65}, "MYT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "AKST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "WIB": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61}, "WESZ": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x68, 0x6c, 0x6f, 0x62, 0x6f, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x57, 0x65, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65}, "UYT": {0x49, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}, "ACDT": {0x65, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x69, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x20, 0x54, 0x69, 0x6d, 0x65}, "CST": {0x69, 0x73, 0x69, 0x6b, 0x68, 0x61, 0x74, 0x68, 0x69, 0x20, 0x73, 0x61, 0x73, 0x65, 0x2d, 0x4e, 0x6f, 0x72, 0x74, 0x68, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x65, 0x73, 0x69, 0x6a, 0x77, 0x61, 0x79, 0x65, 0x6c, 0x65, 0x6b, 0x69, 0x6c, 0x65}},
	}
}

// Locale returns the current translators string locale
func (zu *zu) Locale() string {
	return zu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'zu'
func (zu *zu) PluralsCardinal() []locales.PluralRule {
	return zu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'zu'
func (zu *zu) PluralsOrdinal() []locales.PluralRule {
	return zu.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'zu'
func (zu *zu) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'zu'
func (zu *zu) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'zu'
func (zu *zu) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := zu.CardinalPluralRule(num1, v1)
	end := zu.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'zu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(zu.decimal) + len(zu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'zu' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (zu *zu) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(zu.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, zu.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zu.currencies[currency]
	l := len(s) + len(zu.decimal) + len(zu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
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
		b = append(b, zu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, zu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'zu'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := zu.currencies[currency]
	l := len(s) + len(zu.decimal) + len(zu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, zu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, zu.group[0])
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

		for j := len(zu.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, zu.currencyNegativePrefix[j])
		}

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
			b = append(b, zu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, zu.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, zu.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, zu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, zu.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, zu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zu.periodsAbbreviated[0]...)
	} else {
		b = append(b, zu.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zu.periodsAbbreviated[0]...)
	} else {
		b = append(b, zu.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zu.periodsAbbreviated[0]...)
	} else {
		b = append(b, zu.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, zu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, zu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, zu.periodsAbbreviated[0]...)
	} else {
		b = append(b, zu.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := zu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateTimeShort(t time.Time) []byte {

	b := zu.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, zu.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateTimeMedium(t time.Time) []byte {

	b := zu.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, zu.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateTimeLong(t time.Time) []byte {

	b := zu.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, zu.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'zu'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (zu *zu) FmtDateTimeFull(t time.Time) []byte {

	b := zu.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, zu.FmtTimeFull(t)...)

	return b
}
