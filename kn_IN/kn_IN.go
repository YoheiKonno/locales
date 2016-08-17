package kn_IN

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type kn_IN struct {
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
}

// New returns a new instance of translator for the 'kn_IN' locale
func New() locales.Translator {
	return &kn_IN{
		locale:                 "kn_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xb2, 0x9c, 0xe0, 0xb2, 0xa8}, {0xe0, 0xb2, 0xab, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9a, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0x8f, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf}, {0xe0, 0xb2, 0xae, 0xe0, 0xb3, 0x87}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x82, 0xe0, 0xb2, 0xa8, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb2, 0xe0, 0xb3, 0x88}, {0xe0, 0xb2, 0x86, 0xe0, 0xb2, 0x97}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9f, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82}, {0xe0, 0xb2, 0x85, 0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9f, 0xe0, 0xb3, 0x8b}, {0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xb5, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82}, {0xe0, 0xb2, 0xa1, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe0, 0xb2, 0x9c}, {0xe0, 0xb2, 0xab, 0xe0, 0xb3, 0x86}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0xbe}, {0xe0, 0xb2, 0x8f}, {0xe0, 0xb2, 0xae, 0xe0, 0xb3, 0x87}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x82}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0x86}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x86}, {0xe0, 0xb2, 0x85}, {0xe0, 0xb2, 0xa8}, {0xe0, 0xb2, 0xa1, 0xe0, 0xb2, 0xbf}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xb2, 0x9c, 0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf}, {0xe0, 0xb2, 0xab, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9a, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0x8f, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb2, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0xae, 0xe0, 0xb3, 0x87}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x82, 0xe0, 0xb2, 0xa8, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0x9c, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb2, 0xe0, 0xb3, 0x88}, {0xe0, 0xb2, 0x86, 0xe0, 0xb2, 0x97, 0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9f, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9f, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0xac, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0x85, 0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0x9f, 0xe0, 0xb3, 0x8b, 0xe0, 0xb2, 0xac, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xb5, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0xac, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d}, {0xe0, 0xb2, 0xa1, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x86, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0xac, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xb2, 0xad, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xa8, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8b, 0xe0, 0xb2, 0xae}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0x97, 0xe0, 0xb2, 0xb3}, {0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xa7}, {0xe0, 0xb2, 0x97, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xbf}},
		daysNarrow:             [][]uint8{{0xe0, 0xb2, 0xad, 0xe0, 0xb2, 0xbe}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8b}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0x82}, {0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0x97, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb6}},
		daysShort:              [][]uint8{{0xe0, 0xb2, 0xad, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xa8, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8b, 0xe0, 0xb2, 0xae}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0x97, 0xe0, 0xb2, 0xb3}, {0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xa7}, {0xe0, 0xb2, 0x97, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x81}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xbf}},
		daysWide:               [][]uint8{{0xe0, 0xb2, 0xad, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xa8, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8b, 0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xae, 0xe0, 0xb2, 0x82, 0xe0, 0xb2, 0x97, 0xe0, 0xb2, 0xb3, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xac, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xa7, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0x97, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb3, 0x81, 0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}, {0xe0, 0xb2, 0xb6, 0xe0, 0xb2, 0xa8, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x82}, {0xe0, 0xb2, 0x85}},
		periodsWide:            [][]uint8{{0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x82, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb5, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb9, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xa8}, {0xe0, 0xb2, 0x85, 0xe0, 0xb2, 0xaa, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbe, 0xe0, 0xb2, 0xb9, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xa8}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf, 0x2e, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x82}, {0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf, 0x2e, 0xe0, 0xb2, 0xb6}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xa4, 0x20, 0xe0, 0xb2, 0xaa, 0xe0, 0xb3, 0x82, 0xe0, 0xb2, 0xb0, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb5}, {0xe0, 0xb2, 0x95, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xb0, 0xe0, 0xb2, 0xbf, 0xe0, 0xb2, 0xb8, 0xe0, 0xb3, 0x8d, 0xe0, 0xb2, 0xa4, 0x20, 0xe0, 0xb2, 0xb6, 0xe0, 0xb2, 0x95}}}
}

// Locale returns the current translators string locale
func (kn *kn_IN) Locale() string {
	return kn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kn_IN'
func (kn *kn_IN) PluralsCardinal() []locales.PluralRule {
	return kn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kn_IN'
func (kn *kn_IN) PluralsOrdinal() []locales.PluralRule {
	return kn.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kn_IN'
func (kn *kn_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kn_IN'
func (kn *kn_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kn_IN'
func (kn *kn_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := kn.CardinalPluralRule(num1, v1)
	end := kn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kn_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kn *kn_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kn.decimal) + len(kn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kn_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kn *kn_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(kn.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kn.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kn_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kn *kn_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kn.currencies[currency]
	l := len(s) + len(kn.decimal) + len(kn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
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
		b = append(b, kn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kn_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (kn *kn_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kn.currencies[currency]
	l := len(s) + len(kn.decimal) + len(kn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kn.group[0])
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

		for j := len(kn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kn.currencyNegativePrefix[j])
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
			b = append(b, kn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, kn.currencyNegativeSuffix...)
	}

	return b
}
