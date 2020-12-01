using DelimitedFiles
using Base.Iterators

input_array = readdlm("input", Int)

# Part 1 https://adventofcode.com/2018/day/1#part1
println(sum(input_array))

# Part 2 https://adventofcode.com/2018/day/1#part2

dict = Dict{Int, Int}(0 => 1)
current_frequency = 0
for frequency in cycle(input_array)
    global current_frequency
    current_frequency += frequency
    dict[current_frequency] = get(dict, current_frequency, 0) + 1
    if dict[current_frequency] > 1
        println(current_frequency)
        break
    end
end
