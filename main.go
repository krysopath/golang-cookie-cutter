package main

type Ingredient struct {
	Name string
	ID   int64
}

type IngredientList struct{}

type Cookie interface {
	Bake([]Ingredient) (int64, bool)
	Eat() (bool, error)
}

func main() {
	delicious := []Ingredient{
		Ingredient{Name: "Sugar"},
		Ingredient{Name: "Flour"},
		Ingredient{Name: "Milk"},
		Ingredient{Name: "Egg"},
	}

	var cookie Cookie

}
