package externalapi

import (
	"testing"
)

type testHashToCompare struct {
	hash           *DomainHash
	expectedResult bool
}

type testHashStruct struct {
	baseHash          *DomainHash
	hashesToCompareTo []testHashToCompare
}

func initTestDomainHashForEqual() []*testHashStruct {
	tests := []*testHashStruct{
		{
			baseHash: nil,
			hashesToCompareTo: []testHashToCompare{
				{
					hash:           nil,
					expectedResult: true,
				}, {
					hash: NewDomainHashFromByteArray(&[DomainHashSize]byte{
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
					expectedResult: false,
				},
			},
		}, {
			baseHash: NewDomainHashFromByteArray(&[DomainHashSize]byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}),
			hashesToCompareTo: []testHashToCompare{
				{
					hash:           nil,
					expectedResult: false,
				}, {
					hash: NewDomainHashFromByteArray(&[DomainHashSize]byte{
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
					expectedResult: false,
				}, {
					hash: NewDomainHashFromByteArray(&[DomainHashSize]byte{
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
						0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF}),
					expectedResult: true,
				},
			},
		},
	}
	return tests
}

func TestDomainHash_Equal(t *testing.T) {
	hashTests := initTestDomainHashForEqual()
	for i, test := range hashTests {
		for j, subTest := range test.hashesToCompareTo {
			result1 := test.baseHash.Equal(subTest.hash)
			if result1 != subTest.expectedResult {
				t.Fatalf("Test #%d:%d: Expected %t but got %t", i, j, subTest.expectedResult, result1)
			}
			result2 := subTest.hash.Equal(test.baseHash)
			if result2 != subTest.expectedResult {
				t.Fatalf("Test #%d:%d: Expected %t but got %t", i, j, subTest.expectedResult, result2)
			}
		}
	}
}
