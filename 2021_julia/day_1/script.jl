using DelimitedFiles
input_array = readdlm("input", Int)

# Part 1 https://adventofcode.com/2021/day/1

global counter = 0

for (index, value) in enumerate(input_array)
	if index == 1
	  continue
	end
	if value > input_array[index-1]
		global counter += 1
	end
end

println(counter)