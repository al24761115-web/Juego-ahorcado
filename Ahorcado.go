package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var estados = []string{
	"  O",
	"  O\n  |",
	"  O\n /|",
	"  O\n /|\\",
	"  O\n /|\\\n /",
	"  O\n /|\\\n / \\",
}

func main() {
	palabras := []string{"linux", "kombat", "mortal", "debo", "paguen", "renta"}

	rand.Seed(time.Now().UnixNano()) // el unix nano es para que sea de manera aleatoria  la palabra que vas a adivinar
	secreta := palabras[rand.Intn(len(palabras))]

	progreso := make([]string, len(secreta))
	for i := range progreso {
		progreso[i] = "_"
	}

	fallos := 0
	maxVidas := 6

	fmt.Println(" AHORCADO ")

	for fallos < maxVidas {
		fmt.Printf("\nPalabra: %s  (Te quedan  %d vidas)\n", strings.Join(progreso, " "), maxVidas-fallos)
		fmt.Print("dame una letra: ")

		var input string
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if len(input) != 1 {
			fmt.Println("Tranquilo saltamontes, dije una")
			continue
		}

		acierto := false
		for i, letra := range secreta {
			if string(letra) == input {
				progreso[i] = input
				acierto = true
			}
		}

		if !acierto {
			fmt.Println(estados[fallos])
			fallos++
			fmt.Printf("Tsss tas chavo chavo, la '%s' no esta.\n", input)
		}

		if strings.Join(progreso, "") == secreta {
			fmt.Printf("\nCorreeeectooo, la palabra es: %s\n", secreta)
			return
		}
	}

	fmt.Printf("\nEstas ahorcado.  La palabra correcta es: %s\n", secreta)
}
