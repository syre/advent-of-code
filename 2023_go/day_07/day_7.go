package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

type HandType int
const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

// keys and values stolen from exp/maps package
func keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func count(arr []int, value int) int{
    freq := make(map[int]int)
    for _ , num :=  range arr {
        freq[num] = freq[num]+1
    }
    return freq[value]
}

func calculate_hand(hand string) HandType {
	count_map := map[string]int {}
	for _, char := range hand {
		count_map[string(char)] = strings.Count(hand, string(char))
	}
	values := values(count_map)

	if slices.Contains(values, 5) {
		return FIVE_OF_A_KIND
	} else if slices.Contains(values, 4) {
		return FOUR_OF_A_KIND
	} else if slices.Contains(values, 3) && slices.Contains(values, 2) {
		return FULL_HOUSE
	} else if slices.Contains(values, 3) {
		return THREE_OF_A_KIND
	} else if count(values, 2) == 2 {
		return TWO_PAIR
	} else if count(values, 2) == 1 {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}
}

func calculate_hand_with_joker(hand string) HandType {
	count_map := map[string]int {}
	for _, char := range hand {
		count_map[string(char)] = strings.Count(hand, string(char))
	}
	// pop joker values
	joker_count, ok := count_map["J"]
    if ok {
        delete(count_map, "J")
    }
    // edge case of full jokers, just return five of a kind
    if joker_count == 5 {
    	return FIVE_OF_A_KIND
    }
	values := values(count_map)
	values[slices.Index(values, slices.Max(values))] += joker_count

	if slices.Contains(values, 5) {
		return FIVE_OF_A_KIND
	} else if slices.Contains(values, 4) {
		return FOUR_OF_A_KIND
	} else if slices.Contains(values, 3) && slices.Contains(values, 2) {
		return FULL_HOUSE
	} else if slices.Contains(values, 3) {
		return THREE_OF_A_KIND
	} else if count(values, 2) == 2 {
		return TWO_PAIR
	} else if count(values, 2) == 1 {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}
}


func part_one() {
	input, _ := os.ReadFile("input")
	hand_bid_map := map[string]int {}
	input_splitted := strings.Split(string(input), "\n")
	card_strengths := []string {"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	// load into hand_bid_map - assumes no duplicate hands
	for _, line := range input_splitted {
		line_splitted := strings.Split(line, " ")
		int_value, err := strconv.Atoi(string(line_splitted[1]))
		if err == nil {
			hand_bid_map[string(line_splitted[0])] = int_value
		}
	}
	fmt.Println(hand_bid_map)

	hand_sort := func(a, b string) int {
		a_hand := calculate_hand(a)
		b_hand := calculate_hand(b)
		if a_hand < b_hand {
			return -1
		} else if a_hand > b_hand {
			return 1
		} else {
			// use second ordering rule
			for i := 0; i < len(a); i++ {
				if a[i] == b[i] {
					continue
				}
				if slices.Index(card_strengths, string(a[i])) < slices.Index(card_strengths, string(b[i])) {
					return 1
				} else {
					return -1
				}
			}
			return 0
		}
	}
	hands_list := keys(hand_bid_map)
	slices.SortFunc(hands_list, hand_sort)
	fmt.Println(hands_list)

	total_winnings := 0
	// we use the index as hand rank
	for hand_pos, hand := range hands_list {
		hand_rank := hand_pos + 1
		fmt.Println(hand_rank, hand, hand_bid_map[hand])
		total_winnings+= hand_rank * hand_bid_map[hand]
	}

	fmt.Println(total_winnings)
}

func part_two() {
	input, _ := os.ReadFile("input")
	hand_bid_map := map[string]int {}
	input_splitted := strings.Split(string(input), "\n")
	card_strengths := []string {"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

	// load into hand_bid_map - assumes no duplicate hands
	for _, line := range input_splitted {
		line_splitted := strings.Split(line, " ")
		int_value, err := strconv.Atoi(string(line_splitted[1]))
		if err == nil {
			hand_bid_map[string(line_splitted[0])] = int_value
		}
	}
	fmt.Println(hand_bid_map)

	hand_sort := func(a, b string) int {
		a_hand := calculate_hand_with_joker(a)
		b_hand := calculate_hand_with_joker(b)
		if a_hand < b_hand {
			return -1
		} else if a_hand > b_hand {
			return 1
		} else {
			// use second ordering rule
			for i := 0; i < len(a); i++ {
				if a[i] == b[i] {
					continue
				}
				if slices.Index(card_strengths, string(a[i])) < slices.Index(card_strengths, string(b[i])) {
					return 1
				} else {
					return -1
				}
			}
			return 0
		}
	}
	hands_list := keys(hand_bid_map)
	slices.SortFunc(hands_list, hand_sort)
	fmt.Println(hands_list)

	total_winnings := 0
	// we use the index+1 as hand rank
	for hand_pos, hand := range hands_list {
		hand_rank := hand_pos + 1
		fmt.Println(hand_rank, hand, hand_bid_map[hand])
		total_winnings+= hand_rank * hand_bid_map[hand]
	}

	fmt.Println(total_winnings)
}


func main() {
	//part_one()
	part_two()
}