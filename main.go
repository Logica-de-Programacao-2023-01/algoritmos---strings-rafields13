package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// 9. Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.

	// 10. Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.

	// 11. Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.

	// 12. Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não.

	// 13. Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").

	// 14. Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").

	// 15. Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).

	// 16. Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.

	// 17. Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).

	// 18. Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).

	// 19. Solicite ao usuário uma string e imprima a mesma string invertida.

	// 20. Solicite ao usuário uma string e informe se ela é está em camelCase e quantas palavras possuí. CamelCase é quando a primeira letra de cada palavra é maiúscula, e as demais são minúsculas. Exemplo: "estaStringEstaEmCamelCase".

	// Desafio bônus. Escreva um programa que receba uma string e um padrão (outro string) e retorne todos os índices em que o padrão ocorre na string. Informe ao usuário o resultado.

}
