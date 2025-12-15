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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-bali-documents/wiki
*/
package module

import (
	lan "github.com/bali-nebula/go-assembly-language/v3"
	age "github.com/bali-nebula/go-bali-instructions/v3/agents"
	ins "github.com/bali-nebula/go-bali-instructions/v3/instructions"
)

// TYPE ALIASES

// Agents

type (
	DeflatorClassLike  = age.DeflatorClassLike
	InflatorClassLike  = age.InflatorClassLike
	ProcessorClassLike = age.ProcessorClassLike
	ValidatorClassLike = age.ValidatorClassLike
	VisitorClassLike   = age.VisitorClassLike
)

type (
	DeflatorLike  = age.DeflatorLike
	InflatorLike  = age.InflatorLike
	ProcessorLike = age.ProcessorLike
	ValidatorLike = age.ValidatorLike
	VisitorLike   = age.VisitorLike
)

type (
	Methodical = age.Methodical
)

// Instructions

type (
	Action   = ins.Action
	Modifier = ins.Modifier
)

const (
	NoteAction                     = ins.NoteAction
	SkipAction                     = ins.SkipAction
	JumpAction                     = ins.JumpAction
	PushAction                     = ins.PushAction
	PullActionAction               = ins.PullActionAction
	LoadAction                     = ins.LoadAction
	SaveAction                     = ins.SaveAction
	DropAction                     = ins.DropAction
	CallAction                     = ins.CallAction
	SendAction                     = ins.SendAction
	OnAnyModifier                  = ins.OnAnyModifier
	OnEmptyModifier                = ins.OnEmptyModifier
	OnNoneModifier                 = ins.OnNoneModifier
	OnFalseModifier                = ins.OnFalseModifier
	HandlerModifier                = ins.HandlerModifier
	LiteralModifier                = ins.LiteralModifier
	ConstantModifier               = ins.ConstantModifier
	ArgumentModifier               = ins.ArgumentModifier
	ExceptionModifier              = ins.ExceptionModifier
	ComponentModifier              = ins.ComponentModifier
	ResultModifier                 = ins.ResultModifier
	DocumentModifier               = ins.DocumentModifier
	DraftModifier                  = ins.DraftModifier
	MessageModifier                = ins.MessageModifier
	VariableModifier               = ins.VariableModifier
	With0ArgumentsModifier         = ins.With0ArgumentsModifier
	With1ArgumentModifier          = ins.With1ArgumentModifier
	With2ArgumentsModifier         = ins.With2ArgumentsModifier
	With3ArgumentsModifier         = ins.With3ArgumentsModifier
	ComponentWithArgumentsModifier = ins.ComponentWithArgumentsModifier
	DocumentWithArgumentsModifier  = ins.DocumentWithArgumentsModifier
)

type (
	ArgumentClassLike    = ins.ArgumentClassLike
	AssemblyClassLike    = ins.AssemblyClassLike
	CallClassLike        = ins.CallClassLike
	ConstantClassLike    = ins.ConstantClassLike
	DropClassLike        = ins.DropClassLike
	HandlerClassLike     = ins.HandlerClassLike
	InstructionClassLike = ins.InstructionClassLike
	JumpClassLike        = ins.JumpClassLike
	LiteralClassLike     = ins.LiteralClassLike
	LoadClassLike        = ins.LoadClassLike
	NoteClassLike        = ins.NoteClassLike
	PullClassLike        = ins.PullClassLike
	PushClassLike        = ins.PushClassLike
	SaveClassLike        = ins.SaveClassLike
	SendClassLike        = ins.SendClassLike
	SkipClassLike        = ins.SkipClassLike
)

type (
	ArgumentLike    = ins.ArgumentLike
	AssemblyLike    = ins.AssemblyLike
	CallLike        = ins.CallLike
	ConstantLike    = ins.ConstantLike
	DropLike        = ins.DropLike
	HandlerLike     = ins.HandlerLike
	InstructionLike = ins.InstructionLike
	JumpLike        = ins.JumpLike
	LiteralLike     = ins.LiteralLike
	LoadLike        = ins.LoadLike
	NoteLike        = ins.NoteLike
	PullLike        = ins.PullLike
	PushLike        = ins.PushLike
	SaveLike        = ins.SaveLike
	SendLike        = ins.SendLike
	SkipLike        = ins.SkipLike
)

// CLASS ACCESSORS

// Agents

func DeflatorClass() DeflatorClassLike {
	return age.DeflatorClass()
}

func Deflator() DeflatorLike {
	return DeflatorClass().Deflator()
}

func InflatorClass() InflatorClassLike {
	return age.InflatorClass()
}

func Inflator() InflatorLike {
	return InflatorClass().Inflator()
}

func ProcessorClass() ProcessorClassLike {
	return age.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ValidatorClass() ValidatorClassLike {
	return age.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return age.VisitorClass()
}

func Visitor(
	processor age.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// Instructions

func ArgumentClass() ArgumentClassLike {
	return ins.ArgumentClass()
}

func Argument(
	symbol string,
) ArgumentLike {
	return ArgumentClass().Argument(
		symbol,
	)
}

func AssemblyClass() AssemblyClassLike {
	return ins.AssemblyClass()
}

func Assembly() AssemblyLike {
	return AssemblyClass().Assembly()
}

func CallClass() CallClassLike {
	return ins.CallClass()
}

func Call(
	symbol string,
	cardinality ins.Modifier,
) CallLike {
	return CallClass().Call(
		symbol,
		cardinality,
	)
}

func ConstantClass() ConstantClassLike {
	return ins.ConstantClass()
}

func Constant(
	symbol string,
) ConstantLike {
	return ConstantClass().Constant(
		symbol,
	)
}

func DropClass() DropClassLike {
	return ins.DropClass()
}

func Drop(
	component ins.Modifier,
	symbol string,
) DropLike {
	return DropClass().Drop(
		component,
		symbol,
	)
}

func HandlerClass() HandlerClassLike {
	return ins.HandlerClass()
}

func Handler(
	label string,
) HandlerLike {
	return HandlerClass().Handler(
		label,
	)
}

func InstructionClass() InstructionClassLike {
	return ins.InstructionClass()
}

func Instruction(
	optionalPrefix string,
	action any,
) InstructionLike {
	return InstructionClass().Instruction(
		optionalPrefix,
		action,
	)
}

func JumpClass() JumpClassLike {
	return ins.JumpClass()
}

func Jump(
	label string,
	condition ins.Modifier,
) JumpLike {
	return JumpClass().Jump(
		label,
		condition,
	)
}

func LiteralClass() LiteralClassLike {
	return ins.LiteralClass()
}

func Literal(
	quoted string,
) LiteralLike {
	return LiteralClass().Literal(
		quoted,
	)
}

func LoadClass() LoadClassLike {
	return ins.LoadClass()
}

func Load(
	component ins.Modifier,
	symbol string,
) LoadLike {
	return LoadClass().Load(
		component,
		symbol,
	)
}

func NoteClass() NoteClassLike {
	return ins.NoteClass()
}

func Note(
	description string,
) NoteLike {
	return NoteClass().Note(
		description,
	)
}

func PullClass() PullClassLike {
	return ins.PullClass()
}

func Pull(
	value ins.Modifier,
) PullLike {
	return PullClass().Pull(
		value,
	)
}

func PushClass() PushClassLike {
	return ins.PushClass()
}

func Push(
	source any,
) PushLike {
	return PushClass().Push(
		source,
	)
}

func SaveClass() SaveClassLike {
	return ins.SaveClass()
}

func Save(
	component ins.Modifier,
	symbol string,
) SaveLike {
	return SaveClass().Save(
		component,
		symbol,
	)
}

func SendClass() SendClassLike {
	return ins.SendClass()
}

func Send(
	symbol string,
	destination ins.Modifier,
) SendLike {
	return SendClass().Send(
		symbol,
		destination,
	)
}

func SkipClass() SkipClassLike {
	return ins.SkipClass()
}

func Skip() SkipLike {
	return SkipClass().Skip()
}

// GLOBAL FUNCTIONS

// Instructions

func ParseAssembly(
	source string,
) AssemblyLike {
	var inflator = Inflator()
	var parser = lan.Parser()
	var assembly = inflator.InflateAssembly(parser.ParseSource(source))
	return assembly
}

func FormatAssembly(
	assembly AssemblyLike,
) string {
	var deflator = Deflator()
	var formatter = lan.Formatter()
	var source = formatter.FormatAssembly(deflator.DeflateAssembly(assembly))
	return source
}

// Agents
