package main

import (
	// "math/rand"
	// "sync"
	// "fmt"
	// "time"
)

// type monk struct {
//     monastery string
//     energy  int
// }

// func fight(monk1 *monk, monk2 *monk) int {
// 	time.Sleep(time.Millisecond * 300)
//     if monk1.energy > monk2.energy {
// 		fmt.Println(monk1.monastery + " wins")
// 		return 1
// 	} else {
// 		fmt.Println(monk2.monastery + " wins")
// 		return 2
// 	}
// }

func main() {
	// rand.Seed(time.Now().UnixNano())
	// Guan_Yin := []monk{}
	// Guan_Yan := []monk{}
	// for i := 0; i < 20; i++ {
	// 	Guan_Yin = append(Guan_Yin, monk{"Guan-Yin", rand.Intn(90)+10})
	// 	Guan_Yan = append(Guan_Yan, monk{"Guan-Yan", rand.Intn(90)+10})
	// }

	// fighters := []monk{}
	// for i := 0; i < len(Guan_Yin); i++ {
	// 	fighters = append(fighters, Guan_Yin[i])
	// 	fighters = append(fighters, Guan_Yan[i])
	// }

	// for len(fighters) > 2 {
	// 	// fmt.Println(fighters)
	// 	var wg sync.WaitGroup 
	// 	wg.Add(len(fighters)/2)
	// 	for i := 0; i < len(fighters)-1; i += 2 {
	// 		go func(monk1 *monk, monk2 *monk) {
	// 			defer wg.Done()
	// 			winner := fight(monk1, monk2)
	// 			if winner == 1 {
	// 				monk1.energy = 0
	// 			} else {
	// 				monk2.energy = 0
	// 			}
	// 		  }(&fighters[i], &fighters[i+1])
	// 	}
	// 	wg.Wait()
	// 	temp_fighters := []monk{}
	// 	for i := 0; i < len(fighters); i++ {
	// 		if fighters[i].energy > 0 {
	// 			temp_fighters = append(temp_fighters, fighters[i])
	// 		}
	// 	}
	// 	fighters = temp_fighters
	// 	fmt.Println("")
	// }

	// fmt.Println("Finally, " + fighters[0].monastery + " wins")
}