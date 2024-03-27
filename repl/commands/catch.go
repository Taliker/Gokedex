package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Taliker/Gokedex/internal/api"
)

func CommandCatch(config *Config, pokemonName string) error {
	catchRate, err := api.GetCatchRate(pokemonName, &config.Cache)
	if err != nil {
		return err
	}
	rando := rand.Intn(catchRate + 1)
	fmt.Println("You threw a Pokeball...")
	time.Sleep(300 * time.Millisecond)
	if rando <= 1 {
		fmt.Println("Wiggle...")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Wiggle...")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Wiggle...")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("You caught the Pokemon!")
	} else {
		rando = rand.Intn(3)
		for i := 0; i < 3; i++ {
			fmt.Println("Wiggle...")
			time.Sleep(300 * time.Millisecond)
			if i == rando {
				fmt.Println("The Pokemon broke free!")
				break
			}
		}
	}
	return nil
}
