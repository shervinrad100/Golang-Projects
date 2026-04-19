package main

type heterogeneousList struct {
	previous *heterogeneousList
	val      any
}

type homogeneousList[T any] struct { // this will require us to fix the types in this list.
	// You cannot create a list without telling us what you're putting in it
	previous *homogeneousList[T]
	val      T
}

func main() {
	heteroitem0 := heterogeneousList{
		val:      1,
		previous: nil,
	}
	heteroitem1 := heterogeneousList{
		val:      "hello",
		previous: &heteroitem0,
	}
	// note that the values in the heterogeneousList are different types
	homoItem0 := homogeneousList[int]{
		val:      1,
		previous: nil,
	}
	homoItem1 := homogeneousList[string]{ // you cannot mix types here because you've strictly required previous to be of type homogeneousList[T]
		val:      "hello",
		previous: homoItem0,
	}
}
