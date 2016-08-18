package yav

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type yav struct {
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

// New returns a new instance of translator for the 'yav' locale
func New() locales.Translator {
	return &yav{
		locale:                 "yav",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6f, 0x2e, 0x31}, {0x6f, 0x2e, 0x32}, {0x6f, 0x2e, 0x33}, {0x6f, 0x2e, 0x34}, {0x6f, 0x2e, 0x35}, {0x6f, 0x2e, 0x36}, {0x6f, 0x2e, 0x37}, {0x6f, 0x2e, 0x38}, {0x6f, 0x2e, 0x39}, {0x6f, 0x2e, 0x31, 0x30}, {0x6f, 0x2e, 0x31, 0x31}, {0x6f, 0x2e, 0x31, 0x32}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x70, 0x69, 0x6b, 0xc3, 0xad, 0x74, 0xc3, 0xad, 0x6b, 0xc3, 0xad, 0x74, 0x69, 0x65, 0x2c, 0x20, 0x6f, 0xc3, 0xb3, 0x6c, 0xc3, 0xad, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0x75, 0x74, 0xc3, 0xba, 0x61, 0x6e}, {0x73, 0x69, 0xc9, 0x9b, 0x79, 0xc9, 0x9b, 0xcc, 0x81, 0x2c, 0x20, 0x6f, 0xc3, 0xb3, 0x6c, 0x69, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0xc3, 0xa1, 0x6e, 0x64, 0xc3, 0xad, 0xc9, 0x9b}, {0xc9, 0x94, 0x6e, 0x73, 0xc3, 0xba, 0x6d, 0x62, 0xc9, 0x94, 0x6c, 0x2c, 0x20, 0x6f, 0xc3, 0xb3, 0x6c, 0x69, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0xc3, 0xa1, 0x74, 0xc3, 0xa1, 0x74, 0xc3, 0xba, 0xc9, 0x9b}, {0x6d, 0x65, 0x73, 0x69, 0xc5, 0x8b, 0x2c, 0x20, 0x6f, 0xc3, 0xb3, 0x6c, 0x69, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0xc3, 0xa9, 0x6e, 0x69, 0x65}, {0x65, 0x6e, 0x73, 0x69, 0x6c, 0x2c, 0x20, 0x6f, 0xc3, 0xb3, 0x6c, 0x69, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0xc3, 0xa1, 0x74, 0xc3, 0xa1, 0x6e, 0x75, 0xc9, 0x9b}, {0xc9, 0x94, 0x73, 0xc9, 0x94, 0x6e}, {0x65, 0x66, 0x75, 0x74, 0x65}, {0x70, 0x69, 0x73, 0x75, 0x79, 0xc3, 0xba}, {0x69, 0x6d, 0xc9, 0x9b, 0xc5, 0x8b, 0x20, 0x69, 0x20, 0x70, 0x75, 0xc9, 0x94, 0x73}, {0x69, 0x6d, 0xc9, 0x9b, 0xc5, 0x8b, 0x20, 0x69, 0x20, 0x70, 0x75, 0x74, 0xc3, 0xba, 0x6b, 0x2c, 0x6f, 0xc3, 0xb3, 0x6c, 0x69, 0x20, 0xc3, 0xba, 0x20, 0x6b, 0xc3, 0xa1, 0x74, 0xc3, 0xad, 0xc9, 0x9b}, {0x6d, 0x61, 0x6b, 0x61, 0x6e, 0x64, 0x69, 0x6b, 0xc9, 0x9b}, {0x70, 0x69, 0x6c, 0xc9, 0x94, 0x6e, 0x64, 0xc9, 0x94, 0xcc, 0x81}},
		daysAbbreviated:        [][]uint8{{0x73, 0x64}, {0x6d, 0x64}, {0x6d, 0x77}, {0x65, 0x74}, {0x6b, 0x6c}, {0x66, 0x6c}, {0x73, 0x73}},
		daysNarrow:             [][]uint8{{0x73}, {0x6d}, {0x6d}, {0x65}, {0x6b}, {0x66}, {0x73}},
		daysWide:               [][]uint8{{0x73, 0xc9, 0x94, 0xcc, 0x81, 0x6e, 0x64, 0x69, 0xc9, 0x9b}, {0x6d, 0xc3, 0xb3, 0x6e, 0x64, 0x69, 0x65}, {0x6d, 0x75, 0xc3, 0xa1, 0x6e, 0x79, 0xc3, 0xa1, 0xc5, 0x8b, 0x6d, 0xc3, 0xb3, 0x6e, 0x64, 0x69, 0x65}, {0x6d, 0x65, 0x74, 0xc3, 0xba, 0x6b, 0x70, 0xc3, 0xad, 0xc3, 0xa1, 0x70, 0xc9, 0x9b}, {0x6b, 0xc3, 0xba, 0x70, 0xc3, 0xa9, 0x6c, 0x69, 0x6d, 0x65, 0x74, 0xc3, 0xba, 0x6b, 0x70, 0x69, 0x61, 0x70, 0xc9, 0x9b}, {0x66, 0x65, 0x6c, 0xc3, 0xa9, 0x74, 0x65}, {0x73, 0xc3, 0xa9, 0x73, 0x65, 0x6c, 0xc3, 0xa9}},
		periodsAbbreviated:     [][]uint8{{0x6b, 0x69, 0xc9, 0x9b, 0x6d, 0xc9, 0x9b, 0xcc, 0x81, 0xc9, 0x9b, 0x6d}, {0x6b, 0x69, 0x73, 0xc9, 0x9b, 0xcc, 0x81, 0x6e, 0x64, 0xc9, 0x9b}},
		periodsWide:            [][]uint8{{0x6b, 0x69, 0xc9, 0x9b, 0x6d, 0xc9, 0x9b, 0xcc, 0x81, 0xc9, 0x9b, 0x6d}, {0x6b, 0x69, 0x73, 0xc9, 0x9b, 0xcc, 0x81, 0x6e, 0x64, 0xc9, 0x9b}},
		erasAbbreviated:        [][]uint8{{0x6b, 0x2e, 0x59, 0x2e}, {0x2b, 0x4a, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x6b, 0x61, 0x74, 0x69, 0x6b, 0x75, 0x70, 0xc3, 0xad, 0x65, 0x6e, 0x20, 0x59, 0xc3, 0xa9, 0x73, 0x75, 0x73, 0x65}, {0xc3, 0xa9, 0x6b, 0xc3, 0xa9, 0x6c, 0xc3, 0xa9, 0x6d, 0x6b, 0xc3, 0xba, 0x6e, 0x75, 0x70, 0xc3, 0xad, 0xc3, 0xa9, 0x6e, 0x20, 0x6e}},
		timezones:              map[string][]uint8{"UYST": {0x55, 0x59, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "WIT": {0x57, 0x49, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CST": {0x43, 0x53, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "AST": {0x41, 0x53, 0x54}, "MST": {0x4d, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "CLT": {0x43, 0x4c, 0x54}, "EST": {0x45, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "ADT": {0x41, 0x44, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "BOT": {0x42, 0x4f, 0x54}, "IST": {0x49, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "MYT": {0x4d, 0x59, 0x54}, "PST": {0x50, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "CDT": {0x43, 0x44, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "BT": {0x42, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "JDT": {0x4a, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "CAT": {0x43, 0x41, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WAT": {0x57, 0x41, 0x54}, "HAT": {0x48, 0x41, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "GFT": {0x47, 0x46, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "SRT": {0x53, 0x52, 0x54}},
	}
}

// Locale returns the current translators string locale
func (yav *yav) Locale() string {
	return yav.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yav'
func (yav *yav) PluralsCardinal() []locales.PluralRule {
	return yav.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yav'
func (yav *yav) PluralsOrdinal() []locales.PluralRule {
	return yav.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yav'
func (yav *yav) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yav'
func (yav *yav) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yav'
func (yav *yav) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yav' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(yav.decimal) + len(yav.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yav.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(yav.group) - 1; j >= 0; j-- {
					b = append(b, yav.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(yav.minus) - 1; j >= 0; j-- {
			b = append(b, yav.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yav' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yav *yav) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(yav.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yav.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(yav.minus) - 1; j >= 0; j-- {
			b = append(b, yav.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yav.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yav.currencies[currency]
	l := len(s) + len(yav.decimal) + len(yav.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yav.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(yav.group) - 1; j >= 0; j-- {
					b = append(b, yav.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(yav.minus) - 1; j >= 0; j-- {
			b = append(b, yav.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yav.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, yav.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yav'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yav.currencies[currency]
	l := len(s) + len(yav.decimal) + len(yav.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yav.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(yav.group) - 1; j >= 0; j-- {
					b = append(b, yav.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(yav.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, yav.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yav.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, yav.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, yav.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yav.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yav.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, yav.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, yav.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yav.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}

// FmtDateTimeShort returns the short date & time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateTimeShort(t time.Time) []byte {

	b := yav.FmtDateShort(t)
	b = append(b, ' ')
	b = append(b, yav.FmtTimeShort(t)...)

	return b
}

// FmtDateTimeMedium returns the medium date & time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateTimeMedium(t time.Time) []byte {

	b := yav.FmtDateMedium(t)
	b = append(b, ' ')
	b = append(b, yav.FmtTimeMedium(t)...)

	return b
}

// FmtDateTimeLong returns the long date & time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateTimeLong(t time.Time) []byte {

	b := yav.FmtDateLong(t)
	b = append(b, ' ')
	b = append(b, yav.FmtTimeLong(t)...)

	return b
}

// FmtDateTimeFull returns the full date & time representation of 't' for 'yav'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yav *yav) FmtDateTimeFull(t time.Time) []byte {

	b := yav.FmtDateFull(t)
	b = append(b, ' ')
	b = append(b, yav.FmtTimeFull(t)...)

	return b
}
