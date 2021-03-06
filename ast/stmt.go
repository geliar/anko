package ast

// Stmt provides all of interfaces for statement.
type Stmt interface {
	Pos
}

// StmtImpl provide commonly implementations for Stmt..
type StmtImpl struct {
	PosImpl // PosImpl provide Pos() function.
}

// ExprStmt provide expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

// IfStmt provide "if/else" statement.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   []Stmt
	ElseIf []Stmt // This is array of IfStmt
	Else   []Stmt
}

// TryStmt provide "try/catch/finally" statement.
type TryStmt struct {
	StmtImpl
	Try     []Stmt
	Var     string
	Catch   []Stmt
	Finally []Stmt
}

// ForStmt provide "for in" expression statement.
type ForStmt struct {
	StmtImpl
	Vars  []string
	Value Expr
	Stmts []Stmt
}

// CForStmt provide C-style "for (;;)" expression statement.
type CForStmt struct {
	StmtImpl
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmts []Stmt
}

// LoopStmt provide "for expr" expression statement.
type LoopStmt struct {
	StmtImpl
	Expr  Expr
	Stmts []Stmt
}

// BreakStmt provide "break" expression statement.
type BreakStmt struct {
	StmtImpl
}

// ContinueStmt provide "continue" expression statement.
type ContinueStmt struct {
	StmtImpl
}

// ReturnStmt provide "return" expression statement.
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

// ThrowStmt provide "throw" expression statement.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

// ModuleStmt provide "module" expression statement.
type ModuleStmt struct {
	StmtImpl
	Name  string
	Stmts []Stmt
}

// SwitchStmt provide switch statement.
type SwitchStmt struct {
	StmtImpl
	Expr Expr
	Body Stmt
}

// SwitchBodyStmt provide switch case statements and default statement.
type SwitchBodyStmt struct {
	StmtImpl
	Cases   []Stmt
	Default []Stmt
}

// SwitchCaseStmt provide switch case statement.
type SwitchCaseStmt struct {
	StmtImpl
	Exprs []Expr
	Stmts []Stmt
}

// VarStmt provide statement to let variables in current scope.
type VarStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}

// LetsStmt provide multiple statement of let.
type LetsStmt struct {
	StmtImpl
	LHSS     []Expr
	Operator string
	RHSS     []Expr
}

// LetMapItemStmt provide statement of let for map item.
type LetMapItemStmt struct {
	StmtImpl
	LHSS []Expr
	RHS  Expr
}

// GoroutineStmt provide statement of groutine.
type GoroutineStmt struct {
	StmtImpl
	Expr Expr
}
