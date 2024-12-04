using DelimitedFiles

input = read("input", String)

# Example 'a' == 'a', 'A' == 'A'
function is_same_type(a, b)
    return a == b
end

# Example 'a' == 'A', 'A' == 'a'
function is_different_polarity(a, b)
    return !is_same_type(a, b) && (uppercase(a) == b || uppercase(b) == a)
end

# aAdaDdbDdAcCaCBAcCcaDACcDd -> dabAaCBAcaDA -> dabCBAcaDA
function create_reactions(input)
    input_indexes = Vector{Int}()
    current_index = 1
    while(current_index < length(input))
        if is_different_polarity(input[current_index], input[current_index+1])
            current_index += 2
        else
            push!(input_indexes, current_index)
            current_index += 1
        end
    end
    if current_index == length(input)
        push!(input_indexes, current_index)
    end
    return input[input_indexes]
end

# Part 1 https://adventofcode.com/2018/day/5#part1
input_length = length(input)
input = create_reactions(input)
revised_length = length(input)
while(input_length != revised_length)
    global revised_length
    global input
    global input_length
    input_length = revised_length
    input = create_reactions(input)
    revised_length = length(input)
    println(input_length, " ", revised_length)
end

# Part 2 https://adventofcode.com/2018/day/5#part2
counter_dict = Dict{Char, Int}(i=>0 for i in "abcdefghijklmnopqrstuvwxyz")
for key in keys(counter_dict)
    global revised_length
    global input
    global input_length
    new_input = replace(input, [key, uppercase(key)] => "")
    input_length = length(new_input)
    new_input = create_reactions(new_input)
    revised_length = length(new_input)
    while(input_length != revised_length)
        input_length = revised_length
        new_input = create_reactions(new_input)
        revised_length = length(new_input)
        println(input_length, " ", revised_length)
    end
    counter_dict[key] = revised_length
end

println(counter_dict)
print(min(values(counter_dict)))