/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	uti "github.com/craterdog/go-essential-utilities/v8"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	cmp "math/cmplx"
	sts "strings"
	syn "sync"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	uti.MakeDirectory(testDirectory)
	var filenames = uti.ReadDirectory(testDirectory)
	for _, filename := range filenames {
		if sts.HasSuffix(filename, ".bali") {
			filename = testDirectory + filename
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var document = doc.ParseDocument(source)
			var formatted = doc.FormatDocument(document)
			ass.Equal(t, source, formatted)
		}
	}
}

func TestParameterAccess(t *tes.T) {
	var source = `[ ]`
	var component = doc.ParseComponent(source)
	var key = doc.Symbol("$type")
	var parameter = component.GetParameter(key)
	ass.Equal(t, nil, parameter)

	source = `[ ]($type: "foo")`
	component = doc.ParseComponent(source)
	parameter = component.GetParameter(key)
	ass.Equal(t, "\"foo\"", doc.FormatComponent(parameter))

	source = `[ ]($type: "foo" $hype: /bar $skype: none)`
	component = doc.ParseComponent(source)
	key = doc.Symbol("$type")
	parameter = component.GetParameter(key)
	ass.Equal(t, "\"foo\"", doc.FormatComponent(parameter))
	key = doc.Symbol("$hype")
	parameter = component.GetParameter(key)
	ass.Equal(t, "/bar", doc.FormatComponent(parameter))
	key = doc.Symbol("$skype")
	parameter = component.GetParameter(key)
	ass.Equal(t, "none", doc.FormatComponent(parameter))
}

func TestSubcomponentAccess(t *tes.T) {
	var source = `[ ]`
	var component = doc.ParseComponent(source)
	var index = 1
	var content = component.GetSubcomponent(index)
	ass.Equal(t, nil, content)

	var component2 = doc.ParseComponent("$new")
	index = 0 // Append a new content.
	component.SetSubcomponent(component2, index)
	index = 1
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$new", doc.FormatComponent(content))
	content = component.RemoveSubcomponent(index)
	ass.Equal(t, "$new", doc.FormatComponent(content))
	content = component.GetSubcomponent(index)
	ass.Equal(t, nil, content)

	source = `[
    $alpha
    $beta
    $gamma
]`
	component = doc.ParseComponent(source)
	index = 1
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$alpha", doc.FormatComponent(content))
	index = 2
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$beta", doc.FormatComponent(content))
	index = 3
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$gamma", doc.FormatComponent(content))
	component2 = doc.ParseComponent("$delta")
	component.SetSubcomponent(component2, index)
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$delta", doc.FormatComponent(content))
	index = 0
	component2 = doc.ParseComponent("$epsilon")
	component.SetSubcomponent(component2, index)
	index = 4
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$epsilon", doc.FormatComponent(content))
	content = component.RemoveSubcomponent(index)
	ass.Equal(t, "$epsilon", doc.FormatComponent(content))
	index = -1
	content = component.GetSubcomponent(index)
	ass.Equal(t, "$delta", doc.FormatComponent(content))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	component = doc.ParseComponent(source)
	var key = doc.Symbol("alpha")
	content = component.GetSubcomponent(key)
	ass.Equal(t, "\"1\"", doc.FormatComponent(content))
	key = doc.Symbol("beta")
	content = component.GetSubcomponent(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(content))
	key = doc.Symbol("gamma")
	content = component.GetSubcomponent(key)
	ass.Equal(t, "\"3\"", doc.FormatComponent(content))
	component2 = doc.ParseComponent("\"5\"")
	component.SetSubcomponent(component2, key)
	content = component.GetSubcomponent(key)
	ass.Equal(t, "\"5\"", doc.FormatComponent(content))
	content = component.RemoveSubcomponent(key)
	ass.Equal(t, "\"5\"", doc.FormatComponent(content))
	key = doc.Symbol("beta")
	content = component.GetSubcomponent(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(content))

	source = `[
    $items: [
        1
        2
        3
    ]
    $attributes: [
        $alpha: "1"
        $beta: "2"
        $gamma: "3"
    ]
]`
	component = doc.ParseComponent(source)
	key = doc.Symbol("items")
	index = 2
	content = component.GetSubcomponent(key, index)
	ass.Equal(t, "2", doc.FormatComponent(content))
	key = doc.Symbol("attributes")
	var key2 = doc.Symbol("gamma")
	content = component.GetSubcomponent(key, key2)
	ass.Equal(t, "\"3\"", doc.FormatComponent(content))
	key2 = doc.Symbol("delta")
	component2 = doc.ParseComponent("\"4\"")
	component.SetSubcomponent(component2, key, key2)
	content = component.GetSubcomponent(key, key2)
	ass.Equal(t, "\"4\"", doc.FormatComponent(content))
	content = component.RemoveSubcomponent(key, key2)
	ass.Equal(t, "\"4\"", doc.FormatComponent(content))
	key2 = doc.Symbol("beta")
	content = component.GetSubcomponent(key, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(content))

	source = `[
    [
        1
        2
        [
            ~pi: 3.14
            ~tau: 6.28
        ]
    ]
    [
        $alpha: [
            'a'
            'b'
            'c'
        ]
        $beta: "2"
        $gamma: "3"
    ]
]`
	component = doc.ParseComponent(source)
	index = 1
	var index2 = 2
	content = component.GetSubcomponent(index, index2)
	ass.Equal(t, "2", doc.FormatComponent(content))
	index = 2
	key2 = doc.Symbol("beta")
	content = component.GetSubcomponent(index, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(content))
	index = 1
	index2 = 3
	var key3 = doc.Angle("~tau")
	content = component.GetSubcomponent(index, index2, key3)
	ass.Equal(t, "6.28", doc.FormatComponent(content))
	component2 = doc.ParseComponent("~τ")
	component.SetSubcomponent(component2, index, index2, key3)
	content = component.GetSubcomponent(index, index2, key3)
	ass.Equal(t, "~τ", doc.FormatComponent(content))
	index = 2
	key2 = doc.Symbol("alpha")
	var index3 = -1
	component.RemoveSubcomponent(index, key2, index3)
	content = component.GetSubcomponent(index, key2, index3)
	ass.Equal(t, "'b'", doc.FormatComponent(content))
	index = 3
	content = component.GetSubcomponent(index)
	ass.Equal(t, nil, content)
	index = 2
	key2 = doc.Symbol("delta")
	content = component.GetSubcomponent(index, key2)
	ass.Equal(t, nil, content)
}

func TestModuleFunctions(t *tes.T) {
	doc.Collator[any]()
	doc.Collator[any](8)
	var sorter = doc.Sorter[any]()
	doc.Sorter[any](sorter.GetRanker())
	doc.List[string]()
	var list = doc.List[string]([]string{"A"})
	doc.List[string](list)
	doc.ListClass[string]().Concatenate(list, list)
	var association = doc.Association[string, int]("A", 1)
	var catalog = doc.Catalog[string, int]()
	doc.Catalog[string, int]([]doc.AssociationLike[string, int]{association})
	doc.Catalog[string, int](catalog.AsMap())
	doc.Catalog[string, int](catalog)
	doc.CatalogClass[string, int]().Extract(catalog, list)
	doc.CatalogClass[string, int]().Merge(catalog, catalog)
	doc.Queue[string]()
	doc.Queue[string](8)
	var queue = doc.Queue[string](list.AsArray())
	doc.Queue[string](queue)
	var group = new(syn.WaitGroup)
	defer group.Wait()
	var queues = doc.QueueClass[string]().Fork(group, queue, 2)
	doc.QueueClass[string]().Split(group, queue, 2)
	doc.QueueClass[string]().Join(group, queues)
	queue.CloseChannel()
	var set = doc.Set[string]()
	doc.Set[string](set.GetCollator())
	doc.Set[string](set.AsArray())
	doc.Set[string](set)
	doc.SetClass[string]().And(set, set)
	doc.SetClass[string]().Ior(set, set)
	doc.SetClass[string]().San(set, set)
	doc.SetClass[string]().Xor(set, set)
	doc.Stack[string]()
	doc.Stack[string](8)
	doc.Stack[string](list.AsArray())
	doc.Stack[string](list)
}

// Pure Structure
type Fuz struct {
	Bar string
}

func TestRank(t *tes.T) {
	ass.Equal(t, "LesserRank", doc.LesserRank.String())
	ass.Equal(t, "EqualRank", doc.EqualRank.String())
	ass.Equal(t, "GreaterRank", doc.GreaterRank.String())
}

func TestCompareMaximum(t *tes.T) {
	var collator = doc.Collator[any](1)
	var list = doc.List[any]([]any{"foo", []int{1, 2, 3}})
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 1", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.CompareValues(list, list)
}

func TestRankMaximum(t *tes.T) {
	var collator = doc.Collator[any](1)
	var list = doc.List[any]([]any{"foo", []int{1, 2, 3}})
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The maximum traversal depth was exceeded: 1", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	collator.RankValues(list, list)
}

func TestIteratorsWithLists(t *tes.T) {
	var list = doc.List[int]([]int{1, 2, 3, 4, 5})
	list = doc.List[int](list)
	var iterator = list.GetIterator()
	ass.False(t, iterator.IsEmpty())
	ass.True(t, iterator.GetSize() == 5)
	ass.True(t, iterator.GetSlot() == 0)
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetPrevious())
	iterator.SetSlot(2)
	ass.True(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 3, iterator.GetNext())
	ass.False(t, iterator.IsEmpty())
	ass.True(t, iterator.GetSize() == 5)
	ass.True(t, iterator.GetSlot() == 3)
	iterator.ToEnd()
	ass.True(t, iterator.HasPrevious())
	ass.False(t, iterator.HasNext())
	ass.Equal(t, 5, iterator.GetPrevious())
	iterator.ToStart()
	ass.False(t, iterator.HasPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, 1, iterator.GetNext())
}

func TestSortingEmpty(t *tes.T) {
	var collator = doc.Collator[any]()
	var ranker = collator.RankValues
	var sorter = doc.Sorter[any](ranker)
	var empty = []any{}
	sorter.SortValues(empty)
}

func TestSortingIntegers(t *tes.T) {
	var collator = doc.Collator[int]()
	var ranker = collator.RankValues
	var sorter = doc.Sorter[int](ranker)
	var unsorted = []int{4, 3, 1, 5, 2}
	var sorted = []int{1, 2, 3, 4, 5}
	sorter.SortValues(unsorted)
	ass.Equal(t, sorted, unsorted)
}

func TestSortingStrings(t *tes.T) {
	var collator = doc.Collator[string]()
	var ranker = collator.RankValues
	var sorter = doc.Sorter[string](ranker)
	var unsorted = []string{"alpha", "beta", "gamma", "delta"}
	var sorted = []string{"alpha", "beta", "delta", "gamma"}
	sorter.SortValues(unsorted)
	ass.Equal(t, sorted, unsorted)
}

// COLLECTION

func TestCatalogConstructors(t *tes.T) {
	doc.Catalog[rune, int64]()
	doc.Catalog[rune, int64]([]doc.AssociationLike[rune, int64]{})
	doc.Catalog[rune, int64](map[rune]int64{})
	var sequence = doc.Catalog[rune, int64](map[rune]int64{
		'a': 1,
		'b': 2,
		'c': 3,
	})
	var catalog = doc.Catalog[rune, int64](sequence)
	ass.Equal(t, sequence.AsArray(), catalog.AsArray())
}

func TestCatalogsWithPrimitivesAndStrings(t *tes.T) {
	var catalog = doc.Catalog[any, string]()

	var binaryString = `'>
    0123456789abcdefghijk
<'`
	var binary = doc.Binary(binaryString)
	catalog.SetValue(binary, binaryString)
	ass.Equal(t, catalog.GetValue(binary), binary.AsSource())

	var nameString = "/foo/5bar"
	var name = doc.Name(nameString)
	catalog.SetValue(name, nameString)
	ass.Equal(t, catalog.GetValue(name), name.AsSource())

	var narrativeString = `">
    This is a narrative.
<"`
	var narrative = doc.Narrative(narrativeString)
	catalog.SetValue(narrative, narrativeString)
	ass.Equal(t, catalog.GetValue(narrative), narrative.AsSource())

	var patternString = `"b[aeiou]g"?`
	var pattern = doc.Pattern(patternString)
	catalog.SetValue(pattern, patternString)
	ass.Equal(t, catalog.GetValue(pattern), pattern.AsSource())

	var quoteString = `"To be or not to be..."`
	var quote = doc.Quote(quoteString)
	catalog.SetValue(quote, quoteString)
	ass.Equal(t, catalog.GetValue(quote), quote.AsSource())

	var tag = doc.Tag(16)
	var tagString = tag.AsSource()
	catalog.SetValue(tag, tagString)
	ass.Equal(t, catalog.GetValue(tag), tag.AsSource())

	var versionString = "v1.2.3"
	var version = doc.Version(versionString)
	catalog.SetValue(version, versionString)
	ass.Equal(t, catalog.GetValue(version), version.AsSource())
}

func TestCatalogsWithStringsAndIntegers(t *tes.T) {
	var catalogCollator = doc.Collator[doc.CatalogLike[string, int]]()
	var keys = doc.List[string]([]string{"foo", "bar"})
	var association1 = doc.Association("foo", 1)
	var association2 = doc.Association("bar", 2)
	var association3 = doc.Association("baz", 3)
	var catalog = doc.Catalog[string, int]()
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
	ass.Equal(t, []string{}, catalog.GetKeys().AsArray())
	ass.Equal(t, []doc.AssociationLike[string, int]{}, catalog.AsArray())
	var iterator = catalog.GetIterator()
	ass.False(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	catalog.SortValues()
	catalog.ShuffleValues()
	catalog.RemoveAll()
	catalog.SetValue(association1.GetKey(), association1.GetValue())
	ass.False(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 1)
	catalog.SetValue(association2.GetKey(), association2.GetValue())
	catalog.SetValue(association3.GetKey(), association3.GetValue())
	ass.True(t, catalog.GetSize() == 3)
	var catalog2 = doc.Catalog[string, int](catalog)
	ass.True(t, catalogCollator.CompareValues(catalog, catalog2))
	var m = doc.Catalog[string, int](map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	})
	var associationCollator = doc.Collator[doc.AssociationLike[string, int]]()
	var catalog3 = doc.Catalog[string, int](m)
	catalog2.SortValues()
	catalog3.SortValuesWithRanker(associationCollator.RankValues)
	ass.True(t, catalogCollator.CompareValues(catalog2, catalog3))
	iterator = catalog.GetIterator()
	ass.True(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	ass.Equal(t, association1, iterator.GetNext())
	ass.True(t, iterator.HasPrevious())
	iterator.ToEnd()
	ass.False(t, iterator.HasNext())
	ass.True(t, iterator.HasPrevious())
	ass.Equal(t, association3, iterator.GetPrevious())
	ass.True(t, iterator.HasNext())
	ass.Equal(t, []string{"foo", "bar", "baz"}, catalog.GetKeys().AsArray())
	ass.Equal(t, 3, int(catalog.GetValue("baz")))
	catalog.SetValue("bar", 5)
	ass.Equal(t, []int{1, 5}, catalog.GetValues(keys).AsArray())
	catalog.SortValues()
	ass.Equal(t, []string{"bar", "baz", "foo"}, catalog.GetKeys().AsArray())
	catalog.ReverseValues()
	ass.Equal(t, []string{"foo", "baz", "bar"}, catalog.GetKeys().AsArray())
	catalog.ReverseValues()
	ass.Equal(t, []int{1, 5}, catalog.RemoveValues(keys).AsArray())
	ass.True(t, catalog.GetSize() == 1)
	ass.Equal(t, 3, int(catalog.RemoveValue("baz")))
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
	catalog.RemoveAll()
	ass.True(t, catalog.IsEmpty())
	ass.True(t, catalog.GetSize() == 0)
}

func TestCatalogsWithMerge(t *tes.T) {
	var collator = doc.Collator[doc.CatalogLike[string, int]]()
	var association1 = doc.Association("foo", 1)
	var association2 = doc.Association("bar", 2)
	var association3 = doc.Association("baz", 3)
	var catalog1 = doc.Catalog[string, int]()
	catalog1.SetValue(association1.GetKey(), association1.GetValue())
	catalog1.SetValue(association2.GetKey(), association2.GetValue())
	var catalog2 = doc.Catalog[string, int]()
	catalog2.SetValue(association2.GetKey(), association2.GetValue())
	catalog2.SetValue(association3.GetKey(), association3.GetValue())
	var catalog3 = doc.CatalogClass[string, int]().Merge(catalog1, catalog2)
	var catalog4 = doc.Catalog[string, int]()
	catalog4.SetValue(association1.GetKey(), association1.GetValue())
	catalog4.SetValue(association2.GetKey(), association2.GetValue())
	catalog4.SetValue(association3.GetKey(), association3.GetValue())
	ass.True(t, collator.CompareValues(catalog3, catalog4))
}

func TestCatalogsWithExtract(t *tes.T) {
	var keys = doc.List[string]([]string{"foo", "baz"})
	var association1 = doc.Association("foo", 1)
	var association2 = doc.Association("bar", 2)
	var association3 = doc.Association("baz", 3)
	var catalog1 = doc.Catalog[string, int]()
	catalog1.SetValue(association1.GetKey(), association1.GetValue())
	catalog1.SetValue(association2.GetKey(), association2.GetValue())
	catalog1.SetValue(association3.GetKey(), association3.GetValue())
	var catalog2 = doc.CatalogClass[string, int]().Extract(catalog1, keys)
	var catalog3 = doc.Catalog[string, int]()
	catalog3.SetValue(association1.GetKey(), association1.GetValue())
	catalog3.SetValue(association3.GetKey(), association3.GetValue())
	var collator = doc.Collator[doc.CatalogLike[string, int]]()
	ass.True(t, collator.CompareValues(catalog2, catalog3))
	var catalog4 = doc.Catalog[string, int]([]doc.AssociationLike[string, int]{
		association1,
		association2,
		association3,
	})
	ass.True(t, collator.CompareValues(catalog1, catalog4))
}

func TestCatalogsWithEmptyCatalogs(t *tes.T) {
	var keys = doc.List[int]()
	var catalog1 = doc.Catalog[int, string]()
	var catalog2 = doc.Catalog[int, string]()
	var catalog3 = doc.CatalogClass[int, string]().Merge(catalog1, catalog2)
	var catalog4 = doc.CatalogClass[int, string]().Extract(catalog1, keys)
	var collator = doc.Collator[doc.CatalogLike[int, string]]()
	ass.True(t, collator.CompareValues(catalog1, catalog2))
	ass.True(t, collator.CompareValues(catalog2, catalog3))
	ass.True(t, collator.CompareValues(catalog3, catalog4))
	ass.True(t, collator.CompareValues(catalog4, catalog1))
}

func TestIntervalIterators(t *tes.T) {
	var glyphs = doc.Interval[doc.GlyphLike](
		doc.Exclusive,
		doc.Glyph(65),
		doc.Glyph(70),
		doc.Exclusive,
	)
	var iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'B', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'E', iterator.GetPrevious().AsIntrinsic())

	glyphs = doc.Interval[doc.GlyphLike](
		doc.Inclusive,
		doc.Glyph(65),
		doc.Glyph(70),
		doc.Exclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'A', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'E', iterator.GetPrevious().AsIntrinsic())

	glyphs = doc.Interval[doc.GlyphLike](
		doc.Exclusive,
		doc.Glyph(65),
		doc.Glyph(70),
		doc.Inclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'B', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'F', iterator.GetPrevious().AsIntrinsic())

	glyphs = doc.Interval[doc.GlyphLike](
		doc.Inclusive,
		doc.Glyph(65),
		doc.Glyph(70),
		doc.Inclusive,
	)
	iterator = glyphs.GetIterator()
	ass.Equal(t, glyphs.GetSize(), iterator.GetSize())
	ass.Equal(t, 'A', iterator.GetNext().AsIntrinsic())
	iterator.ToEnd()
	ass.Equal(t, 'F', iterator.GetPrevious().AsIntrinsic())

}

func TestListConstructors(t *tes.T) {
	doc.List[int64]()
	var sequence = doc.List[int64]([]int64{1, 2, 3})
	var list = doc.List[int64](sequence)
	ass.Equal(t, sequence.AsArray(), list.AsArray())
}

func TestListsWithStrings(t *tes.T) {
	var collator = doc.Collator[doc.ListLike[string]]()
	var foo = doc.List[string]([]string{"foo"})
	var bar = doc.List[string]([]string{"bar"})
	var baz = doc.List[string]([]string{"baz"})
	var foz = doc.List[string]([]string{"foz"})
	var barbaz = doc.List[string]([]string{"bar", "baz"})
	var bazbaz = doc.List[string]([]string{"baz", "baz"})
	var foobar = doc.List[string]([]string{"foo", "bar"})
	var baxbaz = doc.List[string]([]string{"bax", "baz"})
	var baxbez = doc.List[string]([]string{"bax", "bez"})
	var barfoobax = doc.List[string]([]string{"bar", "foo", "bax"})
	var foobazbar = doc.List[string]([]string{"foo", "baz", "bar"})
	var foobarbaz = doc.List[string]([]string{"foo", "bar", "baz"})
	var barbazfoo = doc.List[string]([]string{"bar", "baz", "foo"})
	var list = doc.List[string]()
	ass.True(t, list.IsEmpty())
	ass.True(t, list.GetSize() == 0)
	ass.False(t, list.ContainsValue("bax"))
	ass.Equal(t, []string{}, list.AsArray())
	var iterator = list.GetIterator()
	ass.False(t, iterator.HasNext())
	ass.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	list.ShuffleValues()
	list.SortValues()
	list.RemoveAll()                      //       [ ]
	list.AppendValue("foo")               //       ["foo"]
	ass.False(t, list.IsEmpty())          //       ["foo"]
	ass.True(t, list.GetSize() == 1)      //       ["foo"]
	ass.Equal(t, "foo", list.GetValue(1)) //       ["foo"]
	list.AppendValues(barbaz)             //       ["foo", "bar", "baz"]
	ass.True(t, list.GetSize() == 3)      //       ["foo", "bar", "baz"]
	ass.Equal(t, "foo", list.GetValue(1)) //       ["foo", "bar", "baz"]
	ass.True(t, collator.CompareValues(doc.List[string](list.AsArray()), list))
	ass.Equal(t, barbaz.AsArray(), list.GetValues(2, 3).AsArray())
	ass.Equal(t, foo.AsArray(), list.GetValues(1, 1).AsArray())
	var list2 = doc.List[string](list)
	ass.True(t, collator.CompareValues(list, list2))
	var array = doc.List[string]([]string{"foo", "bar", "baz"})
	var list3 = doc.List[string](array)
	list2.SortValues()
	list3.SortValues()
	ass.True(t, collator.CompareValues(list2, list3))
	iterator = list.GetIterator()               // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())             // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasPrevious())        // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", iterator.GetNext())     // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())         // ["foo", "bar", "baz"]
	iterator.ToEnd()                            // ["foo", "bar", "baz"]
	ass.False(t, iterator.HasNext())            // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasPrevious())         // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", iterator.GetPrevious()) // ["foo", "bar", "baz"]
	ass.True(t, iterator.HasNext())             // ["foo", "bar", "baz"]
	list.ShuffleValues()                        // [ ?, ?, ? ]
	list.RemoveAll()                            // [ ]
	ass.True(t, list.IsEmpty())                 // [ ]
	ass.True(t, list.GetSize() == 0)            // [ ]
	list.InsertValue(0, "baz")                  // ["baz"]
	ass.True(t, list.GetSize() == 1)            // ["baz"]
	ass.Equal(t, "baz", list.GetValue(-1))      // ["baz"]
	list.InsertValues(0, foobar)                // ["foo", "bar", "baz"]
	ass.True(t, list.GetSize() == 3)            // ["foo", "bar", "baz"]
	ass.Equal(t, "foo", list.GetValue(-3))      // ["foo", "bar", "baz"]
	ass.Equal(t, "bar", list.GetValue(-2))      // ["foo", "bar", "baz"]
	ass.Equal(t, "baz", list.GetValue(-1))      // ["foo", "bar", "baz"]
	list.ReverseValues()                        // ["baz", "bar", "foo"]
	ass.Equal(t, "foo", list.GetValue(-1))      // ["baz", "bar", "foo"]
	ass.Equal(t, "bar", list.GetValue(-2))      // ["baz", "bar", "foo"]
	ass.Equal(t, "baz", list.GetValue(-3))      // ["baz", "bar", "foo"]
	list.ReverseValues()                        // ["foo", "bar", "baz"]
	ass.True(t, list.GetIndex("foz") == 0)      // ["foo", "bar", "baz"]
	ass.True(t, list.GetIndex("baz") == 3)      // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsValue("baz"))      // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsValue("bax"))     // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAny(baxbaz))       // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAny(baxbez))      // ["foo", "bar", "baz"]
	ass.True(t, list.ContainsAll(barbaz))       // ["foo", "bar", "baz"]
	ass.False(t, list.ContainsAll(baxbaz))      // ["foo", "bar", "baz"]
	list.SetValue(3, "bax")                     // ["foo", "bar", "bax"]
	list.InsertValues(3, baz)                   // ["foo", "bar", "bax", "baz"]
	ass.True(t, list.GetSize() == 4)            // ["foo", "bar", "bax", "baz"]
	ass.Equal(t, "baz", list.GetValue(4))       // ["foo", "bar", "bax", "baz"]
	list.InsertValue(4, "bar")                  // ["foo", "bar", "bax", "baz", "bar"]
	ass.True(t, list.GetSize() == 5)            // ["foo", "bar", "bax", "baz", "bar"]
	ass.Equal(t, "bar", list.GetValue(5))       // ["foo", "bar", "bax", "baz", "bar"]
	list.InsertValue(2, "foo")                  // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.True(t, list.GetSize() == 6)            // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bar", list.GetValue(2))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "foo", list.GetValue(3))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, "bax", list.GetValue(4))       // ["foo", "bar", "foo", "bax", "baz", "bar"]
	ass.Equal(t, bar.AsArray(), list.GetValues(6, 6).AsArray())
	list.InsertValues(5, baz)             //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.True(t, list.GetSize() == 7)      //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "bax", list.GetValue(4)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", list.GetValue(5)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, "baz", list.GetValue(6)) //       ["foo", "bar", "foo", "bax", "baz", "baz", "bar"]
	ass.Equal(t, barfoobax.AsArray(), list.GetValues(2, -4).AsArray())
	list.SetValues(2, foobazbar) //                        ["foo", "foo", "baz", "bar", "baz", "baz", "bar"]
	ass.Equal(t, foobazbar.AsArray(), list.GetValues(2, -4).AsArray())
	list.SetValues(-1, foz)
	ass.Equal(t, "foz", list.GetValue(-1)) //      ["foo", "foo", "baz", "bar", "baz", "baz", "foz"]
	list.SortValues()                      //      ["bar", "baz", "baz", "baz", "foo", "foo", "foz"]

	ass.Equal(t, bazbaz.AsArray(), list.RemoveValues(2, -5).AsArray()) // ["bar", "baz", "foo", "foo", "foz"]
	ass.Equal(t, barbaz.AsArray(), list.RemoveValues(1, 2).AsArray())  // ["foo", "foo", "foz"]
	ass.Equal(t, "foz", list.RemoveValue(-1))                          // ["foo", "foo"]
	ass.True(t, list.GetSize() == 2)                                   // ["foo", "foo"]
	list.RemoveAll()                                                   // [ ]
	ass.True(t, list.GetSize() == 0)                                   // [ ]
	list.SortValues()                                                  // [ ]
	list.AppendValues(foobarbaz)                                       // ["foo", "bar", "baz"]
	list.SortValues()                                                  // ["bar", "baz", "foo"]
	ass.Equal(t, barbazfoo.AsArray(), list.AsArray())                  // ["bar", "baz", "foo"]
	list.RemoveAll()                                                   // [ ]
	list.AppendValue("foo")                                            // ["foo"]
	list.SortValues()                                                  // ["foo"]
	ass.True(t, list.GetSize() == 1)                                   // ["foo"]
	ass.Equal(t, "foo", list.GetValue(1))                              // ["foo"]
	list.AppendValue("bar")                                            // ["foo", "bar"]
	list.SortValues()                                                  // ["bar", "foo"]
	ass.True(t, list.GetSize() == 2)                                   // ["bar", "foo"]
	ass.Equal(t, "bar", list.GetValue(1))                              // ["bar", "foo"]
}

func TestListsWithConcatenate(t *tes.T) {
	var collator = doc.Collator[doc.ListLike[int]]()
	var onetwothree = doc.List[int]([]int{1, 2, 3})
	var fourfivesix = doc.List[int]([]int{4, 5, 6})
	var onethrusix = doc.List[int]([]int{1, 2, 3, 4, 5, 6})
	var list1 = doc.List[int]()
	list1.AppendValues(onetwothree)
	var list2 = doc.List[int]()
	list2.AppendValues(fourfivesix)
	var list3 = doc.ListClass[int]().Concatenate(list1, list2)
	var list4 = doc.List[int]()
	list4.AppendValues(onethrusix)
	ass.True(t, collator.CompareValues(list3, list4))
}

func TestListsWithEmptyLists(t *tes.T) {
	var collator = doc.Collator[doc.ListLike[int]]()
	var empty = doc.List[int]()
	var list = doc.ListClass[int]().Concatenate(empty, empty)
	ass.True(t, collator.CompareValues(empty, empty))
	ass.True(t, collator.CompareValues(list, empty))
	ass.True(t, collator.CompareValues(empty, list))
	ass.True(t, collator.CompareValues(list, list))
}

func TestQueueConstructors(t *tes.T) {
	doc.Queue[int64]()
	doc.Queue[int64](5)
	var sequence = doc.Queue[int64]([]int64{1, 2, 3})
	var queue = doc.Queue[int64](sequence)
	ass.Equal(t, sequence.AsArray(), queue.AsArray())
}

func TestQueueWithConcurrency(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a specific capacity.
	var queue = doc.Queue[int](12)
	ass.True(t, queue.GetCapacity() == 12)
	ass.True(t, queue.IsEmpty())
	ass.True(t, queue.GetSize() == 0)

	// Add some values to the queue.
	for i := 1; i < 10; i++ {
		queue.AddValue(i)
	}
	ass.True(t, queue.GetSize() == 9)

	// Remove values from the queue in the background.
	group.Add(1)
	go func() {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = queue.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
		queue.RemoveAll()
	}()

	// Add some more values to the queue.
	for i := 10; i < 101; i++ {
		queue.AddValue(i)
	}
	queue.CloseChannel()
}

func TestQueueWithFork(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a fan out of two.
	var input = doc.Queue[int](3)
	var outputs = doc.QueueClass[int]().Fork(group, input, 2)

	// Remove values from the output queues in the background.
	var readOutput = func(output doc.QueueLike[int], name string) {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = output.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
	}
	group.Add(2)
	var iterator = outputs.GetIterator()
	for iterator.HasNext() {
		var output = iterator.GetNext()
		go readOutput(output, "output")
	}

	// Add values to the input queue.
	for i := 1; i < 11; i++ {
		input.AddValue(i)
	}
	input.CloseChannel()
}

func TestQueueWithInvalidFanOut(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var input = doc.Queue[int](3)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The fan out size for a queue must be greater than one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	doc.QueueClass[int]().Fork(group, input, 1) // Should panic here.
}

func TestQueueWithSplitAndJoin(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with a split of five outputs and a join back to one.
	var input = doc.Queue[int](3)
	var split = doc.QueueClass[int]().Split(group, input, 5)
	var output = doc.QueueClass[int]().Join(group, split)

	// Remove values from the output queue in the background.
	group.Add(1)
	go func() {
		defer group.Done()
		var value int
		var ok = true
		for i := 1; ok; i++ {
			value, ok = output.RemoveFirst()
			if ok {
				ass.Equal(t, i, value)
			}
		}
	}()

	// Add values to the input queue.
	for i := 1; i < 21; i++ {
		input.AddValue(i)
	}
	input.CloseChannel()
}

func TestQueueWithInvalidSplit(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var input = doc.Queue[int](3)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The size of the split must be greater than one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	doc.QueueClass[int]().Split(group, input, 1) // Should panic here.
}

func TestQueueWithInvalidJoin(t *tes.T) {
	// Create a wait group for synchronization.
	var group = new(syn.WaitGroup)
	defer group.Wait()

	// Create a new queue with an invalid fan out.
	var inputs = doc.List[doc.QueueLike[int]]()
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "The number of input queues for a join must be at least one.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	doc.QueueClass[int]().Join(group, inputs) // Should panic here.
	defer group.Done()
}

func TestSetConstructors(t *tes.T) {
	var collator = doc.Collator[int64]()
	doc.Set[int64]()
	doc.Set[int64](collator)
	var sequence = doc.Set[int64]([]int64{1, 2, 3})
	var set = doc.Set[int64](sequence)
	ass.Equal(t, sequence.AsArray(), set.AsArray())
}

func TestSetsWithStrings(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[string]]()
	doc.List[string]()
	var empty = []string{}
	var bazbar = doc.List[string]([]string{"baz", "bar"})
	var bazfoo = doc.List[string]([]string{"baz", "foo"})
	var baxbaz = doc.List[string]([]string{"bax", "baz"})
	var baxbez = doc.List[string]([]string{"bax", "bez"})
	var barbaz = doc.List[string]([]string{"bar", "baz"})
	var bar = doc.List[string]([]string{"bar"})
	var set = doc.Set[string]()                                   // [ ]
	ass.True(t, set.IsEmpty())                                    // [ ]
	ass.True(t, set.GetSize() == 0)                               // [ ]
	ass.False(t, set.ContainsValue("bax"))                        // [ ]
	ass.Equal(t, empty, set.AsArray())                            // [ ]
	var iterator = set.GetIterator()                              // [ ]
	ass.False(t, iterator.HasNext())                              // [ ]
	ass.False(t, iterator.HasPrevious())                          // [ ]
	iterator.ToStart()                                            // [ ]
	iterator.ToEnd()                                              // [ ]
	set.RemoveAll()                                               // [ ]
	set.RemoveValue("foo")                                        // [ ]
	set.AddValue("foo")                                           // ["foo"]
	ass.False(t, set.IsEmpty())                                   // ["foo"]
	ass.True(t, set.GetSize() == 1)                               // ["foo"]
	ass.Equal(t, "foo", set.GetValue(1))                          // ["foo"]
	ass.True(t, set.GetIndex("baz") == 0)                         // ["foo"]
	ass.True(t, set.ContainsValue("foo"))                         // ["foo"]
	ass.False(t, set.ContainsValue("bax"))                        // ["foo"]
	set.AddValues(bazbar)                                         // ["bar", "baz", "foo"]
	ass.True(t, set.GetSize() == 3)                               // ["bar", "baz", "foo"]
	ass.True(t, set.GetIndex("baz") == 2)                         // ["bar", "baz", "foo"]
	ass.Equal(t, "bar", set.GetValue(1))                          // ["bar", "baz", "foo"]
	ass.Equal(t, bazfoo.AsArray(), set.GetValues(2, 3).AsArray()) // ["bar", "baz", "foo"]
	ass.Equal(t, bar.AsArray(), set.GetValues(1, 1).AsArray())    // ["bar", "baz", "foo"]
	var set2 = doc.Set[string](set)                               // ["bar", "baz", "foo"]
	ass.True(t, collator.CompareValues(set, set2))                // ["bar", "baz", "foo"]
	var array = doc.List[string]([]string{"foo", "bar", "baz"})   // ["bar", "baz", "foo"]
	var set3 = doc.Set[string](array)                             // ["bar", "baz", "foo"]
	ass.True(t, collator.CompareValues(set2, set3))               // ["bar", "baz", "foo"]
	iterator = set.GetIterator()                                  // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasNext())                               // ["bar", "baz", "foo"]
	ass.False(t, iterator.HasPrevious())                          // ["bar", "baz", "foo"]
	ass.Equal(t, "bar", string(iterator.GetNext()))               // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasPrevious())                           // ["bar", "baz", "foo"]
	iterator.ToEnd()                                              // ["bar", "baz", "foo"]
	ass.False(t, iterator.HasNext())                              // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasPrevious())                           // ["bar", "baz", "foo"]
	ass.Equal(t, "foo", string(iterator.GetPrevious()))           // ["bar", "baz", "foo"]
	ass.True(t, iterator.HasNext())                               // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsValue("baz"))                         // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsValue("bax"))                        // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsAny(baxbaz))                          // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsAny(baxbez))                         // ["bar", "baz", "foo"]
	ass.True(t, set.ContainsAll(barbaz))                          // ["bar", "baz", "foo"]
	ass.False(t, set.ContainsAll(baxbaz))                         // ["bar", "baz", "foo"]
	set.RemoveAll()                                               // [ ]
	ass.True(t, set.IsEmpty())                                    // [ ]
	ass.True(t, set.GetSize() == 0)                               // [ ]
}

func TestSetsWithIntegers(t *tes.T) {
	var array = doc.List[int]([]int{3, 1, 4, 5, 9, 2})
	var set = doc.Set[int]()           // [ ]
	set.AddValues(array)               // [1,2,3,4,5,9]
	ass.False(t, set.IsEmpty())        // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(1) == 1)  // [1,2,3,4,5,9]
	ass.True(t, set.GetValue(-1) == 9) // [1,2,3,4,5,9]
	set.RemoveValue(6)                 // [1,2,3,4,5,9]
	ass.True(t, set.GetSize() == 6)    // [1,2,3,4,5,9]
	set.RemoveValue(3)                 // [1,2,4,5,9]
	ass.True(t, set.GetSize() == 5)    // [1,2,4,5,9]
	ass.True(t, set.GetValue(3) == 4)  // [1,2,4,5,9]
}

func TestSetsWithSets(t *tes.T) {
	var array1 = doc.List[int]([]int{3, 1, 4, 5, 9, 2})
	var array2 = doc.List[int]([]int{7, 1, 4, 5, 9, 2})
	var set1 = doc.Set[int]()
	set1.AddValues(array1)
	var set2 = doc.Set[int]()
	set2.AddValues(array2)
	var set = doc.Set[doc.SetLike[int]]()
	set.AddValue(set1)
	set.AddValue(set2)
	ass.False(t, set.IsEmpty())
	ass.True(t, set.GetSize() == 2)
	ass.Equal(t, set1, set.GetValue(1))
	ass.Equal(t, set2, set.GetValue(-1))
	set.RemoveValue(set1)
	ass.True(t, set.GetSize() == 1)
	set.RemoveAll()
	ass.True(t, set.GetSize() == 0)
}

func TestSetsWithAnd(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[int]]()
	var array1 = doc.List[int]([]int{3, 1, 2})
	var array2 = doc.List[int]([]int{3, 2, 4})
	var array3 = doc.List[int]([]int{3, 2})
	var set1 = doc.Set[int]()
	set1.AddValues(array1)
	var set2 = doc.Set[int]()
	set2.AddValues(array2)
	var set3 = doc.SetClass[int]().And(set1, set2)
	var set4 = doc.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithSan(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[int]]()
	var array1 = doc.List[int]([]int{3, 1, 2})
	var array2 = doc.List[int]([]int{3, 2, 4})
	var array3 = doc.List[int]([]int{1})
	var set1 = doc.Set[int]()
	set1.AddValues(array1)
	var set2 = doc.Set[int]()
	set2.AddValues(array2)
	var set3 = doc.SetClass[int]().San(set1, set2)
	var set4 = doc.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithIor(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[int]]()
	var array1 = doc.List[int]([]int{3, 1, 5})
	var array2 = doc.List[int]([]int{6, 2, 4})
	var array3 = doc.List[int]([]int{1, 3, 5, 6, 2, 4})
	var set1 = doc.Set[int]()
	set1.AddValues(array1)
	var set2 = doc.Set[int]()
	set2.AddValues(array2)
	var set3 = doc.SetClass[int]().Ior(set1, set2)
	ass.True(t, set3.ContainsAll(set1))
	ass.True(t, set3.ContainsAll(set2))
	var set4 = doc.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithXor(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[int]]()
	var array1 = doc.List[int]([]int{2, 3, 1, 5})
	var array2 = doc.List[int]([]int{6, 2, 5, 4})
	var array3 = doc.List[int]([]int{1, 3, 4, 6})
	var set1 = doc.Set[int]()
	set1.AddValues(array1)
	var set2 = doc.Set[int]()
	set2.AddValues(array2)
	var set3 = doc.SetClass[int]().Xor(set1, set2)
	var set4 = doc.Set[int]()
	set4.AddValues(array3)
	ass.True(t, collator.CompareValues(set3, set4))
}

func TestSetsWithEmptySets(t *tes.T) {
	var collator = doc.Collator[doc.SetLike[int]]()
	var set1 = doc.Set[int]()
	var set2 = doc.Set[int]()
	var set3 = doc.SetClass[int]().And(set1, set2)
	var set4 = doc.SetClass[int]().San(set1, set2)
	var set5 = doc.SetClass[int]().Ior(set1, set2)
	var set6 = doc.SetClass[int]().Xor(set1, set2)
	ass.True(t, collator.CompareValues(set3, set4))
	ass.True(t, collator.CompareValues(set4, set5))
	ass.True(t, collator.CompareValues(set5, set6))
	ass.True(t, collator.CompareValues(set6, set1))
}

func TestStackConstructors(t *tes.T) {
	doc.Stack[int64]()
	doc.Stack[int64](5)
	var sequence = doc.Stack[int64]([]int64{1, 2, 3})
	var stack = doc.Stack[int64](sequence)
	ass.Equal(t, sequence.AsArray(), stack.AsArray())
}

func TestStackWithSmallCapacity(t *tes.T) {
	var stack = doc.Stack[int](1)
	stack.AddValue(1)
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "Attempted to add a value onto a stack that has reached its capacity.", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	stack.AddValue(2) // This should panic.
}

func TestEmptyStackRemoval(t *tes.T) {
	var stack = doc.Stack[int]()
	defer func() {
		if e := recover(); e != nil {
			ass.Equal(t, "Attempted to remove a value from an empty stack!", e)
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	stack.RemoveLast() // This should panic.
}

func TestStacksWithStrings(t *tes.T) {
	var stack = doc.Stack[string]()
	ass.True(t, stack.IsEmpty())
	ass.True(t, stack.GetSize() == 0)
	stack.RemoveAll()
	stack.AddValue("foo")
	stack.AddValue("bar")
	stack.AddValue("baz")
	ass.True(t, stack.GetSize() == 3)
	var last = stack.GetLast()
	ass.Equal(t, last, stack.RemoveLast())
	ass.True(t, stack.GetSize() == 2)
	ass.Equal(t, "bar", stack.RemoveLast())
	ass.True(t, stack.GetSize() == 1)
	stack.RemoveAll()
}

// ELEMENT

func TestUnits(t *tes.T) {
	ass.Equal(t, "Degrees", doc.Degrees.String())
	ass.Equal(t, "Radians", doc.Radians.String())
	ass.Equal(t, "Gradians", doc.Gradians.String())
}

func TestZeroAngles(t *tes.T) {
	var v = doc.Angle()
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, doc.AngleClass().Zero())

	v = doc.Angle(2.0 * mat.Pi)
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, doc.AngleClass().Zero())

	v = doc.Angle("~0")
	ass.Equal(t, "~0", v.AsSource())
	ass.Equal(t, v, doc.AngleClass().Zero())

	v = doc.Angle("~τ")
	ass.Equal(t, "~τ", v.AsSource())
	ass.Equal(t, v, doc.AngleClass().Tau())
}

func TestPositiveAngles(t *tes.T) {
	var v = doc.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, doc.AngleClass().Pi())

	v = doc.Angle("~π")
	ass.Equal(t, "~π", v.AsSource())
	ass.Equal(t, v, doc.AngleClass().Pi())
}

func TestNegativeAngles(t *tes.T) {
	var v = doc.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())
	ass.Equal(t, v, doc.AngleClass().Pi())

	v = doc.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var class = doc.AngleClass()
	var v0 = class.Zero()
	var v1 = doc.Angle(mat.Pi * 0.25)
	var v2 = doc.Angle(mat.Pi * 0.5)
	var v3 = doc.Angle(mat.Pi * 0.75)
	var v4 = class.Pi()
	var v5 = doc.Angle(mat.Pi * 1.25)
	var v6 = doc.Angle(mat.Pi * 1.5)
	var v7 = doc.Angle(mat.Pi * 1.75)
	var v8 = class.Tau()

	ass.Equal(t, v4, class.Inverse(v0))
	ass.Equal(t, v5, class.Inverse(v1))
	ass.Equal(t, v6, class.Inverse(v2))
	ass.Equal(t, v7, class.Inverse(v3))
	ass.Equal(t, v0, class.Inverse(v4))
	ass.Equal(t, v4, class.Inverse(v8))

	ass.Equal(t, v1, class.Sum(v0, v1))
	ass.Equal(t, v0, class.Difference(v1, v1))
	ass.Equal(t, v3, class.Sum(v1, v2))
	ass.Equal(t, v1, class.Difference(v3, v2))
	ass.Equal(t, v5, class.Sum(v2, v3))
	ass.Equal(t, v2, class.Difference(v5, v3))
	ass.Equal(t, v7, class.Sum(v3, v4))
	ass.Equal(t, v3, class.Difference(v7, v4))
	ass.Equal(t, v1, class.Sum(v8, v1))
	ass.Equal(t, v0, class.Difference(v8, v8))

	ass.Equal(t, v3, class.Scaled(v1, 3.0))
	ass.Equal(t, v0, class.Scaled(v4, 2.0))
	ass.Equal(t, v4, class.Scaled(v4, -1.0))
	ass.Equal(t, v0, class.Scaled(v8, 1.0))

	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v0)))
	ass.Equal(t, v1, class.ArcCosine(class.Cosine(v1)))
	ass.Equal(t, v2, class.ArcCosine(class.Cosine(v2)))
	ass.Equal(t, v3, class.ArcCosine(class.Cosine(v3)))
	ass.Equal(t, v4, class.ArcCosine(class.Cosine(v4)))
	ass.Equal(t, v0, class.ArcCosine(class.Cosine(v8)))

	ass.Equal(t, v0, class.ArcSine(class.Sine(v0)))
	ass.Equal(t, v1, class.ArcSine(class.Sine(v1)))
	ass.Equal(t, v2, class.ArcSine(class.Sine(v2)))
	ass.Equal(t, v6, class.ArcSine(class.Sine(v6)))
	ass.Equal(t, v7, class.ArcSine(class.Sine(v7)))
	ass.Equal(t, v0, class.ArcSine(class.Sine(v8)))

	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v0), class.Sine(v0)))
	ass.Equal(t, v1, class.ArcTangent(class.Cosine(v1), class.Sine(v1)))
	ass.Equal(t, v2, class.ArcTangent(class.Cosine(v2), class.Sine(v2)))
	ass.Equal(t, v3, class.ArcTangent(class.Cosine(v3), class.Sine(v3)))
	ass.Equal(t, v4, class.ArcTangent(class.Cosine(v4), class.Sine(v4)))
	ass.Equal(t, v5, class.ArcTangent(class.Cosine(v5), class.Sine(v5)))
	ass.Equal(t, v0, class.ArcTangent(class.Cosine(v8), class.Sine(v8)))
}

func TestFalseBooleans(t *tes.T) {
	ass.False(t, doc.BooleanClass().False().AsIntrinsic())
	var v = doc.Boolean(false)
	ass.False(t, v.AsIntrinsic())
	v = doc.Boolean("false")
	ass.Equal(t, "false", v.AsSource())
	ass.Equal(t, v, doc.BooleanClass().False())
}

func TestTrueBooleans(t *tes.T) {
	ass.True(t, doc.BooleanClass().True().AsIntrinsic())
	var v = doc.Boolean(true)
	ass.True(t, v.AsIntrinsic())
	v = doc.Boolean("true")
	ass.Equal(t, "true", v.AsSource())
	ass.Equal(t, v, doc.BooleanClass().True())
}

func TestBooleansLibrary(t *tes.T) {
	var T = doc.Boolean(true)
	var F = doc.Boolean(false)
	var class = doc.BooleanClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Ior(class.San(T, T), class.San(T, T))
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Ior(class.San(T, F), class.San(F, T))
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Ior(class.San(F, T), class.San(T, F))
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Ior(class.San(F, F), class.San(F, F))
	ass.Equal(t, xor, class.Xor(F, F))
}

var DurationClass = doc.DurationClass()

func TestZeroDurations(t *tes.T) {
	var v = doc.Duration(0)
	ass.True(t, v.AsInteger() == 0)
	ass.True(t, v.AsIntrinsic() == 0)
	ass.True(t, v.AsMilliseconds() == 0)
	ass.True(t, v.AsSeconds() == 0)
	ass.True(t, v.AsMinutes() == 0)
	ass.True(t, v.AsHours() == 0)
	ass.True(t, v.AsDays() == 0)
	ass.True(t, v.AsWeeks() == 0)
	ass.True(t, v.AsMonths() == 0)
	ass.True(t, v.AsYears() == 0)
	ass.True(t, v.GetMilliseconds() == 0)
	ass.True(t, v.GetSeconds() == 0)
	ass.True(t, v.GetMinutes() == 0)
	ass.True(t, v.GetHours() == 0)
	ass.True(t, v.GetDays() == 0)
	ass.True(t, v.GetWeeks() == 0)
	ass.True(t, v.GetMonths() == 0)
	ass.True(t, v.GetYears() == 0)
}

func TestStringDurations(t *tes.T) {
	var duration = doc.Duration("~P1Y2M3DT4H5M6S")
	ass.Equal(t, "~P1Y2M3DT4H5M6S", duration.AsSource())
	duration = doc.Duration("~P0W")
	ass.Equal(t, "~P0W", duration.AsSource())
}

func TestDurations(t *tes.T) {
	var v = doc.Duration(60000)
	ass.True(t, v.AsSource() == "~PT1M")
	ass.True(t, v.AsInteger() == 60000)
	ass.True(t, v.AsIntrinsic() == 60000)
	ass.True(t, v.AsMilliseconds() == 60000)
	ass.True(t, v.AsSeconds() == 60)
	ass.True(t, v.AsMinutes() == 1)
	ass.True(t, v.AsHours() == 0.016666666666666666)
	ass.True(t, v.AsDays() == 0.0006944444444444445)
	ass.True(t, v.AsWeeks() == 9.92063492063492e-05)
	ass.True(t, v.AsMonths() == 2.2815891724904232e-05)
	ass.True(t, v.AsYears() == 1.9013243104086858e-06)
	ass.True(t, v.GetMilliseconds() == 0)
	ass.True(t, v.GetSeconds() == 0)
	ass.True(t, v.GetMinutes() == 1)
	ass.True(t, v.GetHours() == 0)
	ass.True(t, v.GetDays() == 0)
	ass.True(t, v.GetWeeks() == 0)
	ass.True(t, v.GetMonths() == 0)
	ass.True(t, v.GetYears() == 0)
}

var GlyphClass = doc.GlyphClass()

func TestGlyphs(t *tes.T) {
	var v = doc.Glyph("'''")
	ass.Equal(t, "'''", v.AsSource())

	v = doc.Glyph('a')
	ass.Equal(t, "'a'", v.AsSource())

	v = doc.Glyph('"')
	ass.Equal(t, `'"'`, v.AsSource())

	v = doc.Glyph('😊')
	ass.Equal(t, "'😊'", v.AsSource())

	v = doc.Glyph('界')
	ass.Equal(t, "'界'", v.AsSource())

	v = doc.Glyph('\'')
	ass.Equal(t, "'''", v.AsSource())

	v = doc.Glyph('\\')
	ass.Equal(t, "'\\'", v.AsSource())

	v = doc.Glyph('\n')
	ass.Equal(t, "'\n'", v.AsSource())

	v = doc.Glyph('\t')
	ass.Equal(t, "'\t'", v.AsSource())
}

var MomentClass = doc.MomentClass()

func TestIntegerMoments(t *tes.T) {
	var v = doc.Moment(1238589296789)
	ass.Equal(t, 1238589296789, v.AsIntrinsic())
	ass.Equal(t, 1238589296789, v.AsInteger())
	ass.Equal(t, 1238589296789.0, v.AsMilliseconds())
	ass.Equal(t, 1238589296.789, v.AsSeconds())
	ass.Equal(t, 20643154.946483333, v.AsMinutes())
	ass.Equal(t, 344052.58244138886, v.AsHours())
	ass.Equal(t, 14335.524268391204, v.AsDays())
	ass.Equal(t, 2047.9320383416004, v.AsWeeks())
	ass.Equal(t, 470.9919881193849, v.AsMonths())
	ass.Equal(t, 39.24933234328208, v.AsYears())
	ass.True(t, v.GetMilliseconds() == 789)
	ass.True(t, v.GetSeconds() == 56)
	ass.True(t, v.GetMinutes() == 34)
	ass.True(t, v.GetHours() == 12)
	ass.True(t, v.GetDays() == 1)
	ass.True(t, v.GetWeeks() == 14)
	ass.True(t, v.GetMonths() == 4)
	ass.True(t, v.GetYears() == 2009)
}

func TestStringMoments(t *tes.T) {
	var v = doc.Moment("<-1-02-03T04:05:06.700>")
	ass.Equal(t, "<-1-02-03T04:05:06.700>", v.AsSource())
}

func TestMomentsLibrary(t *tes.T) {
	var now = doc.Moment()
	var duration = doc.Duration(12345)
	var after = doc.Moment(now.AsInteger() + duration.AsInteger())
	var class = doc.MomentClass()

	ass.Equal(t, duration, class.Duration(now, after))
	ass.Equal(t, duration, class.Duration(after, now))
	ass.Equal(t, after, class.Later(now, duration))
	ass.Equal(t, now, class.Earlier(after, duration))
}

func TestZero(t *tes.T) {
	var v = doc.Number(0 + 0i)
	ass.Equal(t, 0+0i, v.AsIntrinsic())
	ass.True(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "0", v.AsSource())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, v, doc.NumberClass().Zero())
}

func TestInfinity(t *tes.T) {
	var v = doc.Number(cmp.Inf())
	ass.Equal(t, cmp.Inf(), v.AsIntrinsic())
	ass.False(t, v.IsZero())
	ass.True(t, v.IsInfinite())
	ass.True(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.Equal(t, "∞", v.AsSource())
	ass.Equal(t, mat.Inf(1), v.AsFloat())
	ass.Equal(t, mat.Inf(1), v.GetReal())
	ass.Equal(t, mat.Inf(1), v.GetImaginary())
	ass.Equal(t, v, doc.NumberClass().Infinity())
}

func TestUndefined(t *tes.T) {
	var v = doc.Number(cmp.NaN())
	ass.True(t, cmp.IsNaN(v.AsIntrinsic()))
	ass.False(t, v.IsZero())
	ass.False(t, v.IsInfinite())
	ass.False(t, v.IsDefined())
	ass.False(t, v.IsNegative())
	ass.True(t, mat.IsNaN(v.AsFloat()))
	ass.True(t, mat.IsNaN(v.GetReal()))
	ass.True(t, mat.IsNaN(v.GetImaginary()))
}

func TestPositivePureReals(t *tes.T) {
	var v = doc.Number(0.25)
	ass.Equal(t, 0.25+0i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.25, v.AsFloat())
	ass.Equal(t, 0.25, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	var integer = 5
	v = doc.Number(integer)
	ass.Equal(t, 5.0, v.AsFloat())
	var float = 5.0
	v = doc.Number(float)
	ass.Equal(t, 5.0, v.AsFloat())
	v = doc.Number("1.23456789E+100")
	ass.Equal(t, "1.23456789E+100", v.AsSource())
	v = doc.Number("1.23456789E-10")
	ass.Equal(t, "1.23456789E-10", v.AsSource())
}

func TestPositivePureImaginaries(t *tes.T) {
	var v = doc.Number(0.25i)
	ass.Equal(t, 0+0.25i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 0.25, v.GetImaginary())
}

func TestNegativePureReals(t *tes.T) {
	var v = doc.Number(-0.75)
	ass.Equal(t, -0.75+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -0.75, v.AsFloat())
	ass.Equal(t, -0.75, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
}

func TestNegativePureImaginaries(t *tes.T) {
	var v = doc.Number(-0.75i)
	ass.Equal(t, 0-0.75i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -0.75, v.GetImaginary())
}

func TestNumber(t *tes.T) {
	var v = doc.Polar(1.0, mat.Pi)
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetAngle())

	v = doc.Rectangular(3.0, 4.0)
	ass.Equal(t, 3.0+4.0i, v.AsIntrinsic())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 3.0, v.GetReal())
	ass.Equal(t, 4.0, v.GetImaginary())
	ass.Equal(t, 5.0, v.GetMagnitude())

	v = doc.Number("5e^~1i")
	ass.Equal(t, 5.0, v.GetMagnitude())
	ass.Equal(t, 1.0, v.GetAngle())
	ass.Equal(t, "5e^~1i", v.AsPolar())

	v = doc.Number("1e^~πi")
	ass.Equal(t, -1.0+0i, v.AsIntrinsic())
	ass.True(t, v.IsNegative())
	ass.Equal(t, "-1", v.AsSource())
	ass.Equal(t, "1e^~πi", v.AsPolar())
	ass.Equal(t, -1.0, v.AsFloat())
	ass.Equal(t, -1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi, v.GetAngle())

	v = doc.Number("-1.2-3.4i")
	ass.Equal(t, "-1.2-3.4i", v.AsSource())
	ass.Equal(t, -1.2, v.GetReal())
	ass.Equal(t, -3.4, v.GetImaginary())

	v = doc.Number("-π+τi")
	ass.Equal(t, "-π+τi", v.AsSource())
	ass.Equal(t, -3.141592653589793, v.GetReal())
	ass.Equal(t, 6.283185307179586, v.GetImaginary())

	v = doc.Number("undefined")
	ass.Equal(t, "undefined", v.AsSource())
	ass.False(t, v.IsDefined())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("+infinity")
	ass.Equal(t, "+∞", v.AsSource())
	ass.True(t, v.IsMaximum())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("infinity")
	ass.Equal(t, "∞", v.AsSource())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("-infinity")
	ass.Equal(t, "-∞", v.AsSource())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("∞")
	ass.Equal(t, "∞", v.AsSource())
	ass.True(t, v.IsInfinite())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("-∞")
	ass.Equal(t, "-∞", v.AsSource())
	ass.True(t, v.IsMinimum())
	ass.False(t, v.HasMagnitude())

	v = doc.Number("+1")
	ass.Equal(t, "1", v.AsSource())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = doc.Number("1")
	ass.Equal(t, "1", v.AsSource())
	ass.Equal(t, 1.0, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, 0.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = doc.Number("-π")
	ass.Equal(t, "-π", v.AsSource())
	ass.Equal(t, -mat.Pi, v.GetReal())
	ass.Equal(t, mat.Pi, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.True(t, v.IsNegative())

	v = doc.Number("+1i")
	ass.Equal(t, "1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = doc.Number("1i")
	ass.Equal(t, "1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, 1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = doc.Number("-1i")
	ass.Equal(t, "-1i", v.AsSource())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.0, v.GetImaginary())
	ass.Equal(t, 1.0, v.GetMagnitude())
	ass.Equal(t, -mat.Pi/2.0, v.GetAngle())
	ass.True(t, v.HasMagnitude())
	ass.False(t, v.IsNegative())

	v = doc.Number("-1.2345678E+90")
	ass.Equal(t, "-1.2345678E+90", v.AsSource())
	ass.True(t, v.IsNegative())
	ass.Equal(t, -1.2345678e+90, v.GetReal())
	ass.Equal(t, 0.0, v.GetImaginary())

	v = doc.Number("-1.2345678E+90i")
	ass.Equal(t, "-1.2345678E+90i", v.AsSource())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 0.0, v.GetReal())
	ass.Equal(t, -1.2345678e+90, v.GetImaginary())

	v = doc.Number("1.2345678E+90e^~1.2345678E-90i")
	ass.Equal(t, "1.2345678E+90e^~1.2345678E-90i", v.AsPolar())
	ass.False(t, v.IsNegative())
	ass.Equal(t, 1.2345678e+90, v.GetMagnitude())
	ass.Equal(t, 1.2345678e-90, v.GetAngle())
}

func TestNumberLibrary(t *tes.T) {
	var class = doc.NumberClass()
	var zero = class.Zero()
	var i = class.I()
	var minusi = doc.Number(-1i)
	var half = doc.Number(0.5)
	var minushalf = doc.Number(-0.5)
	var one = class.One()
	var minusone = doc.Number(-1)
	var two = doc.Number(2.0)
	var minustwo = doc.Number(-2.0)
	var infinity = class.Infinity()
	var undefined = class.Undefined()

	//	-z
	ass.Equal(t, zero, class.Inverse(zero))
	ass.Equal(t, minushalf, class.Inverse(half))
	ass.Equal(t, minusone, class.Inverse(one))
	ass.Equal(t, minusi, class.Inverse(i))
	ass.Equal(t, infinity, class.Inverse(infinity))
	ass.False(t, class.Inverse(undefined).IsDefined())

	//	z + zero => z
	ass.Equal(t, minusi, class.Sum(minusi, zero))
	ass.Equal(t, minusone, class.Sum(minusone, zero))
	ass.Equal(t, zero, class.Sum(zero, zero))
	ass.Equal(t, one, class.Sum(one, zero))
	ass.Equal(t, i, class.Sum(i, zero))
	ass.Equal(t, infinity, class.Sum(infinity, zero))
	ass.False(t, class.Sum(undefined, zero).IsDefined())

	//	z + infinity => infinity
	ass.Equal(t, infinity, class.Sum(minusi, infinity))
	ass.Equal(t, infinity, class.Sum(minusone, infinity))
	ass.Equal(t, infinity, class.Sum(zero, infinity))
	ass.Equal(t, infinity, class.Sum(one, infinity))
	ass.Equal(t, infinity, class.Sum(i, infinity))
	ass.Equal(t, infinity, class.Sum(infinity, infinity))
	ass.False(t, class.Sum(undefined, infinity).IsDefined())

	//	z - infinity => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(minusi, infinity))
	ass.Equal(t, infinity, class.Difference(minusone, infinity))
	ass.Equal(t, infinity, class.Difference(zero, infinity))
	ass.Equal(t, infinity, class.Difference(one, infinity))
	ass.Equal(t, infinity, class.Difference(i, infinity))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, infinity).IsDefined())

	//	infinity - z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Difference(infinity, minusi))
	ass.Equal(t, infinity, class.Difference(infinity, minusone))
	ass.Equal(t, infinity, class.Difference(infinity, zero))
	ass.Equal(t, infinity, class.Difference(infinity, one))
	ass.Equal(t, infinity, class.Difference(infinity, i))
	ass.False(t, class.Difference(infinity, undefined).IsDefined())

	//	z - z => zero  {z != infinity}
	ass.Equal(t, zero, class.Difference(minusi, minusi))
	ass.Equal(t, zero, class.Difference(minusone, minusone))
	ass.Equal(t, zero, class.Difference(zero, zero))
	ass.Equal(t, zero, class.Difference(one, one))
	ass.Equal(t, zero, class.Difference(i, i))
	ass.False(t, class.Difference(infinity, infinity).IsDefined())
	ass.False(t, class.Difference(undefined, undefined).IsDefined())

	//	z * r
	ass.Equal(t, minusi, class.Scaled(minusi, 1.0))
	ass.Equal(t, minushalf, class.Scaled(minusone, 0.5))
	ass.Equal(t, zero, class.Scaled(zero, 5.0))
	ass.Equal(t, half, class.Scaled(one, 0.5))
	ass.Equal(t, i, class.Scaled(i, 1.0))
	ass.Equal(t, infinity, class.Scaled(infinity, 5.0))
	ass.False(t, class.Scaled(undefined, 5.0).IsDefined())

	//	/z
	ass.Equal(t, infinity, class.Reciprocal(zero))
	ass.Equal(t, two, class.Reciprocal(half))
	ass.Equal(t, one, class.Reciprocal(one))
	ass.Equal(t, minushalf, class.Reciprocal(minustwo))
	ass.Equal(t, minusi, class.Reciprocal(i))
	ass.Equal(t, zero, class.Reciprocal(infinity))
	ass.False(t, class.Reciprocal(undefined).IsDefined())

	//	*z
	ass.Equal(t, zero, class.Conjugate(zero))
	ass.Equal(t, one, class.Conjugate(one))
	ass.Equal(t, minusi, class.Conjugate(i))
	ass.Equal(t, i, class.Conjugate(minusi))
	ass.False(t, class.Conjugate(undefined).IsDefined())

	//	z * zero => zero          {z != infinity}
	ass.Equal(t, zero, class.Product(zero, zero))
	ass.Equal(t, zero, class.Product(one, zero))
	ass.Equal(t, zero, class.Product(i, zero))
	ass.False(t, class.Product(infinity, zero).IsDefined())
	ass.False(t, class.Product(undefined, zero).IsDefined())

	//	z * one => z
	ass.Equal(t, zero, class.Product(zero, one))
	ass.Equal(t, one, class.Product(one, one))
	ass.Equal(t, i, class.Product(i, one))
	ass.Equal(t, infinity, class.Product(infinity, one))
	ass.False(t, class.Product(undefined, one).IsDefined())

	//	z * infinity => infinity  {z != zero}
	ass.False(t, class.Product(zero, infinity).IsDefined())
	ass.Equal(t, infinity, class.Product(one, infinity))
	ass.Equal(t, infinity, class.Product(i, infinity))
	ass.Equal(t, infinity, class.Product(infinity, infinity))

	//	zero / z => zero          {z != zero}
	ass.False(t, class.Quotient(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Quotient(zero, one))
	ass.Equal(t, zero, class.Quotient(zero, i))
	ass.Equal(t, zero, class.Quotient(zero, infinity))
	ass.False(t, class.Quotient(zero, undefined).IsDefined())

	//	z / zero => infinity      {z != zero}
	ass.Equal(t, infinity, class.Quotient(one, zero))
	ass.Equal(t, infinity, class.Quotient(i, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.False(t, class.Quotient(undefined, zero).IsDefined())

	//	z / infinity => zero      {z != infinity}
	ass.Equal(t, zero, class.Quotient(one, infinity))
	ass.Equal(t, zero, class.Quotient(i, infinity))
	ass.False(t, class.Quotient(infinity, infinity).IsDefined())
	ass.False(t, class.Quotient(undefined, infinity).IsDefined())

	//	infinity / z => infinity  {z != infinity}
	ass.Equal(t, infinity, class.Quotient(infinity, zero))
	ass.Equal(t, infinity, class.Quotient(infinity, one))
	ass.Equal(t, infinity, class.Quotient(infinity, i))
	ass.False(t, class.Quotient(infinity, undefined).IsDefined())

	//	y / z
	ass.Equal(t, one, class.Quotient(one, one))
	ass.Equal(t, one, class.Quotient(i, i))
	ass.Equal(t, i, class.Quotient(i, one))
	ass.Equal(t, two, class.Quotient(one, half))
	ass.Equal(t, one, class.Quotient(half, half))

	//	z ^ zero => one           {by definition}
	ass.Equal(t, one, class.Power(minusi, zero))
	ass.Equal(t, one, class.Power(minusone, zero))
	ass.Equal(t, one, class.Power(zero, zero))
	ass.Equal(t, one, class.Power(one, zero))
	ass.Equal(t, one, class.Power(i, zero))
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.False(t, class.Power(undefined, zero).IsDefined())

	//	zero ^ z => zero          {z != zero}
	ass.Equal(t, zero, class.Power(zero, one))
	ass.Equal(t, zero, class.Power(zero, i))
	ass.Equal(t, zero, class.Power(zero, infinity))
	ass.False(t, class.Power(zero, undefined).IsDefined())

	//	z ^ infinity => zero      {|z| < one}
	//	z ^ infinity => one       {|z| = one}
	//	z ^ infinity => infinity  {|z| > one}
	ass.Equal(t, infinity, class.Power(minustwo, infinity))
	ass.Equal(t, one, class.Power(minusi, infinity))
	ass.Equal(t, one, class.Power(minusone, infinity))
	ass.Equal(t, zero, class.Power(minushalf, infinity))
	ass.Equal(t, zero, class.Power(half, infinity))
	ass.Equal(t, one, class.Power(one, infinity))
	ass.Equal(t, one, class.Power(i, infinity))
	ass.Equal(t, infinity, class.Power(two, infinity))

	//	infinity ^ z => infinity  {z != zero}
	ass.Equal(t, one, class.Power(infinity, zero))
	ass.Equal(t, infinity, class.Power(infinity, one))
	ass.Equal(t, infinity, class.Power(infinity, i))
	ass.Equal(t, infinity, class.Power(infinity, infinity))
	ass.False(t, class.Power(infinity, undefined).IsDefined())

	//	one ^ z => one
	ass.Equal(t, one, class.Power(one, one))
	ass.Equal(t, one, class.Power(one, i))
	ass.Equal(t, one, class.Power(one, minusone))
	ass.Equal(t, one, class.Power(one, minusi))

	//	log(zero, z) => zero
	ass.False(t, class.Logarithm(zero, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(zero, i))
	ass.Equal(t, zero, class.Logarithm(zero, one))
	ass.False(t, class.Logarithm(zero, infinity).IsDefined())
	ass.False(t, class.Logarithm(zero, undefined).IsDefined())

	//	log(one, z) => infinity
	ass.Equal(t, infinity, class.Logarithm(one, zero))
	ass.False(t, class.Logarithm(one, one).IsDefined())
	ass.Equal(t, infinity, class.Logarithm(one, infinity))
	ass.False(t, class.Logarithm(one, undefined).IsDefined())

	//	log(infinity, z) => zero
	ass.False(t, class.Logarithm(infinity, zero).IsDefined())
	ass.Equal(t, zero, class.Logarithm(infinity, one))
	ass.False(t, class.Logarithm(infinity, infinity).IsDefined())
	ass.False(t, class.Logarithm(infinity, undefined).IsDefined())
}

func TestZeroPercentages(t *tes.T) {
	var v = doc.Percentage(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestPositivePercentages(t *tes.T) {
	var v = doc.Percentage(25)
	ass.Equal(t, 0.25, v.AsIntrinsic())
	ass.Equal(t, 25.0, v.AsFloat())
}

func TestNegativePercentages(t *tes.T) {
	var v = doc.Percentage(-75)
	ass.Equal(t, -0.75, v.AsIntrinsic())
	ass.Equal(t, -75.0, v.AsFloat())
}

func TestStringPercentages(t *tes.T) {
	var v = doc.Percentage("1.7%")
	//ass.Equal(t, -1.0, v.AsIntrinsic())
	//ass.Equal(t, -100.0, v.AsFloat())
	ass.Equal(t, "1.7%", v.AsSource())
}

func TestBooleanProbabilities(t *tes.T) {
	var v1 = doc.Probability(false)
	ass.Equal(t, 0.0, v1.AsFloat())

	var v2 = doc.Probability(true)
	ass.Equal(t, 1.0, v2.AsFloat())
}

func TestZeroProbabilities(t *tes.T) {
	var v = doc.Probability(0.0)
	ass.Equal(t, 0.0, v.AsFloat())
}

func TestOneProbabilities(t *tes.T) {
	var v = doc.Probability(1.0)
	ass.Equal(t, 1.0, v.AsFloat())
}

func TestRandomProbability(t *tes.T) {
	doc.Probability()
}

func TestStringProbabilities(t *tes.T) {
	var v = doc.Probability("p0")
	ass.Equal(t, 0.0, v.AsIntrinsic())
	ass.Equal(t, 0.0, v.AsFloat())
	ass.Equal(t, "p0", v.AsSource())

	v = doc.Probability("p0.5")
	ass.Equal(t, 0.5, v.AsIntrinsic())
	ass.Equal(t, 0.5, v.AsFloat())
	ass.Equal(t, "p0.5", v.AsSource())

	v = doc.Probability("p1")
	ass.Equal(t, 1.0, v.AsIntrinsic())
	ass.Equal(t, 1.0, v.AsFloat())
	ass.Equal(t, "p1", v.AsSource())
}

func TestOtherProbabilities(t *tes.T) {
	var v1 = doc.Probability(0.25)
	ass.Equal(t, 0.25, v1.AsFloat())

	var v2 = doc.Probability(0.5)
	ass.Equal(t, 0.5, v2.AsFloat())

	var v3 = doc.Probability(0.75)
	ass.Equal(t, 0.75, v3.AsFloat())
}

func TestProbabilitieLibrary(t *tes.T) {
	var T = doc.Probability(0.75)
	var F = doc.Probability(0.25)
	var class = doc.ProbabilityClass()

	var andNot = class.And(class.Not(T), class.Not(T))
	var notIor = class.Not(class.Ior(T, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(T), class.Not(F))
	notIor = class.Not(class.Ior(T, F))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(T))
	notIor = class.Not(class.Ior(F, T))
	ass.Equal(t, andNot, notIor)

	andNot = class.And(class.Not(F), class.Not(F))
	notIor = class.Not(class.Ior(F, F))
	ass.Equal(t, andNot, notIor)

	var san = class.And(T, class.Not(T))
	ass.Equal(t, san, class.San(T, T))

	san = class.And(T, class.Not(F))
	ass.Equal(t, san, class.San(T, F))

	san = class.And(F, class.Not(T))
	ass.Equal(t, san, class.San(F, T))

	san = class.And(F, class.Not(F))
	ass.Equal(t, san, class.San(F, F))

	var xor = class.Probability(class.San(T, T).AsFloat() + class.San(T, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, T))

	xor = class.Probability(class.San(T, F).AsFloat() + class.San(F, T).AsFloat())
	ass.Equal(t, xor, class.Xor(T, F))

	xor = class.Probability(class.San(F, T).AsFloat() + class.San(T, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, T))

	xor = class.Probability(class.San(F, F).AsFloat() + class.San(F, F).AsFloat())
	ass.Equal(t, xor, class.Xor(F, F))
}

func TestResource(t *tes.T) {
	var v = doc.Resource("https://craterdog.com/About.html")
	ass.Equal(t, "https://craterdog.com/About.html", v.AsIntrinsic())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = doc.Resource("<https://craterdog.com/About.html>")
	ass.Equal(t, "<https://craterdog.com/About.html>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = doc.Resource("<mailto:craterdog@google.com>")
	ass.Equal(t, "<mailto:craterdog@google.com>", v.AsSource())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = doc.Resource("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = doc.Resource("<https://craterdog.com/#Home>")
	ass.Equal(t, "<https://craterdog.com/#Home>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = doc.Resource("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", v.AsSource())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestSymbol(t *tes.T) {
	var foobar = "foo-bar"
	var v = doc.Symbol(foobar)
	ass.Equal(t, []rune{'f', 'o', 'o', '-', 'b', 'a', 'r'}, v.AsIntrinsic())

	foobar = "$foo-bar"
	v = doc.Symbol(foobar)
	ass.Equal(t, foobar, v.AsSource())
}

// STRING

func TestEmptyBinary(t *tes.T) {
	var binary = `'><'`
	var v = doc.Binary(binary)
	ass.Equal(t, binary, v.AsSource())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestBinary(t *tes.T) {
	var b1 = `'>
    abcd1234
<'`
	var v = doc.Binary(b1)
	ass.Equal(t, b1, v.AsSource())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 6, int(v.GetSize()))
	ass.Equal(t, v.AsArray(), doc.Binary(v.AsArray()).AsArray())
}

func TestBinaryLibrary(t *tes.T) {
	var b1 = `'>
    abcd
<'`
	var b2 = `'>
    12345678
<'`
	var b3 = `'>
    abcd12345678
<'`
	var v1 = doc.Binary(b1)
	var v2 = doc.Binary(b2)
	var class = doc.BinaryClass()
	ass.Equal(t, b3, class.Concatenate(v1, v2).AsSource())

	v1 = doc.Binary([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	v2 = doc.Binary([]byte{0x03, 0x00, 0x01, 0x02})
	var not = doc.Binary([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb})
	var and = doc.Binary([]byte{0x00, 0x00, 0x00, 0x02, 0x00})
	var sans = doc.Binary([]byte{0x00, 0x01, 0x02, 0x01, 0x04})
	var or = doc.Binary([]byte{0x03, 0x01, 0x03, 0x03, 0x04})
	var xor = doc.Binary([]byte{0x03, 0x01, 0x03, 0x01, 0x04})
	var sans2 = doc.Binary([]byte{0x03, 0x00, 0x01, 0x00, 0x00})

	ass.Equal(t, not, class.Not(v1))
	ass.Equal(t, and, class.And(v1, v2))
	ass.Equal(t, sans, class.San(v1, v2))
	ass.Equal(t, or, class.Ior(v1, v2))
	ass.Equal(t, xor, class.Xor(v1, v2))
	ass.Equal(t, sans2, class.San(v2, v1))
}

func TestName(t *tes.T) {
	var v1 = doc.Name("/bali-nebula/types/abstractions/5String")
	ass.Equal(t, "/bali-nebula/types/abstractions/5String", v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 4, int(v1.GetSize()))
	ass.Equal(t, doc.Folder("bali-nebula"), v1.GetValue(1))
	ass.Equal(t, doc.Folder("5String"), v1.GetValue(-1))
	var v2 = doc.Name(v1.AsArray())
	ass.Equal(t, v1.AsSource(), v2.AsSource())
	var v3 = doc.Name(v1.GetValues(1, 2))
	ass.Equal(t, 1, int(v1.GetIndex("bali-nebula")))
	ass.Equal(t, "/bali-nebula/types", v3.AsSource())
}

func TestNamesLibrary(t *tes.T) {
	var v1 = doc.Name("/bali-nebula/types/abstractions")
	var v2 = doc.Name("/String")
	ass.Equal(
		t,
		"/bali-nebula/types/abstractions/String",
		doc.NameClass().Concatenate(v1, v2).AsSource(),
	)
}

const n0 = `"><"`

const n1 = `">
    abcd本
<"`

const n2 = `">
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

const n3 = `">
    abcd本
    1234
	\">
        This is an embedded narrative.
    <\"
<"`

func TestEmptyNarrative(t *tes.T) {
	var v0 = doc.Narrative(n0)
	ass.Equal(t, n0, v0.AsSource())
	ass.True(t, v0.IsEmpty())
	ass.Equal(t, 0, int(v0.GetSize()))
	ass.Equal(t, 0, len(v0.AsArray()))
}

func TestNarrative(t *tes.T) {
	var v1 = doc.Narrative(n1)
	ass.Equal(t, n1, v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 1, int(v1.GetSize()))

	var v3 = doc.Narrative(n3)
	ass.Equal(t, n3, v3.AsSource())
	ass.False(t, v3.IsEmpty())
	ass.Equal(t, 5, int(v3.GetSize()))

	ass.Equal(t, n3, doc.Narrative(v3.AsArray()).AsSource())
	ass.Equal(t, 5, len(v3.AsArray()))
}

func TestNarrativesLibrary(t *tes.T) {
	var v1 = doc.Narrative(n1)
	var v2 = doc.Narrative(n2)
	var v3 = doc.NarrativeClass().Concatenate(v1, v2)
	ass.Equal(t, v1.GetValue(1), v3.GetValue(1))
	ass.Equal(t, v2.GetValue(-1), v3.GetValue(-1))
	ass.Equal(t, n3, v3.AsSource())
}

func TestNonePattern(t *tes.T) {
	var v = doc.PatternClass().None()
	ass.Equal(t, `none`, v.AsSource())

	v = doc.Pattern(`none`)
	ass.Equal(t, `none`, v.AsSource())
	ass.Equal(t, v, doc.PatternClass().None())

	var text = ""
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "anything at all..."
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestAnyPattern(t *tes.T) {
	var v = doc.PatternClass().Any()
	ass.Equal(t, `any`, v.AsSource())

	v = doc.Pattern(`any`)
	ass.Equal(t, `any`, v.AsSource())
	ass.Equal(t, v, doc.PatternClass().Any())

	var text = ""
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "anything at all..."
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))

	text = "none"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text}, v.GetMatches(text))
}

func TestSomePattern(t *tes.T) {
	var v = doc.Pattern(`"c(.+t)"?`)
	ass.Equal(t, `"c(.+t)"?`, v.AsSource())

	var text = "ct"
	ass.False(t, v.MatchesText(text))
	ass.Equal(t, []string(nil), v.GetMatches(text))

	text = "cat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "caaat"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))

	text = "cot"
	ass.True(t, v.MatchesText(text))
	ass.Equal(t, []string{text, text[1:]}, v.GetMatches(text))
}

func TestEmptyQuote(t *tes.T) {
	var v = doc.Quote([]rune{})
	ass.Equal(t, []rune{}, v.AsIntrinsic())
	ass.True(t, v.IsEmpty())
	ass.Equal(t, 0, int(v.GetSize()))
}

func TestQuote(t *tes.T) {
	var v = doc.Quote(`"abcd本1234"`)
	ass.Equal(t, `"abcd本1234"`, v.AsSource())
	ass.False(t, v.IsEmpty())
	ass.Equal(t, 9, int(v.GetSize()))
	ass.Equal(t, 'a', rune(v.GetValue(1)))
	ass.Equal(t, '4', rune(v.GetValue(-1)))
	ass.Equal(t, `"d本1"`, doc.Quote(v.GetValues(4, 6)).AsSource())
	ass.Equal(t, 8, int(v.GetIndex('3')))
}

func TestQuotesLibrary(t *tes.T) {
	var v1 = doc.Quote(`"abcd本"`)
	var v2 = doc.Quote(`"1234"`)
	ass.Equal(t, `"abcd本1234"`, doc.QuoteClass().Concatenate(v1, v2).AsSource())
}

func TestStringTags(t *tes.T) {
	for size := 8; size < 33; size++ {
		var t1 = doc.Tag(size)
		ass.Equal(t, len(t1.AsSource()), 1+int(mat.Ceil(float64(size)*8.0/5.0)))
		var s1 = t1.AsSource()
		var t2 = doc.Tag(s1)
		ass.Equal(t, t1, t2)
		var s2 = t2.AsSource()
		ass.Equal(t, s1, s2)
		ass.Equal(t, t1.AsArray(), t2.AsArray())
	}
}

func TestVersion(t *tes.T) {
	var v1 = doc.Version("v1.2.3")
	ass.Equal(t, "v1.2.3", v1.AsSource())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, int(v1.GetSize()))
	ass.Equal(t, uint(1), v1.GetValue(1))
	ass.Equal(t, uint(3), v1.GetValue(-1))
	var v3 = doc.Version(v1.GetValues(1, 2))
	ass.Equal(t, 2, int(v1.GetIndex(2)))
	ass.Equal(t, "v1.2", v3.AsSource())
}

func TestVersionsLibrary(t *tes.T) {
	var v1 = doc.Version([]uint{1})
	var v2 = doc.Version([]uint{2, 3})
	var class = doc.VersionClass()

	var v3 = class.Concatenate(v1, v2)
	ass.Equal(t, []uint{1, 2, 3}, v3.AsIntrinsic())
	ass.False(t, class.IsValidNextVersion(v1, v1))
	ass.Equal(t, "v2", class.GetNextVersion(v1, 0).AsSource())
	ass.Equal(t, "v2", class.GetNextVersion(v1, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v1, class.GetNextVersion(v1, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "v1.1", class.GetNextVersion(v1, 3).AsSource())

	ass.False(t, class.IsValidNextVersion(v2, v2))
	ass.Equal(t, "v3", class.GetNextVersion(v2, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 0).AsSource())
	ass.Equal(t, "v2.4", class.GetNextVersion(v2, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "v2.3.1", class.GetNextVersion(v2, 3).AsSource())
	ass.True(t, class.IsValidNextVersion(v2, class.GetNextVersion(v2, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v2, 3), v2))

	ass.False(t, class.IsValidNextVersion(v3, v3))
	ass.Equal(t, "v2", class.GetNextVersion(v3, 1).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 1)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "v1.3", class.GetNextVersion(v3, 2).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 2)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 0).AsSource())
	ass.Equal(t, "v1.2.4", class.GetNextVersion(v3, 3).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 3)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "v1.2.3.1", class.GetNextVersion(v3, 4).AsSource())
	ass.True(t, class.IsValidNextVersion(v3, class.GetNextVersion(v3, 4)))
	ass.False(t, class.IsValidNextVersion(class.GetNextVersion(v3, 4), v3))
}

func TestIntervalConstructors(t *tes.T) {
	var glyphs = doc.Interval[doc.GlyphLike](
		doc.Inclusive,
		doc.Glyph(65),
		doc.Glyph(70),
		doc.Inclusive,
	)
	ass.Equal(t, 6, int(glyphs.GetSize()))
	ass.Equal(t, "['A'..'F']", fmt.Sprintf("%v", glyphs))

	var durations = doc.IntervalClass[doc.DurationLike]().Interval(
		doc.Exclusive,
		doc.Duration(0),
		doc.Duration("~P4W"),
		doc.Inclusive,
	)
	ass.Equal(t, 2419200000, int(durations.GetSize()))
	ass.Equal(t, "(..~P4W]", fmt.Sprintf("%v", durations))

	durations = doc.Interval[doc.DurationLike](
		doc.Inclusive,
		doc.Duration("~P5D"),
		doc.Duration(0),
		doc.Exclusive,
	)
	ass.Equal(t, -432000000, int(durations.GetSize()))
	ass.Equal(t, "[~P5D..)", fmt.Sprintf("%v", durations))

	var moments = doc.Interval[doc.MomentLike](
		doc.Exclusive,
		doc.Moment("<2001-02-03T04:05:06>"),
		doc.Moment("<2001-02-03T04:05:07>"),
		doc.Exclusive,
	)
	ass.Equal(t, 999, int(moments.GetSize()))
	ass.Equal(
		t,
		"(<2001-02-03T04:05:06>..<2001-02-03T04:05:07>)",
		fmt.Sprintf("%v", moments),
	)
}

func TestSpectrumConstructors(t *tes.T) {
	var names = doc.Spectrum[doc.NameLike](
		doc.Inclusive,
		doc.Name("/nebula/classes/abstract"),
		doc.Name("/nebula/types"),
		doc.Inclusive,
	)
	ass.Equal(
		t,
		"[/nebula/classes/abstract../nebula/types]",
		fmt.Sprintf("%v", names),
	)

	var quotes = doc.SpectrumClass[doc.QuoteLike]().Spectrum(
		doc.Inclusive,
		doc.Quote(`"A"`),
		doc.Quote(`"Fe"`),
		doc.Exclusive,
	)
	ass.Equal(
		t,
		`["A".."Fe")`,
		fmt.Sprintf("%v", quotes),
	)

	var versions = doc.Spectrum[doc.VersionLike](
		doc.Exclusive,
		doc.Version("v1.2.3"),
		doc.Version("v2"),
		doc.Exclusive,
	)
	ass.Equal(
		t,
		`(v1.2.3..v2)`,
		fmt.Sprintf("%v", versions),
	)
}

func TestContinuumConstructors(t *tes.T) {
	var numbers = doc.Continuum[doc.NumberLike](
		doc.Exclusive,
		doc.Number(-1.23),
		doc.Number(4.56),
		doc.Exclusive,
	)
	ass.Equal(t, "(-1.23..4.56)", fmt.Sprintf("%v", numbers))

	numbers = doc.Continuum[doc.NumberLike](
		doc.Exclusive,
		doc.NumberClass().Undefined(),
		doc.NumberClass().Zero(),
		doc.Inclusive,
	)
	ass.Equal(t, "(..0]", fmt.Sprintf("%v", numbers))

	numbers = doc.Continuum[doc.NumberLike](
		doc.Inclusive,
		doc.Number(1),
		doc.NumberClass().Undefined(),
		doc.Exclusive,
	)
	ass.Equal(t, "[1..)", fmt.Sprintf("%v", numbers))

	var probabilities = doc.ContinuumClass[doc.ProbabilityLike]().Continuum(
		doc.Inclusive,
		doc.Probability(0.0),
		doc.Probability(1.0),
		doc.Inclusive,
	)
	ass.Equal(t, "[p0..p1]", fmt.Sprintf("%v", probabilities))

	var angles = doc.Continuum[doc.AngleLike](
		doc.Inclusive,
		doc.Angle(0),
		doc.AngleClass().Tau(),
		doc.Exclusive,
	)
	ass.Equal(t, "[~0..~τ)", fmt.Sprintf("%v", angles))
}

var (
	invalid uti.State
	state1  uti.State = "$State1"
	state2  uti.State = "$State2"
	state3  uti.State = "$State3"
)

var (
	initialized uti.Event = "$Initialized"
	processed   uti.Event = "$Processed"
	finalized   uti.Event = "$Finalized"
)

func TestController(t *tes.T) {
	var events = []uti.Event{initialized, processed, finalized}
	var transitions = map[uti.State]uti.Transitions{
		state1: uti.Transitions{state2, invalid, invalid},
		state2: uti.Transitions{invalid, state2, state3},
		state3: uti.Transitions{invalid, invalid, invalid},
	}

	var controller = uti.Controller(events, transitions, state1)
	ass.Equal(t, state1, controller.GetState())
	ass.Equal(t, state2, controller.ProcessEvent(initialized))
	ass.Equal(t, state2, controller.ProcessEvent(processed))
	ass.Equal(t, state3, controller.ProcessEvent(finalized))
	controller.SetState(state1)
	ass.Equal(t, state1, controller.GetState())
}
