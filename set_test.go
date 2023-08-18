package set_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSADD(t *testing.T) {
	type args struct {
		inputs []interface{}
	}
	type want struct {
		expected SetMap
	}
	tests := []struct {
		Name string
		Args args
		Want want
	}{
		{
			Name: "Set integer",
			Args: args{
				inputs: []interface{}{1, 3, 5},
			},
			Want: want{
				expected: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
		},
		{
			Name: "Set integer and string",
			Args: args{
				inputs: []interface{}{1, 3, 5, "hoge", "fuga"},
			},
			Want: want{
				expected: SetMap{
					1:      isMember,
					3:      isMember,
					5:      isMember,
					"hoge": isMember,
					"fuga": isMember,
				},
			},
		},
		{
			Name: "Set duplicate integer",
			Args: args{
				inputs: []interface{}{1, 1, 3, 3, 5, 5},
			},
			Want: want{
				expected: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
		},
		{
			Name: "Set duplicate integer and string",
			Args: args{
				inputs: []interface{}{1, 1, 3, 3, 5, 5, "hoge", "fuga", "hoge", "fuga"},
			},
			Want: want{
				expected: SetMap{
					1:      isMember,
					3:      isMember,
					5:      isMember,
					"hoge": isMember,
					"fuga": isMember,
				},
			},
		},
	}

	for _, test := range tests {
		res := SetMap{}
		res.SADD(test.Args.inputs...)
		assert.Equal(t, test.Want.expected, res)
	}
}

func TestSREM(t *testing.T) {
	type args struct {
		inputs     []interface{}
		initSetMap SetMap
	}
	type want struct {
		expected SetMap
	}
	tests := []struct {
		Name string
		Args args
		Want want
	}{
		{
			Name: "Remove an integer",
			Args: args{
				inputs: []interface{}{1},
				initSetMap: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: SetMap{
					3: isMember,
					5: isMember,
				},
			},
		},
		{
			Name: "Remove integers",
			Args: args{
				inputs: []interface{}{1, 3},
				initSetMap: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: SetMap{
					5: isMember,
				},
			},
		},
		{
			Name: "Remove a string",
			Args: args{
				inputs: []interface{}{"hoge"},
				initSetMap: SetMap{
					"hoge": isMember,
					"fuga": isMember,
					"piyo": isMember,
				},
			},
			Want: want{
				expected: SetMap{
					"fuga": isMember,
					"piyo": isMember,
				},
			},
		},
		{
			Name: "Remove strings",
			Args: args{
				inputs: []interface{}{"hoge", "fuga"},
				initSetMap: SetMap{
					"hoge": isMember,
					"fuga": isMember,
					"piyo": isMember,
				},
			},
			Want: want{
				expected: SetMap{
					"piyo": isMember,
				},
			},
		},
		{
			Name: "Remove integers and strings",
			Args: args{
				inputs: []interface{}{3, 5, "hoge", "fuga"},
				initSetMap: SetMap{
					1:      isMember,
					3:      isMember,
					5:      isMember,
					"hoge": isMember,
					"fuga": isMember,
					"piyo": isMember,
				},
			},
			Want: want{
				expected: SetMap{
					1:      isMember,
					"piyo": isMember,
				},
			},
		},
	}

	for _, test := range tests {
		res := test.Args.initSetMap
		res.SREM(test.Args.inputs...)
		assert.Equal(t, test.Want.expected, res)
	}
}

func TestSISMEMBER(t *testing.T) {
	type args struct {
		inputs     interface{}
		initSetMap SetMap
	}
	type want struct {
		expected bool
	}
	tests := []struct {
		Name string
		Args args
		Want want
	}{
		{
			Name: "Check Exists integer",
			Args: args{
				inputs: 1,
				initSetMap: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: true,
			},
		},
		{
			Name: "Check Exists string",
			Args: args{
				inputs: "hoge",
				initSetMap: SetMap{
					"hoge": isMember,
					"fuga": isMember,
					"piyo": isMember,
				},
			},
			Want: want{
				expected: true,
			},
		},
		{
			Name: "Check Exists not exists integer",
			Args: args{
				inputs: 10,
				initSetMap: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: false,
			},
		},
		{
			Name: "Check Exists not exists string",
			Args: args{
				inputs: "paco",
				initSetMap: SetMap{
					"hoge": isMember,
					"fuga": isMember,
					"piyo": isMember,
				},
			},
			Want: want{
				expected: false,
			},
		},
	}

	for _, test := range tests {
		sm := test.Args.initSetMap
		res := sm.SISMEMBER(test.Args.inputs)
		assert.Equal(t, test.Want.expected, res)
	}
}

func TestSCARD(t *testing.T) {
	type args struct {
		initSetMap SetMap
	}
	type want struct {
		expected int
	}
	tests := []struct {
		Name string
		Args args
		Want want
	}{
		{
			Name: "Length 3",
			Args: args{
				initSetMap: SetMap{
					1: isMember,
					3: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: 3,
			},
		},
		{
			Name: "Length 5",
			Args: args{
				initSetMap: SetMap{
					1: isMember,
					2: isMember,
					3: isMember,
					4: isMember,
					5: isMember,
				},
			},
			Want: want{
				expected: 5,
			},
		},
		{
			Name: "Length 10",
			Args: args{
				initSetMap: SetMap{
					1:  isMember,
					2:  isMember,
					3:  isMember,
					4:  isMember,
					5:  isMember,
					6:  isMember,
					7:  isMember,
					8:  isMember,
					9:  isMember,
					10: isMember,
				},
			},
			Want: want{
				expected: 10,
			},
		},
	}

	for _, test := range tests {
		s := test.Args.initSetMap
		res := s.SCARD()
		assert.Equal(t, test.Want.expected, res)
	}
}
