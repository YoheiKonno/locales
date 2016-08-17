package ru

import (
	"math"
	"strconv"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ru struct {
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

// New returns a new instance of translator for the 'ru' locale
func New() locales.Translator {
	return &ru{
		locale:                 "ru",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0xe2, 0x82, 0xbd}, {0xd1, 0x80, 0x2e}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0xd0, 0xa2, 0xd0, 0x9c, 0xd0, 0xa2}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0xe2, 0x82, 0xb4}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xd1, 0x8f, 0xd0, 0xbd, 0xd0, 0xb2, 0x2e}, {0xd1, 0x84, 0xd0, 0xb5, 0xd0, 0xb2, 0xd1, 0x80, 0x2e}, {0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x80, 0x2e}, {0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80, 0x2e}, {0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x8f}, {0xd0, 0xb8, 0xd1, 0x8e, 0xd0, 0xbd, 0x2e}, {0xd0, 0xb8, 0xd1, 0x8e, 0xd0, 0xbb, 0x2e}, {0xd0, 0xb0, 0xd0, 0xb2, 0xd0, 0xb3, 0x2e}, {0xd1, 0x81, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0x2e}, {0xd0, 0xbe, 0xd0, 0xba, 0xd1, 0x82, 0x2e}, {0xd0, 0xbd, 0xd0, 0xbe, 0xd1, 0x8f, 0xd0, 0xb1, 0x2e}, {0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xba, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xd0, 0xaf}, {0xd0, 0xa4}, {0xd0, 0x9c}, {0xd0, 0x90}, {0xd0, 0x9c}, {0xd0, 0x98}, {0xd0, 0x98}, {0xd0, 0x90}, {0xd0, 0xa1}, {0xd0, 0x9e}, {0xd0, 0x9d}, {0xd0, 0x94}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xd1, 0x8f, 0xd0, 0xbd, 0xd0, 0xb2, 0xd0, 0xb0, 0xd1, 0x80, 0xd1, 0x8f}, {0xd1, 0x84, 0xd0, 0xb5, 0xd0, 0xb2, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8f}, {0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x80, 0xd1, 0x82, 0xd0, 0xb0}, {0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xb5, 0xd0, 0xbb, 0xd1, 0x8f}, {0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x8f}, {0xd0, 0xb8, 0xd1, 0x8e, 0xd0, 0xbd, 0xd1, 0x8f}, {0xd0, 0xb8, 0xd1, 0x8e, 0xd0, 0xbb, 0xd1, 0x8f}, {0xd0, 0xb0, 0xd0, 0xb2, 0xd0, 0xb3, 0xd1, 0x83, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb0}, {0xd1, 0x81, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd1, 0x8f, 0xd0, 0xb1, 0xd1, 0x80, 0xd1, 0x8f}, {0xd0, 0xbe, 0xd0, 0xba, 0xd1, 0x82, 0xd1, 0x8f, 0xd0, 0xb1, 0xd1, 0x80, 0xd1, 0x8f}, {0xd0, 0xbd, 0xd0, 0xbe, 0xd1, 0x8f, 0xd0, 0xb1, 0xd1, 0x80, 0xd1, 0x8f}, {0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xb1, 0xd1, 0x80, 0xd1, 0x8f}},
		daysAbbreviated:        [][]uint8{{0xd0, 0xb2, 0xd1, 0x81}, {0xd0, 0xbf, 0xd0, 0xbd}, {0xd0, 0xb2, 0xd1, 0x82}, {0xd1, 0x81, 0xd1, 0x80}, {0xd1, 0x87, 0xd1, 0x82}, {0xd0, 0xbf, 0xd1, 0x82}, {0xd1, 0x81, 0xd0, 0xb1}},
		daysNarrow:             [][]uint8{{0xd0, 0xb2, 0xd1, 0x81}, {0xd0, 0xbf, 0xd0, 0xbd}, {0xd0, 0xb2, 0xd1, 0x82}, {0xd1, 0x81, 0xd1, 0x80}, {0xd1, 0x87, 0xd1, 0x82}, {0xd0, 0xbf, 0xd1, 0x82}, {0xd1, 0x81, 0xd0, 0xb1}},
		daysShort:              [][]uint8{{0xd0, 0xb2, 0xd1, 0x81}, {0xd0, 0xbf, 0xd0, 0xbd}, {0xd0, 0xb2, 0xd1, 0x82}, {0xd1, 0x81, 0xd1, 0x80}, {0xd1, 0x87, 0xd1, 0x82}, {0xd0, 0xbf, 0xd1, 0x82}, {0xd1, 0x81, 0xd0, 0xb1}},
		daysWide:               [][]uint8{{0xd0, 0xb2, 0xd0, 0xbe, 0xd1, 0x81, 0xd0, 0xba, 0xd1, 0x80, 0xd0, 0xb5, 0xd1, 0x81, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x8c, 0xd0, 0xb5}, {0xd0, 0xbf, 0xd0, 0xbe, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba}, {0xd0, 0xb2, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba}, {0xd1, 0x81, 0xd1, 0x80, 0xd0, 0xb5, 0xd0, 0xb4, 0xd0, 0xb0}, {0xd1, 0x87, 0xd0, 0xb5, 0xd1, 0x82, 0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb3}, {0xd0, 0xbf, 0xd1, 0x8f, 0xd1, 0x82, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x86, 0xd0, 0xb0}, {0xd1, 0x81, 0xd1, 0x83, 0xd0, 0xb1, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x82, 0xd0, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0xd0, 0x94, 0xd0, 0x9f}, {0xd0, 0x9f, 0xd0, 0x9f}},
		periodsNarrow:          [][]uint8{{0xd0, 0x94, 0xd0, 0x9f}, {0xd0, 0x9f, 0xd0, 0x9f}},
		periodsWide:            [][]uint8{{0xd0, 0x94, 0xd0, 0x9f}, {0xd0, 0x9f, 0xd0, 0x9f}},
		erasAbbreviated:        [][]uint8{{0xd0, 0xb4, 0xd0, 0xbe, 0x20, 0xd0, 0xbd, 0x2e, 0x20, 0xd1, 0x8d, 0x2e}, {0xd0, 0xbd, 0x2e, 0x20, 0xd1, 0x8d, 0x2e}},
		erasNarrow:             [][]uint8{{0xd0, 0xb4, 0xd0, 0xbe, 0x20, 0xd0, 0xbd, 0x2e, 0xd1, 0x8d, 0x2e}, {0xd0, 0xbd, 0x2e, 0xd1, 0x8d, 0x2e}},
		erasWide:               [][]uint8{{0xd0, 0xb4, 0xd0, 0xbe, 0x20, 0xd0, 0xa0, 0xd0, 0xbe, 0xd0, 0xb6, 0xd0, 0xb4, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb2, 0xd0, 0xb0, 0x20, 0xd0, 0xa5, 0xd1, 0x80, 0xd0, 0xb8, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xb0}, {0xd0, 0xbe, 0xd1, 0x82, 0x20, 0xd0, 0xa0, 0xd0, 0xbe, 0xd0, 0xb6, 0xd0, 0xb4, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb2, 0xd0, 0xb0, 0x20, 0xd0, 0xa5, 0xd1, 0x80, 0xd0, 0xb8, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xb0}}}
}

// Locale returns the current translators string locale
func (ru *ru) Locale() string {
	return ru.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ru'
func (ru *ru) PluralsCardinal() []locales.PluralRule {
	return ru.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ru'
func (ru *ru) PluralsOrdinal() []locales.PluralRule {
	return ru.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ru'
func (ru *ru) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if v == 0 && iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && iMod10 == 0) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 11 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ru'
func (ru *ru) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ru'
func (ru *ru) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ru.CardinalPluralRule(num1, v1)
	end := ru.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
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

// FmtNumber returns 'num' with digits/precision of 'v' for 'ru' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ru *ru) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ru.decimal) + len(ru.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ru.group) - 1; j >= 0; j-- {
					b = append(b, ru.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ru.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ru' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ru *ru) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ru.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ru.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ru.percentSuffix...)

	b = append(b, ru.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ru'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ru *ru) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ru.currencies[currency]
	l := len(s) + len(ru.decimal) + len(ru.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ru.group) - 1; j >= 0; j-- {
					b = append(b, ru.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ru.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ru.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ru.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ru'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ru *ru) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ru.currencies[currency]
	l := len(s) + len(ru.decimal) + len(ru.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ru.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ru.group) - 1; j >= 0; j-- {
					b = append(b, ru.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ru.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ru.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ru.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ru.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}
