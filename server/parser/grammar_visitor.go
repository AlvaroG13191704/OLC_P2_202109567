// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by GrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by GrammarParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#structStmt.
	VisitStructStmt(ctx *StructStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#structBlock.
	VisitStructBlock(ctx *StructBlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#structStmts.
	VisitStructStmts(ctx *StructStmtsContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructDeclarationWithValueAndType.
	VisitStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructDeclarationWithoutValue.
	VisitStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructDeclarationImplicitValue.
	VisitStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructDeclarationVector.
	VisitStructDeclarationVector(ctx *StructDeclarationVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructFunctionWithoutParams.
	VisitStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructFunctionWithParams.
	VisitStructFunctionWithParams(ctx *StructFunctionWithParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#structCallList.
	VisitStructCallList(ctx *StructCallListContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructCreation.
	VisitStructCreation(ctx *StructCreationContext) interface{}

	// Visit a parse tree produced by GrammarParser#TypeValueDeclaration.
	VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#TypeOptionalValueDeclaration.
	VisitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValueDeclaration.
	VisitValueDeclaration(ctx *ValueDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorOfStructDeclaration.
	VisitVectorOfStructDeclaration(ctx *VectorOfStructDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorDeclaration.
	VisitVectorDeclaration(ctx *VectorDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorOfStructCreation.
	VisitVectorOfStructCreation(ctx *VectorOfStructCreationContext) interface{}

	// Visit a parse tree produced by GrammarParser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValueAssignment.
	VisitValueAssignment(ctx *ValueAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#PlusAssignment.
	VisitPlusAssignment(ctx *PlusAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#MinusAssignment.
	VisitMinusAssignment(ctx *MinusAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorAssignment.
	VisitVectorAssignment(ctx *VectorAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorMinusAssignment.
	VisitVectorMinusAssignment(ctx *VectorMinusAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorPlusAssignment.
	VisitVectorPlusAssignment(ctx *VectorPlusAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#IfElseStmt.
	VisitIfElseStmt(ctx *IfElseStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ElseIfStmt.
	VisitElseIfStmt(ctx *ElseIfStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#switchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#defaultBlock.
	VisitDefaultBlock(ctx *DefaultBlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ForRangeExpr.
	VisitForRangeExpr(ctx *ForRangeExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ForExpr.
	VisitForExpr(ctx *ForExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#forRange.
	VisitForRange(ctx *ForRangeContext) interface{}

	// Visit a parse tree produced by GrammarParser#guardStmt.
	VisitGuardStmt(ctx *GuardStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#FunctionWithoutParams.
	VisitFunctionWithoutParams(ctx *FunctionWithoutParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#FunctionWithParams.
	VisitFunctionWithParams(ctx *FunctionWithParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsEI.
	VisitListFunctionParamsEI(ctx *ListFunctionParamsEIContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsNEI.
	VisitListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsBEI.
	VisitListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsEIVector.
	VisitListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsNEIVector.
	VisitListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#listFunctionParamsBEIVector.
	VisitListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#CallFunctionWithoutParams.
	VisitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#CallFunctionWithParams.
	VisitCallFunctionWithParams(ctx *CallFunctionWithParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#listCallFunctionStmtEI.
	VisitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) interface{}

	// Visit a parse tree produced by GrammarParser#listCallFunctionStmtNEI.
	VisitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) interface{}

	// Visit a parse tree produced by GrammarParser#AppendVector.
	VisitAppendVector(ctx *AppendVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#RemoveLastVector.
	VisitRemoveLastVector(ctx *RemoveLastVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#RemoveAtVector.
	VisitRemoveAtVector(ctx *RemoveAtVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#IsEmptyVector.
	VisitIsEmptyVector(ctx *IsEmptyVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#CountVector.
	VisitCountVector(ctx *CountVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#AccessVectorStruct.
	VisitAccessVectorStruct(ctx *AccessVectorStructContext) interface{}

	// Visit a parse tree produced by GrammarParser#AccessVector.
	VisitAccessVector(ctx *AccessVectorContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructCallFunction.
	VisitStructCallFunction(ctx *StructCallFunctionContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructAttribute.
	VisitStructAttribute(ctx *StructAttributeContext) interface{}

	// Visit a parse tree produced by GrammarParser#SelfFunction.
	VisitSelfFunction(ctx *SelfFunctionContext) interface{}

	// Visit a parse tree produced by GrammarParser#embbededFunc.
	VisitEmbbededFunc(ctx *EmbbededFuncContext) interface{}

	// Visit a parse tree produced by GrammarParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by GrammarParser#intstmt.
	VisitIntstmt(ctx *IntstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#floatstmt.
	VisitFloatstmt(ctx *FloatstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#stringstmt.
	VisitStringstmt(ctx *StringstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#EmbeddedFunctionExpr.
	VisitEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#CallBackExpr.
	VisitCallBackExpr(ctx *CallBackExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#LogicalOperationExpr.
	VisitLogicalOperationExpr(ctx *LogicalOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#NegExpr.
	VisitNegExpr(ctx *NegExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ComparationOperationExpr.
	VisitComparationOperationExpr(ctx *ComparationOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructAsArgument.
	VisitStructAsArgument(ctx *StructAsArgumentContext) interface{}

	// Visit a parse tree produced by GrammarParser#ArithmeticOperationExpr.
	VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#RelationalOperationExpr.
	VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#DigitExpr.
	VisitDigitExpr(ctx *DigitExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#CallFunctionExpr.
	VisitCallFunctionExpr(ctx *CallFunctionExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#BooleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#type.
	VisitType(ctx *TypeContext) interface{}
}
