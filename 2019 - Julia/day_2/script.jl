using DelimitedFiles

# Part 1 https://adventofcode.com/2019/day/2
function calculate_gravity_assist(noun, verb)
    input_array = readdlm("input", ',', Int, '\n')
    index = 1
    # Restore position 1 to value 12 and position 2 with value 2
    input_array[2] = noun
    input_array[3] = verb
    while(true)
        value = input_array[index]

        first_pos = input_array[index + 1] + 1
        second_pos = input_array[index + 2] + 1
        output_pos = input_array[index + 3] + 1
        if value == 1
            input_array[output_pos] = input_array[first_pos] + input_array[second_pos]
        elseif value == 2
            input_array[output_pos] = input_array[first_pos] * input_array[second_pos]
        elseif value == 99
            # print value at first position
            return input_array[1]
        end
        index += 4
    end
end

println(calculate_gravity_assist(12, 2))

# Part 2 https://adventofcode.com/2019/day/2#part2
for noun in 1:99
    for verb in 1:99
        gravity_assist = calculate_gravity_assist(noun, verb)
        if gravity_assist == 19690720
            println(100*noun + verb)
        end
    end
end