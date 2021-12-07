package world

// Triangle defines the triangulation for a cell
type Triangle []int

// Triangulations contains all the possible cell triangulation given the cell Number
var Triangulations = []Triangle{
	{-100}, // Outside
	{0, 3},
	{0, 1},
	{1, 3},
	{1, 2},
	{0, 1, 2, 3}, // Saddle
	{0, 2},
	{2, 3},
	{2, 3},
	{0, 2},
	{1, 2, 0, 3}, // Saddle
	{1, 2},
	{1, 3},
	{0, 1},
	{0, 3},
	{100}, // Inside
}
