package kl_GL

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type kl_GL struct {
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

// New returns a new instance of translator for the 'kl_GL' locale
func New() locales.Translator {
	return &kl_GL{
		locale:             "kl_GL",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		decimal:            []byte{0x2c},
		group:              []byte{0x2e},
		minus:              []byte{0x2d},
		percent:            []byte{0x25},
		perMille:           []byte{0xe2, 0x80, 0xb0},
		timeSeparator:      []byte{0x3a},
		inifinity:          []byte{0xe2, 0x88, 0x9e},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:      []byte{0xc2, 0xa0},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x6a}, {0x6a, 0x75, 0x6e}, {0x6a, 0x75, 0x6c}, {0x61, 0x75, 0x67}, {0x73, 0x65, 0x70}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x63}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:         [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x69}, {0x66, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x69}, {0x6d, 0x61, 0x72, 0x74, 0x73, 0x69}, {0x61, 0x70, 0x72, 0x69, 0x6c, 0x69}, {0x6d, 0x61, 0x6a, 0x69}, {0x6a, 0x75, 0x6e, 0x69}, {0x6a, 0x75, 0x6c, 0x69}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74, 0x75, 0x73, 0x69}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x69}, {0x6f, 0x6b, 0x74, 0x6f, 0x62, 0x65, 0x72, 0x69}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x69}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x69}},
		daysAbbreviated:    [][]uint8{{0x73, 0x61, 0x62}, {0x61, 0x74, 0x61}, {0x6d, 0x61, 0x72}, {0x70, 0x69, 0x6e}, {0x73, 0x69, 0x73}, {0x74, 0x61, 0x6c}, {0x61, 0x72, 0x66}},
		daysNarrow:         [][]uint8{{0x53}, {0x41}, {0x4d}, {0x50}, {0x53}, {0x54}, {0x41}},
		daysShort:          [][]uint8{{0x73, 0x61, 0x62}, {0x61, 0x74, 0x61}, {0x6d, 0x61, 0x72}, {0x70, 0x69, 0x6e}, {0x73, 0x69, 0x73}, {0x74, 0x61, 0x6c}, {0x61, 0x72, 0x66}},
		daysWide:           [][]uint8{{0x73, 0x61, 0x62, 0x61, 0x61, 0x74}, {0x61, 0x74, 0x61, 0x61, 0x73, 0x69, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}, {0x6d, 0x61, 0x72, 0x6c, 0x75, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}, {0x70, 0x69, 0x6e, 0x67, 0x61, 0x73, 0x75, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}, {0x73, 0x69, 0x73, 0x61, 0x6d, 0x61, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}, {0x74, 0x61, 0x6c, 0x6c, 0x69, 0x6d, 0x61, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}, {0x61, 0x72, 0x66, 0x69, 0x6e, 0x69, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x71}},
		periodsAbbreviated: [][]uint8{{0x75, 0x2e, 0x74, 0x2e}, {0x75, 0x2e, 0x6b, 0x2e}},
		periodsWide:        [][]uint8{{0x75, 0x6c, 0x6c, 0x6f, 0x71, 0x65, 0x71, 0x71, 0x61, 0x74, 0x61, 0x2d, 0x74, 0x75, 0x6e, 0x67, 0x61, 0x61}, {0x75, 0x6c, 0x6c, 0x6f, 0x71, 0x65, 0x71, 0x71, 0x61, 0x74, 0x61, 0x2d, 0x6b, 0x69, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x61}},
		erasAbbreviated:    [][]uint8{{0x4b, 0x72, 0x2e, 0x69, 0x6e, 0x2e, 0x73, 0x69, 0x2e}, {0x4b, 0x72, 0x2e, 0x69, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x67, 0x2e}},
		erasNarrow:         [][]uint8{{0x4b, 0x72, 0x2e, 0x73, 0x2e}, {0x4b, 0x72, 0x2e, 0x6b, 0x2e}},
		erasWide:           [][]uint8{{0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x73, 0x69, 0x70, 0x20, 0x69, 0x6e, 0x75, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x61, 0x20, 0x73, 0x69, 0x6f, 0x72, 0x6e, 0x61, 0x67, 0x75, 0x74}, {0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x73, 0x69, 0x70, 0x20, 0x69, 0x6e, 0x75, 0x6e, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x61, 0x20, 0x6b, 0x69, 0x6e, 0x67, 0x6f, 0x72, 0x6e, 0x61, 0x67, 0x75, 0x74}},
		timezones:          map[string][]uint8{"AWST": {0x41, 0x57, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "GFT": {0x47, 0x46, 0x54}, "WIT": {0x57, 0x49, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HAST": {0x48, 0x41, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "SGT": {0x53, 0x47, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "ChST": {0x43, 0x68, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "JDT": {0x4a, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "SRT": {0x53, 0x52, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "BT": {0x42, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "AEST": {0x41, 0x45, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "IST": {0x49, 0x53, 0x54}, "CDT": {0x43, 0x44, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "ART": {0x41, 0x52, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "MST": {0x4d, 0x53, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "AST": {0x41, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "VET": {0x56, 0x45, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "HKT": {0x48, 0x4b, 0x54}, "EAT": {0x45, 0x41, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "CST": {0x43, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "EST": {0x45, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}},
	}
}

// Locale returns the current translators string locale
func (kl *kl_GL) Locale() string {
	return kl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kl_GL'
func (kl *kl_GL) PluralsCardinal() []locales.PluralRule {
	return kl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kl_GL'
func (kl *kl_GL) PluralsOrdinal() []locales.PluralRule {
	return kl.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kl_GL'
func (kl *kl_GL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kl_GL'
func (kl *kl_GL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kl_GL'
func (kl *kl_GL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kl_GL' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kl.decimal) + len(kl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kl_GL' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kl *kl_GL) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kl.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kl.percentSuffix...)

	b = append(b, kl.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kl.currencies[currency]
	l := len(s) + len(kl.decimal) + len(kl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kl.group[0])
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
		b = append(b, kl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kl_GL'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kl.currencies[currency]
	l := len(s) + len(kl.decimal) + len(kl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kl.group[0])
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

		b = append(b, kl.minus[0])

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
			b = append(b, kl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, kl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, kl.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kl.periodsAbbreviated[0]...)
	} else {
		b = append(b, kl.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kl.periodsAbbreviated[0]...)
	} else {
		b = append(b, kl.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kl.periodsAbbreviated[0]...)
	} else {
		b = append(b, kl.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kl.periodsAbbreviated[0]...)
	} else {
		b = append(b, kl.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateTimeShort(t time.Time) []byte {

	b := kl.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, kl.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateTimeMedium(t time.Time) []byte {

	b := kl.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, kl.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateTimeLong(t time.Time) []byte {

	b := kl.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, kl.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'kl_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kl *kl_GL) FmtDateTimeFull(t time.Time) []byte {

	b := kl.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, kl.FmtTimeFull(t)...)

	return b
}
