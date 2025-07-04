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

func PushClass() PushClassLike {
	return pushClass()
}

// Constructor Methods

func (c *pushClass_) Push(
	delimiter string,
	source SourceLike,
) PushLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	var instance = &push_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		source_:    source,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *push_) GetClass() PushClassLike {
	return pushClass()
}

// Attribute Methods

func (v *push_) GetDelimiter() string {
	return v.delimiter_
}

func (v *push_) GetSource() SourceLike {
	return v.source_
}

// PROTECTED INTERFACE

// Instance Structure

type push_ struct {
	// Declare the instance attributes.
	delimiter_ string
	source_    SourceLike
}

// Class Structure

type pushClass_ struct {
	// Declare the class constants.
}

// Class Reference

func pushClass() *pushClass_ {
	return pushClassReference_
}

var pushClassReference_ = &pushClass_{
	// Initialize the class constants.
}
