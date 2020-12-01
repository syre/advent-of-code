using DelimitedFiles

input_array = readdlm("input", '\n', String)

regex = r"position=<[ ]{0,2}(?<p_x>[-\d]*),[ ]{0,2}(?<p_y>[-\d]*)> velocity=<[ ]{0,2}(?<v_x>[-\d]*),[ ]{0,2}(?<v_y>[-\d]*)>"

light_vector = Vector{Tuple}()

for input in input_array
    m = match(regex, input)
    p_x = parse(Int, m[:p_x])
    p_y = parse(Int, m[:p_y])
    v_x = parse(Int, m[:v_x])
    v_y = parse(Int, m[:v_y])
    push!(light_vector, (p_x, p_y, v_x, v_y))
end

function print_sky(sky)
    (rows, cols) = size(sky)
    for x in 1:rows
        for y in 1:cols
            print(sky[x, y])
        end
        println(" ")
    end
end

function print_sky_to_file(sky)
    (rows, cols) = size(sky)
    open("output.txt", "w") do file
        for x in 1:rows
            for y in 1:cols
                write(file, sky[x, y])
            end
            write(file, "\n")
        end
    end
end

# Part 1 - https://adventofcode.com/2018/day/10#part1
max_sky_size_diameter = Inf
for i in 0:100000000
    global max_sky_size_diameter
    global sky
    min_row = 0
    max_row = 0
    min_col = 0
    max_col = 0
    for (p_x, p_y, v_x, v_y) in light_vector
        min_row = min(min_row, p_x + (v_x*i))
        max_row = max(max_row, p_x + (v_x*i))
        min_col = min(min_col, p_y + (v_y*i))
        max_col = max(max_col, p_y + (v_y*i))
    end
    sky_size = (max_row - min_row, max_col - min_col)
    if sky_size[1] * sky_size[2] > max_sky_size_diameter
        break
    end
    max_sky_size_diameter = sky_size[1] * sky_size[2]
    if sky_size[1] > 250 || sky_size[2] > 250
        continue
    end
    sky = fill('.', sky_size)

    for (p_x, p_y, v_x, v_y) in light_vector
        sky[(p_x + (v_x*i)) - min_row, (p_y + (v_y*i)) - min_col] = '#'
    end
    print_sky(sky)
    print_sky_to_file(sky)

    # Part 2 - https://adventofcode.com/2018/day/10#part2
    println(i)
end
