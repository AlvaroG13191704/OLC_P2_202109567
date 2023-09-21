// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGrammarListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGrammarListener) ExitStart(ctx *StartContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseGrammarListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseGrammarListener) ExitStmts(ctx *StmtsContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterStructStmt is called when production structStmt is entered.
func (s *BaseGrammarListener) EnterStructStmt(ctx *StructStmtContext) {}

// ExitStructStmt is called when production structStmt is exited.
func (s *BaseGrammarListener) ExitStructStmt(ctx *StructStmtContext) {}

// EnterStructBlock is called when production structBlock is entered.
func (s *BaseGrammarListener) EnterStructBlock(ctx *StructBlockContext) {}

// ExitStructBlock is called when production structBlock is exited.
func (s *BaseGrammarListener) ExitStructBlock(ctx *StructBlockContext) {}

// EnterStructStmts is called when production structStmts is entered.
func (s *BaseGrammarListener) EnterStructStmts(ctx *StructStmtsContext) {}

// ExitStructStmts is called when production structStmts is exited.
func (s *BaseGrammarListener) ExitStructStmts(ctx *StructStmtsContext) {}

// EnterStructDeclarationWithValueAndType is called when production StructDeclarationWithValueAndType is entered.
func (s *BaseGrammarListener) EnterStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) {
}

// ExitStructDeclarationWithValueAndType is called when production StructDeclarationWithValueAndType is exited.
func (s *BaseGrammarListener) ExitStructDeclarationWithValueAndType(ctx *StructDeclarationWithValueAndTypeContext) {
}

// EnterStructDeclarationWithoutValue is called when production StructDeclarationWithoutValue is entered.
func (s *BaseGrammarListener) EnterStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) {
}

// ExitStructDeclarationWithoutValue is called when production StructDeclarationWithoutValue is exited.
func (s *BaseGrammarListener) ExitStructDeclarationWithoutValue(ctx *StructDeclarationWithoutValueContext) {
}

// EnterStructDeclarationImplicitValue is called when production StructDeclarationImplicitValue is entered.
func (s *BaseGrammarListener) EnterStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) {
}

// ExitStructDeclarationImplicitValue is called when production StructDeclarationImplicitValue is exited.
func (s *BaseGrammarListener) ExitStructDeclarationImplicitValue(ctx *StructDeclarationImplicitValueContext) {
}

// EnterStructDeclarationVector is called when production StructDeclarationVector is entered.
func (s *BaseGrammarListener) EnterStructDeclarationVector(ctx *StructDeclarationVectorContext) {}

// ExitStructDeclarationVector is called when production StructDeclarationVector is exited.
func (s *BaseGrammarListener) ExitStructDeclarationVector(ctx *StructDeclarationVectorContext) {}

// EnterStructFunctionWithoutParams is called when production StructFunctionWithoutParams is entered.
func (s *BaseGrammarListener) EnterStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) {
}

// ExitStructFunctionWithoutParams is called when production StructFunctionWithoutParams is exited.
func (s *BaseGrammarListener) ExitStructFunctionWithoutParams(ctx *StructFunctionWithoutParamsContext) {
}

// EnterStructFunctionWithParams is called when production StructFunctionWithParams is entered.
func (s *BaseGrammarListener) EnterStructFunctionWithParams(ctx *StructFunctionWithParamsContext) {}

// ExitStructFunctionWithParams is called when production StructFunctionWithParams is exited.
func (s *BaseGrammarListener) ExitStructFunctionWithParams(ctx *StructFunctionWithParamsContext) {}

// EnterStructCallList is called when production structCallList is entered.
func (s *BaseGrammarListener) EnterStructCallList(ctx *StructCallListContext) {}

// ExitStructCallList is called when production structCallList is exited.
func (s *BaseGrammarListener) ExitStructCallList(ctx *StructCallListContext) {}

// EnterStructCreation is called when production StructCreation is entered.
func (s *BaseGrammarListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production StructCreation is exited.
func (s *BaseGrammarListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterTypeValueDeclaration is called when production TypeValueDeclaration is entered.
func (s *BaseGrammarListener) EnterTypeValueDeclaration(ctx *TypeValueDeclarationContext) {}

// ExitTypeValueDeclaration is called when production TypeValueDeclaration is exited.
func (s *BaseGrammarListener) ExitTypeValueDeclaration(ctx *TypeValueDeclarationContext) {}

// EnterTypeOptionalValueDeclaration is called when production TypeOptionalValueDeclaration is entered.
func (s *BaseGrammarListener) EnterTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) {
}

// ExitTypeOptionalValueDeclaration is called when production TypeOptionalValueDeclaration is exited.
func (s *BaseGrammarListener) ExitTypeOptionalValueDeclaration(ctx *TypeOptionalValueDeclarationContext) {
}

// EnterValueDeclaration is called when production ValueDeclaration is entered.
func (s *BaseGrammarListener) EnterValueDeclaration(ctx *ValueDeclarationContext) {}

// ExitValueDeclaration is called when production ValueDeclaration is exited.
func (s *BaseGrammarListener) ExitValueDeclaration(ctx *ValueDeclarationContext) {}

// EnterVectorOfStructDeclaration is called when production VectorOfStructDeclaration is entered.
func (s *BaseGrammarListener) EnterVectorOfStructDeclaration(ctx *VectorOfStructDeclarationContext) {}

// ExitVectorOfStructDeclaration is called when production VectorOfStructDeclaration is exited.
func (s *BaseGrammarListener) ExitVectorOfStructDeclaration(ctx *VectorOfStructDeclarationContext) {}

// EnterVectorDeclaration is called when production VectorDeclaration is entered.
func (s *BaseGrammarListener) EnterVectorDeclaration(ctx *VectorDeclarationContext) {}

// ExitVectorDeclaration is called when production VectorDeclaration is exited.
func (s *BaseGrammarListener) ExitVectorDeclaration(ctx *VectorDeclarationContext) {}

// EnterVectorOfStructCreation is called when production VectorOfStructCreation is entered.
func (s *BaseGrammarListener) EnterVectorOfStructCreation(ctx *VectorOfStructCreationContext) {}

// ExitVectorOfStructCreation is called when production VectorOfStructCreation is exited.
func (s *BaseGrammarListener) ExitVectorOfStructCreation(ctx *VectorOfStructCreationContext) {}

// EnterType_declaration is called when production type_declaration is entered.
func (s *BaseGrammarListener) EnterType_declaration(ctx *Type_declarationContext) {}

// ExitType_declaration is called when production type_declaration is exited.
func (s *BaseGrammarListener) ExitType_declaration(ctx *Type_declarationContext) {}

// EnterValueAssignment is called when production ValueAssignment is entered.
func (s *BaseGrammarListener) EnterValueAssignment(ctx *ValueAssignmentContext) {}

// ExitValueAssignment is called when production ValueAssignment is exited.
func (s *BaseGrammarListener) ExitValueAssignment(ctx *ValueAssignmentContext) {}

// EnterPlusAssignment is called when production PlusAssignment is entered.
func (s *BaseGrammarListener) EnterPlusAssignment(ctx *PlusAssignmentContext) {}

// ExitPlusAssignment is called when production PlusAssignment is exited.
func (s *BaseGrammarListener) ExitPlusAssignment(ctx *PlusAssignmentContext) {}

// EnterMinusAssignment is called when production MinusAssignment is entered.
func (s *BaseGrammarListener) EnterMinusAssignment(ctx *MinusAssignmentContext) {}

// ExitMinusAssignment is called when production MinusAssignment is exited.
func (s *BaseGrammarListener) ExitMinusAssignment(ctx *MinusAssignmentContext) {}

// EnterVectorAssignment is called when production VectorAssignment is entered.
func (s *BaseGrammarListener) EnterVectorAssignment(ctx *VectorAssignmentContext) {}

// ExitVectorAssignment is called when production VectorAssignment is exited.
func (s *BaseGrammarListener) ExitVectorAssignment(ctx *VectorAssignmentContext) {}

// EnterVectorMinusAssignment is called when production VectorMinusAssignment is entered.
func (s *BaseGrammarListener) EnterVectorMinusAssignment(ctx *VectorMinusAssignmentContext) {}

// ExitVectorMinusAssignment is called when production VectorMinusAssignment is exited.
func (s *BaseGrammarListener) ExitVectorMinusAssignment(ctx *VectorMinusAssignmentContext) {}

// EnterVectorPlusAssignment is called when production VectorPlusAssignment is entered.
func (s *BaseGrammarListener) EnterVectorPlusAssignment(ctx *VectorPlusAssignmentContext) {}

// ExitVectorPlusAssignment is called when production VectorPlusAssignment is exited.
func (s *BaseGrammarListener) ExitVectorPlusAssignment(ctx *VectorPlusAssignmentContext) {}

// EnterIfElseStmt is called when production IfElseStmt is entered.
func (s *BaseGrammarListener) EnterIfElseStmt(ctx *IfElseStmtContext) {}

// ExitIfElseStmt is called when production IfElseStmt is exited.
func (s *BaseGrammarListener) ExitIfElseStmt(ctx *IfElseStmtContext) {}

// EnterElseIfStmt is called when production ElseIfStmt is entered.
func (s *BaseGrammarListener) EnterElseIfStmt(ctx *ElseIfStmtContext) {}

// ExitElseIfStmt is called when production ElseIfStmt is exited.
func (s *BaseGrammarListener) ExitElseIfStmt(ctx *ElseIfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseGrammarListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseGrammarListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseGrammarListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseGrammarListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseGrammarListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseGrammarListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForRangeExpr is called when production ForRangeExpr is entered.
func (s *BaseGrammarListener) EnterForRangeExpr(ctx *ForRangeExprContext) {}

// ExitForRangeExpr is called when production ForRangeExpr is exited.
func (s *BaseGrammarListener) ExitForRangeExpr(ctx *ForRangeExprContext) {}

// EnterForExpr is called when production ForExpr is entered.
func (s *BaseGrammarListener) EnterForExpr(ctx *ForExprContext) {}

// ExitForExpr is called when production ForExpr is exited.
func (s *BaseGrammarListener) ExitForExpr(ctx *ForExprContext) {}

// EnterForRange is called when production forRange is entered.
func (s *BaseGrammarListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production forRange is exited.
func (s *BaseGrammarListener) ExitForRange(ctx *ForRangeContext) {}

// EnterGuardStmt is called when production guardStmt is entered.
func (s *BaseGrammarListener) EnterGuardStmt(ctx *GuardStmtContext) {}

// ExitGuardStmt is called when production guardStmt is exited.
func (s *BaseGrammarListener) ExitGuardStmt(ctx *GuardStmtContext) {}

// EnterFunctionWithoutParams is called when production FunctionWithoutParams is entered.
func (s *BaseGrammarListener) EnterFunctionWithoutParams(ctx *FunctionWithoutParamsContext) {}

// ExitFunctionWithoutParams is called when production FunctionWithoutParams is exited.
func (s *BaseGrammarListener) ExitFunctionWithoutParams(ctx *FunctionWithoutParamsContext) {}

// EnterFunctionWithParams is called when production FunctionWithParams is entered.
func (s *BaseGrammarListener) EnterFunctionWithParams(ctx *FunctionWithParamsContext) {}

// ExitFunctionWithParams is called when production FunctionWithParams is exited.
func (s *BaseGrammarListener) ExitFunctionWithParams(ctx *FunctionWithParamsContext) {}

// EnterListFunctionParamsEI is called when production listFunctionParamsEI is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsEI(ctx *ListFunctionParamsEIContext) {}

// ExitListFunctionParamsEI is called when production listFunctionParamsEI is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsEI(ctx *ListFunctionParamsEIContext) {}

// EnterListFunctionParamsNEI is called when production listFunctionParamsNEI is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) {}

// ExitListFunctionParamsNEI is called when production listFunctionParamsNEI is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsNEI(ctx *ListFunctionParamsNEIContext) {}

// EnterListFunctionParamsBEI is called when production listFunctionParamsBEI is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) {}

// ExitListFunctionParamsBEI is called when production listFunctionParamsBEI is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsBEI(ctx *ListFunctionParamsBEIContext) {}

// EnterListFunctionParamsEIVector is called when production listFunctionParamsEIVector is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) {
}

// ExitListFunctionParamsEIVector is called when production listFunctionParamsEIVector is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsEIVector(ctx *ListFunctionParamsEIVectorContext) {
}

// EnterListFunctionParamsNEIVector is called when production listFunctionParamsNEIVector is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) {
}

// ExitListFunctionParamsNEIVector is called when production listFunctionParamsNEIVector is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsNEIVector(ctx *ListFunctionParamsNEIVectorContext) {
}

// EnterListFunctionParamsBEIVector is called when production listFunctionParamsBEIVector is entered.
func (s *BaseGrammarListener) EnterListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) {
}

// ExitListFunctionParamsBEIVector is called when production listFunctionParamsBEIVector is exited.
func (s *BaseGrammarListener) ExitListFunctionParamsBEIVector(ctx *ListFunctionParamsBEIVectorContext) {
}

// EnterCallFunctionWithoutParams is called when production CallFunctionWithoutParams is entered.
func (s *BaseGrammarListener) EnterCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) {}

// ExitCallFunctionWithoutParams is called when production CallFunctionWithoutParams is exited.
func (s *BaseGrammarListener) ExitCallFunctionWithoutParams(ctx *CallFunctionWithoutParamsContext) {}

// EnterCallFunctionWithParams is called when production CallFunctionWithParams is entered.
func (s *BaseGrammarListener) EnterCallFunctionWithParams(ctx *CallFunctionWithParamsContext) {}

// ExitCallFunctionWithParams is called when production CallFunctionWithParams is exited.
func (s *BaseGrammarListener) ExitCallFunctionWithParams(ctx *CallFunctionWithParamsContext) {}

// EnterListCallFunctionStmtEI is called when production listCallFunctionStmtEI is entered.
func (s *BaseGrammarListener) EnterListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) {}

// ExitListCallFunctionStmtEI is called when production listCallFunctionStmtEI is exited.
func (s *BaseGrammarListener) ExitListCallFunctionStmtEI(ctx *ListCallFunctionStmtEIContext) {}

// EnterListCallFunctionStmtNEI is called when production listCallFunctionStmtNEI is entered.
func (s *BaseGrammarListener) EnterListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) {}

// ExitListCallFunctionStmtNEI is called when production listCallFunctionStmtNEI is exited.
func (s *BaseGrammarListener) ExitListCallFunctionStmtNEI(ctx *ListCallFunctionStmtNEIContext) {}

// EnterAppendVector is called when production AppendVector is entered.
func (s *BaseGrammarListener) EnterAppendVector(ctx *AppendVectorContext) {}

// ExitAppendVector is called when production AppendVector is exited.
func (s *BaseGrammarListener) ExitAppendVector(ctx *AppendVectorContext) {}

// EnterRemoveLastVector is called when production RemoveLastVector is entered.
func (s *BaseGrammarListener) EnterRemoveLastVector(ctx *RemoveLastVectorContext) {}

// ExitRemoveLastVector is called when production RemoveLastVector is exited.
func (s *BaseGrammarListener) ExitRemoveLastVector(ctx *RemoveLastVectorContext) {}

// EnterRemoveAtVector is called when production RemoveAtVector is entered.
func (s *BaseGrammarListener) EnterRemoveAtVector(ctx *RemoveAtVectorContext) {}

// ExitRemoveAtVector is called when production RemoveAtVector is exited.
func (s *BaseGrammarListener) ExitRemoveAtVector(ctx *RemoveAtVectorContext) {}

// EnterIsEmptyVector is called when production IsEmptyVector is entered.
func (s *BaseGrammarListener) EnterIsEmptyVector(ctx *IsEmptyVectorContext) {}

// ExitIsEmptyVector is called when production IsEmptyVector is exited.
func (s *BaseGrammarListener) ExitIsEmptyVector(ctx *IsEmptyVectorContext) {}

// EnterCountVector is called when production CountVector is entered.
func (s *BaseGrammarListener) EnterCountVector(ctx *CountVectorContext) {}

// ExitCountVector is called when production CountVector is exited.
func (s *BaseGrammarListener) ExitCountVector(ctx *CountVectorContext) {}

// EnterAccessVectorStruct is called when production AccessVectorStruct is entered.
func (s *BaseGrammarListener) EnterAccessVectorStruct(ctx *AccessVectorStructContext) {}

// ExitAccessVectorStruct is called when production AccessVectorStruct is exited.
func (s *BaseGrammarListener) ExitAccessVectorStruct(ctx *AccessVectorStructContext) {}

// EnterAccessVector is called when production AccessVector is entered.
func (s *BaseGrammarListener) EnterAccessVector(ctx *AccessVectorContext) {}

// ExitAccessVector is called when production AccessVector is exited.
func (s *BaseGrammarListener) ExitAccessVector(ctx *AccessVectorContext) {}

// EnterStructCallFunction is called when production StructCallFunction is entered.
func (s *BaseGrammarListener) EnterStructCallFunction(ctx *StructCallFunctionContext) {}

// ExitStructCallFunction is called when production StructCallFunction is exited.
func (s *BaseGrammarListener) ExitStructCallFunction(ctx *StructCallFunctionContext) {}

// EnterStructAttribute is called when production StructAttribute is entered.
func (s *BaseGrammarListener) EnterStructAttribute(ctx *StructAttributeContext) {}

// ExitStructAttribute is called when production StructAttribute is exited.
func (s *BaseGrammarListener) ExitStructAttribute(ctx *StructAttributeContext) {}

// EnterSelfFunction is called when production SelfFunction is entered.
func (s *BaseGrammarListener) EnterSelfFunction(ctx *SelfFunctionContext) {}

// ExitSelfFunction is called when production SelfFunction is exited.
func (s *BaseGrammarListener) ExitSelfFunction(ctx *SelfFunctionContext) {}

// EnterEmbbededFunc is called when production embbededFunc is entered.
func (s *BaseGrammarListener) EnterEmbbededFunc(ctx *EmbbededFuncContext) {}

// ExitEmbbededFunc is called when production embbededFunc is exited.
func (s *BaseGrammarListener) ExitEmbbededFunc(ctx *EmbbededFuncContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseGrammarListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseGrammarListener) ExitExprList(ctx *ExprListContext) {}

// EnterIntstmt is called when production intstmt is entered.
func (s *BaseGrammarListener) EnterIntstmt(ctx *IntstmtContext) {}

// ExitIntstmt is called when production intstmt is exited.
func (s *BaseGrammarListener) ExitIntstmt(ctx *IntstmtContext) {}

// EnterFloatstmt is called when production floatstmt is entered.
func (s *BaseGrammarListener) EnterFloatstmt(ctx *FloatstmtContext) {}

// ExitFloatstmt is called when production floatstmt is exited.
func (s *BaseGrammarListener) ExitFloatstmt(ctx *FloatstmtContext) {}

// EnterStringstmt is called when production stringstmt is entered.
func (s *BaseGrammarListener) EnterStringstmt(ctx *StringstmtContext) {}

// ExitStringstmt is called when production stringstmt is exited.
func (s *BaseGrammarListener) ExitStringstmt(ctx *StringstmtContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseGrammarListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseGrammarListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterEmbeddedFunctionExpr is called when production EmbeddedFunctionExpr is entered.
func (s *BaseGrammarListener) EnterEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) {}

// ExitEmbeddedFunctionExpr is called when production EmbeddedFunctionExpr is exited.
func (s *BaseGrammarListener) ExitEmbeddedFunctionExpr(ctx *EmbeddedFunctionExprContext) {}

// EnterNilExpr is called when production NilExpr is entered.
func (s *BaseGrammarListener) EnterNilExpr(ctx *NilExprContext) {}

// ExitNilExpr is called when production NilExpr is exited.
func (s *BaseGrammarListener) ExitNilExpr(ctx *NilExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseGrammarListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseGrammarListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterCallBackExpr is called when production CallBackExpr is entered.
func (s *BaseGrammarListener) EnterCallBackExpr(ctx *CallBackExprContext) {}

// ExitCallBackExpr is called when production CallBackExpr is exited.
func (s *BaseGrammarListener) ExitCallBackExpr(ctx *CallBackExprContext) {}

// EnterLogicalOperationExpr is called when production LogicalOperationExpr is entered.
func (s *BaseGrammarListener) EnterLogicalOperationExpr(ctx *LogicalOperationExprContext) {}

// ExitLogicalOperationExpr is called when production LogicalOperationExpr is exited.
func (s *BaseGrammarListener) ExitLogicalOperationExpr(ctx *LogicalOperationExprContext) {}

// EnterNegExpr is called when production NegExpr is entered.
func (s *BaseGrammarListener) EnterNegExpr(ctx *NegExprContext) {}

// ExitNegExpr is called when production NegExpr is exited.
func (s *BaseGrammarListener) ExitNegExpr(ctx *NegExprContext) {}

// EnterComparationOperationExpr is called when production ComparationOperationExpr is entered.
func (s *BaseGrammarListener) EnterComparationOperationExpr(ctx *ComparationOperationExprContext) {}

// ExitComparationOperationExpr is called when production ComparationOperationExpr is exited.
func (s *BaseGrammarListener) ExitComparationOperationExpr(ctx *ComparationOperationExprContext) {}

// EnterStructAsArgument is called when production StructAsArgument is entered.
func (s *BaseGrammarListener) EnterStructAsArgument(ctx *StructAsArgumentContext) {}

// ExitStructAsArgument is called when production StructAsArgument is exited.
func (s *BaseGrammarListener) ExitStructAsArgument(ctx *StructAsArgumentContext) {}

// EnterArithmeticOperationExpr is called when production ArithmeticOperationExpr is entered.
func (s *BaseGrammarListener) EnterArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) {}

// ExitArithmeticOperationExpr is called when production ArithmeticOperationExpr is exited.
func (s *BaseGrammarListener) ExitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) {}

// EnterRelationalOperationExpr is called when production RelationalOperationExpr is entered.
func (s *BaseGrammarListener) EnterRelationalOperationExpr(ctx *RelationalOperationExprContext) {}

// ExitRelationalOperationExpr is called when production RelationalOperationExpr is exited.
func (s *BaseGrammarListener) ExitRelationalOperationExpr(ctx *RelationalOperationExprContext) {}

// EnterDigitExpr is called when production DigitExpr is entered.
func (s *BaseGrammarListener) EnterDigitExpr(ctx *DigitExprContext) {}

// ExitDigitExpr is called when production DigitExpr is exited.
func (s *BaseGrammarListener) ExitDigitExpr(ctx *DigitExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseGrammarListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseGrammarListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseGrammarListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseGrammarListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterCallFunctionExpr is called when production CallFunctionExpr is entered.
func (s *BaseGrammarListener) EnterCallFunctionExpr(ctx *CallFunctionExprContext) {}

// ExitCallFunctionExpr is called when production CallFunctionExpr is exited.
func (s *BaseGrammarListener) ExitCallFunctionExpr(ctx *CallFunctionExprContext) {}

// EnterBooleanExpr is called when production BooleanExpr is entered.
func (s *BaseGrammarListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production BooleanExpr is exited.
func (s *BaseGrammarListener) ExitBooleanExpr(ctx *BooleanExprContext) {}

// EnterType is called when production type is entered.
func (s *BaseGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGrammarListener) ExitType(ctx *TypeContext) {}
