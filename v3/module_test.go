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
	not "github.com/bali-nebula/go-bali-instructions/v3"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var assemblyFiles = []string{
	"./testdata/instructions.basm",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, assemblyFile := range assemblyFiles {
		fmt.Printf("   %v\n", assemblyFile)
		var bytes, err = osx.ReadFile(assemblyFile)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var assembly = not.ParseSource(source)
		not.ValidateAssembly(assembly)
		var actual = not.FormatAssembly(assembly)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
