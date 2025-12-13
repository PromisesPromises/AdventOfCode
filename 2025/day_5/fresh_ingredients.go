package main

type FreshIngredientsRange struct {
	lower, upper int
}

func (ids FreshIngredientsRange) isFresh(val int) bool {
	return ids.lower <= val && val <= ids.upper
}

type FreshIngredientsList struct {
	arr []FreshIngredientsRange
}

func (list FreshIngredientsList) isFresh(val int) bool {
	for _, fresh_ingredients := range list.arr {
		if fresh_ingredients.isFresh(val) {
			return true
		}
	}
	return false
}

func (list *FreshIngredientsList) compacted_add(_range FreshIngredientsRange) {
	for _, fresh_ingredients := range list.arr {
		// new add is completely inside an existing range
		if fresh_ingredients.lower <= _range.lower &&
			fresh_ingredients.upper >= _range.upper {
			return
			// start of range is outside, end is inside
		} else if _range.lower <= fresh_ingredients.lower &&
			fresh_ingredients.lower <= _range.upper &&
			_range.upper <= fresh_ingredients.upper {
			list.compacted_add(FreshIngredientsRange{lower: _range.lower, upper: fresh_ingredients.lower - 1})
			return
			// start of range is inside, end is outside
		} else if fresh_ingredients.lower <= _range.lower &&
			_range.lower <= fresh_ingredients.upper &&
			fresh_ingredients.upper <= _range.upper {
			list.compacted_add(FreshIngredientsRange{lower: fresh_ingredients.upper + 1, upper: _range.upper})
			return
			// new add completely contains an existing range
		} else if _range.lower <= fresh_ingredients.lower &&
			fresh_ingredients.upper <= _range.upper {
			list.compacted_add(FreshIngredientsRange{lower: _range.lower, upper: fresh_ingredients.lower - 1})
			list.compacted_add(FreshIngredientsRange{lower: fresh_ingredients.upper + 1, upper: _range.upper})
			return
		}
	}
	list.arr = append(list.arr, _range)
}
