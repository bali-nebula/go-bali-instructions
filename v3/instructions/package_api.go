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

/*
Package "instructions" provides the class definitions for each part of in an
"inflated" Bali Instruction™.  An inflated instruction can be manipulated by the
agents defined in this project.

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
package instructions

import (
	com "github.com/craterdog/go-essential-composites/v8"
)

// TYPE DECLARATIONS

/*
Component is a constrained type representing the possible target components.
*/
type Component uint16

const (
	VariableComponent Component = iota
	DraftComponent
	DocumentComponent
	MessageComponent
)

/*
Condition is a constrained type representing the possible jump conditions.
*/
type Condition uint16

const (
	OnAnyCondition Condition = iota
	OnEmptyCondition
	OnNoneCondition
	OnFalseCondition
)

/*
Destination is a constrained type representing the possible target destinations.
*/
type Destination uint16

const (
	ComponentDestination Destination = iota
	ComponentWithArgumentsDestination
	DocumentDestination
	DocumentWithArgumentsDestination
)

/*
Source is a constrained type representing the possible push sources.
*/
type Source uint16

const (
	LiteralSource Source = iota
	ConstantSource
	ArgumentSource
	HandlerSource
)

/*
Value is a constrained type representing the possible pull values.
*/
type Value uint16

const (
	ComponentValue Value = iota
	ResultValue
	ExceptionValue
	HandlerValue
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ArgumentClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		symbol string,
	) ArgumentLike
}

/*
CallClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
call-like class.
*/
type CallClassLike interface {
	// Constructor Methods
	Call(
		symbol string,
		cardinality uint8,
	) CallLike
}

/*
ConstantClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
constant-like class.
*/
type ConstantClassLike interface {
	// Constructor Methods
	Constant(
		symbol string,
	) ConstantLike
}

/*
DropClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
drop-like class.
*/
type DropClassLike interface {
	// Constructor Methods
	Drop(
		component Component,
		symbol string,
	) DropLike
}

/*
HandlerClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
handler-like class.
*/
type HandlerClassLike interface {
	// Constructor Methods
	Handler(
		label string,
	) HandlerLike
}

/*
InstructionClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete instruction-like class.
*/
type InstructionClassLike interface {
	// Constructor Methods
	Instruction(
		optionalPrefix PrefixLike,
		action any,
	) InstructionLike
}

/*
JumpClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
jump-like class.
*/
type JumpClassLike interface {
	// Constructor Methods
	Jump(
		label string,
		condition Condition,
	) JumpLike
}

/*
LiteralClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
literal-like class.
*/
type LiteralClassLike interface {
	// Constructor Methods
	Literal(
		quoted string,
	) LiteralLike
}

/*
LoadClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
load-like class.
*/
type LoadClassLike interface {
	// Constructor Methods
	Load(
		component Component,
		symbol string,
	) LoadLike
}

/*
MethodClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		instructions com.Sequential[InstructionLike],
	) MethodLike
}

/*
NoteClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
note-like class.
*/
type NoteClassLike interface {
	// Constructor Methods
	Note(
		description string,
	) NoteLike
}

/*
PrefixClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
prefix-like class.
*/
type PrefixClassLike interface {
	// Constructor Methods
	Prefix(
		label string,
	) PrefixLike
}

/*
PullClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
pull-like class.
*/
type PullClassLike interface {
	// Constructor Methods
	Pull(
		value Value,
	) PullLike
}

/*
PushClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
push-like class.
*/
type PushClassLike interface {
	// Constructor Methods
	Push(
		source any,
	) PushLike
}

/*
SaveClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
save-like class.
*/
type SaveClassLike interface {
	// Constructor Methods
	Save(
		component Component,
		symbol string,
	) SaveLike
}

/*
SendClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
send-like class.
*/
type SendClassLike interface {
	// Constructor Methods
	Send(
		symbol string,
		destination Destination,
	) SendLike
}

/*
SkipClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
skip-like class.
*/
type SkipClassLike interface {
	// Constructor Methods
	Skip() SkipLike
}

// INSTANCE DECLARATIONS

/*
ArgumentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
CallLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete call-like class.
*/
type CallLike interface {
	// Principal Methods
	GetClass() CallClassLike
	AsSource() string

	// Attribute Methods
	GetSymbol() string
	GetArgumentCount() uint8
}

/*
ConstantLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete constant-like class.
*/
type ConstantLike interface {
	// Principal Methods
	GetClass() ConstantClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
DropLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete drop-like class.
*/
type DropLike interface {
	// Principal Methods
	GetClass() DropClassLike
	AsSource() string

	// Attribute Methods
	GetComponent() Component
	GetSymbol() string
}

/*
HandlerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete handler-like class.
*/
type HandlerLike interface {
	// Principal Methods
	GetClass() HandlerClassLike

	// Attribute Methods
	GetLabel() string
}

/*
InstructionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete instruction-like class.
*/
type InstructionLike interface {
	// Principal Methods
	GetClass() InstructionClassLike
	AsSource() string

	// Attribute Methods
	GetOptionalPrefix() PrefixLike
	GetAction() any
}

/*
JumpLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete jump-like class.
*/
type JumpLike interface {
	// Principal Methods
	GetClass() JumpClassLike
	AsSource() string

	// Attribute Methods
	GetLabel() string
	GetCondition() Condition
}

/*
LiteralLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete literal-like class.
*/
type LiteralLike interface {
	// Principal Methods
	GetClass() LiteralClassLike

	// Attribute Methods
	GetQuoted() string
}

/*
LoadLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete load-like class.
*/
type LoadLike interface {
	// Principal Methods
	GetClass() LoadClassLike
	AsSource() string

	// Attribute Methods
	GetComponent() Component
	GetSymbol() string
}

/*
MethodLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike
	GetInstructions() com.Sequential[InstructionLike]
}

/*
NoteLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete note-like class.
*/
type NoteLike interface {
	// Principal Methods
	GetClass() NoteClassLike
	AsSource() string

	// Attribute Methods
	GetDescription() string
}

/*
PrefixLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete prefix-like class.
*/
type PrefixLike interface {
	// Principal Methods
	GetClass() PrefixClassLike

	// Attribute Methods
	GetLabel() string
}

/*
PullLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete pull-like class.
*/
type PullLike interface {
	// Principal Methods
	GetClass() PullClassLike
	AsSource() string

	// Attribute Methods
	GetValue() Value
}

/*
PushLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete push-like class.
*/
type PushLike interface {
	// Principal Methods
	GetClass() PushClassLike
	AsSource() string

	// Attribute Methods
	GetSource() any
}

/*
SaveLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete save-like class.
*/
type SaveLike interface {
	// Principal Methods
	GetClass() SaveClassLike
	AsSource() string

	// Attribute Methods
	GetComponent() Component
	GetSymbol() string
}

/*
SendLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete send-like class.
*/
type SendLike interface {
	// Principal Methods
	GetClass() SendClassLike
	AsSource() string

	// Attribute Methods
	GetSymbol() string
	GetDestination() Destination
}

/*
SkipLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete skip-like class.
*/
type SkipLike interface {
	// Principal Methods
	GetClass() SkipClassLike
	AsSource() string
}

// ASPECT DECLARATIONS
