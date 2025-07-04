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

func LiteralClass() LiteralClassLike {
	return literalClass()
}

// Constructor Methods

func (c *literalClass_) Literal(
	delimiter string,
	quoted string,
) LiteralLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(quoted) {
		panic("The \"quoted\" attribute is required by this class.")
	}
	var instance = &literal_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		quoted_:    quoted,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *literal_) GetClass() LiteralClassLike {
	return literalClass()
}

// Attribute Methods

func (v *literal_) GetDelimiter() string {
	return v.delimiter_
}

func (v *literal_) GetQuoted() string {
	return v.quoted_
}

// PROTECTED INTERFACE

// Instance Structure

type literal_ struct {
	// Declare the instance attributes.
	delimiter_ string
	quoted_    string
}

// Class Structure

type literalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func literalClass() *literalClass_ {
	return literalClassReference_
}

var literalClassReference_ = &literalClass_{
	// Initialize the class constants.
}
