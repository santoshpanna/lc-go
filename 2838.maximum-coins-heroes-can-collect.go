// func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
// 	n := len(heroes)
// 	result := make([]int64, n)

// 	for i := 0; i < n; i++ {
// 			maxCoins := int64(0)
// 			for j := 0; j < len(monsters); j++ {
// 					if heroes[i] >= monsters[j] {
// 							maxCoins += int64(coins[j])
// 					}
// 			}
// 			result[i] = maxCoins
// 	}

// 	return result
// }
import (
	"sort"
)

type monster struct {
	strength int
	coin     int
}

func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
	n := len(heroes)
	result := make([]int64, n)

	monsterList := make([]monster, len(monsters))
	for i := range monsters {
		monsterList[i] = monster{strength: monsters[i], coin: coins[i]}
	}

	sort.Slice(monsterList, func(i, j int) bool {
		return monsterList[i].strength < monsterList[j].strength
	})

	prefixCoins := make([]int64, len(monsterList))
	prefixCoins[0] = int64(monsterList[0].coin)
	for i := 1; i < len(monsterList); i++ {
		prefixCoins[i] = prefixCoins[i-1] + int64(monsterList[i].coin)
	}

	for i := 0; i < n; i++ {
		idx := sort.Search(len(monsterList), func(j int) bool {
			return monsterList[j].strength > heroes[i]
		})

		if idx > 0 {
			result[i] = prefixCoins[idx-1]
		} else {
			result[i] = 0
		}
	}

	return result
}
