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
Package "agents" provides classes that act on Bali documents.

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
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
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
	DeflateDocument(
		document doc.DocumentLike,
	) not.DocumentLike

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
	InflateDocument(
		document not.DocumentLike,
	) doc.DocumentLike

	// Aspect Interfaces
	not.Methodical
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
	ValidateDocument(
		document doc.DocumentLike,
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
	VisitDocument(
		document doc.DocumentLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessAngle(
		angle pri.AngleLike,
	)
	ProcessAnnotation(
		annotation string,
	)
	ProcessAssignment(
		assignment doc.Assignment,
	)
	ProcessBinary(
		binary pri.BinaryLike,
	)
	ProcessBoolean(
		boolean pri.BooleanLike,
	)
	ProcessBracket(
		bracket com.Bracket,
	)
	ProcessBytecode(
		bytecode pri.BytecodeLike,
	)
	ProcessComment(
		comment string,
	)
	ProcessDuration(
		duration pri.DurationLike,
	)
	ProcessGlyph(
		glyph pri.GlyphLike,
	)
	ProcessIdentifier(
		identifier string,
	)
	ProcessInverse(
		inverse doc.Inverse,
	)
	ProcessInvoke(
		invoke doc.Invoke,
	)
	ProcessMoment(
		moment pri.MomentLike,
	)
	ProcessName(
		name pri.NameLike,
	)
	ProcessNarrative(
		narrative pri.NarrativeLike,
	)
	ProcessNote(
		note string,
	)
	ProcessNumber(
		number pri.NumberLike,
	)
	ProcessOperator(
		operator doc.Operator,
	)
	ProcessPattern(
		pattern pri.PatternLike,
	)
	ProcessPercentage(
		percentage pri.PercentageLike,
	)
	ProcessProbability(
		probability pri.ProbabilityLike,
	)
	ProcessQuote(
		quote pri.QuoteLike,
	)
	ProcessResource(
		resource pri.ResourceLike,
	)
	ProcessSymbol(
		symbol pri.SymbolLike,
	)
	ProcessTag(
		tag pri.TagLike,
	)
	ProcessVersion(
		version pri.VersionLike,
	)
	PreprocessAcceptClause(
		acceptClause doc.AcceptClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessAcceptClauseSlot(
		acceptClause doc.AcceptClauseLike,
		slot_ uint,
	)
	PostprocessAcceptClause(
		acceptClause doc.AcceptClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArgument(
		argument any,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		argument any,
		slot_ uint,
	)
	PostprocessArgument(
		argument any,
		index_ uint,
		count_ uint,
	)
	PreprocessAttributes(
		attributes doc.AttributesLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributesSlot(
		attributes doc.AttributesLike,
		slot_ uint,
	)
	PostprocessAttributes(
		attributes doc.AttributesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessBreakClause(
		breakClause doc.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessBreakClauseSlot(
		breakClause doc.BreakClauseLike,
		slot_ uint,
	)
	PostprocessBreakClause(
		breakClause doc.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCheckoutClause(
		checkoutClause doc.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessCheckoutClauseSlot(
		checkoutClause doc.CheckoutClauseLike,
		slot_ uint,
	)
	PostprocessCheckoutClause(
		checkoutClause doc.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessComplement(
		complement doc.ComplementLike,
		index_ uint,
		count_ uint,
	)
	ProcessComplementSlot(
		complement doc.ComplementLike,
		slot_ uint,
	)
	PostprocessComplement(
		complement doc.ComplementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessComposite(
		composite doc.Composite,
		index_ uint,
		count_ uint,
	)
	ProcessCompositeSlot(
		composite doc.Composite,
		slot_ uint,
	)
	PostprocessComposite(
		composite doc.Composite,
		index_ uint,
		count_ uint,
	)
	PreprocessContent(
		content doc.ContentLike,
		index_ uint,
		count_ uint,
	)
	ProcessContentSlot(
		content doc.ContentLike,
		slot_ uint,
	)
	PostprocessContent(
		content doc.ContentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessConstraint(
		constraint doc.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstraintSlot(
		constraint doc.ConstraintLike,
		slot_ uint,
	)
	PostprocessConstraint(
		constraint doc.ConstraintLike,
		index_ uint,
		count_ uint,
	)
	PreprocessContinueClause(
		continueClause doc.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessContinueClauseSlot(
		continueClause doc.ContinueClauseLike,
		slot_ uint,
	)
	PostprocessContinueClause(
		continueClause doc.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDiscardClause(
		discardClause doc.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDiscardClauseSlot(
		discardClause doc.DiscardClauseLike,
		slot_ uint,
	)
	PostprocessDiscardClause(
		discardClause doc.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDoClause(
		doClause doc.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDoClauseSlot(
		doClause doc.DoClauseLike,
		slot_ uint,
	)
	PostprocessDoClause(
		doClause doc.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDocument(
		document doc.DocumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessDocumentSlot(
		document doc.DocumentLike,
		slot_ uint,
	)
	PostprocessDocument(
		document doc.DocumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessEntity(
		entity any,
		index_ uint,
		count_ uint,
	)
	ProcessEntitySlot(
		entity any,
		slot_ uint,
	)
	PostprocessEntity(
		entity any,
		index_ uint,
		count_ uint,
	)
	PreprocessExpression(
		expression doc.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionSlot(
		expression doc.ExpressionLike,
		slot_ uint,
	)
	PostprocessExpression(
		expression doc.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunction(
		function doc.FunctionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionSlot(
		function doc.FunctionLike,
		slot_ uint,
	)
	PostprocessFunction(
		function doc.FunctionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessGenerics(
		generics doc.GenericsLike,
		index_ uint,
		count_ uint,
	)
	ProcessGenericsSlot(
		generics doc.GenericsLike,
		slot_ uint,
	)
	PostprocessGenerics(
		generics doc.GenericsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIfClause(
		ifClause doc.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessIfClauseSlot(
		ifClause doc.IfClauseLike,
		slot_ uint,
	)
	PostprocessIfClause(
		ifClause doc.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIndex(
		index any,
		index_ uint,
		count_ uint,
	)
	ProcessIndexSlot(
		index any,
		slot_ uint,
	)
	PostprocessIndex(
		index any,
		index_ uint,
		count_ uint,
	)
	PreprocessInspectClause(
		inspectClause doc.InspectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessInspectClauseSlot(
		inspectClause doc.InspectClauseLike,
		slot_ uint,
	)
	PostprocessInspectClause(
		inspectClause doc.InspectClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInversion(
		inversion doc.InversionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInversionSlot(
		inversion doc.InversionLike,
		slot_ uint,
	)
	PostprocessInversion(
		inversion doc.InversionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessItems(
		items doc.ItemsLike,
		index_ uint,
		count_ uint,
	)
	ProcessItemsSlot(
		items doc.ItemsLike,
		slot_ uint,
	)
	PostprocessItems(
		items doc.ItemsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLetClause(
		letClause doc.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessLetClauseSlot(
		letClause doc.LetClauseLike,
		slot_ uint,
	)
	PostprocessLetClause(
		letClause doc.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLine(
		line any,
		index_ uint,
		count_ uint,
	)
	ProcessLineSlot(
		line any,
		slot_ uint,
	)
	PostprocessLine(
		line any,
		index_ uint,
		count_ uint,
	)
	PreprocessLogical(
		logical any,
		index_ uint,
		count_ uint,
	)
	ProcessLogicalSlot(
		logical any,
		slot_ uint,
	)
	PostprocessLogical(
		logical any,
		index_ uint,
		count_ uint,
	)
	PreprocessMagnitude(
		magnitude doc.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	ProcessMagnitudeSlot(
		magnitude doc.MagnitudeLike,
		slot_ uint,
	)
	PostprocessMagnitude(
		magnitude doc.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMainClause(
		mainClause any,
		index_ uint,
		count_ uint,
	)
	ProcessMainClauseSlot(
		mainClause any,
		slot_ uint,
	)
	PostprocessMainClause(
		mainClause any,
		index_ uint,
		count_ uint,
	)
	PreprocessMatchingClause(
		matchingClause doc.MatchingClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessMatchingClauseSlot(
		matchingClause doc.MatchingClauseLike,
		slot_ uint,
	)
	PostprocessMatchingClause(
		matchingClause doc.MatchingClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMetadata(
		metadata any,
		index_ uint,
		count_ uint,
	)
	ProcessMetadataSlot(
		metadata any,
		slot_ uint,
	)
	PostprocessMetadata(
		metadata any,
		index_ uint,
		count_ uint,
	)
	PreprocessMethod(
		method doc.MethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessMethodSlot(
		method doc.MethodLike,
		slot_ uint,
	)
	PostprocessMethod(
		method doc.MethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNotarizeClause(
		notarizeClause doc.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessNotarizeClauseSlot(
		notarizeClause doc.NotarizeClauseLike,
		slot_ uint,
	)
	PostprocessNotarizeClause(
		notarizeClause doc.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNumerical(
		numerical any,
		index_ uint,
		count_ uint,
	)
	ProcessNumericalSlot(
		numerical any,
		slot_ uint,
	)
	PostprocessNumerical(
		numerical any,
		index_ uint,
		count_ uint,
	)
	PreprocessOnClause(
		onClause doc.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessOnClauseSlot(
		onClause doc.OnClauseLike,
		slot_ uint,
	)
	PostprocessOnClause(
		onClause doc.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrecedence(
		precedence doc.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrecedenceSlot(
		precedence doc.PrecedenceLike,
		slot_ uint,
	)
	PostprocessPrecedence(
		precedence doc.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPredicate(
		predicate doc.PredicateLike,
		index_ uint,
		count_ uint,
	)
	ProcessPredicateSlot(
		predicate doc.PredicateLike,
		slot_ uint,
	)
	PostprocessPredicate(
		predicate doc.PredicateLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrimitive(
		primitive any,
		index_ uint,
		count_ uint,
	)
	ProcessPrimitiveSlot(
		primitive any,
		slot_ uint,
	)
	PostprocessPrimitive(
		primitive any,
		index_ uint,
		count_ uint,
	)
	PreprocessProcedure(
		procedure doc.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	ProcessProcedureSlot(
		procedure doc.ProcedureLike,
		slot_ uint,
	)
	PostprocessProcedure(
		procedure doc.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPublishClause(
		publishClause doc.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessPublishClauseSlot(
		publishClause doc.PublishClauseLike,
		slot_ uint,
	)
	PostprocessPublishClause(
		publishClause doc.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRange(
		range_ doc.RangeLike,
		index_ uint,
		count_ uint,
	)
	ProcessRangeSlot(
		range_ doc.RangeLike,
		slot_ uint,
	)
	PostprocessRange(
		range_ doc.RangeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessReceiveClause(
		receiveClause doc.ReceiveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessReceiveClauseSlot(
		receiveClause doc.ReceiveClauseLike,
		slot_ uint,
	)
	PostprocessReceiveClause(
		receiveClause doc.ReceiveClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRecipient(
		recipient any,
		index_ uint,
		count_ uint,
	)
	ProcessRecipientSlot(
		recipient any,
		slot_ uint,
	)
	PostprocessRecipient(
		recipient any,
		index_ uint,
		count_ uint,
	)
	PreprocessReference(
		reference any,
		index_ uint,
		count_ uint,
	)
	ProcessReferenceSlot(
		reference any,
		slot_ uint,
	)
	PostprocessReference(
		reference any,
		index_ uint,
		count_ uint,
	)
	PreprocessReferent(
		referent doc.ReferentLike,
		index_ uint,
		count_ uint,
	)
	ProcessReferentSlot(
		referent doc.ReferentLike,
		slot_ uint,
	)
	PostprocessReferent(
		referent doc.ReferentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRejectClause(
		rejectClause doc.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRejectClauseSlot(
		rejectClause doc.RejectClauseLike,
		slot_ uint,
	)
	PostprocessRejectClause(
		rejectClause doc.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRetrieveClause(
		rejectClause doc.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRetrieveClauseSlot(
		rejectClause doc.RetrieveClauseLike,
		slot_ uint,
	)
	PostprocessRetrieveClause(
		rejectClause doc.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessReturnClause(
		returnClause doc.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessReturnClauseSlot(
		returnClause doc.ReturnClauseLike,
		slot_ uint,
	)
	PostprocessReturnClause(
		returnClause doc.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessReversible(
		reversible any,
		index_ uint,
		count_ uint,
	)
	ProcessReversibleSlot(
		reversible any,
		slot_ uint,
	)
	PostprocessReversible(
		reversible any,
		index_ uint,
		count_ uint,
	)
	PreprocessSaveClause(
		saveClause doc.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSaveClauseSlot(
		saveClause doc.SaveClauseLike,
		slot_ uint,
	)
	PostprocessSaveClause(
		saveClause doc.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSelectClause(
		selectClause doc.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSelectClauseSlot(
		selectClause doc.SelectClauseLike,
		slot_ uint,
	)
	PostprocessSelectClause(
		selectClause doc.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSendClause(
		sendClause doc.SendClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSendClauseSlot(
		sendClause doc.SendClauseLike,
		slot_ uint,
	)
	PostprocessSendClause(
		sendClause doc.SendClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessStatement(
		statement doc.StatementLike,
		index_ uint,
		count_ uint,
	)
	ProcessStatementSlot(
		statement doc.StatementLike,
		slot_ uint,
	)
	PostprocessStatement(
		statement doc.StatementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSubcomponent(
		subcomponent doc.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessSubcomponentSlot(
		subcomponent doc.SubcomponentLike,
		slot_ uint,
	)
	PostprocessSubcomponent(
		subcomponent doc.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSubject(
		subject any,
		index_ uint,
		count_ uint,
	)
	ProcessSubjectSlot(
		subject any,
		slot_ uint,
	)
	PostprocessSubject(
		subject any,
		index_ uint,
		count_ uint,
	)
	PreprocessThrowClause(
		throwClause doc.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessThrowClauseSlot(
		throwClause doc.ThrowClauseLike,
		slot_ uint,
	)
	PostprocessThrowClause(
		throwClause doc.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessWhileClause(
		whileClause doc.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWhileClauseSlot(
		whileClause doc.WhileClauseLike,
		slot_ uint,
	)
	PostprocessWhileClause(
		whileClause doc.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessWithClause(
		withClause doc.WithClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWithClauseSlot(
		withClause doc.WithClauseLike,
		slot_ uint,
	)
	PostprocessWithClause(
		withClause doc.WithClauseLike,
		index_ uint,
		count_ uint,
	)
}
