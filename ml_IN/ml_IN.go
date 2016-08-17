package ml_IN

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ml_IN struct {
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

// New returns a new instance of translator for the 'ml_IN' locale
func New() locales.Translator {
	return &ml_IN{
		locale:                 "ml_IN",
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
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xb4, 0x9c, 0xe0, 0xb4, 0xa8, 0xe0, 0xb5, 0x81}, {0xe0, 0xb4, 0xab, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb5, 0x81}, {0xe0, 0xb4, 0xae, 0xe0, 0xb4, 0xbe, 0xe0, 0xb5, 0xbc}, {0xe0, 0xb4, 0x8f, 0xe0, 0xb4, 0xaa, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x87, 0xe0, 0xb4, 0xaf, 0xe0, 0xb5, 0x8d}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82, 0xe0, 0xb5, 0xba}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82, 0xe0, 0xb4, 0xb2, 0xe0, 0xb5, 0x88}, {0xe0, 0xb4, 0x93, 0xe0, 0xb4, 0x97}, {0xe0, 0xb4, 0xb8, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xaa, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb4, 0x82}, {0xe0, 0xb4, 0x92, 0xe0, 0xb4, 0x95, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0x9f, 0xe0, 0xb5, 0x8b}, {0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0x82}, {0xe0, 0xb4, 0xa1, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xb8, 0xe0, 0xb4, 0x82}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe0, 0xb4, 0x9c}, {0xe0, 0xb4, 0xab}, {0xe0, 0xb4, 0xae, 0xe0, 0xb4, 0xbe}, {0xe0, 0xb4, 0x8f}, {0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x86}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82}, {0xe0, 0xb4, 0x93}, {0xe0, 0xb4, 0xb8}, {0xe0, 0xb4, 0x92}, {0xe0, 0xb4, 0xa8}, {0xe0, 0xb4, 0xa1, 0xe0, 0xb4, 0xbf}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xb4, 0x9c, 0xe0, 0xb4, 0xa8, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0xab, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0xae, 0xe0, 0xb4, 0xbe, 0xe0, 0xb5, 0xbc, 0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8d}, {0xe0, 0xb4, 0x8f, 0xe0, 0xb4, 0xaa, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf, 0xe0, 0xb5, 0xbd}, {0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x87, 0xe0, 0xb4, 0xaf, 0xe0, 0xb5, 0x8d}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82, 0xe0, 0xb5, 0xba}, {0xe0, 0xb4, 0x9c, 0xe0, 0xb5, 0x82, 0xe0, 0xb4, 0xb2, 0xe0, 0xb5, 0x88}, {0xe0, 0xb4, 0x93, 0xe0, 0xb4, 0x97, 0xe0, 0xb4, 0xb8, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb5, 0x8d}, {0xe0, 0xb4, 0xb8, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xaa, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb1, 0xe0, 0xb4, 0x82, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0xbc}, {0xe0, 0xb4, 0x92, 0xe0, 0xb4, 0x95, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9f, 0xe0, 0xb5, 0x8b, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0xbc}, {0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0x82, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0xbc}, {0xe0, 0xb4, 0xa1, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xb8, 0xe0, 0xb4, 0x82, 0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0xbc}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xb4, 0x9e, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xaf, 0xe0, 0xb5, 0xbc}, {0xe0, 0xb4, 0xa4, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0x99, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0x95, 0xe0, 0xb5, 0xbe}, {0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8a, 0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb5}, {0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xa7, 0xe0, 0xb5, 0xbb}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb4, 0x82}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xb3, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb3, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0xb6, 0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xbf}},
		daysNarrow:             [][]uint8{{0xe0, 0xb4, 0x9e}, {0xe0, 0xb4, 0xa4, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8a}, {0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x81}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x86}, {0xe0, 0xb4, 0xb6}},
		daysShort:              [][]uint8{{0xe0, 0xb4, 0x9e, 0xe0, 0xb4, 0xbe}, {0xe0, 0xb4, 0xa4, 0xe0, 0xb4, 0xbf}, {0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8a}, {0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x81}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x86}, {0xe0, 0xb4, 0xb6}},
		daysWide:               [][]uint8{{0xe0, 0xb4, 0x9e, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xb1, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0xa4, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0x99, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0x95, 0xe0, 0xb4, 0xb3, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0x9a, 0xe0, 0xb5, 0x8a, 0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0xac, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xa7, 0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0xb5, 0xe0, 0xb5, 0x86, 0xe0, 0xb4, 0xb3, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb3, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}, {0xe0, 0xb4, 0xb6, 0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xaf, 0xe0, 0xb4, 0xbe, 0xe0, 0xb4, 0xb4, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0x9a}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xb4, 0x95, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf, 0x2e, 0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x81, 0x2e}, {0xe0, 0xb4, 0x8e, 0xe0, 0xb4, 0xa1, 0xe0, 0xb4, 0xbf}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0xe0, 0xb4, 0x95, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xb0, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xb8, 0xe0, 0xb5, 0x8d, 0xe2, 0x80, 0x8c, 0xe0, 0xb4, 0xa4, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xb5, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xa8, 0xe0, 0xb5, 0x8d, 0x20, 0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x81, 0xe0, 0xb4, 0xae, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xaa, 0xe0, 0xb5, 0x8d}, {0xe0, 0xb4, 0x86, 0xe0, 0xb4, 0xa8, 0xe0, 0xb5, 0x8d, 0xe0, 0xb4, 0xa8, 0xe0, 0xb5, 0x8b, 0x20, 0xe0, 0xb4, 0xa1, 0xe0, 0xb5, 0x8a, 0xe0, 0xb4, 0xae, 0xe0, 0xb4, 0xbf, 0xe0, 0xb4, 0xa8, 0xe0, 0xb4, 0xbf}}}
}

// Locale returns the current translators string locale
func (ml *ml_IN) Locale() string {
	return ml.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ml_IN'
func (ml *ml_IN) PluralsCardinal() []locales.PluralRule {
	return ml.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ml_IN'
func (ml *ml_IN) PluralsOrdinal() []locales.PluralRule {
	return ml.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ml_IN'
func (ml *ml_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ml_IN'
func (ml *ml_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ml_IN'
func (ml *ml_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ml.CardinalPluralRule(num1, v1)
	end := ml.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ml_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ml *ml_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ml.decimal) + len(ml.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, ml.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ml_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ml *ml_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ml.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ml.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ml_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ml *ml_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ml.currencies[currency]
	l := len(s) + len(ml.decimal) + len(ml.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ml.group[0])
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
		b = append(b, ml.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ml.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ml_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ml *ml_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ml.currencies[currency]
	l := len(s) + len(ml.decimal) + len(ml.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ml.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ml.group[0])
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

		for j := len(ml.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ml.currencyNegativePrefix[j])
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
			b = append(b, ml.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ml.currencyNegativeSuffix...)
	}

	return b
}
