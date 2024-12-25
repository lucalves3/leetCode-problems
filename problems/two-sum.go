package problems

// Entrada: nums = [2,7,11,15], alvo = 9
//  Saída: [0,1]
//  Explicação: Como nums[0] + nums[1] == 9, retornamos [0, 1].

// Dado um array de inteiros nums e um inteiro target, retorne os índices dos dois números de modo que a soma deles seja target .

// Você pode assumir que cada entrada teria exatamente uma solução e não pode usar o mesmo elemento duas vezes.

// Você pode retornar a resposta em qualquer ordem.

func Twosum(nums []int, target int) []int {
	result := []int{}
	sum := 0
	for i, v := range nums {
		for i2, v2 := range nums {
			if i2 != i {
				sum = v2 + v
				if sum == target {
					result = []int{i, i2}
				}
			}
		}
	}
	return result
}
