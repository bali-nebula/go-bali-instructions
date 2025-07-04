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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func PrefixClass() PrefixClassLike {
	return prefixClass()
}

// Constructor Methods

func (c *prefixClass_) Prefix(
	label string,
	delimiter string,
) PrefixLike {
	if uti.IsUndefined(label) {
		panic("The \"label\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &prefix_{
		// Initialize the instance attributes.
		label_:     label,
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *prefix_) GetClass() PrefixClassLike {
	return prefixClass()
}

// Attribute Methods

func (v *prefix_) GetLabel() string {
	return v.label_
}

func (v *prefix_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type prefix_ struct {
	// Declare the instance attributes.
	label_     string
	delimiter_ string
}

// Class Structure

type prefixClass_ struct {
	// Declare the class constants.
}

// Class Reference

func prefixClass() *prefixClass_ {
	return prefixClassReference_
}

var prefixClassReference_ = &prefixClass_{
	// Initialize the class constants.
}
