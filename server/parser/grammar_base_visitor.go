// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

type BaseGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructStmt(ctx *StructStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructBlock(ctx *StructBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructStmts(ctx *StructStmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructDeclarationVector(ctx *StructDeclarationVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructFunctionWithParams(ctx *StructFunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructCallList(ctx *StructCallListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructCreation(ctx *StructCreationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValueDeclaration(ctx *ValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorOfStructDeclaration(ctx *VectorOfStructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorDeclaration(ctx *VectorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorOfStructCreation(ctx *VectorOfStructCreationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPlusAssignment(ctx *PlusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMinusAssignment(ctx *MinusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorAssignment(ctx *VectorAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorMinusAssignment(ctx *VectorMinusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorPlusAssignment(ctx *VectorPlusAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIfElseStmt(ctx *IfElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitElseIfStmt(ctx *ElseIfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForRangeExpr(ctx *ForRangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForExpr(ctx *ForExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitGuardStmt(ctx *GuardStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFunctionWithoutParams(ctx *FunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFunctionWithParams(ctx *FunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsEI(ctx *ListFunctionParamsEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCallFunctionWithParams(ctx *CallFunctionWithParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAppendVector(ctx *AppendVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitRemoveLastVector(ctx *RemoveLastVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitRemoveAtVector(ctx *RemoveAtVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIsEmptyVector(ctx *IsEmptyVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCountVector(ctx *CountVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAccessVectorStruct(ctx *AccessVectorStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitAccessVector(ctx *AccessVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructCallFunction(ctx *StructCallFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructAttribute(ctx *StructAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitSelfFunction(ctx *SelfFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitEmbbededFunc(ctx *EmbbededFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIntstmt(ctx *IntstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFloatstmt(ctx *FloatstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStringstmt(ctx *StringstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCallBackExpr(ctx *CallBackExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitLogicalOperationExpr(ctx *LogicalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNegExpr(ctx *NegExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitComparationOperationExpr(ctx *ComparationOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructAsArgument(ctx *StructAsArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDigitExpr(ctx *DigitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCallFunctionExpr(ctx *CallFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
