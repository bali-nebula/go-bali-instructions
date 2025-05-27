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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-bali-instructions/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ActionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete action-like class.
*/
type ActionClassLike interface {
	// Constructor Methods
	Action(
		any_ any,
	) ActionLike
}

/*
ArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		delimiter string,
		symbol string,
	) ArgumentLike
}

/*
AssemblyClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete assembly-like class.
*/
type AssemblyClassLike interface {
	// Constructor Methods
	Assembly(
		instructions col.ListLike[InstructionLike],
	) AssemblyLike
}

/*
CallClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete call-like class.
*/
type CallClassLike interface {
	// Constructor Methods
	Call(
		delimiter string,
		symbol string,
		optionalCardinality CardinalityLike,
	) CallLike
}

/*
CardinalityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete cardinality-like class.
*/
type CardinalityClassLike interface {
	// Constructor Methods
	Cardinality(
		delimiter1 string,
		count string,
		delimiter2 string,
	) CardinalityLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		any_ any,
	) ComponentLike
}

/*
ConditionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete condition-like class.
*/
type ConditionClassLike interface {
	// Constructor Methods
	Condition(
		any_ any,
	) ConditionLike
}

/*
ConditionallyClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete conditionally-like class.
*/
type ConditionallyClassLike interface {
	// Constructor Methods
	Conditionally(
		delimiter string,
		condition ConditionLike,
	) ConditionallyLike
}

/*
ConstantClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructor Methods
	Constant(
		delimiter string,
		symbol string,
	) ConstantLike
}

/*
DestinationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete destination-like class.
*/
type DestinationClassLike interface {
	// Constructor Methods
	Destination(
		any_ any,
	) DestinationLike
}

/*
DropClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete drop-like class.
*/
type DropClassLike interface {
	// Constructor Methods
	Drop(
		delimiter string,
		component ComponentLike,
		symbol string,
	) DropLike
}

/*
HandlerClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete handler-like class.
*/
type HandlerClassLike interface {
	// Constructor Methods
	Handler(
		delimiter string,
		label string,
	) HandlerLike
}

/*
InstructionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete instruction-like class.
*/
type InstructionClassLike interface {
	// Constructor Methods
	Instruction(
		optionalPrefix PrefixLike,
		action ActionLike,
	) InstructionLike
}

/*
JumpClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete jump-like class.
*/
type JumpClassLike interface {
	// Constructor Methods
	Jump(
		delimiter1 string,
		delimiter2 string,
		label string,
		optionalConditionally ConditionallyLike,
	) JumpLike
}

/*
LiteralClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete literal-like class.
*/
type LiteralClassLike interface {
	// Constructor Methods
	Literal(
		delimiter string,
		quoted string,
	) LiteralLike
}

/*
LoadClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete load-like class.
*/
type LoadClassLike interface {
	// Constructor Methods
	Load(
		delimiter string,
		component ComponentLike,
		symbol string,
	) LoadLike
}

/*
NoopClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete noop-like class.
*/
type NoopClassLike interface {
	// Constructor Methods
	Noop(
		delimiter string,
	) NoopLike
}

/*
NoteClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete note-like class.
*/
type NoteClassLike interface {
	// Constructor Methods
	Note(
		delimiter string,
		explanation string,
	) NoteLike
}

/*
ParameterizedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameterized-like class.
*/
type ParameterizedClassLike interface {
	// Constructor Methods
	Parameterized(
		delimiter1 string,
		delimiter2 string,
	) ParameterizedLike
}

/*
PrefixClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructor Methods
	Prefix(
		label string,
		delimiter string,
	) PrefixLike
}

/*
PullClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete pull-like class.
*/
type PullClassLike interface {
	// Constructor Methods
	Pull(
		delimiter string,
		value ValueLike,
	) PullLike
}

/*
PushClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete push-like class.
*/
type PushClassLike interface {
	// Constructor Methods
	Push(
		delimiter string,
		source SourceLike,
	) PushLike
}

/*
SaveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete save-like class.
*/
type SaveClassLike interface {
	// Constructor Methods
	Save(
		delimiter string,
		component ComponentLike,
		symbol string,
	) SaveLike
}

/*
SendClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete send-like class.
*/
type SendClassLike interface {
	// Constructor Methods
	Send(
		delimiter1 string,
		symbol string,
		delimiter2 string,
		destination DestinationLike,
		optionalParameterized ParameterizedLike,
	) SendLike
}

/*
SourceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete source-like class.
*/
type SourceClassLike interface {
	// Constructor Methods
	Source(
		any_ any,
	) SourceLike
}

/*
ValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Value(
		any_ any,
	) ValueLike
}

// INSTANCE DECLARATIONS

/*
ActionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete action-like class.
*/
type ActionLike interface {
	// Principal Methods
	GetClass() ActionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetSymbol() string
}

/*
AssemblyLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete assembly-like class.
*/
type AssemblyLike interface {
	// Principal Methods
	GetClass() AssemblyClassLike

	// Attribute Methods
	GetInstructions() col.ListLike[InstructionLike]
}

/*
CallLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete call-like class.
*/
type CallLike interface {
	// Principal Methods
	GetClass() CallClassLike

	// Attribute Methods
	GetDelimiter() string
	GetSymbol() string
	GetOptionalCardinality() CardinalityLike
}

/*
CardinalityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete cardinality-like class.
*/
type CardinalityLike interface {
	// Principal Methods
	GetClass() CardinalityClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetCount() string
	GetDelimiter2() string
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetAny() any
}

/*
ConditionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete condition-like class.
*/
type ConditionLike interface {
	// Principal Methods
	GetClass() ConditionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ConditionallyLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete conditionally-like class.
*/
type ConditionallyLike interface {
	// Principal Methods
	GetClass() ConditionallyClassLike

	// Attribute Methods
	GetDelimiter() string
	GetCondition() ConditionLike
}

/*
ConstantLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Principal Methods
	GetClass() ConstantClassLike

	// Attribute Methods
	GetDelimiter() string
	GetSymbol() string
}

/*
DestinationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete destination-like class.
*/
type DestinationLike interface {
	// Principal Methods
	GetClass() DestinationClassLike

	// Attribute Methods
	GetAny() any
}

/*
DropLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete drop-like class.
*/
type DropLike interface {
	// Principal Methods
	GetClass() DropClassLike

	// Attribute Methods
	GetDelimiter() string
	GetComponent() ComponentLike
	GetSymbol() string
}

/*
HandlerLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete handler-like class.
*/
type HandlerLike interface {
	// Principal Methods
	GetClass() HandlerClassLike

	// Attribute Methods
	GetDelimiter() string
	GetLabel() string
}

/*
InstructionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete instruction-like class.
*/
type InstructionLike interface {
	// Principal Methods
	GetClass() InstructionClassLike

	// Attribute Methods
	GetOptionalPrefix() PrefixLike
	GetAction() ActionLike
}

/*
JumpLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete jump-like class.
*/
type JumpLike interface {
	// Principal Methods
	GetClass() JumpClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetLabel() string
	GetOptionalConditionally() ConditionallyLike
}

/*
LiteralLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete literal-like class.
*/
type LiteralLike interface {
	// Principal Methods
	GetClass() LiteralClassLike

	// Attribute Methods
	GetDelimiter() string
	GetQuoted() string
}

/*
LoadLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete load-like class.
*/
type LoadLike interface {
	// Principal Methods
	GetClass() LoadClassLike

	// Attribute Methods
	GetDelimiter() string
	GetComponent() ComponentLike
	GetSymbol() string
}

/*
NoopLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete noop-like class.
*/
type NoopLike interface {
	// Principal Methods
	GetClass() NoopClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
NoteLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete note-like class.
*/
type NoteLike interface {
	// Principal Methods
	GetClass() NoteClassLike

	// Attribute Methods
	GetDelimiter() string
	GetExplanation() string
}

/*
ParameterizedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameterized-like class.
*/
type ParameterizedLike interface {
	// Principal Methods
	GetClass() ParameterizedClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
PrefixLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Principal Methods
	GetClass() PrefixClassLike

	// Attribute Methods
	GetLabel() string
	GetDelimiter() string
}

/*
PullLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete pull-like class.
*/
type PullLike interface {
	// Principal Methods
	GetClass() PullClassLike

	// Attribute Methods
	GetDelimiter() string
	GetValue() ValueLike
}

/*
PushLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete push-like class.
*/
type PushLike interface {
	// Principal Methods
	GetClass() PushClassLike

	// Attribute Methods
	GetDelimiter() string
	GetSource() SourceLike
}

/*
SaveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete save-like class.
*/
type SaveLike interface {
	// Principal Methods
	GetClass() SaveClassLike

	// Attribute Methods
	GetDelimiter() string
	GetComponent() ComponentLike
	GetSymbol() string
}

/*
SendLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete send-like class.
*/
type SendLike interface {
	// Principal Methods
	GetClass() SendClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetSymbol() string
	GetDelimiter2() string
	GetDestination() DestinationLike
	GetOptionalParameterized() ParameterizedLike
}

/*
SourceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete source-like class.
*/
type SourceLike interface {
	// Principal Methods
	GetClass() SourceClassLike

	// Attribute Methods
	GetAny() any
}

/*
ValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete value-like class.
*/
type ValueLike interface {
	// Principal Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetAny() any
}

// ASPECT DECLARATIONS
