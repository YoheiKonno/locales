package ti_ET

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ti_ET struct {
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

// New returns a new instance of translator for the 'ti_ET' locale
func New() locales.Translator {
	return &ti_ET{
		locale:             "ti_ET",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		decimal:            []byte{},
		group:              []byte{},
		minus:              []byte{},
		percent:            []byte{},
		perMille:           []byte{},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0xe1, 0x8c, 0x83, 0xe1, 0x8a, 0x95, 0xe1, 0x8b, 0xa9}, {0xe1, 0x8d, 0x8c, 0xe1, 0x89, 0xa5, 0xe1, 0x88, 0xa9}, {0xe1, 0x88, 0x9b, 0xe1, 0x88, 0xad, 0xe1, 0x89, 0xbd}, {0xe1, 0x8a, 0xa4, 0xe1, 0x8d, 0x95, 0xe1, 0x88, 0xa8}, {0xe1, 0x88, 0x9c, 0xe1, 0x8b, 0xad}, {0xe1, 0x8c, 0x81, 0xe1, 0x8a, 0x95}, {0xe1, 0x8c, 0x81, 0xe1, 0x88, 0x8b, 0xe1, 0x8b, 0xad}, {0xe1, 0x8a, 0xa6, 0xe1, 0x8c, 0x88, 0xe1, 0x88, 0xb5}, {0xe1, 0x88, 0xb4, 0xe1, 0x8d, 0x95, 0xe1, 0x89, 0xb4}, {0xe1, 0x8a, 0xa6, 0xe1, 0x8a, 0xad, 0xe1, 0x89, 0xb0}, {0xe1, 0x8a, 0x96, 0xe1, 0x89, 0xac, 0xe1, 0x88, 0x9d}, {0xe1, 0x8b, 0xb2, 0xe1, 0x88, 0xb4, 0xe1, 0x88, 0x9d}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0xe1, 0x8c, 0x83}, {0xe1, 0x8d, 0x8c}, {0xe1, 0x88, 0x9b}, {0xe1, 0x8a, 0xa4}, {0xe1, 0x88, 0x9c}, {0xe1, 0x8c, 0x81}, {0xe1, 0x8c, 0x81}, {0xe1, 0x8a, 0xa6}, {0xe1, 0x88, 0xb4}, {0xe1, 0x8a, 0xa6}, {0xe1, 0x8a, 0x96}, {0xe1, 0x8b, 0xb2}},
		monthsWide:         [][]uint8{[]uint8(nil), {0xe1, 0x8c, 0x83, 0xe1, 0x8a, 0x95, 0xe1, 0x8b, 0xa9, 0xe1, 0x8b, 0x88, 0xe1, 0x88, 0xaa}, {0xe1, 0x8d, 0x8c, 0xe1, 0x89, 0xa5, 0xe1, 0x88, 0xa9, 0xe1, 0x8b, 0x88, 0xe1, 0x88, 0xaa}, {0xe1, 0x88, 0x9b, 0xe1, 0x88, 0xad, 0xe1, 0x89, 0xbd}, {0xe1, 0x8a, 0xa4, 0xe1, 0x8d, 0x95, 0xe1, 0x88, 0xa8, 0xe1, 0x88, 0x8d}, {0xe1, 0x88, 0x9c, 0xe1, 0x8b, 0xad}, {0xe1, 0x8c, 0x81, 0xe1, 0x8a, 0x95}, {0xe1, 0x8c, 0x81, 0xe1, 0x88, 0x8b, 0xe1, 0x8b, 0xad}, {0xe1, 0x8a, 0xa6, 0xe1, 0x8c, 0x88, 0xe1, 0x88, 0xb5, 0xe1, 0x89, 0xb5}, {0xe1, 0x88, 0xb4, 0xe1, 0x8d, 0x95, 0xe1, 0x89, 0xb4, 0xe1, 0x88, 0x9d, 0xe1, 0x89, 0xa0, 0xe1, 0x88, 0xad}, {0xe1, 0x8a, 0xa6, 0xe1, 0x8a, 0xad, 0xe1, 0x89, 0xb0, 0xe1, 0x8b, 0x8d, 0xe1, 0x89, 0xa0, 0xe1, 0x88, 0xad}, {0xe1, 0x8a, 0x96, 0xe1, 0x89, 0xac, 0xe1, 0x88, 0x9d, 0xe1, 0x89, 0xa0, 0xe1, 0x88, 0xad}, {0xe1, 0x8b, 0xb2, 0xe1, 0x88, 0xb4, 0xe1, 0x88, 0x9d, 0xe1, 0x89, 0xa0, 0xe1, 0x88, 0xad}},
		daysNarrow:         [][]uint8{{0xe1, 0x88, 0xb0}, {0xe1, 0x88, 0xb0}, {0xe1, 0x88, 0xa0}, {0xe1, 0x88, 0xa8}, {0xe1, 0x8a, 0x83}, {0xe1, 0x8b, 0x93}, {0xe1, 0x89, 0x80}},
		daysWide:           [][]uint8{{0xe1, 0x88, 0xb0, 0xe1, 0x8a, 0x95, 0xe1, 0x89, 0xa0, 0xe1, 0x89, 0xb5}, {0xe1, 0x88, 0xb0, 0xe1, 0x8a, 0x91, 0xe1, 0x8b, 0xad}, {0xe1, 0x88, 0xa0, 0xe1, 0x88, 0x89, 0xe1, 0x88, 0xb5}, {0xe1, 0x88, 0xa8, 0xe1, 0x89, 0xa1, 0xe1, 0x8b, 0x95}, {0xe1, 0x8a, 0x83, 0xe1, 0x88, 0x99, 0xe1, 0x88, 0xb5}, {0xe1, 0x8b, 0x93, 0xe1, 0x88, 0xad, 0xe1, 0x89, 0xa2}, {0xe1, 0x89, 0x80, 0xe1, 0x8b, 0xb3, 0xe1, 0x88, 0x9d}},
		periodsAbbreviated: [][]uint8{{0xe1, 0x8a, 0x95, 0xe1, 0x8c, 0x89, 0xe1, 0x88, 0x86, 0x20, 0xe1, 0x88, 0xb0, 0xe1, 0x8b, 0x93, 0xe1, 0x89, 0xb0}, {0xe1, 0x8b, 0xb5, 0xe1, 0x88, 0x95, 0xe1, 0x88, 0xad, 0x20, 0xe1, 0x88, 0xb0, 0xe1, 0x8b, 0x93, 0xe1, 0x89, 0xb5}},
		periodsWide:        [][]uint8{{0xe1, 0x8a, 0x95, 0xe1, 0x8c, 0x89, 0xe1, 0x88, 0x86, 0x20, 0xe1, 0x88, 0xb0, 0xe1, 0x8b, 0x93, 0xe1, 0x89, 0xb0}, {0xe1, 0x8b, 0xb5, 0xe1, 0x88, 0x95, 0xe1, 0x88, 0xad, 0x20, 0xe1, 0x88, 0xb0, 0xe1, 0x8b, 0x93, 0xe1, 0x89, 0xb5}},
		erasAbbreviated:    [][]uint8{{0xe1, 0x8b, 0x93, 0x2f, 0xe1, 0x8b, 0x93}, {0xe1, 0x8b, 0x93, 0x2f, 0xe1, 0x88, 0x9d}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{[]uint8(nil), []uint8(nil)}}
}

// Locale returns the current translators string locale
func (ti *ti_ET) Locale() string {
	return ti.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ti_ET'
func (ti *ti_ET) PluralsCardinal() []locales.PluralRule {
	return ti.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ti_ET'
func (ti *ti_ET) PluralsOrdinal() []locales.PluralRule {
	return ti.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ti_ET'
func (ti *ti_ET) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ti_ET'
func (ti *ti_ET) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ti_ET'
func (ti *ti_ET) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ti_ET' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ti *ti_ET) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ti_ET' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ti *ti_ET) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ti_ET'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ti *ti_ET) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ti.currencies[currency]
	l := len(s) + len(ti.decimal) + len(ti.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ti.decimal) - 1; j >= 0; j-- {
				b = append(b, ti.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ti.group) - 1; j >= 0; j-- {
					b = append(b, ti.group[j])
				}

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
		for j := len(ti.minus) - 1; j >= 0; j-- {
			b = append(b, ti.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ti.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ti_ET'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ti *ti_ET) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ti.currencies[currency]
	l := len(s) + len(ti.decimal) + len(ti.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ti.decimal) - 1; j >= 0; j-- {
				b = append(b, ti.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ti.group) - 1; j >= 0; j-- {
					b = append(b, ti.group[j])
				}

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

		for j := len(ti.minus) - 1; j >= 0; j-- {
			b = append(b, ti.minus[j])
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
			b = append(b, ti.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}
