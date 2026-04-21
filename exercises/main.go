package main

import (
	introduction "github.com/ExploitEngineer/go-the-ritual/exercises/00-introduction"
	variables "github.com/ExploitEngineer/go-the-ritual/exercises/01-variables-and-constants"
	datatypes "github.com/ExploitEngineer/go-the-ritual/exercises/02-data-types"
	compositetypes "github.com/ExploitEngineer/go-the-ritual/exercises/03-composite-types"
	structs "github.com/ExploitEngineer/go-the-ritual/exercises/04-structs"
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
	datatypes.BooleanLogic()
	datatypes.RuneInspection()
	datatypes.StringLengthTrap()
	datatypes.TypeConversion()
	datatypes.ComplexNumber()

	compositetypes.ArrayVsSlice()
	compositetypes.SliceReferenceTrap()
	compositetypes.SliceAppendGrowth()
	compositetypes.MakeVsLiteral()
	compositetypes.MapBasicOperation()
	compositetypes.CommaOkIdiom()
	compositetypes.MissingKeyPitfall()
	compositetypes.SliceLengthVsCapacity()
	compositetypes.ArrayToSlice()

	structs.BasicStructCreation()
	structs.MultipleInitializationStyles()
	structs.PointerVsValues()
	structs.NestedStructs()
	structs.JsonTags()
	structs.UnexpectedFieldsTrap()
	structs.StructComparison()
	structs.ZeroValueStruct()
	structs.RealMiniSystem()
}
