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

func JumpClass() JumpClassLike {
	return jumpClass()
}

// Constructor Methods

func (c *jumpClass_) Jump(
	delimiter1 string,
	delimiter2 string,
	label string,
	optionalConditionally ConditionallyLike,
) JumpLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(label) {
		panic("The \"label\" attribute is required by this class.")
	}
	var instance = &jump_{
		// Initialize the instance attributes.
		delimiter1_:            delimiter1,
		delimiter2_:            delimiter2,
		label_:                 label,
		optionalConditionally_: optionalConditionally,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *jump_) GetClass() JumpClassLike {
	return jumpClass()
}

// Attribute Methods

func (v *jump_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *jump_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *jump_) GetLabel() string {
	return v.label_
}

func (v *jump_) GetOptionalConditionally() ConditionallyLike {
	return v.optionalConditionally_
}

// PROTECTED INTERFACE

// Instance Structure

type jump_ struct {
	// Declare the instance attributes.
	delimiter1_            string
	delimiter2_            string
	label_                 string
	optionalConditionally_ ConditionallyLike
}

// Class Structure

type jumpClass_ struct {
	// Declare the class constants.
}

// Class Reference

func jumpClass() *jumpClass_ {
	return jumpClassReference_
}

var jumpClassReference_ = &jumpClass_{
	// Initialize the class constants.
}
