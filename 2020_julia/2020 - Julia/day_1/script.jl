using DelimitedFiles
input_array = readdlm("input", Int)

# Part 1 https://adventofcode.com/2020/day/1

for (i, j) in Iterators.product(input_array, input_array)
    if i + j == 2020
        println(i*j)
        break
    end
end

# Part 2 https://adventofcode.com/2020/day/1

for (i, j, k) in Iterators.product(input_array, input_array, input_array)
    if i + j + k == 2020
        println(i*j*k)
        break
    end
end