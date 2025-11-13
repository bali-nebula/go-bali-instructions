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
	age "github.com/bali-nebula/go-bali-documents/v3/agents"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uri "net/url"
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

// Documents

type (
	Assignment = doc.Assignment
	Inverse    = doc.Inverse
	Invoke     = doc.Invoke
	Operator   = doc.Operator
)

const (
	DefaultEquals  = doc.DefaultEquals
	AssignEquals   = doc.AssignEquals
	PlusEquals     = doc.PlusEquals
	MinusEquals    = doc.MinusEquals
	TimesEquals    = doc.TimesEquals
	DivideEquals   = doc.DivideEquals
	ChainEquals    = doc.ChainEquals
	Additive       = doc.Additive
	Multiplicative = doc.Multiplicative
	Conjugate      = doc.Conjugate
	Synchronous    = doc.Synchronous
	Asynchronous   = doc.Asynchronous
	Less           = doc.Less
	Equal          = doc.Equal
	More           = doc.More
	Is             = doc.Is
	Matches        = doc.Matches
	And            = doc.And
	San            = doc.San
	Ior            = doc.Ior
	Xor            = doc.Xor
	Plus           = doc.Plus
	Minus          = doc.Minus
	Times          = doc.Times
	Divide         = doc.Divide
	Remainder      = doc.Remainder
	Power          = doc.Power
	Chain          = doc.Chain
)

type (
	AcceptClauseClassLike   = doc.AcceptClauseClassLike
	AttributesClassLike     = doc.AttributesClassLike
	BreakClauseClassLike    = doc.BreakClauseClassLike
	CheckoutClauseClassLike = doc.CheckoutClauseClassLike
	ComplementClassLike     = doc.ComplementClassLike
	ComponentClassLike      = doc.ComponentClassLike
	ContentClassLike        = doc.ContentClassLike
	ConstraintClassLike     = doc.ConstraintClassLike
	ContinueClauseClassLike = doc.ContinueClauseClassLike
	DiscardClauseClassLike  = doc.DiscardClauseClassLike
	DoClauseClassLike       = doc.DoClauseClassLike
	DocumentClassLike       = doc.DocumentClassLike
	ExpressionClassLike     = doc.ExpressionClassLike
	FunctionClassLike       = doc.FunctionClassLike
	GenericsClassLike       = doc.GenericsClassLike
	IfClauseClassLike       = doc.IfClauseClassLike
	InspectClauseClassLike  = doc.InspectClauseClassLike
	InversionClassLike      = doc.InversionClassLike
	ItemsClassLike          = doc.ItemsClassLike
	LetClauseClassLike      = doc.LetClauseClassLike
	MagnitudeClassLike      = doc.MagnitudeClassLike
	MatchingClauseClassLike = doc.MatchingClauseClassLike
	MethodClassLike         = doc.MethodClassLike
	NotarizeClauseClassLike = doc.NotarizeClauseClassLike
	OnClauseClassLike       = doc.OnClauseClassLike
	PrecedenceClassLike     = doc.PrecedenceClassLike
	PredicateClassLike      = doc.PredicateClassLike
	ProcedureClassLike      = doc.ProcedureClassLike
	PublishClauseClassLike  = doc.PublishClauseClassLike
	RangeClassLike          = doc.RangeClassLike
	ReceiveClauseClassLike  = doc.ReceiveClauseClassLike
	ReferentClassLike       = doc.ReferentClassLike
	RejectClauseClassLike   = doc.RejectClauseClassLike
	RetrieveClauseClassLike = doc.RetrieveClauseClassLike
	ReturnClauseClassLike   = doc.ReturnClauseClassLike
	SaveClauseClassLike     = doc.SaveClauseClassLike
	SelectClauseClassLike   = doc.SelectClauseClassLike
	SendClauseClassLike     = doc.SendClauseClassLike
	StatementClassLike      = doc.StatementClassLike
	SubcomponentClassLike   = doc.SubcomponentClassLike
	ThrowClauseClassLike    = doc.ThrowClauseClassLike
	WhileClauseClassLike    = doc.WhileClauseClassLike
	WithClauseClassLike     = doc.WithClauseClassLike
)

type (
	AcceptClauseLike   = doc.AcceptClauseLike
	AttributesLike     = doc.AttributesLike
	BreakClauseLike    = doc.BreakClauseLike
	CheckoutClauseLike = doc.CheckoutClauseLike
	ComplementLike     = doc.ComplementLike
	ComponentLike      = doc.ComponentLike
	ContentLike        = doc.ContentLike
	ConstraintLike     = doc.ConstraintLike
	ContinueClauseLike = doc.ContinueClauseLike
	DiscardClauseLike  = doc.DiscardClauseLike
	DoClauseLike       = doc.DoClauseLike
	DocumentLike       = doc.DocumentLike
	ExpressionLike     = doc.ExpressionLike
	FunctionLike       = doc.FunctionLike
	GenericsLike       = doc.GenericsLike
	IfClauseLike       = doc.IfClauseLike
	InspectClauseLike  = doc.InspectClauseLike
	InversionLike      = doc.InversionLike
	ItemsLike          = doc.ItemsLike
	LetClauseLike      = doc.LetClauseLike
	MagnitudeLike      = doc.MagnitudeLike
	MatchingClauseLike = doc.MatchingClauseLike
	MethodLike         = doc.MethodLike
	NotarizeClauseLike = doc.NotarizeClauseLike
	OnClauseLike       = doc.OnClauseLike
	PrecedenceLike     = doc.PrecedenceLike
	PredicateLike      = doc.PredicateLike
	ProcedureLike      = doc.ProcedureLike
	PublishClauseLike  = doc.PublishClauseLike
	RangeLike          = doc.RangeLike
	ReceiveClauseLike  = doc.ReceiveClauseLike
	ReferentLike       = doc.ReferentLike
	RejectClauseLike   = doc.RejectClauseLike
	RetrieveClauseLike = doc.RetrieveClauseLike
	ReturnClauseLike   = doc.ReturnClauseLike
	SaveClauseLike     = doc.SaveClauseLike
	SelectClauseLike   = doc.SelectClauseLike
	SendClauseLike     = doc.SendClauseLike
	StatementLike      = doc.StatementLike
	SubcomponentLike   = doc.SubcomponentLike
	ThrowClauseLike    = doc.ThrowClauseLike
	WhileClauseLike    = doc.WhileClauseLike
	WithClauseLike     = doc.WithClauseLike
)

type (
	Composite = doc.Composite
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

// Documents

func AcceptClauseClass() AcceptClauseClassLike {
	return doc.AcceptClauseClass()
}

func AcceptClause(
	message doc.ExpressionLike,
	bag doc.ExpressionLike,
) AcceptClauseLike {
	return AcceptClauseClass().AcceptClause(
		message,
		bag,
	)
}

func AttributesClass() AttributesClassLike {
	return doc.AttributesClass()
}

func Attributes(
	associations com.CatalogLike[any, doc.ContentLike],
) AttributesLike {
	return AttributesClass().Attributes(
		associations,
	)
}

func BreakClauseClass() BreakClauseClassLike {
	return doc.BreakClauseClass()
}

func BreakClause() BreakClauseLike {
	return BreakClauseClass().BreakClause()
}

func CheckoutClauseClass() CheckoutClauseClassLike {
	return doc.CheckoutClauseClass()
}

func CheckoutClause(
	recipient any,
	optionalAtLevel doc.ExpressionLike,
	location doc.ExpressionLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		location,
	)
}

func ComplementClass() ComplementClassLike {
	return doc.ComplementClass()
}

func Complement(
	reversible any,
) ComplementLike {
	return ComplementClass().Complement(
		reversible,
	)
}

func ComponentClass() ComponentClassLike {
	return doc.ComponentClass()
}

func Component(
	entity any,
	optionalGenerics doc.GenericsLike,
) ComponentLike {
	return ComponentClass().Component(
		entity,
		optionalGenerics,
	)
}

func ContentClass() ContentClassLike {
	return doc.ContentClass()
}

func Content(
	composite doc.Composite,
	optionalNote string,
) ContentLike {
	return ContentClass().Content(
		composite,
		optionalNote,
	)
}

func ConstraintClass() ConstraintClassLike {
	return doc.ConstraintClass()
}

func Constraint(
	metadata any,
	optionalGenerics doc.GenericsLike,
) ConstraintLike {
	return ConstraintClass().Constraint(
		metadata,
		optionalGenerics,
	)
}

func ContinueClauseClass() ContinueClauseClassLike {
	return doc.ContinueClauseClass()
}

func ContinueClause() ContinueClauseLike {
	return ContinueClauseClass().ContinueClause()
}

func DiscardClauseClass() DiscardClauseClassLike {
	return doc.DiscardClauseClass()
}

func DiscardClause(
	citation doc.ExpressionLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		citation,
	)
}

func DoClauseClass() DoClauseClassLike {
	return doc.DoClauseClass()
}

func DoClause(
	method doc.MethodLike,
) DoClauseLike {
	return DoClauseClass().DoClause(
		method,
	)
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func Document(
	optionalAnnotation string,
	composite doc.Composite,
) DocumentLike {
	return DocumentClass().Document(
		optionalAnnotation,
		composite,
	)
}

func ExpressionClass() ExpressionClassLike {
	return doc.ExpressionClass()
}

func Expression(
	subject any,
	predicates com.Sequential[doc.PredicateLike],
) ExpressionLike {
	return ExpressionClass().Expression(
		subject,
		predicates,
	)
}

func FunctionClass() FunctionClassLike {
	return doc.FunctionClass()
}

func Function(
	identifier string,
	arguments com.Sequential[any],
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		arguments,
	)
}

func GenericsClass() GenericsClassLike {
	return doc.GenericsClass()
}

func Generics(
	parameters com.CatalogLike[pri.SymbolLike, doc.ConstraintLike],
) GenericsLike {
	return GenericsClass().Generics(
		parameters,
	)
}

func IfClauseClass() IfClauseClassLike {
	return doc.IfClauseClass()
}

func IfClause(
	condition doc.ExpressionLike,
	procedure doc.ProcedureLike,
) IfClauseLike {
	return IfClauseClass().IfClause(
		condition,
		procedure,
	)
}

func InspectClauseClass() InspectClauseClassLike {
	return doc.InspectClauseClass()
}

func InspectClause(
	recipient any,
	location doc.ExpressionLike,
) InspectClauseLike {
	return InspectClauseClass().InspectClause(
		recipient,
		location,
	)
}

func InversionClass() InversionClassLike {
	return doc.InversionClass()
}

func Inversion(
	inverse doc.Inverse,
	numerical any,
) InversionLike {
	return InversionClass().Inversion(
		inverse,
		numerical,
	)
}

func ItemsClass() ItemsClassLike {
	return doc.ItemsClass()
}

func Items(
	contents com.Sequential[doc.ContentLike],
) ItemsLike {
	return ItemsClass().Items(
		contents,
	)
}

func LetClauseClass() LetClauseClassLike {
	return doc.LetClauseClass()
}

func LetClause(
	recipient any,
	assignment doc.Assignment,
	expression doc.ExpressionLike,
) LetClauseLike {
	return LetClauseClass().LetClause(
		recipient,
		assignment,
		expression,
	)
}

func MagnitudeClass() MagnitudeClassLike {
	return doc.MagnitudeClass()
}

func Magnitude(
	expression doc.ExpressionLike,
) MagnitudeLike {
	return MagnitudeClass().Magnitude(
		expression,
	)
}

func MatchingClauseClass() MatchingClauseClassLike {
	return doc.MatchingClauseClass()
}

func MatchingClause(
	template doc.ExpressionLike,
	procedure doc.ProcedureLike,
) MatchingClauseLike {
	return MatchingClauseClass().MatchingClause(
		template,
		procedure,
	)
}

func MethodClass() MethodClassLike {
	return doc.MethodClass()
}

func Method(
	target string,
	invoke doc.Invoke,
	identifier string,
	arguments com.Sequential[any],
) MethodLike {
	return MethodClass().Method(
		target,
		invoke,
		identifier,
		arguments,
	)
}

func NotarizeClauseClass() NotarizeClauseClassLike {
	return doc.NotarizeClauseClass()
}

func NotarizeClause(
	draft doc.ExpressionLike,
	location doc.ExpressionLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		draft,
		location,
	)
}

func OnClauseClass() OnClauseClassLike {
	return doc.OnClauseClass()
}

func OnClause(
	symbol pri.SymbolLike,
	matchingClauses com.Sequential[doc.MatchingClauseLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		symbol,
		matchingClauses,
	)
}

func PrecedenceClass() PrecedenceClassLike {
	return doc.PrecedenceClass()
}

func Precedence(
	expression doc.ExpressionLike,
) PrecedenceLike {
	return PrecedenceClass().Precedence(
		expression,
	)
}

func PredicateClass() PredicateClassLike {
	return doc.PredicateClass()
}

func Predicate(
	operator doc.Operator,
	expression doc.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		operator,
		expression,
	)
}

func ProcedureClass() ProcedureClassLike {
	return doc.ProcedureClass()
}

func Procedure(
	lines com.Sequential[any],
) ProcedureLike {
	return ProcedureClass().Procedure(
		lines,
	)
}

func PublishClauseClass() PublishClauseClassLike {
	return doc.PublishClauseClass()
}

func PublishClause(
	message doc.ExpressionLike,
) PublishClauseLike {
	return PublishClauseClass().PublishClause(
		message,
	)
}

func RangeClass() RangeClassLike {
	return doc.RangeClass()
}

func Range(
	left com.Bracket,
	first any,
	last any,
	right com.Bracket,
) RangeLike {
	return RangeClass().Range(
		left,
		first,
		last,
		right,
	)
}

func ReceiveClauseClass() ReceiveClauseClassLike {
	return doc.ReceiveClauseClass()
}

func ReceiveClause(
	recipient any,
	bag doc.ExpressionLike,
) ReceiveClauseLike {
	return ReceiveClauseClass().ReceiveClause(
		recipient,
		bag,
	)
}

func ReferentClass() ReferentClassLike {
	return doc.ReferentClass()
}

func Referent(
	reference any,
) ReferentLike {
	return ReferentClass().Referent(
		reference,
	)
}

func RejectClauseClass() RejectClauseClassLike {
	return doc.RejectClauseClass()
}

func RejectClause(
	message doc.ExpressionLike,
	bag doc.ExpressionLike,
) RejectClauseLike {
	return RejectClauseClass().RejectClause(
		message,
		bag,
	)
}

func RetrieveClauseClass() RetrieveClauseClassLike {
	return doc.RetrieveClauseClass()
}

func RetrieveClause(
	recipient any,
	citation doc.ExpressionLike,
) RetrieveClauseLike {
	return RetrieveClauseClass().RetrieveClause(
		recipient,
		citation,
	)
}

func ReturnClauseClass() ReturnClauseClassLike {
	return doc.ReturnClauseClass()
}

func ReturnClause(
	result doc.ExpressionLike,
) ReturnClauseLike {
	return ReturnClauseClass().ReturnClause(
		result,
	)
}

func SaveClauseClass() SaveClauseClassLike {
	return doc.SaveClauseClass()
}

func SaveClause(
	draft doc.ExpressionLike,
	recipient any,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		draft,
		recipient,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return doc.SelectClauseClass()
}

func SelectClause(
	template doc.ExpressionLike,
	matchingClauses com.Sequential[doc.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		template,
		matchingClauses,
	)
}

func SendClauseClass() SendClauseClassLike {
	return doc.SendClauseClass()
}

func SendClause(
	message doc.ExpressionLike,
	bag doc.ExpressionLike,
) SendClauseLike {
	return SendClauseClass().SendClause(
		message,
		bag,
	)
}

func StatementClass() StatementClassLike {
	return doc.StatementClass()
}

func Statement(
	mainClause any,
	optionalOnClause doc.OnClauseLike,
) StatementLike {
	return StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
}

func SubcomponentClass() SubcomponentClassLike {
	return doc.SubcomponentClass()
}

func Subcomponent(
	identifier string,
	indexes com.Sequential[any],
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		indexes,
	)
}

func ThrowClauseClass() ThrowClauseClassLike {
	return doc.ThrowClauseClass()
}

func ThrowClause(
	exception doc.ExpressionLike,
) ThrowClauseLike {
	return ThrowClauseClass().ThrowClause(
		exception,
	)
}

func WhileClauseClass() WhileClauseClassLike {
	return doc.WhileClauseClass()
}

func WhileClause(
	condition doc.ExpressionLike,
	procedure doc.ProcedureLike,
) WhileClauseLike {
	return WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
}

func WithClauseClass() WithClauseClassLike {
	return doc.WithClauseClass()
}

func WithClause(
	symbol pri.SymbolLike,
	sequence doc.ExpressionLike,
	procedure doc.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		symbol,
		sequence,
		procedure,
	)
}

// GLOBAL FUNCTIONS

// Components

func ParseComponent(
	source string,
) Composite {
	var document = ParseDocument(source)
	return document.GetComposite()
}

func ParseDocument(
	source string,
) DocumentLike {
	var inflator = Inflator()
	var parser = not.Parser()
	var document = inflator.InflateDocument(parser.ParseSource(source))
	return document
}

func FormatComponent(
	value any,
) string {
	var composite Composite
	switch actual := value.(type) {
	case Composite:
		composite = actual
	case ConstraintLike:
		var entity = actual.GetMetadata()
		var generics = actual.GetOptionalGenerics()
		composite = Component(entity, generics)
	case DocumentLike:
		composite = actual.GetComposite()
	case ContentLike:
		composite = actual.GetComposite()
	default:
		composite = Component(actual, nil)
	}
	var source = FormatDocument(Document("", composite))
	return source[:len(source)-1] // Remove the trailing newline.
}

func FormatDocument(
	document DocumentLike,
) string {
	var deflator = Deflator()
	var formatter = not.Formatter()
	var source = formatter.FormatDocument(deflator.DeflateDocument(document))
	return source
}

func FormatInstructions(
	instructions com.Sequential[InstructionLike],
) string {
	var result sts.Builder
	result.WriteString(
		`Address   Bytes   Instruction                Mnemonic
--------------------------------------------------------------------
`,
	)
	var iterator = instructions.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var address = fmt.Sprintf("[x%03x]", iterator.GetSlot())
		var bytes = fmt.Sprintf("x%04x", instruction.AsIntrinsic())
		var operation = instruction.GetOperation() >> 13
		var modifier = instruction.GetModifier() >> 11
		var operand = instruction.OperandAsString()
		for len(operand) < 5 {
			operand = " " + operand
		}
		if len(operand) < 6 {
			operand += " "
		}
		var bytecode = fmt.Sprintf("%d %d %s", operation, modifier, operand)
		var description = instruction.AsString()
		var line = fmt.Sprintf(
			"%s:   %s   %s   %s\n",
			address,
			bytes,
			bytecode,
			description,
		)
		result.WriteString(line)
	}
	return result.String()
}

// Agents

type (
	Rank = com.Rank
)

const (
	LesserRank  = com.LesserRank
	EqualRank   = com.EqualRank
	GreaterRank = com.GreaterRank
)

type (
	RankingFunction[V any] = com.RankingFunction[V]
)

type (
	CollatorClassLike[V any] = com.CollatorClassLike[V]
	SorterClassLike[V any]   = com.SorterClassLike[V]
)

type (
	CollatorLike[V any] = com.CollatorLike[V]
	SorterLike[V any]   = com.SorterLike[V]
)

func CollatorClass[V any]() CollatorClassLike[V] {
	return com.CollatorClass[V]()
}

func Collator[V any](
	value ...any,
) CollatorLike[V] {
	if len(value) == 0 {
		return CollatorClass[V]().Collator()
	}
	switch actual := value[0].(type) {
	case int:
		return CollatorClass[V]().CollatorWithMaximumDepth(uint(actual))
	case uint:
		return CollatorClass[V]().CollatorWithMaximumDepth(actual)
	default:
		return CollatorClass[V]().Collator()
	}
}

func SorterClass[V any]() SorterClassLike[V] {
	return com.SorterClass[V]()
}

func Sorter[V any](
	value ...any,
) SorterLike[V] {
	if len(value) == 0 {
		return SorterClass[V]().Sorter()
	}
	switch actual := value[0].(type) {
	case RankingFunction[V]:
		return SorterClass[V]().SorterWithRanker(actual)
	default:
		return SorterClass[V]().Sorter()
	}
}

// Elements

type (
	Units = pri.Units
)

const (
	Degrees  = pri.Degrees
	Radians  = pri.Radians
	Gradians = pri.Gradians
)

type (
	AngleClassLike       = pri.AngleClassLike
	BooleanClassLike     = pri.BooleanClassLike
	DurationClassLike    = pri.DurationClassLike
	GlyphClassLike       = pri.GlyphClassLike
	MomentClassLike      = pri.MomentClassLike
	NumberClassLike      = pri.NumberClassLike
	PercentageClassLike  = pri.PercentageClassLike
	ProbabilityClassLike = pri.ProbabilityClassLike
	ResourceClassLike    = pri.ResourceClassLike
)

type (
	AngleLike       = pri.AngleLike
	BooleanLike     = pri.BooleanLike
	DurationLike    = pri.DurationLike
	GlyphLike       = pri.GlyphLike
	MomentLike      = pri.MomentLike
	NumberLike      = pri.NumberLike
	PercentageLike  = pri.PercentageLike
	ProbabilityLike = pri.ProbabilityLike
	ResourceLike    = pri.ResourceLike
)

type (
	Continuous = pri.Continuous
	Discrete   = pri.Discrete
	Factored   = pri.Factored
	Polarized  = pri.Polarized
	Temporal   = pri.Temporal
)

func AngleClass() AngleClassLike {
	return pri.AngleClass()
}

func Angle(
	value ...any,
) AngleLike {
	if len(value) == 0 {
		return AngleClass().Zero()
	}
	switch actual := value[0].(type) {
	case string:
		return AngleClass().AngleFromSource(actual)
	case float64:
		return AngleClass().Angle(actual)
	default:
		return AngleClass().Zero()
	}
}

func BooleanClass() BooleanClassLike {
	return pri.BooleanClass()
}

func Boolean(
	value ...any,
) BooleanLike {
	if len(value) == 0 {
		return BooleanClass().False()
	}
	switch actual := value[0].(type) {
	case string:
		return BooleanClass().BooleanFromSource(actual)
	case bool:
		return BooleanClass().Boolean(actual)
	default:
		return BooleanClass().False()
	}
}

func DurationClass() DurationClassLike {
	return pri.DurationClass()
}

func Duration(
	value ...any,
) DurationLike {
	if len(value) == 0 {
		return DurationClass().Duration(0)
	}
	switch actual := value[0].(type) {
	case string:
		return DurationClass().DurationFromSource(actual)
	case int:
		if actual < 0 {
			actual = -actual
		}
		return DurationClass().Duration(uint(actual))
	case uint:
		return DurationClass().Duration(actual)
	default:
		return DurationClass().Duration(0)
	}
}

func GlyphClass() GlyphClassLike {
	return pri.GlyphClass()
}

func Glyph(
	value ...any,
) GlyphLike {
	if len(value) == 0 {
		return GlyphClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return GlyphClass().GlyphFromSource(actual)
	case rune:
		return GlyphClass().Glyph(actual)
	case int:
		return GlyphClass().GlyphFromInteger(actual)
	default:
		return GlyphClass().Undefined()
	}
}

func MomentClass() MomentClassLike {
	return pri.MomentClass()
}

func Moment(
	value ...any,
) MomentLike {
	if len(value) == 0 {
		return MomentClass().Now()
	}
	switch actual := value[0].(type) {
	case string:
		return MomentClass().MomentFromSource(actual)
	case int:
		return MomentClass().Moment(actual)
	default:
		return MomentClass().Now()
	}
}

func NumberClass() NumberClassLike {
	return pri.NumberClass()
}

func Number(
	value ...any,
) NumberLike {
	if len(value) == 0 {
		return NumberClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return NumberClass().NumberFromSource(actual)
	case complex128:
		return NumberClass().Number(actual)
	case int:
		return NumberClass().Number(complex(float64(actual), 0))
	case float64:
		return NumberClass().Number(complex(actual, 0))
	default:
		return NumberClass().Undefined()
	}
}

func Polar(
	magnitude float64,
	phase float64,
) NumberLike {
	return NumberClass().NumberFromPolar(magnitude, phase)
}

func Rectangular(
	x float64,
	y float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(x, y)
}

func PercentageClass() PercentageClassLike {
	return pri.PercentageClass()
}

func Percentage(
	value ...any,
) PercentageLike {
	if len(value) == 0 {
		return PercentageClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return PercentageClass().PercentageFromSource(actual)
	case int:
		return PercentageClass().PercentageFromInteger(actual)
	case float64:
		return PercentageClass().Percentage(actual)
	default:
		return PercentageClass().Undefined()
	}
}

func ProbabilityClass() ProbabilityClassLike {
	return pri.ProbabilityClass()
}

func Probability(
	value ...any,
) ProbabilityLike {
	if len(value) == 0 {
		return ProbabilityClass().Random()
	}
	switch actual := value[0].(type) {
	case string:
		return ProbabilityClass().ProbabilityFromSource(actual)
	case bool:
		return ProbabilityClass().ProbabilityFromBoolean(actual)
	case float64:
		return ProbabilityClass().Probability(actual)
	default:
		return ProbabilityClass().Random()
	}
}

func ResourceClass() ResourceClassLike {
	return pri.ResourceClass()
}

func Resource(
	value ...any,
) ResourceLike {
	if len(value) == 0 {
		return ResourceClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] == '<' {
			return ResourceClass().ResourceFromSource(actual)
		} else {
			return ResourceClass().Resource(actual)
		}
	case *uri.URL:
		return ResourceClass().ResourceFromUri(actual)
	default:
		return ResourceClass().Undefined()
	}
}

// Sequences

type (
	Folder = pri.Folder
)

type (
	BinaryClassLike    = pri.BinaryClassLike
	BytecodeClassLike  = pri.BytecodeClassLike
	NameClassLike      = pri.NameClassLike
	NarrativeClassLike = pri.NarrativeClassLike
	PatternClassLike   = pri.PatternClassLike
	QuoteClassLike     = pri.QuoteClassLike
	SymbolClassLike    = pri.SymbolClassLike
	TagClassLike       = pri.TagClassLike
	VersionClassLike   = pri.VersionClassLike
)

type (
	BinaryLike    = pri.BinaryLike
	BytecodeLike  = pri.BytecodeLike
	NameLike      = pri.NameLike
	NarrativeLike = pri.NarrativeLike
	PatternLike   = pri.PatternLike
	QuoteLike     = pri.QuoteLike
	SymbolLike    = pri.SymbolLike
	TagLike       = pri.TagLike
	VersionLike   = pri.VersionLike
)

type (
	Accessible[V any] = pri.Accessible[V]
	Searchable[V any] = pri.Searchable[V]
	Sequential[V any] = pri.Sequential[V]
	Ordered[V any]    = pri.Ordered[V]
)

func BinaryClass() BinaryClassLike {
	return pri.BinaryClass()
}

func Binary(
	value ...any,
) BinaryLike {
	if len(value) == 0 {
		return BinaryClass().Binary([]byte{})
	}
	switch actual := value[0].(type) {
	case string:
		return BinaryClass().BinaryFromSource(actual)
	case []byte:
		return BinaryClass().Binary(actual)
	case Sequential[byte]:
		return BinaryClass().BinaryFromSequence(actual)
	default:
		return BinaryClass().Binary([]byte{})
	}
}

func BytecodeClass() BytecodeClassLike {
	return pri.BytecodeClass()
}

func Bytecode(
	value ...any,
) BytecodeLike {
	if len(value) == 0 {
		return BytecodeClass().Bytecode([]uint16{0})
	}
	switch actual := value[0].(type) {
	case string:
		return BytecodeClass().BytecodeFromSource(actual)
	case []uint16:
		return BytecodeClass().Bytecode(actual)
	default:
		return BytecodeClass().Bytecode([]uint16{0})
	}
}

func NameClass() NameClassLike {
	return pri.NameClass()
}

func Name(
	value ...any,
) NameLike {
	if len(value) == 0 {
		return NameClass().Name([]Folder{})
	}
	switch actual := value[0].(type) {
	case string:
		return NameClass().NameFromSource(actual)
	case []Folder:
		return NameClass().Name(actual)
	case Sequential[Folder]:
		return NameClass().NameFromSequence(actual)
	default:
		return NameClass().Name([]Folder{})
	}
}

func NarrativeClass() NarrativeClassLike {
	return pri.NarrativeClass()
}

func Narrative(
	value ...any,
) NarrativeLike {
	if len(value) == 0 {
		return NarrativeClass().Narrative([]string{})
	}
	switch actual := value[0].(type) {
	case string:
		return NarrativeClass().NarrativeFromSource(actual)
	case []string:
		return NarrativeClass().Narrative(actual)
	case Sequential[string]:
		return NarrativeClass().NarrativeFromSequence(actual)
	default:
		return NarrativeClass().Narrative([]string{})
	}
}

func PatternClass() PatternClassLike {
	return pri.PatternClass()
}

func Pattern(
	value ...any,
) PatternLike {
	if len(value) == 0 {
		return PatternClass().None()
	}
	switch actual := value[0].(type) {
	case string:
		return PatternClass().PatternFromSource(actual)
	case []rune:
		return PatternClass().Pattern(actual)
	case Sequential[rune]:
		return PatternClass().PatternFromSequence(actual)
	default:
		return PatternClass().None()
	}
}

func QuoteClass() QuoteClassLike {
	return pri.QuoteClass()
}

func Quote(
	value ...any,
) QuoteLike {
	if len(value) == 0 {
		return QuoteClass().QuoteFromSource(`""`)
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] == '"' {
			return QuoteClass().QuoteFromSource(actual)
		} else {
			return QuoteClass().QuoteFromSource(`"` + actual + `"`)
		}
	case []rune:
		return QuoteClass().Quote(actual)
	case Sequential[rune]:
		return QuoteClass().QuoteFromSequence(actual)
	default:
		return QuoteClass().QuoteFromSource(`""`)
	}
}

func SymbolClass() SymbolClassLike {
	return pri.SymbolClass()
}

func Symbol(
	value ...any,
) SymbolLike {
	if len(value) == 0 {
		return SymbolClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] != '$' {
			actual = "$" + actual
		}
		return SymbolClass().SymbolFromSource(actual)
	case []rune:
		return SymbolClass().Symbol(actual)
	default:
		return SymbolClass().Undefined()
	}
}

func TagClass() TagClassLike {
	return pri.TagClass()
}

func Tag(
	value ...any,
) TagLike {
	if len(value) == 0 {
		return TagClass().TagWithSize(20)
	}
	switch actual := value[0].(type) {
	case string:
		return TagClass().TagFromSource(actual)
	case []byte:
		return TagClass().Tag(actual)
	case Sequential[byte]:
		return TagClass().TagFromSequence(actual)
	case int:
		return TagClass().TagWithSize(uint(actual))
	case uint:
		return TagClass().TagWithSize(actual)
	default:
		return TagClass().TagWithSize(20)
	}
}

func VersionClass() VersionClassLike {
	return pri.VersionClass()
}

func Version(
	value ...any,
) VersionLike {
	if len(value) == 0 {
		return VersionClass().Version([]uint{1})
	}
	switch actual := value[0].(type) {
	case string:
		return VersionClass().VersionFromSource(actual)
	case []uint:
		return VersionClass().Version(actual)
	case Sequential[uint]:
		return VersionClass().VersionFromSequence(actual)
	default:
		return VersionClass().Version([]uint{})
	}
}

// Ranges

type (
	Bracket = com.Bracket
)

const (
	Inclusive = com.Inclusive
	Exclusive = com.Exclusive
)

type (
	ContinuumClassLike[V com.Continuous] = com.ContinuumClassLike[V]
	IntervalClassLike[V com.Discrete]    = com.IntervalClassLike[V]
	SpectrumClassLike[V com.Ordered[V]]  = com.SpectrumClassLike[V]
)

type (
	ContinuumLike[V com.Continuous] = com.ContinuumLike[V]
	IntervalLike[V com.Discrete]    = com.IntervalLike[V]
	SpectrumLike[V com.Ordered[V]]  = com.SpectrumLike[V]
)

type (
	Bounded[V any] = com.Bounded[V]
)

func ContinuumClass[V Continuous]() ContinuumClassLike[V] {
	return com.ContinuumClass[V]()
}

func Continuum[V Continuous](
	left Bracket,
	minimum V,
	maximum V,
	right Bracket,
) ContinuumLike[V] {
	return ContinuumClass[V]().Continuum(
		left,
		minimum,
		maximum,
		right,
	)
}

func IntervalClass[V Discrete]() IntervalClassLike[V] {
	return com.IntervalClass[V]()
}

func Interval[V Discrete](
	left Bracket,
	minimum V,
	maximum V,
	right Bracket,
) IntervalLike[V] {
	return IntervalClass[V]().Interval(
		left,
		minimum,
		maximum,
		right,
	)
}

func SpectrumClass[V Ordered[V]]() SpectrumClassLike[V] {
	return com.SpectrumClass[V]()
}

func Spectrum[V Ordered[V]](
	left Bracket,
	minimum V,
	maximum V,
	right Bracket,
) SpectrumLike[V] {
	return SpectrumClass[V]().Spectrum(
		left,
		minimum,
		maximum,
		right,
	)
}

// Collections

type (
	AssociationClassLike[K comparable, V any] = com.AssociationClassLike[K, V]
	CatalogClassLike[K comparable, V any]     = com.CatalogClassLike[K, V]
	ListClassLike[V any]                      = com.ListClassLike[V]
	QueueClassLike[V any]                     = com.QueueClassLike[V]
	SetClassLike[V any]                       = com.SetClassLike[V]
	StackClassLike[V any]                     = com.StackClassLike[V]
)

type (
	AssociationLike[K comparable, V any] = com.AssociationLike[K, V]
	CatalogLike[K comparable, V any]     = com.CatalogLike[K, V]
	ListLike[V any]                      = com.ListLike[V]
	QueueLike[V any]                     = com.QueueLike[V]
	SetLike[V any]                       = com.SetLike[V]
	StackLike[V any]                     = com.StackLike[V]
)

type (
	Associative[K comparable, V any] = com.Associative[K, V]
	Elastic[V any]                   = com.Elastic[V]
	Fifo[V any]                      = com.Fifo[V]
	Lifo[V any]                      = com.Lifo[V]
	Malleable[V any]                 = com.Malleable[V]
	Sortable[V any]                  = com.Sortable[V]
	Synchronized                     = com.Synchronized
	Updatable[V any]                 = com.Updatable[V]
)

func AssociationClass[K comparable, V any]() AssociationClassLike[K, V] {
	return com.AssociationClass[K, V]()
}

func Association[K comparable, V any](
	key K,
	value V,
) AssociationLike[K, V] {
	return AssociationClass[K, V]().Association(
		key,
		value,
	)
}

func CatalogClass[K comparable, V any]() CatalogClassLike[K, V] {
	return com.CatalogClass[K, V]()
}

func Catalog[K comparable, V any](
	value ...any,
) CatalogLike[K, V] {
	if len(value) == 0 {
		return com.Catalog[K, V]()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(CatalogLike[K, V])
	case []AssociationLike[K, V]:
		return com.CatalogFromArray(actual)
	case map[K]V:
		return com.CatalogFromMap(actual)
	case Sequential[AssociationLike[K, V]]:
		return com.CatalogFromSequence(actual)
	default:
		return com.Catalog[K, V]()
	}
}

func ListClass[V any]() ListClassLike[V] {
	return com.ListClass[V]()
}

func List[V any](
	value ...any,
) ListLike[V] {
	if len(value) == 0 {
		return ListClass[V]().List()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(ListLike[V])
	case []V:
		return ListClass[V]().ListFromArray(actual)
	case Sequential[V]:
		return ListClass[V]().ListFromSequence(actual)
	default:
		return ListClass[V]().List()
	}
}

func QueueClass[V any]() QueueClassLike[V] {
	return com.QueueClass[V]()
}

func Queue[V any](
	value ...any,
) QueueLike[V] {
	if len(value) == 0 {
		return QueueClass[V]().Queue()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(QueueLike[V])
	case int:
		return QueueClass[V]().QueueWithCapacity(uint(actual))
	case uint:
		return QueueClass[V]().QueueWithCapacity(actual)
	case []V:
		return QueueClass[V]().QueueFromArray(actual)
	case Sequential[V]:
		return QueueClass[V]().QueueFromSequence(actual)
	default:
		return QueueClass[V]().Queue()
	}
}

func SetClass[V any]() SetClassLike[V] {
	return com.SetClass[V]()
}

func Set[V any](
	value ...any,
) SetLike[V] {
	if len(value) == 0 {
		return SetClass[V]().Set()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(SetLike[V])
	case CollatorLike[V]:
		return SetClass[V]().SetWithCollator(actual)
	case []V:
		return SetClass[V]().SetFromArray(actual)
	case Sequential[V]:
		return SetClass[V]().SetFromSequence(actual)
	default:
		return SetClass[V]().Set()
	}
}

func StackClass[V any]() StackClassLike[V] {
	return com.StackClass[V]()
}

func Stack[V any](
	value ...any,
) StackLike[V] {
	if len(value) == 0 {
		return StackClass[V]().Stack()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(StackLike[V])
	case int:
		return StackClass[V]().StackWithCapacity(uint(actual))
	case uint:
		return StackClass[V]().StackWithCapacity(actual)
	case []V:
		return StackClass[V]().StackFromArray(actual)
	case Sequential[V]:
		return StackClass[V]().StackFromSequence(actual)
	default:
		return StackClass[V]().Stack()
	}
}
