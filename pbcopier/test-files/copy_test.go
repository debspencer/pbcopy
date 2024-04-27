package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func TestCopy(t *testing.T) {
	bit := true
	i32 := int32(1)
	i64 := int64(1)
	f32 := float32(1.0)
	f64 := float64(1.0)
	str := "one"
	on := Base_on

	orig := &Base{
		Bit:            true,
		Smallint:       1,
		Largeint:       1,
		Smallfloat:     1.0,
		Largefloat:     1.0,
		Text:           "one",
		Child:          &Child{Kids: "goats"},
		Onoff:          Base_on,
		Optbit:         &bit,
		Optsmallint:    &i32,
		Optlargeint:    &i64,
		Optsmallfloat:  &f32,
		Optlargefloat:  &f64,
		Opttext:        &str,
		Optchild:       &Child{Kids: "lambs"},
		Optonoff:       &on,
		Listbit:        []bool{false, true},
		Listsmallint:   []int32{0, 1},
		Listlargeint:   []int64{0, 1},
		Listsmallfloat: []float32{1.0, 2.0},
		Listlargefloat: []float64{1.0, 2.0},
		Listtext:       []string{"one", "two"},
		Listchild:      []*Child{{Kids: "bunnies"}, {Kids: "kittens"}},
		Listonoff:      []Base_OnOff{Base_off, Base_on},
		Mapsmallint:    map[int32]int32{1: 2},
		Maplargeint:    map[int64]int64{1: 2},
		Maptext:        map[string]string{"one": "two"},
		Mapchild:       map[string]*Child{"dogs": {Kids: "puppies"}},
	}

	deep1 := orig.DeepCopy()
	deep2 := orig.DeepCopy()

	shallow1 := orig.ShallowCopy()
	shallow2 := orig.ShallowCopy()

	compare := func(t *testing.T, o1, o2 interface{}) {
		err := compareObj(o1, o2)
		if err != nil {
			t.Error(err.Error())
		}
	}

	t.Run("basic copy", func(t *testing.T) {
		compare(t, orig, deep1)
		compare(t, orig, deep2)

		compare(t, orig, shallow1)
		compare(t, orig, shallow2)
	})

	t.Run("modify deep", func(t *testing.T) {
		deep2.Smallint = 0
		*deep2.Optsmallint = 0
		compare(t, orig, deep1)

		deep1.Smallint = 0
		*deep1.Optsmallint = 0
		compare(t, deep1, deep2)
	})

	t.Run("modify shallow safe", func(t *testing.T) {
		shallow2.Smallint = 0
		compare(t, orig, shallow1)

		shallow1.Smallint = 0
		compare(t, shallow1, shallow2)
	})

	t.Run("modify shallow usafe", func(t *testing.T) {
		// make new copies
		shallow1 := orig.ShallowCopy()
		shallow2 := orig.ShallowCopy()

		compare(t, orig, shallow1)
		compare(t, orig, shallow2)

		*shallow2.Optsmallint = 0 // shallow will update them all

		compare(t, orig, shallow1)
		compare(t, orig, shallow2)
	})
}

func compareObj(o1, o2 interface{}) error {
	j1, err := json.Marshal(o1)
	if err != nil {
		return fmt.Errorf("Can not Marshal: o1 %v (%#v)", err, o1)
	}
	j2, err := json.Marshal(o2)
	if err != nil {
		return fmt.Errorf("Can not Marshal: o1 %v (%#v)", err, o2)
	}

	differ := gojsondiff.New()
	diff, err := differ.Compare(j1, j2)
	if err != nil {
		return fmt.Errorf("Error: %v\nj1: (%d) '%s'\nj2: (%d) '%s'", err,
			len(j1), string(j1), len(j2), string(j2))
	}
	if diff.Modified() {
		var gotJson map[string]interface{}
		json.Unmarshal(j1, &gotJson)
		formatter := formatter.NewAsciiFormatter(gotJson, formatter.AsciiFormatterConfig{ShowArrayIndex: true})
		diffString, _ := formatter.Format(diff)
		return fmt.Errorf("Diff failed: %s\nj1: (%d) '%s'\nj2: (%d) '%s'", diffString,
			len(j1), string(j1), len(j2), string(j2))
	}

	return nil
}
