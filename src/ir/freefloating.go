package ir

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
)

func (n *Assign) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignBitwiseAnd) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignBitwiseOr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignBitwiseXor) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignCoalesce) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignConcat) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignDiv) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignMinus) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignMod) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignMul) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignPlus) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignPow) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignReference) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignShiftLeft) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *AssignShiftRight) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BitwiseAndExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BitwiseOrExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BitwiseXorExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BooleanAndExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BooleanOrExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *CoalesceExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ConcatExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *DivExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *EqualExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *GreaterExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *GreaterOrEqualExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *IdenticalExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *LogicalAndExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *LogicalOrExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *LogicalXorExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *MinusExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ModExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *MulExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NotEqualExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NotIdenticalExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PlusExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PowExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ShiftLeftExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ShiftRightExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *SmallerExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *SmallerOrEqualExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *SpaceshipExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TypeCastExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UnsetCastExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ArrayExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ArrayDimFetchExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ArrayItemExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ArrowFunctionExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BitwiseNotExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BooleanNotExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassConstFetchExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *CloneExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClosureExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClosureUseExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ConstFetchExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *EmptyExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ErrorSuppressExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *EvalExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ExitExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *FunctionCallExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *IncludeExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *IncludeOnceExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *InstanceOfExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *IssetExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ListExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *MethodCallExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NewExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ParenExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PostDecExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PostIncExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PreDecExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PreIncExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PrintExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PropertyFetchExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ReferenceExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *RequireExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *RequireOnceExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ShellExecExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *StaticCallExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *StaticPropertyFetchExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TernaryExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UnaryMinusExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UnaryPlusExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *YieldExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *YieldFromExpr) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *FullyQualifiedName) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Name) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NamePart) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *RelativeName) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Argument) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ArgumentList) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Identifier) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Nullable) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Parameter) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Root) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *SimpleVar) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Var) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Dnumber) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Encapsed) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *EncapsedStringPart) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Heredoc) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *Lnumber) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *MagicConstant) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *String) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *BreakStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *CaseStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *CaseListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *CatchStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassConstListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassExtendsStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassImplementsStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ClassMethodStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ConstListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ConstantStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ContinueStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *DeclareStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *DefaultStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *DoStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *EchoStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ElseStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ElseIfStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ExpressionStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *FinallyStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ForStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ForeachStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *FunctionStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *GlobalStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *GotoStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *GroupUseStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *HaltCompilerStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *IfStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *InlineHTMLStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *InterfaceStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *InterfaceExtendsStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *LabelStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NamespaceStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *NopStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PropertyStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *PropertyListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ReturnStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *StaticStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *StaticVarStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *StmtList) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *SwitchStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *ThrowStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitAdaptationListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitMethodRefStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitUseStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitUseAliasStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TraitUsePrecedenceStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *TryStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UnsetStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UseStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *UseListStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }

func (n *WhileStmt) GetFreeFloating() *freefloating.Collection { return &n.FreeFloating }
