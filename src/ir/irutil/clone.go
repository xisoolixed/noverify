// Code generated by `codegen.go`. DO NOT EDIT.
package irutil

import (
	"fmt"
	"github.com/VKCOM/noverify/src/ir"
)

func NodeClone(x ir.Node) ir.Node {
	if x == nil {
		return nil
	}
	switch x := x.(type) {
	case *ir.Argument:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ArgumentList:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Arguments))
			for i := range x.Arguments {
				sliceClone[i] = NodeClone(x.Arguments[i]).(ir.Node)
			}
			clone.Arguments = sliceClone
		}
		return &clone
	case *ir.ArrayDimFetchExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Dim != nil {
			clone.Dim = NodeClone(x.Dim).(ir.Node)
		}
		return &clone
	case *ir.ArrayExpr:
		clone := *x
		{
			sliceClone := make([]*ir.ArrayItemExpr, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*ir.ArrayItemExpr)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *ir.ArrayItemExpr:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(ir.Node)
		}
		if x.Val != nil {
			clone.Val = NodeClone(x.Val).(ir.Node)
		}
		return &clone
	case *ir.ArrowFunctionExpr:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Params))
			for i := range x.Params {
				sliceClone[i] = NodeClone(x.Params[i]).(ir.Node)
			}
			clone.Params = sliceClone
		}
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(ir.Node)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.Assign:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignBitwiseAnd:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignBitwiseOr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignBitwiseXor:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignCoalesce:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignConcat:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignDiv:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignMinus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignMod:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignMul:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignPlus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignPow:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignReference:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignShiftLeft:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.AssignShiftRight:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(ir.Node)
		}
		return &clone
	case *ir.BitwiseAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.BitwiseNotExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.BitwiseOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.BitwiseXorExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.BooleanAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.BooleanNotExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.BooleanOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.BreakStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.CaseListStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Cases))
			for i := range x.Cases {
				sliceClone[i] = NodeClone(x.Cases[i]).(ir.Node)
			}
			clone.Cases = sliceClone
		}
		return &clone
	case *ir.CaseStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.CatchStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Types))
			for i := range x.Types {
				sliceClone[i] = NodeClone(x.Types[i]).(ir.Node)
			}
			clone.Types = sliceClone
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.ClassConstFetchExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(ir.Node)
		}
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*ir.Identifier)
		}
		return &clone
	case *ir.ClassConstListStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		{
			sliceClone := make([]ir.Node, len(x.Consts))
			for i := range x.Consts {
				sliceClone[i] = NodeClone(x.Consts[i]).(ir.Node)
			}
			clone.Consts = sliceClone
		}
		return &clone
	case *ir.ClassExtendsStmt:
		clone := *x
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(ir.Node)
		}
		return &clone
	case *ir.ClassImplementsStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.InterfaceNames))
			for i := range x.InterfaceNames {
				sliceClone[i] = NodeClone(x.InterfaceNames[i]).(ir.Node)
			}
			clone.InterfaceNames = sliceClone
		}
		return &clone
	case *ir.ClassMethodStmt:
		clone := *x
		if x.MethodName != nil {
			clone.MethodName = NodeClone(x.MethodName).(*ir.Identifier)
		}
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		{
			sliceClone := make([]ir.Node, len(x.Params))
			for i := range x.Params {
				sliceClone[i] = NodeClone(x.Params[i]).(ir.Node)
			}
			clone.Params = sliceClone
		}
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(ir.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.ClassStmt:
		clone := *x
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(*ir.Identifier)
		}
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*ir.ArgumentList)
		}
		if x.Extends != nil {
			clone.Extends = NodeClone(x.Extends).(*ir.ClassExtendsStmt)
		}
		if x.Implements != nil {
			clone.Implements = NodeClone(x.Implements).(*ir.ClassImplementsStmt)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.CloneExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ClosureExpr:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Params))
			for i := range x.Params {
				sliceClone[i] = NodeClone(x.Params[i]).(ir.Node)
			}
			clone.Params = sliceClone
		}
		if x.ClosureUse != nil {
			clone.ClosureUse = NodeClone(x.ClosureUse).(*ir.ClosureUseExpr)
		}
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.ClosureUseExpr:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Uses))
			for i := range x.Uses {
				sliceClone[i] = NodeClone(x.Uses[i]).(ir.Node)
			}
			clone.Uses = sliceClone
		}
		return &clone
	case *ir.CoalesceExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.ConcatExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.ConstFetchExpr:
		clone := *x
		if x.Constant != nil {
			clone.Constant = NodeClone(x.Constant).(ir.Node)
		}
		return &clone
	case *ir.ConstListStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Consts))
			for i := range x.Consts {
				sliceClone[i] = NodeClone(x.Consts[i]).(ir.Node)
			}
			clone.Consts = sliceClone
		}
		return &clone
	case *ir.ConstantStmt:
		clone := *x
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*ir.Identifier)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ContinueStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.DeclareStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Consts))
			for i := range x.Consts {
				sliceClone[i] = NodeClone(x.Consts[i]).(ir.Node)
			}
			clone.Consts = sliceClone
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.DefaultStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.DivExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.Dnumber:
		clone := *x
		return &clone
	case *ir.DoStmt:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		return &clone
	case *ir.EchoStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Exprs))
			for i := range x.Exprs {
				sliceClone[i] = NodeClone(x.Exprs[i]).(ir.Node)
			}
			clone.Exprs = sliceClone
		}
		return &clone
	case *ir.ElseIfStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.ElseStmt:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.EmptyExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.Encapsed:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.EncapsedStringPart:
		clone := *x
		return &clone
	case *ir.EqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.ErrorSuppressExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.EvalExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ExitExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ExpressionStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.FinallyStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.ForStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Init))
			for i := range x.Init {
				sliceClone[i] = NodeClone(x.Init[i]).(ir.Node)
			}
			clone.Init = sliceClone
		}
		{
			sliceClone := make([]ir.Node, len(x.Cond))
			for i := range x.Cond {
				sliceClone[i] = NodeClone(x.Cond[i]).(ir.Node)
			}
			clone.Cond = sliceClone
		}
		{
			sliceClone := make([]ir.Node, len(x.Loop))
			for i := range x.Loop {
				sliceClone[i] = NodeClone(x.Loop[i]).(ir.Node)
			}
			clone.Loop = sliceClone
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.ForeachStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(ir.Node)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.FullyQualifiedName:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.FunctionCallExpr:
		clone := *x
		if x.Function != nil {
			clone.Function = NodeClone(x.Function).(ir.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*ir.ArgumentList)
		}
		return &clone
	case *ir.FunctionStmt:
		clone := *x
		if x.FunctionName != nil {
			clone.FunctionName = NodeClone(x.FunctionName).(*ir.Identifier)
		}
		{
			sliceClone := make([]ir.Node, len(x.Params))
			for i := range x.Params {
				sliceClone[i] = NodeClone(x.Params[i]).(ir.Node)
			}
			clone.Params = sliceClone
		}
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.GlobalStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Vars))
			for i := range x.Vars {
				sliceClone[i] = NodeClone(x.Vars[i]).(ir.Node)
			}
			clone.Vars = sliceClone
		}
		return &clone
	case *ir.GotoStmt:
		clone := *x
		if x.Label != nil {
			clone.Label = NodeClone(x.Label).(*ir.Identifier)
		}
		return &clone
	case *ir.GreaterExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.GreaterOrEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.GroupUseStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(ir.Node)
		}
		if x.Prefix != nil {
			clone.Prefix = NodeClone(x.Prefix).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.UseList))
			for i := range x.UseList {
				sliceClone[i] = NodeClone(x.UseList[i]).(ir.Node)
			}
			clone.UseList = sliceClone
		}
		return &clone
	case *ir.HaltCompilerStmt:
		clone := *x
		return &clone
	case *ir.Heredoc:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.IdenticalExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.Identifier:
		clone := *x
		return &clone
	case *ir.IfStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.ElseIf))
			for i := range x.ElseIf {
				sliceClone[i] = NodeClone(x.ElseIf[i]).(ir.Node)
			}
			clone.ElseIf = sliceClone
		}
		if x.Else != nil {
			clone.Else = NodeClone(x.Else).(ir.Node)
		}
		return &clone
	case *ir.IncludeExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.IncludeOnceExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.InlineHTMLStmt:
		clone := *x
		return &clone
	case *ir.InstanceOfExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(ir.Node)
		}
		return &clone
	case *ir.InterfaceExtendsStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.InterfaceNames))
			for i := range x.InterfaceNames {
				sliceClone[i] = NodeClone(x.InterfaceNames[i]).(ir.Node)
			}
			clone.InterfaceNames = sliceClone
		}
		return &clone
	case *ir.InterfaceStmt:
		clone := *x
		if x.InterfaceName != nil {
			clone.InterfaceName = NodeClone(x.InterfaceName).(*ir.Identifier)
		}
		if x.Extends != nil {
			clone.Extends = NodeClone(x.Extends).(*ir.InterfaceExtendsStmt)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.IssetExpr:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Variables))
			for i := range x.Variables {
				sliceClone[i] = NodeClone(x.Variables[i]).(ir.Node)
			}
			clone.Variables = sliceClone
		}
		return &clone
	case *ir.LabelStmt:
		clone := *x
		if x.LabelName != nil {
			clone.LabelName = NodeClone(x.LabelName).(*ir.Identifier)
		}
		return &clone
	case *ir.ListExpr:
		clone := *x
		{
			sliceClone := make([]*ir.ArrayItemExpr, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*ir.ArrayItemExpr)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *ir.Lnumber:
		clone := *x
		return &clone
	case *ir.LogicalAndExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.LogicalOrExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.LogicalXorExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.MagicConstant:
		clone := *x
		return &clone
	case *ir.MethodCallExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method).(ir.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*ir.ArgumentList)
		}
		return &clone
	case *ir.MinusExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.ModExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.MulExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.Name:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.NamePart:
		clone := *x
		return &clone
	case *ir.NamespaceStmt:
		clone := *x
		if x.NamespaceName != nil {
			clone.NamespaceName = NodeClone(x.NamespaceName).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.NewExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(ir.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*ir.ArgumentList)
		}
		return &clone
	case *ir.NopStmt:
		clone := *x
		return &clone
	case *ir.NotEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.NotIdenticalExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.Nullable:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.Parameter:
		clone := *x
		if x.VariableType != nil {
			clone.VariableType = NodeClone(x.VariableType).(ir.Node)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.DefaultValue != nil {
			clone.DefaultValue = NodeClone(x.DefaultValue).(ir.Node)
		}
		return &clone
	case *ir.ParenExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.PlusExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.PostDecExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		return &clone
	case *ir.PostIncExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		return &clone
	case *ir.PowExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.PreDecExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		return &clone
	case *ir.PreIncExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		return &clone
	case *ir.PrintExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.PropertyFetchExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property).(ir.Node)
		}
		return &clone
	case *ir.PropertyListStmt:
		clone := *x
		{
			sliceClone := make([]*ir.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*ir.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.Type != nil {
			clone.Type = NodeClone(x.Type).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Properties))
			for i := range x.Properties {
				sliceClone[i] = NodeClone(x.Properties[i]).(ir.Node)
			}
			clone.Properties = sliceClone
		}
		return &clone
	case *ir.PropertyStmt:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ReferenceExpr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(ir.Node)
		}
		return &clone
	case *ir.RelativeName:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.RequireExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.RequireOnceExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.ReturnStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.Root:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.ShellExecExpr:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Parts))
			for i := range x.Parts {
				sliceClone[i] = NodeClone(x.Parts[i]).(ir.Node)
			}
			clone.Parts = sliceClone
		}
		return &clone
	case *ir.ShiftLeftExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.ShiftRightExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.SimpleVar:
		clone := *x
		return &clone
	case *ir.SmallerExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.SmallerOrEqualExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.SpaceshipExpr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(ir.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(ir.Node)
		}
		return &clone
	case *ir.StaticCallExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(ir.Node)
		}
		if x.Call != nil {
			clone.Call = NodeClone(x.Call).(ir.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*ir.ArgumentList)
		}
		return &clone
	case *ir.StaticPropertyFetchExpr:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(ir.Node)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property).(ir.Node)
		}
		return &clone
	case *ir.StaticStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Vars))
			for i := range x.Vars {
				sliceClone[i] = NodeClone(x.Vars[i]).(ir.Node)
			}
			clone.Vars = sliceClone
		}
		return &clone
	case *ir.StaticVarStmt:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*ir.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.StmtList:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.String:
		clone := *x
		return &clone
	case *ir.SwitchStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		if x.CaseList != nil {
			clone.CaseList = NodeClone(x.CaseList).(*ir.CaseListStmt)
		}
		return &clone
	case *ir.TernaryExpr:
		clone := *x
		if x.Condition != nil {
			clone.Condition = NodeClone(x.Condition).(ir.Node)
		}
		if x.IfTrue != nil {
			clone.IfTrue = NodeClone(x.IfTrue).(ir.Node)
		}
		if x.IfFalse != nil {
			clone.IfFalse = NodeClone(x.IfFalse).(ir.Node)
		}
		return &clone
	case *ir.ThrowStmt:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.TraitAdaptationListStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Adaptations))
			for i := range x.Adaptations {
				sliceClone[i] = NodeClone(x.Adaptations[i]).(ir.Node)
			}
			clone.Adaptations = sliceClone
		}
		return &clone
	case *ir.TraitMethodRefStmt:
		clone := *x
		if x.Trait != nil {
			clone.Trait = NodeClone(x.Trait).(ir.Node)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method).(*ir.Identifier)
		}
		return &clone
	case *ir.TraitStmt:
		clone := *x
		if x.TraitName != nil {
			clone.TraitName = NodeClone(x.TraitName).(*ir.Identifier)
		}
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *ir.TraitUseAliasStmt:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref).(ir.Node)
		}
		if x.Modifier != nil {
			clone.Modifier = NodeClone(x.Modifier).(ir.Node)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*ir.Identifier)
		}
		return &clone
	case *ir.TraitUsePrecedenceStmt:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Insteadof))
			for i := range x.Insteadof {
				sliceClone[i] = NodeClone(x.Insteadof[i]).(ir.Node)
			}
			clone.Insteadof = sliceClone
		}
		return &clone
	case *ir.TraitUseStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Traits))
			for i := range x.Traits {
				sliceClone[i] = NodeClone(x.Traits[i]).(ir.Node)
			}
			clone.Traits = sliceClone
		}
		if x.TraitAdaptationList != nil {
			clone.TraitAdaptationList = NodeClone(x.TraitAdaptationList).(ir.Node)
		}
		return &clone
	case *ir.TryStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(ir.Node)
			}
			clone.Stmts = sliceClone
		}
		{
			sliceClone := make([]ir.Node, len(x.Catches))
			for i := range x.Catches {
				sliceClone[i] = NodeClone(x.Catches[i]).(ir.Node)
			}
			clone.Catches = sliceClone
		}
		if x.Finally != nil {
			clone.Finally = NodeClone(x.Finally).(ir.Node)
		}
		return &clone
	case *ir.TypeCastExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.UnaryMinusExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.UnaryPlusExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.UnsetCastExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.UnsetStmt:
		clone := *x
		{
			sliceClone := make([]ir.Node, len(x.Vars))
			for i := range x.Vars {
				sliceClone[i] = NodeClone(x.Vars[i]).(ir.Node)
			}
			clone.Vars = sliceClone
		}
		return &clone
	case *ir.UseListStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(ir.Node)
		}
		{
			sliceClone := make([]ir.Node, len(x.Uses))
			for i := range x.Uses {
				sliceClone[i] = NodeClone(x.Uses[i]).(ir.Node)
			}
			clone.Uses = sliceClone
		}
		return &clone
	case *ir.UseStmt:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(*ir.Identifier)
		}
		if x.Use != nil {
			clone.Use = NodeClone(x.Use).(ir.Node)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*ir.Identifier)
		}
		return &clone
	case *ir.Var:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	case *ir.WhileStmt:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(ir.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(ir.Node)
		}
		return &clone
	case *ir.YieldExpr:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(ir.Node)
		}
		if x.Value != nil {
			clone.Value = NodeClone(x.Value).(ir.Node)
		}
		return &clone
	case *ir.YieldFromExpr:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(ir.Node)
		}
		return &clone
	}
	panic(fmt.Sprintf(`unhandled type %T`, x))
}
