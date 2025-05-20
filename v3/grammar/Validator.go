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
│                 This class file was automatically generated.                 │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/bali-nebula/go-bali-instructions/v3/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateAssembly(
	assembly ast.AssemblyLike,
) {
	v.visitor_.VisitAssembly(assembly)
}

// Methodical Methods

func (v *validator_) ProcessCount(
	count string,
) {
	v.validateToken(count, CountToken)
}

func (v *validator_) ProcessExplanation(
	explanation string,
) {
	v.validateToken(explanation, ExplanationToken)
}

func (v *validator_) ProcessLabel(
	label string,
) {
	v.validateToken(label, LabelToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessQuoted(
	quoted string,
) {
	v.validateToken(quoted, QuotedToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) ProcessSymbol(
	symbol string,
) {
	v.validateToken(symbol, SymbolToken)
}

func (v *validator_) PreprocessAction(
	action ast.ActionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAction(
	action ast.ActionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArgument(
	argument ast.ArgumentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAssembly(
	assembly ast.AssemblyLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAssembly(
	assembly ast.AssemblyLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCall(
	call ast.CallLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCall(
	call ast.CallLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCondition(
	condition ast.ConditionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCondition(
	condition ast.ConditionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConditionally(
	conditionally ast.ConditionallyLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConditionally(
	conditionally ast.ConditionallyLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstant(
	constant ast.ConstantLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstant(
	constant ast.ConstantLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessContext(
	context ast.ContextLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessContext(
	context ast.ContextLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDestination(
	destination ast.DestinationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDestination(
	destination ast.DestinationLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDrop(
	drop ast.DropLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDrop(
	drop ast.DropLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessHandler(
	handler ast.HandlerLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessHandler(
	handler ast.HandlerLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInstruction(
	instruction ast.InstructionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInstruction(
	instruction ast.InstructionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessItem(
	item ast.ItemLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessItem(
	item ast.ItemLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessJump(
	jump ast.JumpLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessJump(
	jump ast.JumpLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLiteral(
	literal ast.LiteralLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLiteral(
	literal ast.LiteralLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLoad(
	load ast.LoadLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLoad(
	load ast.LoadLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoop(
	noop ast.NoopLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNoop(
	noop ast.NoopLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNote(
	note ast.NoteLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNote(
	note ast.NoteLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrefix(
	prefix ast.PrefixLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrefix(
	prefix ast.PrefixLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPull(
	pull ast.PullLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPull(
	pull ast.PullLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPush(
	push ast.PushLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPush(
	push ast.PushLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSave(
	save ast.SaveLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSave(
	save ast.SaveLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSend(
	send ast.SendLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSend(
	send ast.SendLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSource(
	source ast.SourceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSource(
	source ast.SourceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
