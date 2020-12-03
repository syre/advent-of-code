f = open("input")
input_array = readlines(f)

# Part 1 & 2 https://adventofcode.com/2020/day/2

function is_valid_password_part1(lower, upper, char, password)
    return lower <= count(i->(string(i)==char), password) <= upper
end

function is_valid_password_part2(first_pos, second_pos, char, password)
    first_pos_correct = string(password[first_pos]) == char
    second_pos_correct = string(password[second_pos]) == char
    return any([first_pos_correct, second_pos_correct]) && !all([first_pos_correct, second_pos_correct])
end

num_valid_passwords_part1 = 0
num_valid_passwords_part2 = 0
for line in input_array
    regex_match = match(r"([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)", line)
    lower, upper, char, password = regex_match.captures
    lower = parse(Int, lower)
    upper = parse(Int, upper)
    if is_valid_password_part1(lower, upper, char, password)
        global num_valid_passwords_part1 += 1
    end
    if is_valid_password_part2(lower, upper, char, password)
        global num_valid_passwords_part2 += 1
    end
end

println(num_valid_passwords_part1)
println(num_valid_passwords_part2)
