package main

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    // ScanRunes to split on each utf-8 rune.
    scanner.Split(bufio.ScanRunes)
    var result []int
    for scanner.Scan() {

        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func main() {

	// Initialization.
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input), "\n")

	int_matrix := make([][]int, 0, len(lines))

	for _, line := range lines {
		ints, _ := ReadInts(strings.NewReader(line))
		int_matrix = append(int_matrix, ints)
	}
	for x := 0; x < 100; x++ {
		// increase stage
		for i := 0; i < len(int_matrix); i++ {
			for j := 0; j < len(int_matrix[0]); j++ {
				int_matrix[i][j] += 1
			}
		}
		// flash stage
		for i := 0; i < len(int_matrix); i++ {
			for j := 0; j < len(int_matrix[0]); j++ {
				octopus_value = int_matrix[i][j]
				if octopus_value < 9 {
					continue
				}

				if i+1 < len(int_matrix) {
					top_neighbor := int_matrix[i+1][j]
					if top_neighbor < value {
						continue
					}
				}
				if i-1 >= 0 {
					bottom_neighbor := int_matrix[i-1][j]
					if bottom_neighbor < value {
						continue
					}
				}
				if j+1 < len(int_matrix[0]) {
					right_neighbor := int_matrix[i][j+1]
					if right_neighbor < value {
						continue
					}
				}
				if j-1 >= 0 {
					left_neighbor := int_matrix[i][j-1]
					if left_neighbor < value {
						continue
					}
				}

			}
		}
		// energy depleted stage
	}
}