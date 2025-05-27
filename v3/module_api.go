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
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-bali-instructions/wiki
*/
package module

import (
	ast "github.com/bali-nebula/go-bali-instructions/v3/ast"
	gra "github.com/bali-nebula/go-bali-instructions/v3/grammar"
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE ALIASES

// Ast

type (
	ActionClassLike        = ast.ActionClassLike
	ArgumentClassLike      = ast.ArgumentClassLike
	AssemblyClassLike      = ast.AssemblyClassLike
	CallClassLike          = ast.CallClassLike
	CardinalityClassLike   = ast.CardinalityClassLike
	ComponentClassLike     = ast.ComponentClassLike
	ConditionClassLike     = ast.ConditionClassLike
	ConditionallyClassLike = ast.ConditionallyClassLike
	ConstantClassLike      = ast.ConstantClassLike
	DestinationClassLike   = ast.DestinationClassLike
	DropClassLike          = ast.DropClassLike
	HandlerClassLike       = ast.HandlerClassLike
	InstructionClassLike   = ast.InstructionClassLike
	JumpClassLike          = ast.JumpClassLike
	LiteralClassLike       = ast.LiteralClassLike
	LoadClassLike          = ast.LoadClassLike
	NoopClassLike          = ast.NoopClassLike
	NoteClassLike          = ast.NoteClassLike
	ParameterizedClassLike = ast.ParameterizedClassLike
	PrefixClassLike        = ast.PrefixClassLike
	PullClassLike          = ast.PullClassLike
	PushClassLike          = ast.PushClassLike
	SaveClassLike          = ast.SaveClassLike
	SendClassLike          = ast.SendClassLike
	SourceClassLike        = ast.SourceClassLike
	ValueClassLike         = ast.ValueClassLike
)

type (
	ActionLike        = ast.ActionLike
	ArgumentLike      = ast.ArgumentLike
	AssemblyLike      = ast.AssemblyLike
	CallLike          = ast.CallLike
	CardinalityLike   = ast.CardinalityLike
	ComponentLike     = ast.ComponentLike
	ConditionLike     = ast.ConditionLike
	ConditionallyLike = ast.ConditionallyLike
	ConstantLike      = ast.ConstantLike
	DestinationLike   = ast.DestinationLike
	DropLike          = ast.DropLike
	HandlerLike       = ast.HandlerLike
	InstructionLike   = ast.InstructionLike
	JumpLike          = ast.JumpLike
	LiteralLike       = ast.LiteralLike
	LoadLike          = ast.LoadLike
	NoopLike          = ast.NoopLike
	NoteLike          = ast.NoteLike
	ParameterizedLike = ast.ParameterizedLike
	PrefixLike        = ast.PrefixLike
	PullLike          = ast.PullLike
	PushLike          = ast.PushLike
	SaveLike          = ast.SaveLike
	SendLike          = ast.SendLike
	SourceLike        = ast.SourceLike
	ValueLike         = ast.ValueLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	CountToken       = gra.CountToken
	DelimiterToken   = gra.DelimiterToken
	ExplanationToken = gra.ExplanationToken
	LabelToken       = gra.LabelToken
	NewlineToken     = gra.NewlineToken
	QuotedToken      = gra.QuotedToken
	SpaceToken       = gra.SpaceToken
	SymbolToken      = gra.SymbolToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func ActionClass() ActionClassLike {
	return ast.ActionClass()
}

func Action(
	any_ any,
) ActionLike {
	return ActionClass().Action(
		any_,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	delimiter string,
	symbol string,
) ArgumentLike {
	return ArgumentClass().Argument(
		delimiter,
		symbol,
	)
}

func AssemblyClass() AssemblyClassLike {
	return ast.AssemblyClass()
}

func Assembly(
	instructions col.ListLike[ast.InstructionLike],
) AssemblyLike {
	return AssemblyClass().Assembly(
		instructions,
	)
}

func CallClass() CallClassLike {
	return ast.CallClass()
}

func Call(
	delimiter string,
	symbol string,
	optionalCardinality ast.CardinalityLike,
) CallLike {
	return CallClass().Call(
		delimiter,
		symbol,
		optionalCardinality,
	)
}

func CardinalityClass() CardinalityClassLike {
	return ast.CardinalityClass()
}

func Cardinality(
	delimiter1 string,
	count string,
	delimiter2 string,
) CardinalityLike {
	return CardinalityClass().Cardinality(
		delimiter1,
		count,
		delimiter2,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	any_ any,
) ComponentLike {
	return ComponentClass().Component(
		any_,
	)
}

func ConditionClass() ConditionClassLike {
	return ast.ConditionClass()
}

func Condition(
	any_ any,
) ConditionLike {
	return ConditionClass().Condition(
		any_,
	)
}

func ConditionallyClass() ConditionallyClassLike {
	return ast.ConditionallyClass()
}

func Conditionally(
	delimiter string,
	condition ast.ConditionLike,
) ConditionallyLike {
	return ConditionallyClass().Conditionally(
		delimiter,
		condition,
	)
}

func ConstantClass() ConstantClassLike {
	return ast.ConstantClass()
}

func Constant(
	delimiter string,
	symbol string,
) ConstantLike {
	return ConstantClass().Constant(
		delimiter,
		symbol,
	)
}

func DestinationClass() DestinationClassLike {
	return ast.DestinationClass()
}

func Destination(
	any_ any,
) DestinationLike {
	return DestinationClass().Destination(
		any_,
	)
}

func DropClass() DropClassLike {
	return ast.DropClass()
}

func Drop(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) DropLike {
	return DropClass().Drop(
		delimiter,
		component,
		symbol,
	)
}

func HandlerClass() HandlerClassLike {
	return ast.HandlerClass()
}

func Handler(
	delimiter string,
	label string,
) HandlerLike {
	return HandlerClass().Handler(
		delimiter,
		label,
	)
}

func InstructionClass() InstructionClassLike {
	return ast.InstructionClass()
}

func Instruction(
	optionalPrefix ast.PrefixLike,
	action ast.ActionLike,
) InstructionLike {
	return InstructionClass().Instruction(
		optionalPrefix,
		action,
	)
}

func JumpClass() JumpClassLike {
	return ast.JumpClass()
}

func Jump(
	delimiter1 string,
	delimiter2 string,
	label string,
	optionalConditionally ast.ConditionallyLike,
) JumpLike {
	return JumpClass().Jump(
		delimiter1,
		delimiter2,
		label,
		optionalConditionally,
	)
}

func LiteralClass() LiteralClassLike {
	return ast.LiteralClass()
}

func Literal(
	delimiter string,
	quoted string,
) LiteralLike {
	return LiteralClass().Literal(
		delimiter,
		quoted,
	)
}

func LoadClass() LoadClassLike {
	return ast.LoadClass()
}

func Load(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) LoadLike {
	return LoadClass().Load(
		delimiter,
		component,
		symbol,
	)
}

func NoopClass() NoopClassLike {
	return ast.NoopClass()
}

func Noop(
	delimiter string,
) NoopLike {
	return NoopClass().Noop(
		delimiter,
	)
}

func NoteClass() NoteClassLike {
	return ast.NoteClass()
}

func Note(
	delimiter string,
	explanation string,
) NoteLike {
	return NoteClass().Note(
		delimiter,
		explanation,
	)
}

func ParameterizedClass() ParameterizedClassLike {
	return ast.ParameterizedClass()
}

func Parameterized(
	delimiter1 string,
	delimiter2 string,
) ParameterizedLike {
	return ParameterizedClass().Parameterized(
		delimiter1,
		delimiter2,
	)
}

func PrefixClass() PrefixClassLike {
	return ast.PrefixClass()
}

func Prefix(
	label string,
	delimiter string,
) PrefixLike {
	return PrefixClass().Prefix(
		label,
		delimiter,
	)
}

func PullClass() PullClassLike {
	return ast.PullClass()
}

func Pull(
	delimiter string,
	value ast.ValueLike,
) PullLike {
	return PullClass().Pull(
		delimiter,
		value,
	)
}

func PushClass() PushClassLike {
	return ast.PushClass()
}

func Push(
	delimiter string,
	source ast.SourceLike,
) PushLike {
	return PushClass().Push(
		delimiter,
		source,
	)
}

func SaveClass() SaveClassLike {
	return ast.SaveClass()
}

func Save(
	delimiter string,
	component ast.ComponentLike,
	symbol string,
) SaveLike {
	return SaveClass().Save(
		delimiter,
		component,
		symbol,
	)
}

func SendClass() SendClassLike {
	return ast.SendClass()
}

func Send(
	delimiter1 string,
	symbol string,
	delimiter2 string,
	destination ast.DestinationLike,
	optionalParameterized ast.ParameterizedLike,
) SendLike {
	return SendClass().Send(
		delimiter1,
		symbol,
		delimiter2,
		destination,
		optionalParameterized,
	)
}

func SourceClass() SourceClassLike {
	return ast.SourceClass()
}

func Source(
	any_ any,
) SourceLike {
	return SourceClass().Source(
		any_,
	)
}

func ValueClass() ValueClassLike {
	return ast.ValueClass()
}

func Value(
	any_ any,
) ValueLike {
	return ValueClass().Value(
		any_,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens col.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS

func FormatAssembly(
	syntax AssemblyLike,
) string {
	var formatter = Formatter()
	return formatter.FormatAssembly(syntax)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) AssemblyLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateAssembly(
	syntax AssemblyLike,
) {
	var validator = Validator()
	validator.ValidateAssembly(syntax)
}
