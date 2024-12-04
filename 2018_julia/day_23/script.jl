using DelimitedFiles

input_array = readdlm("input", '\n', String)

regex = r"pos=<(?<x>-?\d*),(?<y>-?\d*),(?<z>-?\d*)>, r=(?<r>\d*)"

function calculate_distance((x_1, y_1, z_1), (x_2, y_2, z_2))
    return abs(x_1-x_2) + abs(y_1-y_2) + abs(z_1-z_2)
end

# Part 1 - https://adventofcode.com/2018/day/23#part1

# Generate a dict of nano bot range -> number of bots in range
nano_bot_dict = Dict{Int, Int}()
nano_bots = Vector{Tuple}()
for input in input_array
    m = match(regex, input)
    x_1 = parse(Int, m[:x])
    y_1 = parse(Int, m[:y])
    z_1 = parse(Int, m[:z])
    r_1 = parse(Int, m[:r])
    push!(nano_bots, (x_1, y_1, z_1, r_1))
    nano_bot_dict[r_1] = 0
    for input in input_array
        m = match(regex, input)
        x_2 = parse(Int, m[:x])
        y_2 = parse(Int, m[:y])
        z_2 = parse(Int, m[:z])
        r_2 = parse(Int, m[:r])

        if calculate_distance((x_1, y_1, z_1), (x_2, y_2, z_2)) <= r_1
            nano_bot_dict[r_1] += 1
        end
    end
end

max_range = maximum(keys(nano_bot_dict))
println(max_range)
println(nano_bot_dict[max_range])

# Part 2 - https://adventofcode.com/2018/day/23#part2 UNFINISHED

min_x = Inf
max_x = -Inf
min_y = Inf
max_y = -Inf
min_z = Inf
max_z = -Inf

for bot in nano_bots
    global min_x
    global max_x
    global min_y
    global max_y
    global min_z
    global max_z

    x_1 = bot[1]
    y_1 = bot[2]
    z_1 = bot[3]
    r_1 = bot[4]

    min_x = min(min_x, x_1)
    max_x = max(max_x, x_1)
    min_y = min(min_y, y_1)
    max_y = max(max_y, y_1)
    min_z = min(min_z, z_1)
    max_z = max(max_z, z_1)
end

max_in_range = 0
println(min_x, " ", max_x)
println(min_y, " ", max_y)
println(min_z, " ", max_z)

for x in min_x:max_x
    println(x)
    for y in min_y:max_y
        for z in min_z:max_z
            for bot in nano_bots

            end
        end
    end
end