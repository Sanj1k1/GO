package handler

type CSPlayer struct {
	FullName string `json:"name"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
}

var players = []CSPlayer{
	{FullName: "Valerii Vakhovskyi", Age: 21, Role: "Entry Fragger"},
	{FullName: "Justinas Lekavicius", Age: 24, Role: "Rifler"},
	{FullName: "Aleksi Virolainen", Age: 26, Role: "In-game Leader"},
	{FullName: "Mihai Ivan", Age: 24, Role: "Entry Fragger"},
	{FullName: "Ihor Zhdanov", Age: 19, Role: "AWPer"},
}
