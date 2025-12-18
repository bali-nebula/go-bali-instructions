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
Package "agents" provides classes that act on Bali Assembly Language™.

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package agents

import (
	lan "github.com/bali-nebula/go-assembly-language/v3"
	ins "github.com/bali-nebula/go-bali-instructions/v3/instructions"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
DeflatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete deflator-like class.
*/
type DeflatorClassLike interface {
	// Constructor Methods
	Deflator() DeflatorLike
}

/*
InflatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete inflator-like class.
*/
type InflatorClassLike interface {
	// Constructor Methods
	Inflator() InflatorLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
DeflatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete deflator-like class.
*/
type DeflatorLike interface {
	// Principal Methods
	GetClass() DeflatorClassLike
	DeflateAssembly(
		assembly ins.AssemblyLike,
	) lan.AssemblyLike

	// Aspect Interfaces
	Methodical
}

/*
InflatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete inflator-like class.
*/
type InflatorLike interface {
	// Principal Methods
	GetClass() InflatorClassLike
	InflateAssembly(
		assembly lan.AssemblyLike,
	) ins.AssemblyLike

	// Aspect Interfaces
	lan.Methodical
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateAssembly(
		assembly ins.AssemblyLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitAssembly(
		assembly ins.AssemblyLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessDescription(
		description string,
	)
	ProcessLabel(
		label string,
	)
	ProcessModifier(
		modifier ins.Modifier,
	)
	ProcessQuoted(
		quoted string,
	)
	ProcessSymbol(
		symbol string,
	)
	PreprocessAction(
		action any,
		index_ uint,
		count_ uint,
	)
	PostprocessAction(
		action any,
		index_ uint,
		count_ uint,
	)
	ProcessActionSlot(
		action any,
		slot_ uint,
	)
	PreprocessArgument(
		argument ins.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArgument(
		argument ins.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		argument ins.ArgumentLike,
		slot_ uint,
	)
	PreprocessAssembly(
		assembly ins.AssemblyLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAssembly(
		assembly ins.AssemblyLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssemblySlot(
		assembly ins.AssemblyLike,
		slot_ uint,
	)
	PreprocessCall(
		call ins.CallLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCall(
		call ins.CallLike,
		index_ uint,
		count_ uint,
	)
	ProcessCallSlot(
		call ins.CallLike,
		slot_ uint,
	)
	PreprocessConstant(
		constant ins.ConstantLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstant(
		constant ins.ConstantLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstantSlot(
		constant ins.ConstantLike,
		slot_ uint,
	)
	PreprocessDrop(
		drop ins.DropLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDrop(
		drop ins.DropLike,
		index_ uint,
		count_ uint,
	)
	ProcessDropSlot(
		drop ins.DropLike,
		slot_ uint,
	)
	PreprocessHandler(
		handler ins.HandlerLike,
		index_ uint,
		count_ uint,
	)
	PostprocessHandler(
		handler ins.HandlerLike,
		index_ uint,
		count_ uint,
	)
	ProcessHandlerSlot(
		handler ins.HandlerLike,
		slot_ uint,
	)
	PreprocessInstruction(
		instruction ins.InstructionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInstruction(
		instruction ins.InstructionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInstructionSlot(
		instruction ins.InstructionLike,
		slot_ uint,
	)
	PreprocessJump(
		jump ins.JumpLike,
		index_ uint,
		count_ uint,
	)
	PostprocessJump(
		jump ins.JumpLike,
		index_ uint,
		count_ uint,
	)
	ProcessJumpSlot(
		jump ins.JumpLike,
		slot_ uint,
	)
	PreprocessLiteral(
		literal ins.LiteralLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLiteral(
		literal ins.LiteralLike,
		index_ uint,
		count_ uint,
	)
	ProcessLiteralSlot(
		literal ins.LiteralLike,
		slot_ uint,
	)
	PreprocessLoad(
		load ins.LoadLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLoad(
		load ins.LoadLike,
		index_ uint,
		count_ uint,
	)
	ProcessLoadSlot(
		load ins.LoadLike,
		slot_ uint,
	)
	PreprocessNote(
		note ins.NoteLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNote(
		note ins.NoteLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoteSlot(
		note ins.NoteLike,
		slot_ uint,
	)
	PreprocessPrefix(
		prefix ins.PrefixLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrefix(
		prefix ins.PrefixLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrefixSlot(
		prefix ins.PrefixLike,
		slot_ uint,
	)
	PreprocessPull(
		pull ins.PullLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPull(
		pull ins.PullLike,
		index_ uint,
		count_ uint,
	)
	ProcessPullSlot(
		pull ins.PullLike,
		slot_ uint,
	)
	PreprocessPush(
		push ins.PushLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPush(
		push ins.PushLike,
		index_ uint,
		count_ uint,
	)
	ProcessPushSlot(
		push ins.PushLike,
		slot_ uint,
	)
	PreprocessSave(
		save ins.SaveLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSave(
		save ins.SaveLike,
		index_ uint,
		count_ uint,
	)
	ProcessSaveSlot(
		save ins.SaveLike,
		slot_ uint,
	)
	PreprocessSend(
		send ins.SendLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSend(
		send ins.SendLike,
		index_ uint,
		count_ uint,
	)
	ProcessSendSlot(
		send ins.SendLike,
		slot_ uint,
	)
	PreprocessSkip(
		skip ins.SkipLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSkip(
		skip ins.SkipLike,
		index_ uint,
		count_ uint,
	)
	ProcessSkipSlot(
		skip ins.SkipLike,
		slot_ uint,
	)
	PreprocessSource(
		source any,
		index_ uint,
		count_ uint,
	)
	PostprocessSource(
		source any,
		index_ uint,
		count_ uint,
	)
	ProcessSourceSlot(
		source any,
		slot_ uint,
	)
}
