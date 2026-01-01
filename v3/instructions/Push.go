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

import (
	fmt "fmt"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func PushClass() PushClassLike {
	return pushClass()
}

// Constructor Methods

func (c *pushClass_) Push(
	source any,
) PushLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	var instance = &push_{
		// Initialize the instance attributes.
		source_: source,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *push_) GetClass() PushClassLike {
	return pushClass()
}

func (v *push_) AsSource() string {
	var source = "PUSH "
	switch actual := v.source_.(type) {
	case LiteralLike:
		source += "LITERAL" + actual.GetQuoted()
	case ConstantLike:
		source += "CONSTANT" + actual.GetSymbol()
	case ArgumentLike:
		source += "ARGUMENT" + actual.GetSymbol()
	case HandlerLike:
		source += "HANDLER" + actual.GetLabel()
	default:
		var message = fmt.Sprintf(
			"An invalid source was found: %v(%T)",
			v.source_,
			v.source_,
		)
		panic(message)
	}
	return source
}

// Attribute Methods

func (v *push_) GetSource() any {
	return v.source_
}

// PROTECTED INTERFACE

// Instance Structure

type push_ struct {
	// Declare the instance attributes.
	source_ any
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
