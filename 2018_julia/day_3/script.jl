using DelimitedFiles

input_array = readdlm("input", '\n', String)

fabric_size = (1000, 1000)

fabric = fill('.', fabric_size)

regex = r"^#(?<id>\d+) @ (?<inches_from_left_edge>\d+),(?<inches_from_top_edge>\d+): (?<rectangle_width>\d+)x(?<rectangle_height>\d+)"
input_dict = Dict{String, Tuple}()
for entry in input_array
    m = match(regex, entry)
    id = m[:id]
    inches_from_left_edge = parse(Int, m[:inches_from_left_edge])
    inches_from_top_edge = parse(Int, m[:inches_from_top_edge])
    rectangle_width = parse(Int, m[:rectangle_width])
    rectangle_height = parse(Int, m[:rectangle_height])
    for x in 1:rectangle_width
        for y in 1:rectangle_height
            if fabric[inches_from_left_edge+x, inches_from_top_edge+y] == '.'
                fabric[inches_from_left_edge+x, inches_from_top_edge+y] = id[1]
            else
                fabric[inches_from_left_edge+x, inches_from_top_edge+y] = 'X'
            end
        end
    end
    input_dict[id] = (inches_from_left_edge, inches_from_top_edge, rectangle_width, rectangle_height)
end
# Part 1 https://adventofcode.com/2018/day/3#part1
inches_sum = 0
for x in 1:fabric_size[1]
    for y in 1:fabric_size[2]
        global inches_sum
        if fabric[x, y] == 'X'
            inches_sum += 1
        end
    end
end
println(inches_sum)
# Part 2 https://adventofcode.com/2018/day/3#part2
overlaps_dict = Dict{String, Bool}()
for kv in input_dict
    id = kv[1]
    inches_from_left_edge = kv[2][1]
    inches_from_top_edge = kv[2][2]
    rectangle_width = kv[2][3]
    rectangle_height = kv[2][4]
    overlaps = false
    for x in 1:rectangle_width
        for y in 1:rectangle_height
            if fabric[inches_from_left_edge+x, inches_from_top_edge+y] == 'X' && overlaps == false
                overlaps = true
            end
        end
    end
    overlaps_dict[id] = overlaps
end
for kv in overlaps_dict
    if kv[2] == false
        println(kv[1], " ", kv[2])
    end
end