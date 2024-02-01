package function

var Variable int

func Set_Global(open int) int {
	Variable = open
	return Variable
}
func Show_Global() int {

	return Variable
}
