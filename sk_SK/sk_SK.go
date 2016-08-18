package sk_SK

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sk_SK struct {
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

// New returns a new instance of translator for the 'sk_SK' locale
func New() locales.Translator {
	return &sk_SK{
		locale:                 "sk_SK",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0xc3, 0xa1, 0x6a}, {0x6a, 0xc3, 0xba, 0x6e}, {0x6a, 0xc3, 0xba, 0x6c}, {0x61, 0x75, 0x67}, {0x73, 0x65, 0x70}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x6a}, {0x66}, {0x6d}, {0x61}, {0x6d}, {0x6a}, {0x6a}, {0x61}, {0x73}, {0x6f}, {0x6e}, {0x64}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0xc3, 0xa1, 0x72, 0x61}, {0x66, 0x65, 0x62, 0x72, 0x75, 0xc3, 0xa1, 0x72, 0x61}, {0x6d, 0x61, 0x72, 0x63, 0x61}, {0x61, 0x70, 0x72, 0xc3, 0xad, 0x6c, 0x61}, {0x6d, 0xc3, 0xa1, 0x6a, 0x61}, {0x6a, 0xc3, 0xba, 0x6e, 0x61}, {0x6a, 0xc3, 0xba, 0x6c, 0x61}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74, 0x61}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x61}, {0x6f, 0x6b, 0x74, 0xc3, 0xb3, 0x62, 0x72, 0x61}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x61}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x61}},
		daysAbbreviated:        [][]uint8{{0x6e, 0x65}, {0x70, 0x6f}, {0x75, 0x74}, {0x73, 0x74}, {0xc5, 0xa1, 0x74}, {0x70, 0x69}, {0x73, 0x6f}},
		daysNarrow:             [][]uint8{{0x6e}, {0x70}, {0x75}, {0x73}, {0xc5, 0xa1}, {0x70}, {0x73}},
		daysShort:              [][]uint8{{0x6e, 0x65}, {0x70, 0x6f}, {0x75, 0x74}, {0x73, 0x74}, {0xc5, 0xa1, 0x74}, {0x70, 0x69}, {0x73, 0x6f}},
		daysWide:               [][]uint8{{0x6e, 0x65, 0x64, 0x65, 0xc4, 0xbe, 0x61}, {0x70, 0x6f, 0x6e, 0x64, 0x65, 0x6c, 0x6f, 0x6b}, {0x75, 0x74, 0x6f, 0x72, 0x6f, 0x6b}, {0x73, 0x74, 0x72, 0x65, 0x64, 0x61}, {0xc5, 0xa1, 0x74, 0x76, 0x72, 0x74, 0x6f, 0x6b}, {0x70, 0x69, 0x61, 0x74, 0x6f, 0x6b}, {0x73, 0x6f, 0x62, 0x6f, 0x74, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x70, 0x72, 0x65, 0x64, 0x20, 0x4b, 0x72, 0x2e}, {0x70, 0x6f, 0x20, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x70, 0x72, 0x65, 0x64, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x6f, 0x6d}, {0x70, 0x6f, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x6f, 0x76, 0x69}},
		timezones:              map[string][]uint8{"WAST": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "BOT": {0x62, 0x6f, 0x6c, 0xc3, 0xad, 0x76, 0x69, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ChST": {0x63, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CAT": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "UYT": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GFT": {0x66, 0x72, 0x61, 0x6e, 0x63, 0xc3, 0xba, 0x7a, 0x73, 0x6b, 0x6f, 0x67, 0x75, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CLST": {0xc4, 0x8d, 0x69, 0x6c, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SGT": {0x73, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "JST": {0x6a, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACST": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SRT": {0x73, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "UYST": {0x75, 0x72, 0x75, 0x67, 0x75, 0x61, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WIB": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x73, 0x6b, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ECT": {0x65, 0x6b, 0x76, 0xc3, 0xa1, 0x64, 0x6f, 0x72, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EST": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WART": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HKT": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "NZST": {0x6e, 0x6f, 0x76, 0x6f, 0x7a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CHADT": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AWST": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AWDT": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "SAST": {0x6a, 0x75, 0x68, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MEZ": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WIT": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x73, 0x6b, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ADT": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HAT": {0x6e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WITA": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa9, 0x7a, 0x73, 0x6b, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AKDT": {0x61, 0x6c, 0x6a, 0x61, 0xc5, 0xa1, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EAT": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "VET": {0x76, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MESZ": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AEST": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WEZ": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WESZ": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "LHST": {0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73, 0x20, 0x6f, 0x73, 0x74, 0x72, 0x6f, 0x76, 0x61, 0x20, 0x6c, 0x6f, 0x72, 0x64, 0x61, 0x20, 0x48, 0x6f, 0x77, 0x61}, "IST": {0x69, 0x6e, 0x64, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HAST": {0x68, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "AKST": {0x61, 0x6c, 0x6a, 0x61, 0xc5, 0xa1, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ART": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HNT": {0x6e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "TMT": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0xc3, 0xa9, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "BT": {0x62, 0x68, 0x75, 0x74, 0xc3, 0xa1, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CLT": {0xc4, 0x8d, 0x69, 0x6c, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "EDT": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "COST": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "PDT": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x74, 0x69, 0x63, 0x68, 0x6f, 0x6d, 0x6f, 0x72, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WARST": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "PST": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x74, 0x69, 0x63, 0x68, 0x6f, 0x6d, 0x6f, 0x72, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "LHDT": {0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73, 0x20, 0x6f, 0x73, 0x74, 0x72, 0x6f, 0x76, 0x61, 0x20, 0x6c, 0x6f, 0x72, 0x64, 0x61, 0x20, 0x48, 0x6f, 0x77, 0x61}, "ACWST": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MDT": {0x4d, 0x44, 0x54}, "AEDT": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "AST": {0x61, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACWDT": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "NZDT": {0x6e, 0x6f, 0x76, 0x6f, 0x7a, 0xc3, 0xa9, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CHAST": {0x63, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "TMST": {0x74, 0x75, 0x72, 0x6b, 0x6d, 0xc3, 0xa9, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MST": {0x4d, 0x53, 0x54}, "CST": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x6e, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HADT": {0x68, 0x61, 0x76, 0x61, 0x6a, 0x73, 0x6b, 0x6f, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ARST": {0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0xc3, 0xad, 0x6e, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "OESZ": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GYT": {0x67, 0x75, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "GMT": {0x67, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "JDT": {0x6a, 0x61, 0x70, 0x6f, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "COT": {0x6b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "ACDT": {0x73, 0x74, 0x72, 0x65, 0x64, 0x6f, 0x61, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x73, 0x6b, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "WAT": {0x7a, 0xc3, 0xa1, 0x70, 0x61, 0x64, 0x6f, 0x61, 0x66, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "MYT": {0x6d, 0x61, 0x6c, 0x61, 0x6a, 0x7a, 0x69, 0x6a, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "OEZ": {0x76, 0xc3, 0xbd, 0x63, 0x68, 0x6f, 0x64, 0x6f, 0x65, 0x75, 0x72, 0xc3, 0xb3, 0x70, 0x73, 0x6b, 0x79, 0x20, 0xc5, 0xa1, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "HKST": {0x68, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x73, 0x6b, 0xc3, 0xbd, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}, "CDT": {0x73, 0x65, 0x76, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x6b, 0xc3, 0xbd, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0xc3, 0xa1, 0x6c, 0x6e, 0x79, 0x20, 0x6c, 0x65, 0x74, 0x6e, 0xc3, 0xbd, 0x20, 0xc4, 0x8d, 0x61, 0x73}},
	}
}

// Locale returns the current translators string locale
func (sk *sk_SK) Locale() string {
	return sk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sk_SK'
func (sk *sk_SK) PluralsCardinal() []locales.PluralRule {
	return sk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sk_SK'
func (sk *sk_SK) PluralsOrdinal() []locales.PluralRule {
	return sk.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sk_SK'
func (sk *sk_SK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sk_SK'
func (sk *sk_SK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sk_SK'
func (sk *sk_SK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sk.CardinalPluralRule(num1, v1)
	end := sk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sk_SK' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sk.decimal) + len(sk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sk_SK' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sk *sk_SK) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(sk.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sk.percentSuffix...)

	b = append(b, sk.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sk.currencies[currency]
	l := len(s) + len(sk.decimal) + len(sk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sk_SK'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sk.currencies[currency]
	l := len(s) + len(sk.decimal) + len(sk.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sk.group) - 1; j >= 0; j-- {
					b = append(b, sk.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(sk.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sk.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, sk.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, sk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateTimeShort(t time.Time) []byte {

	b := sk.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, sk.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateTimeMedium(t time.Time) []byte {

	b := sk.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, sk.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateTimeLong(t time.Time) []byte {

	b := sk.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, sk.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'sk_SK'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (sk *sk_SK) FmtDateTimeFull(t time.Time) []byte {

	b := sk.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, sk.FmtTimeFull(t)...)

	return b
}
