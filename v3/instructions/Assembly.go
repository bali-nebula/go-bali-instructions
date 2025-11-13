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

package instructions

import (
	com "github.com/craterdog/go-essential-composites/v8"
)

// CLASS INTERFACE

// Access Function

func AssemblyClass() AssemblyClassLike {
	return assemblyClass()
}

// Constructor Methods

func (c *assemblyClass_) Assembly() AssemblyLike {
	var instance = &assembly_{
		// Initialize the instance attributes.
		literals_:     com.Set[string](),
		constants_:    com.Set[string](),
		arguments_:    com.Set[string](),
		variables_:    com.Set[string](),
		messages_:     com.Set[string](),
		addresses_:    com.Catalog[string, uint16](),
		instructions_: com.List[InstructionLike](),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *assembly_) GetClass() AssemblyClassLike {
	return assemblyClass()
}

func (v *assembly_) AddLiteral(literal string) {
	v.literals_.AddValue(literal)
}

func (v *assembly_) AddConstant(constant string) {
	v.constants_.AddValue(constant)
}

func (v *assembly_) AddArgument(argument string) {
	v.arguments_.AddValue(argument)
}

func (v *assembly_) AddVariable(variable string) {
	v.variables_.AddValue(variable)
}

func (v *assembly_) AddMessage(message string) {
	v.messages_.AddValue(message)
}

func (v *assembly_) AddAddress(label string, address uint16) {
	v.addresses_.SetValue(label, address)
}

func (v *assembly_) AddInstruction(instruction InstructionLike) {
	v.instructions_.AppendValue(instruction)
}

func (v *assembly_) GetLiterals() com.Accessible[string] {
	return v.literals_
}

func (v *assembly_) GetConstants() com.Accessible[string] {
	return v.constants_
}

func (v *assembly_) GetArguments() com.Accessible[string] {
	return v.arguments_
}

func (v *assembly_) GetVariables() com.Accessible[string] {
	return v.variables_
}

func (v *assembly_) GetMessages() com.Accessible[string] {
	return v.messages_
}

func (v *assembly_) GetAddresses() com.Associative[string, uint16] {
	return v.addresses_
}

func (v *assembly_) GetInstructions() com.Sequential[InstructionLike] {
	return v.instructions_
}

// PROTECTED INTERFACE

// Instance Structure

type assembly_ struct {
	// Declare the instance attributes.
	literals_     com.SetLike[string]
	constants_    com.SetLike[string]
	arguments_    com.SetLike[string]
	variables_    com.SetLike[string]
	messages_     com.SetLike[string]
	addresses_    com.CatalogLike[string, uint16]
	instructions_ com.ListLike[InstructionLike]
}

// Class Structure

type assemblyClass_ struct {
	// Declare the class constants.
}

// Class Reference

func assemblyClass() *assemblyClass_ {
	return assemblyClassReference_
}

var assemblyClassReference_ = &assemblyClass_{
	// Initialize the class constants.
}
