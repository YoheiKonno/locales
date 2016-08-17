package bn_IN

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type bn_IN struct {
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

// New returns a new instance of translator for the 'bn_IN' locale
func New() locales.Translator {
	return &bn_IN{
		locale:                 "bn_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 5, 6},
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xa6, 0x9c, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xa8, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xab, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9a}, {0xe0, 0xa6, 0x8f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0x87}, {0xe0, 0xa6, 0x86, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xad, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xa1, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe0, 0xa6, 0x9c, 0xe0, 0xa6, 0xbe}, {0xe0, 0xa6, 0xab, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xbe}, {0xe0, 0xa6, 0x8f}, {0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81}, {0xe0, 0xa6, 0x86}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x85}, {0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0xa1, 0xe0, 0xa6, 0xbf}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xa6, 0x9c, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xa8, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xab, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9a}, {0xe0, 0xa6, 0x8f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0x87}, {0xe0, 0xa6, 0x86, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xad, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xa1, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xb0}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbf}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xae}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0x99, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa7}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb9, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xaa, 0xe0, 0xa6, 0xa4, 0xe0, 0xa6, 0xbf}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xbf}},
		daysNarrow:             [][]uint8{{0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b}, {0xe0, 0xa6, 0xae}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81}, {0xe0, 0xa6, 0xb6}},
		daysShort:              [][]uint8{{0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x83}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0x83}},
		daysWide:               [][]uint8{{0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0x99, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa7, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb9, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xaa, 0xe0, 0xa6, 0xa4, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x82, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0xaa, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xa6, 0x96, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x82, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac}, {0xe0, 0xa6, 0x96, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb7, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa6}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0xe0, 0xa6, 0x96, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x82, 0xe0, 0xa6, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac}, {0xe0, 0xa6, 0x96, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb7, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa6}}}
}

// Locale returns the current translators string locale
func (bn *bn_IN) Locale() string {
	return bn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bn_IN'
func (bn *bn_IN) PluralsCardinal() []locales.PluralRule {
	return bn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bn_IN'
func (bn *bn_IN) PluralsOrdinal() []locales.PluralRule {
	return bn.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bn_IN'
func (bn *bn_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bn_IN'
func (bn *bn_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 5 || n == 7 || n == 8 || n == 9 || n == 10 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bn_IN'
func (bn *bn_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := bn.CardinalPluralRule(num1, v1)
	end := bn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bn_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bn *bn_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bn.decimal) + len(bn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(bn.decimal) - 1; j >= 0; j-- {
				b = append(b, bn.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(bn.group) - 1; j >= 0; j-- {
					b = append(b, bn.group[j])
				}

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
		for j := len(bn.minus) - 1; j >= 0; j-- {
			b = append(b, bn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bn_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bn *bn_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bn.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(bn.decimal) - 1; j >= 0; j-- {
				b = append(b, bn.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(bn.minus) - 1; j >= 0; j-- {
			b = append(b, bn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bn.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bn_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bn *bn_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bn.currencies[currency]
	l := len(s) + len(bn.decimal) + len(bn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(bn.decimal) - 1; j >= 0; j-- {
				b = append(b, bn.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(bn.group) - 1; j >= 0; j-- {
					b = append(b, bn.group[j])
				}

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
		for j := len(bn.minus) - 1; j >= 0; j-- {
			b = append(b, bn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bn_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bn *bn_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bn.currencies[currency]
	l := len(s) + len(bn.decimal) + len(bn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(bn.decimal) - 1; j >= 0; j-- {
				b = append(b, bn.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(bn.group) - 1; j >= 0; j-- {
					b = append(b, bn.group[j])
				}

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

		for j := len(bn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, bn.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bn.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return b
}
