package accumulate

func Accumulate(collection []string,transformation func(string) string)[]string{
	var result []string
	for i :=0;i<len(collection);i++{
		result = append(result,transformation(collection[i]))
	}
	return result
}