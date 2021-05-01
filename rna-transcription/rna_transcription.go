package strand

func ToRNA(dna string) string {
	complement := map[rune]rune{
		'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U',
	}
	runes := []rune(dna)
	for i := 0; i < len(runes); i++ {
		runes[i] = complement[runes[i]]
	}

	return string(runes)
}
