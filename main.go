package main

	import (

		"bufio"
		"fmt"
		"os"
		"sort"
		"strconv"
		"strings"
		"unicode"

	)

		func main() {

			// 1. Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas. Informe ao usuário o resultado.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str := scanner.Text()

				fmt.Println(strings.ToUpper(str))

			// 2. Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				str1 := strings.ReplaceAll(str0, " ", "")

				fmt.Println(str1)

			// 3. Escreva um programa que receba uma string e um caractere e substitua todas as ocorrências desse caractere na string por outro caractere também especificado pelo usuário. Informe ao usuário o resultado.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				fmt.Print("Usuário, digite um caractere que quer substituir, por favor: ")
				scanner.Scan()

				r := scanner.Text()

				fmt.Print("Usuário, digite um caractere que substituirá o caractere escolhido anteriormente, por favor: ")
				scanner.Scan()

				c := scanner.Text()

				str1 := strings.ReplaceAll(str0, r, c)

				fmt.Println(str1)

			// 4. Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				fmt.Print("Usuário, digite um outro texto, por favor: ")
				scanner.Scan()

				str1 := scanner.Text()

				if str0 == str1 {

					fmt.Print("Os textos são iguais.")

				} else {

				     fmt.Print("Os textos são diferentes.")

				  }

			// 5. Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante. Informe ao usuário se é ou não.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str := scanner.Text()

				if strings.Contains(str, ".") {

					fmt.Print("O texto é um número válido em ponto flutuante.")

				} else {

					 fmt.Print("O texto não é um número válido em ponto flutuante.")

				  }

			// 6. Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				str1 := strings.Count(str0, " ") + 1

				fmt.Printf("O texto tem %d palavra(s).", str1)

			// 7. Solicite ao usuário uma string e informe se ela contém pelo menos um número.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str := scanner.Text()

				if strings.Contains(str, "0") {

					fmt.Print("O texto contém pelo menos um número.")

				} else if strings.Contains(str, "1") {

					  fmt.Print("O texto contém pelo menos um número.")

				  } else if strings.Contains(str, "2") {

					   fmt.Print("O texto contém pelo menos um número.")

				    } else if strings.Contains(str, "3") {

						fmt.Print("O texto contém pelo menos um número.")

					  } else if strings.Contains(str, "4") {

						 fmt.Print("O texto contém pelo menos um número.")

						} else if strings.Contains(str, "5") {

						  fmt.Print("O texto contém pelo menos um número.")

						  } else if strings.Contains(str, "6") {

						   fmt.Print("O texto contém pelo menos um número.")

							} else if strings.Contains(str, "7") {

							fmt.Print("O texto contém pelo menos um número.")

							  } else if strings.Contains(str, "8") {

							 fmt.Print("O texto contém pelo menos um número.")

								} else if strings.Contains(str, "9") {

							  fmt.Print("O texto contém pelo menos um número.")

								  } else {

							   fmt.Print("O texto não contém número(s).")

									}

			// 8. Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.

				func main() {

					scanner := bufio.NewScanner(os.Stdin)

					fmt.Print("Usuário, digite um texto, por favor: ")
					scanner.Scan()

					str := scanner.Text()

					rts := reverse(str)

					fmt.Println(rts)

				}

				func reverse(s string) string {

					runes := []rune(s)

					for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {

						runes[i], runes[j] = runes[j], runes[i]

					}

					return string(runes)

				}

			// 9. Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				fmt.Print("Usuário, digite um caractere que quer substituir, por favor: ")
				scanner.Scan()

				r := scanner.Text()

				fmt.Print("Usuário, digite um caractere que substituirá o caractere escolhido anteriormente, por favor: ")
				scanner.Scan()

				c := scanner.Text()

				str1 := strings.ReplaceAll(str0, r, c)

				fmt.Println(str1)

			// 10. Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				fmt.Print("Usuário, digite um outro texto, por favor: ")
				scanner.Scan()

				str1 := scanner.Text()

				str0 = strings.ToLower(strings.ReplaceAll(str0, " ", ""))

				str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))

				if len(str0) != len(str1) {

					fmt.Println("Não são anagramas.")

					return

				}

				str0r := []rune(str0)

				str1r := []rune(str1)

				sort.Slice(str0r, func(i, j int) bool {

					return str0r[i] < str0r[j]

				} )

				sort.Slice(str1r, func(i, j int) bool {

					return str1r[i] < str1r[j]

				} )

				str0s := string(str0r)

				str1s := string(str1r)

				if str0s == str1s {

					fmt.Println("São anagramas!")

				} else {

					 fmt.Println("Não são anagramas.")

				  }

			// 11. Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				vogais := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}

				resultado := strings.Builder{}

				for _, letra := range str0 {

					if !vogais[strings.ToLower(string(letra))] {

						resultado.WriteRune(letra)

					}

				}

				fmt.Printf("Resultado: %s\n", resultado.String())

			// 12. Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não.

				func main() {

					scanner := bufio.NewScanner(os.Stdin)

					fmt.Print("Usuário, digite um texto, por favor: ")
					scanner.Scan()

					str := scanner.Text()

					rts := reverse(str)

					fmt.Println(rts)

					if str == rts {

						fmt.Println("É um palíndromo!")

					} else {

						 fmt.Println("Não é um palíndromo.")

					  }

				}

				func reverse(s string) string {

					runes := []rune(s)

					for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {

						runes[i], runes[j] = runes[j], runes[i]

					}

					return string(runes)

				}

			// 13. Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str := scanner.Text()

				nums := strings.Split(str, " ")

				var anterior int

				atualMaior := true

				for i, numStr := range nums {

					x, err := strconv.Atoi(numStr)

					if err != nil {

						fmt.Println("A sequência contém caracteres inválidos.")

						return

					}

					if i > 0 && x <= anterior {

						atualMaior = false

						break

					}

					anterior = x

				}

				if atualMaior {

					fmt.Println("A sequência é numérica crescente.")

				} else {

					 fmt.Println("A sequência não é numérica crescente.")

				  }

			// 14. Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").

				var input string

				fmt.Print("Digite uma sequência numérica: ")
				fmt.Scan(&input)

				prev := -1

				for _, c := range input {

					num, err := strconv.Atoi(string(c))

					if err != nil {

						fmt.Println("A sequência contém caracteres não numéricos!")

						return

					}

					if num >= prev && prev != -1 {

						fmt.Println("Não é uma sequência numérica decrescente!")

						return

					}

					prev = num

				}

				fmt.Println("É uma sequência numérica decrescente!")

			// 15. Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str := scanner.Text()

				str = strings.ReplaceAll(str, "a", "*")
				str = strings.ReplaceAll(str, "e", "*")
				str = strings.ReplaceAll(str, "i", "*")
				str = strings.ReplaceAll(str, "o", "*")
				str = strings.ReplaceAll(str, "u", "*")

				fmt.Println("String resultante:", str)

			// 16. Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.

				scanner := bufio.NewScanner(os.Stdin)

				fmt.Print("Usuário, digite um texto, por favor: ")
				scanner.Scan()

				str0 := scanner.Text()

				fmt.Print("Usuário, digite um outro texto, por favor: ")
				scanner.Scan()

				str1 := scanner.Text()

				if strings.Contains(str0, str1) {

					fmt.Println("O segundo texto é um sub-texto do primeiro")

				} else {

					 fmt.Println("O segundo texto não é um sub-texto do primeiro")

				  }

			// 17. Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).

				var str string

				fmt.Print("Digite uma string: ")
				fmt.Scan(&str)

				freq := make(map[rune]int)

				for _, char := range str {

					freq[char]++

				}

				fmt.Print("Letras únicas: ")

				for char, count := range freq {

					if count == 1 {

						fmt.Printf("%c ", char)

					}

				}

			// 18. Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).

				var input string

				fmt.Println("Digite uma string:")
				fmt.Scanln(&input)

				isNumber := true

				for _, char := range input {

					_, err := strconv.Atoi(string(char))

					if err != nil {

						isNumber = false

						break
					}

				}

				if isNumber {

					fmt.Println("A string contém somente números.")

				} else {

					 fmt.Println("A string não contém somente números.")
				  }

			// 19. Solicite ao usuário uma string e imprima a mesma string invertida.

				func main() {

					scanner := bufio.NewScanner(os.Stdin)

					fmt.Print("Usuário, digite um texto, por favor: ")
					scanner.Scan()

					str := scanner.Text()

					rts := reverse(str)

					fmt.Println(rts)

				}

				func reverse(s string) string {

					runes := []rune(s)

					for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {

						runes[i], runes[j] = runes[j], runes[i]

					}

					return string(runes)

				}

			// 20. Solicite ao usuário uma string e informe se ela é está em CamelCase e quantas palavras possui. CamelCase é quando a primeira letra de cada palavra é maiúscula, e as demais são minúsculas. Exemplo: "EstaStringEstáEmCamelCase".

				var input string

				fmt.Println("Digite uma string em camelCase:")
				fmt.Scanln(&input)

				words := 1

				for i, r := range input {

					if i == 0 { // se for a primeira letra da string, verifica se é maiúscula

						if !unicode.IsUpper(r) {

							fmt.Println("A string não está em camelCase.")

							return

						}

						continue

					}

				if unicode.IsSpace(r) { // se encontrar um espaço, conta uma nova palavra

					words++

				} else if !unicode.IsLower(r) { // se encontrar uma letra maiúscula que não seja a primeira, não está em camelCase

					 fmt.Println("A string não está em camelCase.")

					 return
				  }

				}

				fmt.Printf("A string está em camelCase e possui %d palavras.\n", words)

			// Desafio bônus. Escreva um programa que receba uma string e um padrão (outro string) e retorne todos os índices em que o padrão ocorre na string. Informe ao usuário o resultado.

				var str, pattern string

				fmt.Println("Digite uma string:")
				fmt.Scanln(&str)

				fmt.Println("Digite um padrão:")
				fmt.Scanln(&pattern)

				indexes := []int{}

				lastIndex := -1

				for {

					index := strings.Index(str[lastIndex+1:], pattern)

					if index == -1 {

						break

					}

					index += lastIndex + 1

					lastIndex = index

					indexes = append(indexes, index)

				}

				if len(indexes) > 0 {

					fmt.Println("O padrão ocorre nas seguintes posições:", indexes)

				} else {

					 fmt.Println("O padrão não ocorre na string.")

				  }

		}
