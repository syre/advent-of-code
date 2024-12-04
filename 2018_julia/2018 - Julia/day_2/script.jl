using DelimitedFiles
using Base.Iterators

input_array = readdlm("input", String)

# Part 1 https://adventofcode.com/2018/day/2#part1
appears_two_times = 0
appears_three_times = 0

for box_id in input_array
    global appears_two_times
    global appears_three_times
    count_dict = Dict{Char, Int}()
    for char in box_id
        count_dict[char] = get(count_dict, char, 0) + 1
    end
    appears_three_times_filled = false
    appears_two_times_filled = false
    for value in values(count_dict)
        if value == 2 && appears_two_times_filled == false
            appears_two_times += 1
            appears_two_times_filled = true
        end
        if value == 3 && appears_three_times_filled == false
            appears_three_times += 1
            appears_three_times_filled = true
        end
    end
end
println(appears_two_times*appears_three_times)

# Part 2 https://adventofcode.com/2018/day/2#part2
# This could be n logn instead of n^2
for (index, box_id) in enumerate(input_array)
    num_diffs = 0
    for compared_box_id in input_array[1:end .!= index,:]
        num_diffs = 0
        zipped = zip(box_id, compared_box_id)
        for (first_zip, second_zip) in zipped
            if num_diffs > 1
                break
            end
            if first_zip != second_zip
                num_diffs += 1
            end
        end
        if num_diffs == 1
            println(box_id)
            println(compared_box_id)
        end
    end
end