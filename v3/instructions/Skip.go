/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package instructions

import ()

// CLASS INTERFACE

// Access Function

func SkipClass() SkipClassLike {
	return skipClass()
}

// Constructor Methods

func (c *skipClass_) Skip() SkipLike {
	var instance = &skip_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *skip_) GetClass() SkipClassLike {
	return skipClass()
}

func (v *skip_) AsSource() string {
	return "SKIP"
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type skip_ struct {
	// Declare the instance attributes.
}

// Class Structure

type skipClass_ struct {
	// Declare the class constants.
}

// Class Reference

func skipClass() *skipClass_ {
	return skipClassReference_
}

var skipClassReference_ = &skipClass_{
	// Initialize the class constants.
}
