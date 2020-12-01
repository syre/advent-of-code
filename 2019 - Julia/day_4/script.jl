range_start = 265275
range_end = 781584

# Part 1 https://adventofcode.com/2019/day/4

function is_increasing(number)
    number_string = string(number)
    for index in 1:length(number_string) - 1
       if number_string[index + 1] < number_string[index]
        return false
       end
    end
    return true
end

function has_six_digits(number)
    return length(string(number)) == 6
end

function contains_double(number)
    number_string = string(number)
    return occursin(r"([\d])\1", number_string)
end

function is_password(number)
    return is_increasing(number) && has_six_digits(number) && contains_double(number)
end

password_counter = 0
for number in range_start:range_end
    global password_counter
    is_a_password = is_password(number)
    if is_a_password
        password_counter += 1
        println(number, " is a password!")
    end
end

println(password_counter)

# Part 2 https://adventofcode.com/2019/day/4#part2

function contains_double_not_part_of_group(number)
    number_string = string(number)
    double_match = eachmatch(r"(([\d])\2{1,})", number_string)
    if double_match == nothing
        return false
    end
    for capture in collect(double_match)
        if length(capture[1]) == 2
            return true
        end
    end
    return false
end

function is_password(number)
    return is_increasing(number) && has_six_digits(number) && contains_double_not_part_of_group(number)
end

password_counter = 0
for number in range_start:range_end
    global password_counter
    is_a_password = is_password(number)
    if is_a_password
        password_counter += 1
        println(number, " is a password!")
    end
end

println(password_counter)