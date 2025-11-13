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
	ins "github.com/bali-nebula/go-bali-instructions/v3/instructions"
	lan "github.com/bali-nebula/go-assembly-language/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func InflatorClass() InflatorClassLike {
	return inflatorClass()
}

// Constructor Methods

func (c *inflatorClass_) Inflator() InflatorLike {
	var instance = &inflator_{
		// Initialize the instance attributes.
		stack_: com.StackWithCapacity[any](256),

		// Initialize the inherited aspects.
		Methodical: not.ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *inflator_) GetClass() InflatorClassLike {
	return inflatorClass()
}

func (v *inflator_) InflateDocument(
	document not.DocumentLike,
) doc.DocumentLike {
	not.VisitorClass().Visitor(v).VisitDocument(document)
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the inflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
	return v.stack_.RemoveLast().(doc.DocumentLike)
}

// Attribute Methods

// not.Methodical Methods

func (v *inflator_) ProcessAngle(
	angle string,
) {
	v.stack_.AddValue(pri.AngleClass().AngleFromSource(angle))
}

func (v *inflator_) ProcessBinary(
	binary string,
) {
	v.stack_.AddValue(pri.BinaryClass().BinaryFromSource(binary))
}

func (v *inflator_) ProcessBoolean(
	boolean string,
) {
	v.stack_.AddValue(pri.BooleanClass().BooleanFromSource(boolean))
}

func (v *inflator_) ProcessBytecode(
	bytecode string,
) {
	v.stack_.AddValue(pri.BytecodeClass().BytecodeFromSource(bytecode))
}

func (v *inflator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *inflator_) ProcessDuration(
	duration string,
) {
	v.stack_.AddValue(pri.DurationClass().DurationFromSource(duration))
}

func (v *inflator_) ProcessGlyph(
	glyph string,
) {
	v.stack_.AddValue(pri.GlyphClass().GlyphFromSource(glyph))
}

func (v *inflator_) ProcessIdentifier(
	identifier string,
) {
	v.stack_.AddValue(identifier)
}

func (v *inflator_) ProcessMoment(
	moment string,
) {
	v.stack_.AddValue(pri.MomentClass().MomentFromSource(moment))
}

func (v *inflator_) ProcessName(
	name string,
) {
	v.stack_.AddValue(pri.NameClass().NameFromSource(name))
}

func (v *inflator_) ProcessNarrative(
	narrative string,
) {
	v.stack_.AddValue(pri.NarrativeClass().NarrativeFromSource(narrative))
}

func (v *inflator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *inflator_) ProcessNumber(
	number string,
) {
	v.stack_.AddValue(pri.NumberClass().NumberFromSource(number))
}

func (v *inflator_) ProcessPattern(
	pattern string,
) {
	v.stack_.AddValue(pri.PatternClass().PatternFromSource(pattern))
}

func (v *inflator_) ProcessPercentage(
	percentage string,
) {
	v.stack_.AddValue(pri.PercentageClass().PercentageFromSource(percentage))
}

func (v *inflator_) ProcessProbability(
	probability string,
) {
	v.stack_.AddValue(pri.ProbabilityClass().ProbabilityFromSource(probability))
}

func (v *inflator_) ProcessQuote(
	quote string,
) {
	v.stack_.AddValue(pri.QuoteClass().QuoteFromSource(quote))
}

func (v *inflator_) ProcessResource(
	resource string,
) {
	v.stack_.AddValue(pri.ResourceClass().ResourceFromSource(resource))
}

func (v *inflator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(pri.SymbolClass().SymbolFromSource(symbol))
}

func (v *inflator_) ProcessTag(
	tag string,
) {
	v.stack_.AddValue(pri.TagClass().TagFromSource(tag))
}

func (v *inflator_) ProcessVersion(
	version string,
) {
	v.stack_.AddValue(pri.VersionClass().VersionFromSource(version))
}

func (v *inflator_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.AcceptClauseClass().AcceptClause(message, bag))
}

func (v *inflator_) PostprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	var operator = assignment.GetAny().(string)
	switch operator {
	case "?=":
		v.stack_.AddValue(doc.DefaultEquals)
	case ":=":
		v.stack_.AddValue(doc.AssignEquals)
	case "+=":
		v.stack_.AddValue(doc.PlusEquals)
	case "-=":
		v.stack_.AddValue(doc.MinusEquals)
	case "*=":
		v.stack_.AddValue(doc.TimesEquals)
	case "/=":
		v.stack_.AddValue(doc.DivideEquals)
	case "&=":
		v.stack_.AddValue(doc.ChainEquals)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operator,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var catalog = com.Catalog[any, doc.ContentLike]()
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var content = v.stack_.RemoveLast().(doc.ContentLike)
		var primitive = v.stack_.RemoveLast()
		catalog.SetValue(primitive, content)
		iterator.GetNext()
	}
	catalog.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.AttributesClass().Attributes(catalog))
}

func (v *inflator_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.BreakClauseClass().BreakClause())
}

func (v *inflator_) ProcessCheckoutClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		if uti.IsUndefined(checkoutClause.GetOptionalAtLevel()) {
			var atLevel doc.ExpressionLike
			v.stack_.AddValue(atLevel)
		}
	}
}

func (v *inflator_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	var atLevel doc.ExpressionLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		atLevel = optional.(doc.ExpressionLike)
	}
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.CheckoutClauseClass().CheckoutClause(recipient, atLevel, location),
	)
}

func (v *inflator_) PostprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	var logical = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ComplementClass().Complement(logical))
}

func (v *inflator_) ProcessComponentSlot(
	component not.ComponentLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(component.GetOptionalGenerics()) {
			var generics doc.GenericsLike
			v.stack_.AddValue(generics)
		}
	}
}

func (v *inflator_) PostprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var generics doc.GenericsLike
	var optional = v.stack_.RemoveLast()
	var entity = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		generics = optional.(doc.GenericsLike)
		switch actual := entity.(type) {
		case doc.ItemsLike:
			var contents = v.getContents(actual, generics)
			entity = doc.ItemsClass().Items(contents)
		}
	}
	v.stack_.AddValue(doc.ComponentClass().Component(entity, generics))
}

func (v *inflator_) ProcessContentSlot(
	content not.ContentLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(content.GetOptionalNote()) {
			var note string
			v.stack_.AddValue(note)
		}
	}
}

func (v *inflator_) PostprocessContent(
	content not.ContentLike,
	index_ uint,
	count_ uint,
) {
	var note = v.stack_.RemoveLast().(string)
	var composite = v.stack_.RemoveLast().(doc.Composite)
	v.stack_.AddValue(
		doc.ContentClass().Content(
			composite,
			note,
		),
	)
}

func (v *inflator_) ProcessConstraintSlot(
	constraint not.ConstraintLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(constraint.GetOptionalGenerics()) {
			var generics doc.GenericsLike
			v.stack_.AddValue(generics)
		}
	}
}

func (v *inflator_) PostprocessConstraint(
	constraint not.ConstraintLike,
	index_ uint,
	count_ uint,
) {
	var generics doc.GenericsLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		generics = optional.(doc.GenericsLike)
	}
	var metadata = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ConstraintClass().Constraint(metadata, generics))
}

func (v *inflator_) PostprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.ContinueClauseClass().ContinueClause())
}

func (v *inflator_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.DiscardClauseClass().DiscardClause(location))
}

func (v *inflator_) PostprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(doc.MethodLike)
	v.stack_.AddValue(doc.DoClauseClass().DoClause(method))
}

func (v *inflator_) PreprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var comment string
	var heading = document.GetOptionalHeading()
	if uti.IsUndefined(heading) {
		// We only add it if it is not defined, otherwise ProcessComment adds it.
		v.stack_.AddValue(comment)
	}
}

func (v *inflator_) PostprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var composite = v.stack_.RemoveLast().(doc.Composite)
	var comment = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(doc.DocumentClass().Document(comment, composite))
}

func (v *inflator_) PostprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[doc.PredicateLike]()
	var predicates = expression.GetPredicates()
	var iterator = predicates.GetIterator()
	for iterator.HasNext() {
		var predicate = v.stack_.RemoveLast().(doc.PredicateLike)
		list.AppendValue(predicate)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var subject = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ExpressionClass().Expression(subject, list))
}

func (v *inflator_) PostprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var arguments = function.GetArguments()
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = v.stack_.RemoveLast()
		list.AppendValue(argument)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(doc.FunctionClass().Function(identifier, list))
}

func (v *inflator_) PostprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.IfClauseClass().IfClause(condition, procedure))
}

func (v *inflator_) PostprocessInverse(
	inverse not.InverseLike,
	index_ uint,
	count_ uint,
) {
	var operator = inverse.GetAny().(string)
	switch operator {
	case "-":
		v.stack_.AddValue(doc.Additive)
	case "/":
		v.stack_.AddValue(doc.Multiplicative)
	case "*":
		v.stack_.AddValue(doc.Conjugate)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operator,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	var numerical = v.stack_.RemoveLast()
	var inverse = v.stack_.RemoveLast().(doc.Inverse)
	v.stack_.AddValue(doc.InversionClass().Inversion(inverse, numerical))
}

func (v *inflator_) PostprocessInvoke(
	invoke not.InvokeLike,
	index_ uint,
	count_ uint,
) {
	var operator = invoke.GetAny().(string)
	switch operator {
	case "<-":
		v.stack_.AddValue(doc.Synchronous)
	case "<~":
		v.stack_.AddValue(doc.Asynchronous)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operator,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	var contents = com.List[doc.ContentLike]()
	var iterator = items.GetContents().GetIterator()
	for iterator.HasNext() {
		var content = v.stack_.RemoveLast().(doc.ContentLike)
		contents.AppendValue(content)
		iterator.GetNext()
	}
	contents.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.ItemsClass().Items(contents))
}

func (v *inflator_) PostprocessLeft(
	left not.LeftLike,
	index_ uint,
	count_ uint,
) {
	var bracket = left.GetAny().(string)
	switch bracket {
	case "[":
		v.stack_.AddValue(com.Inclusive)
	case "(":
		v.stack_.AddValue(com.Exclusive)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			bracket,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(doc.Assignment)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.LetClauseClass().LetClause(recipient, assignment, expression),
	)
}

func (v *inflator_) PostprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.MagnitudeClass().Magnitude(expression))
}

func (v *inflator_) PostprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var template = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.MatchingClauseClass().MatchingClause(template, procedure),
	)
}

func (v *inflator_) PostprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var arguments = method.GetArguments()
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = v.stack_.RemoveLast()
		list.AppendValue(argument)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	var invoke = v.stack_.RemoveLast().(doc.Invoke)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		doc.MethodClass().Method(target, invoke, identifier, list),
	)
}

func (v *inflator_) PostprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.NotarizeClauseClass().NotarizeClause(draft, location),
	)
}

func (v *inflator_) PostprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[doc.MatchingClauseLike]()
	var matchingClauses = onClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var symbol = v.stack_.RemoveLast().(pri.SymbolLike)
	v.stack_.AddValue(
		doc.OnClauseClass().OnClause(symbol, list),
	)
}

func (v *inflator_) PostprocessOperator(
	operator not.OperatorLike,
	index_ uint,
	count_ uint,
) {
	var wrapper any
	switch actual := operator.GetAny().(type) {
	case not.ComparisonLike:
		wrapper = actual.GetAny()
	case not.LogicalLike:
		wrapper = actual.GetAny()
	case not.ArithmeticLike:
		wrapper = actual.GetAny()
	case not.LexicalLike:
		wrapper = actual.GetAny()
	default:
		var message = fmt.Sprintf(
			"Found a value of an unexpected type in a switch statement: %v(%T)",
			actual,
			actual,
		)
		panic(message)
	}
	var actual = wrapper.(string)
	switch actual {
	case "<":
		v.stack_.AddValue(doc.Less)
	case "=":
		v.stack_.AddValue(doc.Equal)
	case ">":
		v.stack_.AddValue(doc.More)
	case "is":
		v.stack_.AddValue(doc.Is)
	case "matches":
		v.stack_.AddValue(doc.Matches)
	case "and":
		v.stack_.AddValue(doc.And)
	case "san":
		v.stack_.AddValue(doc.San)
	case "ior":
		v.stack_.AddValue(doc.Ior)
	case "xor":
		v.stack_.AddValue(doc.Xor)
	case "+":
		v.stack_.AddValue(doc.Plus)
	case "-":
		v.stack_.AddValue(doc.Minus)
	case "*":
		v.stack_.AddValue(doc.Times)
	case "/":
		v.stack_.AddValue(doc.Divide)
	case "%":
		v.stack_.AddValue(doc.Remainder)
	case "^":
		v.stack_.AddValue(doc.Power)
	case "&":
		v.stack_.AddValue(doc.Chain)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			actual,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessGenerics(
	generics not.GenericsLike,
	index_ uint,
	count_ uint,
) {
	var catalog = com.Catalog[pri.SymbolLike, doc.ConstraintLike]()
	var parameters = generics.GetParameters()
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var constraint = v.stack_.RemoveLast().(doc.ConstraintLike)
		var symbol = v.stack_.RemoveLast().(pri.SymbolLike)
		catalog.SetValue(symbol, constraint)
		iterator.GetNext()
	}
	catalog.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.GenericsClass().Generics(catalog))
}

func (v *inflator_) PostprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PrecedenceClass().Precedence(expression))
}

func (v *inflator_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var operator = v.stack_.RemoveLast().(doc.Operator)
	v.stack_.AddValue(doc.PredicateClass().Predicate(operator, expression))
}

func (v *inflator_) PostprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var lines = procedure.GetLines()
	var iterator = lines.GetIterator()
	for iterator.HasNext() {
		var line = v.stack_.RemoveLast()
		list.AppendValue(line)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.ProcedureClass().Procedure(list))
}

func (v *inflator_) PostprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var event = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PublishClauseClass().PublishClause(event))
}

func (v *inflator_) PostprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right = v.stack_.RemoveLast().(com.Bracket)
	var last = v.stack_.RemoveLast()
	var first = v.stack_.RemoveLast()
	var left = v.stack_.RemoveLast().(com.Bracket)
	v.stack_.AddValue(doc.RangeClass().Range(left, first, last, right))
}

func (v *inflator_) PostprocessReceiveClause(
	receiveClause not.ReceiveClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ReceiveClauseClass().ReceiveClause(recipient, bag))
}

func (v *inflator_) PostprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	var indirect = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ReferentClass().Referent(indirect))
}

func (v *inflator_) PostprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.RejectClauseClass().RejectClause(message, bag))
}

func (v *inflator_) PostprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var result = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ReturnClauseClass().ReturnClause(result))
}

func (v *inflator_) PostprocessRight(
	right not.RightLike,
	index_ uint,
	count_ uint,
) {
	var bracket = right.GetAny().(string)
	switch bracket {
	case "]":
		v.stack_.AddValue(com.Inclusive)
	case ")":
		v.stack_.AddValue(com.Exclusive)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			bracket,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var recipient = v.stack_.RemoveLast()
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SaveClauseClass().SaveClause(draft, recipient))
}

func (v *inflator_) PostprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[doc.MatchingClauseLike]()
	var matchingClauses = selectClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SelectClauseClass().SelectClause(expression, list))
}

func (v *inflator_) PostprocessSendClause(
	sendClause not.SendClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SendClauseClass().SendClause(message, bag))
}

func (v *inflator_) ProcessStatementSlot(
	statement not.StatementLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(statement.GetOptionalOnClause()) {
			var onClause doc.OnClauseLike
			v.stack_.AddValue(onClause)
		}
	}
}

func (v *inflator_) PostprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	var onClause doc.OnClauseLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		onClause = optional.(doc.OnClauseLike)
	}
	var mainClause = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.StatementClass().Statement(mainClause, onClause))
}

func (v *inflator_) PostprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var indexes = subcomponent.GetIndexes()
	var iterator = indexes.GetIterator()
	for iterator.HasNext() {
		var index = v.stack_.RemoveLast()
		list.AppendValue(index)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(doc.SubcomponentClass().Subcomponent(identifier, list))
}

func (v *inflator_) PostprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var exception = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ThrowClauseClass().ThrowClause(exception))
}

func (v *inflator_) PostprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.WhileClauseClass().WhileClause(condition, procedure))
}

func (v *inflator_) PostprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var sequence = v.stack_.RemoveLast().(doc.ExpressionLike)
	var variable = v.stack_.RemoveLast().(pri.SymbolLike)
	v.stack_.AddValue(
		doc.WithClauseClass().WithClause(variable, sequence, procedure),
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *inflator_) getContents(
	items doc.ItemsLike,
	generics doc.GenericsLike,
) com.Sequential[doc.ContentLike] {
	var contents = items.GetContents()
	var dummy = doc.ComponentClass().Component(
		pri.PatternClass().None(),
		generics,
	)
	var parameter = dummy.GetParameter(pri.SymbolClass().SymbolFromSource("$type"))
	switch entity := parameter.GetEntity().(type) {
	case pri.ResourceLike:
		switch entity.AsSource() {
		case "<bali:/types/collections/Queue:v3>":
			contents = com.QueueFromSequence(contents)
		case "<bali:/types/collections/Set:v3>":
			contents = com.SetFromSequence(contents)
		case "<bali:/types/collections/Stack:v3>":
			contents = com.StackFromSequence(contents)
		}
	}
	return contents
}

// Instance Structure

type inflator_ struct {
	// Declare the instance attributes.
	stack_ com.StackLike[any]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type inflatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inflatorClass() *inflatorClass_ {
	return inflatorClassReference_
}

var inflatorClassReference_ = &inflatorClass_{
	// Initialize the class constants.
}
