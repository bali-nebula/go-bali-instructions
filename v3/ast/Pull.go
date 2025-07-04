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

func PullClass() PullClassLike {
	return pullClass()
}

// Constructor Methods

func (c *pullClass_) Pull(
	delimiter string,
	value ValueLike,
) PullLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &pull_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		value_:     value,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *pull_) GetClass() PullClassLike {
	return pullClass()
}

// Attribute Methods

func (v *pull_) GetDelimiter() string {
	return v.delimiter_
}

func (v *pull_) GetValue() ValueLike {
	return v.value_
}

// PROTECTED INTERFACE

// Instance Structure

type pull_ struct {
	// Declare the instance attributes.
	delimiter_ string
	value_     ValueLike
}

// Class Structure

type pullClass_ struct {
	// Declare the class constants.
}

// Class Reference

func pullClass() *pullClass_ {
	return pullClassReference_
}

var pullClassReference_ = &pullClass_{
	// Initialize the class constants.
}
