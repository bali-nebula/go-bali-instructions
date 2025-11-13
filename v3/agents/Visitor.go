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

package agents

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document doc.DocumentLike,
) {
	v.processor_.PreprocessDocument(
		document,
		0,
		0,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		0,
		0,
	)
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause doc.AcceptClauseLike,
) {
	var message = acceptClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		0,
		0,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessAcceptClauseSlot(
		acceptClause,
		1,
	)

	var bag = acceptClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		0,
		0,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitArgument(
	argument any,
) {
	switch actual := argument.(type) {
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		v.visitPrimitive(argument)
	}
}

func (v *visitor_) visitAttributes(
	attributes doc.AttributesLike,
) {
	var associationsIndex uint
	var associations = attributes.GetAssociations().GetIterator()
	var associationsCount = uint(associations.GetSize())
	for associations.HasNext() {
		associationsIndex++
		var association = associations.GetNext()
		var primitive = association.GetKey()
		v.visitPrimitive(primitive)
		var content = association.GetValue()
		v.processor_.PreprocessContent(
			content,
			associationsIndex,
			associationsCount,
		)
		v.visitContent(content)
		v.processor_.PostprocessContent(
			content,
			associationsIndex,
			associationsCount,
		)
	}
}

func (v *visitor_) visitBreakClause(
	breakClause doc.BreakClauseLike,
) {
}

func (v *visitor_) visitCheckoutClause(
	checkoutClause doc.CheckoutClauseLike,
) {
	var recipient = checkoutClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		1,
	)

	var optionalAtLevel = checkoutClause.GetOptionalAtLevel()
	if uti.IsDefined(optionalAtLevel) {
		v.processor_.PreprocessExpression(
			optionalAtLevel,
			0,
			0,
		)
		v.visitExpression(optionalAtLevel)
		v.processor_.PostprocessExpression(
			optionalAtLevel,
			0,
			0,
		)
	}

	// Visit slot 2 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		2,
	)

	var location = checkoutClause.GetLocation()
	v.processor_.PreprocessExpression(
		location,
		0,
		0,
	)
	v.visitExpression(location)
	v.processor_.PostprocessExpression(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitComplement(
	complement doc.ComplementLike,
) {
	var reversible = complement.GetReversible()
	v.processor_.PreprocessReversible(
		reversible,
		0,
		0,
	)
	v.visitReversible(reversible)
	v.processor_.PostprocessReversible(
		reversible,
		0,
		0,
	)
}

func (v *visitor_) visitComposite(
	composite doc.Composite,
) {
	var entity = composite.GetEntity()
	v.processor_.PreprocessEntity(
		entity,
		0,
		0,
	)
	v.visitEntity(entity)
	v.processor_.PostprocessEntity(
		entity,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessCompositeSlot(
		composite,
		1,
	)

	var optionalGenerics = composite.GetOptionalGenerics()
	if uti.IsDefined(optionalGenerics) {
		v.processor_.PreprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
		v.visitGenerics(optionalGenerics)
		v.processor_.PostprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
	}
}

func (v *visitor_) visitContent(
	content doc.ContentLike,
) {
	var composite = content.GetComposite()
	v.processor_.PreprocessComposite(
		composite,
		0,
		0,
	)
	v.visitComposite(composite)
	v.processor_.PostprocessComposite(
		composite,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessContentSlot(
		content,
		1,
	)

	var optionalNote = content.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitConstraint(
	constraint doc.ConstraintLike,
) {
	var metadata = constraint.GetMetadata()
	v.processor_.PreprocessMetadata(
		metadata,
		0,
		0,
	)
	v.visitMetadata(metadata)
	v.processor_.PostprocessMetadata(
		metadata,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessConstraintSlot(
		constraint,
		1,
	)

	var optionalGenerics = constraint.GetOptionalGenerics()
	if uti.IsDefined(optionalGenerics) {
		v.processor_.PreprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
		v.visitGenerics(optionalGenerics)
		v.processor_.PostprocessGenerics(
			optionalGenerics,
			0,
			0,
		)
	}
}

func (v *visitor_) visitContinueClause(
	continueClause doc.ContinueClauseLike,
) {
}

func (v *visitor_) visitDiscardClause(
	discardClause doc.DiscardClauseLike,
) {
	var citation = discardClause.GetCitation()
	v.processor_.PreprocessExpression(
		citation,
		0,
		0,
	)
	v.visitExpression(citation)
	v.processor_.PostprocessExpression(
		citation,
		0,
		0,
	)
}

func (v *visitor_) visitDoClause(
	doClause doc.DoClauseLike,
) {
	var method = doClause.GetMethod()
	v.processor_.PreprocessMethod(
		method,
		0,
		0,
	)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(
		method,
		0,
		0,
	)
}

func (v *visitor_) visitDocument(
	document doc.DocumentLike,
) {
	var comment = document.GetOptionalComment()
	if uti.IsDefined(comment) {
		v.processor_.ProcessComment(comment)
	}
	var composite = document.GetComposite()
	v.processor_.PreprocessComposite(
		composite,
		0,
		0,
	)
	v.visitComposite(composite)
	v.processor_.PostprocessComposite(
		composite,
		0,
		0,
	)
}

func (v *visitor_) visitEntity(
	entity any,
) {
	switch actual := entity.(type) {
	case doc.RangeLike:
		v.processor_.PreprocessRange(
			actual,
			0,
			0,
		)
		v.visitRange(actual)
		v.processor_.PostprocessRange(
			actual,
			0,
			0,
		)
	case doc.AttributesLike:
		v.processor_.PreprocessAttributes(
			actual,
			0,
			0,
		)
		v.visitAttributes(actual)
		v.processor_.PostprocessAttributes(
			actual,
			0,
			0,
		)
	case doc.ItemsLike:
		v.processor_.PreprocessItems(
			actual,
			0,
			0,
		)
		v.visitItems(actual)
		v.processor_.PostprocessItems(
			actual,
			0,
			0,
		)
	case doc.ProcedureLike:
		v.processor_.PreprocessProcedure(
			actual,
			0,
			0,
		)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(
			actual,
			0,
			0,
		)
	default:
		v.processor_.PreprocessPrimitive(
			entity,
			0,
			0,
		)
		v.visitPrimitive(entity)
		v.processor_.PostprocessPrimitive(
			entity,
			0,
			0,
		)
	}
}

func (v *visitor_) visitExpression(
	expression doc.ExpressionLike,
) {
	var subject = expression.GetSubject()
	v.processor_.PreprocessSubject(
		subject,
		0,
		0,
	)
	v.visitSubject(subject)
	v.processor_.PostprocessSubject(
		subject,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessExpressionSlot(
		expression,
		1,
	)

	var predicatesIndex uint
	var predicates = expression.GetPredicates().GetIterator()
	var predicatesCount = uint(predicates.GetSize())
	for predicates.HasNext() {
		predicatesIndex++
		var predicate = predicates.GetNext()
		v.processor_.PreprocessPredicate(
			predicate,
			predicatesIndex,
			predicatesCount,
		)
		v.visitPredicate(predicate)
		v.processor_.PostprocessPredicate(
			predicate,
			predicatesIndex,
			predicatesCount,
		)
	}
}

func (v *visitor_) visitFunction(
	function doc.FunctionLike,
) {
	var identifier = function.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionSlot(
		function,
		1,
	)

	var argumentsIndex uint
	var arguments = function.GetArguments().GetIterator()
	var argumentsCount = uint(arguments.GetSize())
	for arguments.HasNext() {
		argumentsIndex++
		var argument = arguments.GetNext()
		v.processor_.PreprocessArgument(
			argument,
			argumentsIndex,
			argumentsCount,
		)
		v.visitArgument(argument)
		v.processor_.PostprocessArgument(
			argument,
			argumentsIndex,
			argumentsCount,
		)
	}
}

func (v *visitor_) visitGenerics(
	generics doc.GenericsLike,
) {
	var parametersIndex uint
	var parameters = generics.GetParameters().GetIterator()
	var parametersCount = uint(parameters.GetSize())
	for parameters.HasNext() {
		parametersIndex++
		var parameter = parameters.GetNext()
		var symbol = parameter.GetKey()
		v.processor_.ProcessSymbol(symbol)
		var constraint = parameter.GetValue()
		v.processor_.PreprocessConstraint(
			constraint,
			parametersIndex,
			parametersCount,
		)
		v.visitConstraint(constraint)
		v.processor_.PostprocessConstraint(
			constraint,
			parametersIndex,
			parametersCount,
		)
	}
}

func (v *visitor_) visitIfClause(
	ifClause doc.IfClauseLike,
) {
	var condition = ifClause.GetCondition()
	v.processor_.PreprocessExpression(
		condition,
		0,
		0,
	)
	v.visitExpression(condition)
	v.processor_.PostprocessExpression(
		condition,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessIfClauseSlot(
		ifClause,
		1,
	)

	var procedure = ifClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitIndex(
	index any,
) {
	switch actual := index.(type) {
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		v.processor_.PreprocessPrimitive(
			index,
			0,
			0,
		)
		v.visitPrimitive(index)
		v.processor_.PostprocessPrimitive(
			index,
			0,
			0,
		)
	}
}

func (v *visitor_) visitInspectClause(
	inspectClause doc.InspectClauseLike,
) {
	var recipient = inspectClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessInspectClauseSlot(
		inspectClause,
		1,
	)

	var location = inspectClause.GetLocation()
	v.processor_.PreprocessExpression(
		location,
		0,
		0,
	)
	v.visitExpression(location)
	v.processor_.PostprocessExpression(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitInversion(
	inversion doc.InversionLike,
) {
	var inverse = inversion.GetInverse()
	v.processor_.ProcessInverse(inverse)

	// Visit slot 1 between terms.
	v.processor_.ProcessInversionSlot(
		inversion,
		1,
	)

	var numerical = inversion.GetNumerical()
	v.processor_.PreprocessNumerical(
		numerical,
		0,
		0,
	)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(
		numerical,
		0,
		0,
	)
}

func (v *visitor_) visitItems(
	items doc.ItemsLike,
) {
	var contentsIndex uint
	var contents = items.GetContents().GetIterator()
	var contentsCount = uint(contents.GetSize())
	for contents.HasNext() {
		contentsIndex++
		var content = contents.GetNext()
		v.processor_.PreprocessContent(
			content,
			contentsIndex,
			contentsCount,
		)
		v.visitContent(content)
		v.processor_.PostprocessContent(
			content,
			contentsIndex,
			contentsCount,
		)
	}
}

func (v *visitor_) visitLetClause(
	letClause doc.LetClauseLike,
) {
	var recipient = letClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		1,
	)

	var assignment = letClause.GetAssignment()
	v.processor_.ProcessAssignment(assignment)

	// Visit slot 2 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		2,
	)

	var expression = letClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitLine(
	line any,
) {
	switch actual := line.(type) {
	case doc.StatementLike:
		v.processor_.PreprocessStatement(
			actual,
			0,
			0,
		)
		v.visitStatement(actual)
		v.processor_.PostprocessStatement(
			actual,
			0,
			0,
		)
	case string:
		v.processor_.ProcessAnnotation(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitMagnitude(
	magnitude doc.MagnitudeLike,
) {
	var expression = magnitude.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitMainClause(
	mainClause any,
) {
	switch actual := mainClause.(type) {
	case doc.IfClauseLike:
		v.processor_.PreprocessIfClause(
			actual,
			0,
			0,
		)
		v.visitIfClause(actual)
		v.processor_.PostprocessIfClause(
			actual,
			0,
			0,
		)
	case doc.SelectClauseLike:
		v.processor_.PreprocessSelectClause(
			actual,
			0,
			0,
		)
		v.visitSelectClause(actual)
		v.processor_.PostprocessSelectClause(
			actual,
			0,
			0,
		)
	case doc.WhileClauseLike:
		v.processor_.PreprocessWhileClause(
			actual,
			0,
			0,
		)
		v.visitWhileClause(actual)
		v.processor_.PostprocessWhileClause(
			actual,
			0,
			0,
		)
	case doc.WithClauseLike:
		v.processor_.PreprocessWithClause(
			actual,
			0,
			0,
		)
		v.visitWithClause(actual)
		v.processor_.PostprocessWithClause(
			actual,
			0,
			0,
		)
	case doc.ContinueClauseLike:
		v.processor_.PreprocessContinueClause(
			actual,
			0,
			0,
		)
		v.visitContinueClause(actual)
		v.processor_.PostprocessContinueClause(
			actual,
			0,
			0,
		)
	case doc.BreakClauseLike:
		v.processor_.PreprocessBreakClause(
			actual,
			0,
			0,
		)
		v.visitBreakClause(actual)
		v.processor_.PostprocessBreakClause(
			actual,
			0,
			0,
		)
	case doc.ReturnClauseLike:
		v.processor_.PreprocessReturnClause(
			actual,
			0,
			0,
		)
		v.visitReturnClause(actual)
		v.processor_.PostprocessReturnClause(
			actual,
			0,
			0,
		)
	case doc.ThrowClauseLike:
		v.processor_.PreprocessThrowClause(
			actual,
			0,
			0,
		)
		v.visitThrowClause(actual)
		v.processor_.PostprocessThrowClause(
			actual,
			0,
			0,
		)
	case doc.SendClauseLike:
		v.processor_.PreprocessSendClause(
			actual,
			0,
			0,
		)
		v.visitSendClause(actual)
		v.processor_.PostprocessSendClause(
			actual,
			0,
			0,
		)
	case doc.ReceiveClauseLike:
		v.processor_.PreprocessReceiveClause(
			actual,
			0,
			0,
		)
		v.visitReceiveClause(actual)
		v.processor_.PostprocessReceiveClause(
			actual,
			0,
			0,
		)
	case doc.AcceptClauseLike:
		v.processor_.PreprocessAcceptClause(
			actual,
			0,
			0,
		)
		v.visitAcceptClause(actual)
		v.processor_.PostprocessAcceptClause(
			actual,
			0,
			0,
		)
	case doc.RejectClauseLike:
		v.processor_.PreprocessRejectClause(
			actual,
			0,
			0,
		)
		v.visitRejectClause(actual)
		v.processor_.PostprocessRejectClause(
			actual,
			0,
			0,
		)
	case doc.PublishClauseLike:
		v.processor_.PreprocessPublishClause(
			actual,
			0,
			0,
		)
		v.visitPublishClause(actual)
		v.processor_.PostprocessPublishClause(
			actual,
			0,
			0,
		)
	case doc.SaveClauseLike:
		v.processor_.PreprocessSaveClause(
			actual,
			0,
			0,
		)
		v.visitSaveClause(actual)
		v.processor_.PostprocessSaveClause(
			actual,
			0,
			0,
		)
	case doc.RetrieveClauseLike:
		v.processor_.PreprocessRetrieveClause(
			actual,
			0,
			0,
		)
		v.visitRetrieveClause(actual)
		v.processor_.PostprocessRetrieveClause(
			actual,
			0,
			0,
		)
	case doc.DiscardClauseLike:
		v.processor_.PreprocessDiscardClause(
			actual,
			0,
			0,
		)
		v.visitDiscardClause(actual)
		v.processor_.PostprocessDiscardClause(
			actual,
			0,
			0,
		)
	case doc.NotarizeClauseLike:
		v.processor_.PreprocessNotarizeClause(
			actual,
			0,
			0,
		)
		v.visitNotarizeClause(actual)
		v.processor_.PostprocessNotarizeClause(
			actual,
			0,
			0,
		)
	case doc.InspectClauseLike:
		v.processor_.PreprocessInspectClause(
			actual,
			0,
			0,
		)
		v.visitInspectClause(actual)
		v.processor_.PostprocessInspectClause(
			actual,
			0,
			0,
		)
	case doc.CheckoutClauseLike:
		v.processor_.PreprocessCheckoutClause(
			actual,
			0,
			0,
		)
		v.visitCheckoutClause(actual)
		v.processor_.PostprocessCheckoutClause(
			actual,
			0,
			0,
		)
	case doc.DoClauseLike:
		v.processor_.PreprocessDoClause(
			actual,
			0,
			0,
		)
		v.visitDoClause(actual)
		v.processor_.PostprocessDoClause(
			actual,
			0,
			0,
		)
	case doc.LetClauseLike:
		v.processor_.PreprocessLetClause(
			actual,
			0,
			0,
		)
		v.visitLetClause(actual)
		v.processor_.PostprocessLetClause(
			actual,
			0,
			0,
		)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitMatchingClause(
	matchingClause doc.MatchingClauseLike,
) {
	var template = matchingClause.GetTemplate()
	v.processor_.PreprocessExpression(
		template,
		0,
		0,
	)
	v.visitExpression(template)
	v.processor_.PostprocessExpression(
		template,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		1,
	)

	var procedure = matchingClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitMetadata(
	metadata any,
) {
	switch actual := metadata.(type) {
	case doc.RangeLike:
		v.processor_.PreprocessRange(
			actual,
			0,
			0,
		)
		v.visitRange(actual)
		v.processor_.PostprocessRange(
			actual,
			0,
			0,
		)
	default:
		v.processor_.PreprocessPrimitive(
			metadata,
			0,
			0,
		)
		v.visitPrimitive(metadata)
		v.processor_.PostprocessPrimitive(
			metadata,
			0,
			0,
		)
	}
}

func (v *visitor_) visitMethod(
	method doc.MethodLike,
) {
	var target = method.GetTarget()
	v.processor_.ProcessIdentifier(target)

	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		1,
	)

	var invoke = method.GetInvoke()
	v.processor_.ProcessInvoke(invoke)

	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		2,
	)

	var identifier = method.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		3,
	)

	var argumentsIndex uint
	var arguments = method.GetArguments().GetIterator()
	var argumentsCount = uint(arguments.GetSize())
	for arguments.HasNext() {
		argumentsIndex++
		var argument = arguments.GetNext()
		v.processor_.PreprocessArgument(
			argument,
			argumentsIndex,
			argumentsCount,
		)
		v.visitArgument(argument)
		v.processor_.PostprocessArgument(
			argument,
			argumentsIndex,
			argumentsCount,
		)
	}
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
) {
	var draft = notarizeClause.GetDraft()
	v.processor_.PreprocessExpression(
		draft,
		0,
		0,
	)
	v.visitExpression(draft)
	v.processor_.PostprocessExpression(
		draft,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		1,
	)

	var location = notarizeClause.GetLocation()
	v.processor_.PreprocessExpression(
		location,
		0,
		0,
	)
	v.visitExpression(location)
	v.processor_.PostprocessExpression(
		location,
		0,
		0,
	)
}

func (v *visitor_) visitNumerical(
	numerical any,
) {
	switch actual := numerical.(type) {
	case doc.Composite:
		v.processor_.PreprocessComposite(
			actual,
			0,
			0,
		)
		v.visitComposite(actual)
		v.processor_.PostprocessComposite(
			actual,
			0,
			0,
		)
	case doc.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case doc.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case doc.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case doc.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			0,
			0,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			0,
			0,
		)
	case doc.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			0,
			0,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			0,
			0,
		)
	case doc.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case doc.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitOnClause(
	onClause doc.OnClauseLike,
) {
	var symbol = onClause.GetSymbol()
	v.processor_.ProcessSymbol(symbol)

	// Visit slot 1 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		1,
	)

	var matchingClausesIndex uint
	var matchingClauses = onClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var clause = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(clause)
		v.processor_.PostprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitPrecedence(
	precedence doc.PrecedenceLike,
) {
	var expression = precedence.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitPredicate(
	predicate doc.PredicateLike,
) {
	var operator = predicate.GetOperator()
	v.processor_.ProcessOperator(operator)

	// Visit slot 1 between terms.
	v.processor_.ProcessPredicateSlot(
		predicate,
		1,
	)

	var expression = predicate.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		0,
		0,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		0,
		0,
	)
}

func (v *visitor_) visitPrimitive(
	primitive any,
) {
	switch actual := primitive.(type) {
	case pri.AngleLike:
		v.processor_.ProcessAngle(actual)
	case pri.BooleanLike:
		v.processor_.ProcessBoolean(actual)
	case pri.DurationLike:
		v.processor_.ProcessDuration(actual)
	case pri.GlyphLike:
		v.processor_.ProcessGlyph(actual)
	case pri.MomentLike:
		v.processor_.ProcessMoment(actual)
	case pri.NumberLike:
		v.processor_.ProcessNumber(actual)
	case pri.PercentageLike:
		v.processor_.ProcessPercentage(actual)
	case pri.ProbabilityLike:
		v.processor_.ProcessProbability(actual)
	case pri.ResourceLike:
		v.processor_.ProcessResource(actual)
	case pri.SymbolLike:
		v.processor_.ProcessSymbol(actual)
	case pri.BinaryLike:
		v.processor_.ProcessBinary(actual)
	case pri.BytecodeLike:
		v.processor_.ProcessBytecode(actual)
	case pri.NameLike:
		v.processor_.ProcessName(actual)
	case pri.NarrativeLike:
		v.processor_.ProcessNarrative(actual)
	case pri.PatternLike:
		v.processor_.ProcessPattern(actual)
	case pri.QuoteLike:
		v.processor_.ProcessQuote(actual)
	case pri.TagLike:
		v.processor_.ProcessTag(actual)
	case pri.VersionLike:
		v.processor_.ProcessVersion(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitProcedure(
	procedure doc.ProcedureLike,
) {
	var linesIndex uint
	var lines = procedure.GetLines().GetIterator()
	var linesCount = uint(lines.GetSize())
	for lines.HasNext() {
		linesIndex++
		var line = lines.GetNext()
		v.processor_.PreprocessLine(
			line,
			linesIndex,
			linesCount,
		)
		v.visitLine(line)
		v.processor_.PostprocessLine(
			line,
			linesIndex,
			linesCount,
		)
	}
}

func (v *visitor_) visitPublishClause(
	publishClause doc.PublishClauseLike,
) {
	var message = publishClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		0,
		0,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		0,
		0,
	)
}

func (v *visitor_) visitRange(
	range_ doc.RangeLike,
) {
	var left = range_.GetLeft()
	v.processor_.ProcessBracket(left)

	// Visit slot 1 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		1,
	)

	var first = range_.GetFirst()
	v.processor_.PreprocessPrimitive(
		first,
		0,
		0,
	)
	v.visitPrimitive(first)
	v.processor_.PostprocessPrimitive(
		first,
		0,
		0,
	)

	// Visit slot 2 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		2,
	)

	var last = range_.GetLast()
	v.processor_.PreprocessPrimitive(
		last,
		0,
		0,
	)
	v.visitPrimitive(last)
	v.processor_.PostprocessPrimitive(
		last,
		0,
		0,
	)

	// Visit slot 3 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		3,
	)

	var right = range_.GetRight()
	v.processor_.ProcessBracket(right)
}

func (v *visitor_) visitReceiveClause(
	receiveClause doc.ReceiveClauseLike,
) {
	var recipient = receiveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessReceiveClauseSlot(
		receiveClause,
		1,
	)

	var bag = receiveClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		0,
		0,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitRecipient(
	recipient any,
) {
	switch actual := recipient.(type) {
	case doc.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case pri.SymbolLike:
		v.processor_.ProcessSymbol(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitReference(
	reference any,
) {
	switch actual := reference.(type) {
	case doc.Composite:
		v.processor_.PreprocessComposite(
			actual,
			0,
			0,
		)
		v.visitComposite(actual)
		v.processor_.PostprocessComposite(
			actual,
			0,
			0,
		)
	case doc.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case doc.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case doc.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case doc.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitReferent(
	referent doc.ReferentLike,
) {
	var reference = referent.GetReference()
	v.processor_.PreprocessReference(
		reference,
		0,
		0,
	)
	v.visitReference(reference)
	v.processor_.PostprocessReference(
		reference,
		0,
		0,
	)
}

func (v *visitor_) visitRejectClause(
	rejectClause doc.RejectClauseLike,
) {
	var message = rejectClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		0,
		0,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		1,
	)

	var bag = rejectClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		0,
		0,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitRetrieveClause(
	retrieveClause doc.RetrieveClauseLike,
) {
	var recipient = retrieveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		1,
	)

	var citation = retrieveClause.GetCitation()
	v.processor_.PreprocessExpression(
		citation,
		0,
		0,
	)
	v.visitExpression(citation)
	v.processor_.PostprocessExpression(
		citation,
		0,
		0,
	)
}

func (v *visitor_) visitReturnClause(
	returnClause doc.ReturnClauseLike,
) {
	var result = returnClause.GetResult()
	v.processor_.PreprocessExpression(
		result,
		0,
		0,
	)
	v.visitExpression(result)
	v.processor_.PostprocessExpression(
		result,
		0,
		0,
	)
}

func (v *visitor_) visitReversible(
	reversible any,
) {
	switch actual := reversible.(type) {
	case doc.Composite:
		v.processor_.PreprocessComposite(
			actual,
			0,
			0,
		)
		v.visitComposite(actual)
		v.processor_.PostprocessComposite(
			actual,
			0,
			0,
		)
	case doc.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case doc.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case doc.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case doc.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			0,
			0,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			0,
			0,
		)
	case doc.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case doc.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitSaveClause(
	saveClause doc.SaveClauseLike,
) {
	var draft = saveClause.GetDraft()
	v.processor_.PreprocessExpression(
		draft,
		0,
		0,
	)
	v.visitExpression(draft)
	v.processor_.PostprocessExpression(
		draft,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		1,
	)

	var recipient = saveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		0,
		0,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		0,
		0,
	)
}

func (v *visitor_) visitSelectClause(
	selectClause doc.SelectClauseLike,
) {
	var template = selectClause.GetTemplate()
	v.processor_.PreprocessExpression(
		template,
		0,
		0,
	)
	v.visitExpression(template)
	v.processor_.PostprocessExpression(
		template,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		1,
	)

	var matchingClausesIndex uint
	var matchingClauses = selectClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var clause = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(clause)
		v.processor_.PostprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitSendClause(
	sendClause doc.SendClauseLike,
) {
	var message = sendClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		0,
		0,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessSendClauseSlot(
		sendClause,
		1,
	)

	var bag = sendClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		0,
		0,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		0,
		0,
	)
}

func (v *visitor_) visitStatement(
	statement doc.StatementLike,
) {
	var mainClause = statement.GetMainClause()
	v.processor_.PreprocessMainClause(
		mainClause,
		0,
		0,
	)
	v.visitMainClause(mainClause)
	v.processor_.PostprocessMainClause(
		mainClause,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessStatementSlot(
		statement,
		1,
	)

	var optionalOnClause = statement.GetOptionalOnClause()
	if uti.IsDefined(optionalOnClause) {
		v.processor_.PreprocessOnClause(
			optionalOnClause,
			0,
			0,
		)
		v.visitOnClause(optionalOnClause)
		v.processor_.PostprocessOnClause(
			optionalOnClause,
			0,
			0,
		)
	}
}

func (v *visitor_) visitSubcomponent(
	subcomponent doc.SubcomponentLike,
) {
	var identifier = subcomponent.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 1 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		1,
	)

	var indexesIndex uint
	var indexes = subcomponent.GetIndexes().GetIterator()
	var indexesCount = uint(indexes.GetSize())
	for indexes.HasNext() {
		indexesIndex++
		var index = indexes.GetNext()
		v.processor_.PreprocessIndex(
			index,
			indexesIndex,
			indexesCount,
		)
		v.visitIndex(index)
		v.processor_.PostprocessIndex(
			index,
			indexesIndex,
			indexesCount,
		)
	}
}

func (v *visitor_) visitSubject(
	subject any,
) {
	switch actual := subject.(type) {
	case doc.Composite:
		v.processor_.PreprocessComposite(
			actual,
			0,
			0,
		)
		v.visitComposite(actual)
		v.processor_.PostprocessComposite(
			actual,
			0,
			0,
		)
	case doc.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			0,
			0,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			0,
			0,
		)
	case doc.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			0,
			0,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			0,
			0,
		)
	case doc.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			0,
			0,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			0,
			0,
		)
	case doc.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			0,
			0,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			0,
			0,
		)
	case doc.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			0,
			0,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			0,
			0,
		)
	case doc.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			0,
			0,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			0,
			0,
		)
	case doc.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			0,
			0,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			0,
			0,
		)
	case doc.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			0,
			0,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			0,
			0,
		)
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
}

func (v *visitor_) visitThrowClause(
	throwClause doc.ThrowClauseLike,
) {
	var exception = throwClause.GetException()
	v.processor_.PreprocessExpression(
		exception,
		0,
		0,
	)
	v.visitExpression(exception)
	v.processor_.PostprocessExpression(
		exception,
		0,
		0,
	)
}

func (v *visitor_) visitWhileClause(
	whileClause doc.WhileClauseLike,
) {
	var condition = whileClause.GetCondition()
	v.processor_.PreprocessExpression(
		condition,
		0,
		0,
	)
	v.visitExpression(condition)
	v.processor_.PostprocessExpression(
		condition,
		0,
		0,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		1,
	)

	var procedure = whileClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

func (v *visitor_) visitWithClause(
	withClause doc.WithClauseLike,
) {
	var symbol = withClause.GetSymbol()
	v.processor_.ProcessSymbol(symbol)

	// Visit slot 1 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		1,
	)

	var sequence = withClause.GetSequence()
	v.processor_.PreprocessExpression(
		sequence,
		0,
		0,
	)
	v.visitExpression(sequence)
	v.processor_.PostprocessExpression(
		sequence,
		0,
		0,
	)

	// Visit slot 2 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		2,
	)

	var procedure = withClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		0,
		0,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		0,
		0,
	)
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
