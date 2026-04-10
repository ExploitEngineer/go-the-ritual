package main

import (
	introduction "github.com/ExploitEngineer/go-the-ritual/exercises/00-introduction"
	variables "github.com/ExploitEngineer/go-the-ritual/exercises/01-variables-and-constants"
	datatypes "github.com/ExploitEngineer/go-the-ritual/exercises/02-data-types"
)

func main() {
	introduction.HelloWorld()
	introduction.SingleLinePrint()
	introduction.MultiLineOutput()
	introduction.ExploreFormatting()
	introduction.ImportErrorTest()
	introduction.RunVsBuild()
	introduction.SyntaxErrorDebugging()
	introduction.PredictOutput()
	introduction.BuildMiniIntro()

	variables.DeclareVariablesInMultipleWays()
	variables.PredictZeroValues()
	variables.BuildUserCard()
	variables.ConstantsPractice()
	variables.IotaEnum()
	variables.ScopeTest()
	variables.ShadowingChallange()
	variables.ModifyConstantError()
	variables.BlockScope()

	datatypes.IntegerTypeShowcase()
	datatypes.OverflowExperiment()
	datatypes.FloatPrecisionTest()
}
