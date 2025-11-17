package main

import "fmt"

type Matrix [][]float64

func main() {
	var choice, size int
	fmt.Println("Калькулятор матриц 2x2 и 3x3")

	for {
		printMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Введите размер матрицы (2/3): ")
			fmt.Scan(&size)
			if !validateSize(size) {
				continue
			}
			a := inputMatrix(size)
			b := inputMatrix(size)
			result := addMatrices(a, b)
			printMatrix(result)

		case 2:
			fmt.Print("Введите размер матрицы (2/3): ")
			fmt.Scan(&size)
			if !validateSize(size) {
				continue
			}
			a := inputMatrix(size)
			var scalar float64
			fmt.Print("Введите число: ")
			fmt.Scan(&scalar)
			result := multiplyByScalar(a, scalar)
			printMatrix(result)

		case 3:
			fmt.Print("Введите размер матрицы (2/3): ")
			fmt.Scan(&size)
			if !validateSize(size) {
				continue
			}
			a := inputMatrix(size)
			b := inputMatrix(size)
			result := multiplyMatrices(a, b)
			printMatrix(result)

		case 4:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}

func printMenu() {
	fmt.Println("\n1. Сложение матриц")
	fmt.Println("2. Умножение матрицы на число")
	fmt.Println("3. Умножение матриц")
	fmt.Println("4. Выход")
	fmt.Print("Выберите операцию: ")
}

func validateSize(s int) bool {
	if s != 2 && s != 3 {
		fmt.Println("Неверный размер. Допустимо: 2 или 3")
		return false
	}
	return true
}

func inputMatrix(size int) Matrix {
	m := make(Matrix, size)
	fmt.Printf("Введите матрицу %dx%d:\n", size, size)
	for i := 0; i < size; i++ {
		m[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			fmt.Printf("Элемент [%d][%d]: ", i+1, j+1)
			fmt.Scan(&m[i][j])
		}
	}
	return m
}

func printMatrix(m Matrix) {
	fmt.Println("Результат:")
	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%6.2f", val)
		}
		fmt.Println()
	}
}

func addMatrices(a, b Matrix) Matrix {
	size := len(a)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func multiplyByScalar(m Matrix, scalar float64) Matrix {
	size := len(m)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = m[i][j] * scalar
		}
	}
	return result
}

func multiplyMatrices(a, b Matrix) Matrix {
	size := len(a)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}
