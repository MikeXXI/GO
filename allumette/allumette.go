package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func dessin_allumette(a int) {
	fmt.Printf("%v : ", a)
	for i := 0; i < a; i++ {
		fmt.Printf("|")
	}
	fmt.Printf("\n")
}

func main() {

	rand.Seed(time.Now().UnixNano())

	allumette := rand.Intn(15) + 1
	dessin_allumette(allumette)
	choix := 0

	var text int

	//Initialiser une seule fois
	scanner := bufio.NewScanner(os.Stdin)

	// A éxécuter en boucle
	for allumette > 0 {
		// Action util
		fmt.Println("Votre choix?")
		for choix > 3 || choix < 1 {
			scanner.Scan()                         //Permet d'écouter l'entrée standard
			text, _ = strconv.Atoi(scanner.Text()) // Permet de récuperer le text entrée par l'utilisateur
			choix = text
		}
		allumette = allumette - choix
		if allumette < 0 {
			fmt.Println("Vous avez tirer la dernière allumette :", allumette)
			break
		}
		dessin_allumette(allumette)
		choix = 0
		// Action AI
		choix_AI := rand.Intn(3) + 1
		fmt.Println("AI a pris :", choix_AI)
		allumette = allumette - choix_AI
		if allumette < 0 {
			fmt.Println("AI a tirer la dernière allumette :", allumette)
			break
		}
		dessin_allumette(allumette)

	}
}
