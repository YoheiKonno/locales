package da_GL

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type da_GL struct {
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

// New returns a new instance of translator for the 'da_GL' locale
func New() locales.Translator {
	return &da_GL{
		locale:                 "da_GL",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x2e},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x61, 0x70, 0x72, 0x2e}, {0x6d, 0x61, 0x6a}, {0x6a, 0x75, 0x6e, 0x2e}, {0x6a, 0x75, 0x6c, 0x2e}, {0x61, 0x75, 0x67, 0x2e}, {0x73, 0x65, 0x70, 0x2e}, {0x6f, 0x6b, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4a}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x4a}, {0x4a}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6a, 0x61, 0x6e, 0x75, 0x61, 0x72}, {0x66, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72}, {0x6d, 0x61, 0x72, 0x74, 0x73}, {0x61, 0x70, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x6a}, {0x6a, 0x75, 0x6e, 0x69}, {0x6a, 0x75, 0x6c, 0x69}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x6f, 0x6b, 0x74, 0x6f, 0x62, 0x65, 0x72}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x65, 0x72}},
		daysAbbreviated:        [][]uint8{{0x73, 0xc3, 0xb8, 0x6e, 0x2e}, {0x6d, 0x61, 0x6e, 0x2e}, {0x74, 0x69, 0x72, 0x2e}, {0x6f, 0x6e, 0x73, 0x2e}, {0x74, 0x6f, 0x72, 0x2e}, {0x66, 0x72, 0x65, 0x2e}, {0x6c, 0xc3, 0xb8, 0x72, 0x2e}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x54}, {0x4f}, {0x54}, {0x46}, {0x4c}},
		daysShort:              [][]uint8{{0x73, 0xc3, 0xb8}, {0x6d, 0x61}, {0x74, 0x69}, {0x6f, 0x6e}, {0x74, 0x6f}, {0x66, 0x72}, {0x6c, 0xc3, 0xb8}},
		daysWide:               [][]uint8{{0x73, 0xc3, 0xb8, 0x6e, 0x64, 0x61, 0x67}, {0x6d, 0x61, 0x6e, 0x64, 0x61, 0x67}, {0x74, 0x69, 0x72, 0x73, 0x64, 0x61, 0x67}, {0x6f, 0x6e, 0x73, 0x64, 0x61, 0x67}, {0x74, 0x6f, 0x72, 0x73, 0x64, 0x61, 0x67}, {0x66, 0x72, 0x65, 0x64, 0x61, 0x67}, {0x6c, 0xc3, 0xb8, 0x72, 0x64, 0x61, 0x67}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0x66, 0x2e, 0x4b, 0x72, 0x2e}, {0x65, 0x2e, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x66, 0x4b, 0x72}, {0x65, 0x4b, 0x72}},
		erasWide:               [][]uint8{{0x66, 0x2e, 0x4b, 0x72, 0x2e}, {0x65, 0x2e, 0x4b, 0x72, 0x2e}},
		timezones:              map[string][]uint8{"PDT": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "MESZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x2d, 0x74, 0x69, 0x64}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "AWST": {0x56, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "MST": {0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "CST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "CDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "HADT": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "ACWST": {0x56, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x67, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "WART": {0x56, 0x65, 0x73, 0x74, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "WITA": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "HAT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "MEZ": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "AEST": {0xc3, 0x98, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "CLST": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "EST": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x2d, 0x74, 0x69, 0x64}, "EAT": {0xc3, 0x98, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "JST": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "COT": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "CAT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "WAST": {0x56, 0x65, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "SAST": {0x53, 0x79, 0x64, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x2d, 0x74, 0x69, 0x64}, "NZST": {0x4e, 0x65, 0x77, 0x7a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "WARST": {0x56, 0x65, 0x73, 0x74, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "WIT": {0xc3, 0x98, 0x73, 0x74, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "WEZ": {0x56, 0x65, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "AWDT": {0x56, 0x65, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "GMT": {0x47, 0x4d, 0x54}, "NZDT": {0x4e, 0x65, 0x77, 0x7a, 0x65, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "∅∅∅": {0x41, 0x63, 0x72, 0x65, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "WAT": {0x56, 0x65, 0x73, 0x74, 0x61, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "MDT": {0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "COST": {0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "HNT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x2d, 0x74, 0x69, 0x64}, "ACWDT": {0x56, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x67, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "OEZ": {0xc3, 0x98, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "WESZ": {0x56, 0x65, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "ECT": {0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65, 0x61, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "HAST": {0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "WIB": {0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "OESZ": {0xc3, 0x98, 0x73, 0x74, 0x65, 0x75, 0x72, 0x6f, 0x70, 0xc3, 0xa6, 0x69, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "ACST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "AEDT": {0xc3, 0x98, 0x73, 0x74, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "CLT": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x73, 0x6b, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}, "EDT": {0x45, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "ACDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x61, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x73, 0x6b, 0x20, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x2d, 0x73, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x74, 0x69, 0x64}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x73, 0x6b, 0x20, 0x74, 0x69, 0x64}, "PST": {0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2d, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x74, 0x69, 0x64}},
	}
}

// Locale returns the current translators string locale
func (da *da_GL) Locale() string {
	return da.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'da_GL'
func (da *da_GL) PluralsCardinal() []locales.PluralRule {
	return da.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'da_GL'
func (da *da_GL) PluralsOrdinal() []locales.PluralRule {
	return da.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'da_GL'
func (da *da_GL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'da_GL'
func (da *da_GL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'da_GL'
func (da *da_GL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := da.CardinalPluralRule(num1, v1)
	end := da.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'da_GL' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(da.decimal) + len(da.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'da_GL' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (da *da_GL) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(da.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, da.percentSuffix...)

	b = append(b, da.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := da.currencies[currency]
	l := len(s) + len(da.decimal) + len(da.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, da.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, da.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'da_GL'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := da.currencies[currency]
	l := len(s) + len(da.decimal) + len(da.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, da.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, da.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, da.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, da.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, da.daysWide[t.Day()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x6e, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, da.periodsAbbreviated[0]...)
	} else {
		b = append(b, da.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, da.periodsAbbreviated[0]...)
	} else {
		b = append(b, da.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, da.periodsAbbreviated[0]...)
	} else {
		b = append(b, da.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, da.periodsAbbreviated[0]...)
	} else {
		b = append(b, da.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := da.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateTimeShort(t time.Time) []byte {

	b := da.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, da.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateTimeMedium(t time.Time) []byte {

	b := da.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, da.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateTimeLong(t time.Time) []byte {

	b := da.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, da.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'da_GL'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (da *da_GL) FmtDateTimeFull(t time.Time) []byte {

	b := da.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, da.FmtTimeFull(t)...)

	return b
}
