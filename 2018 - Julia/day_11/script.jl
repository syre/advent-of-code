grid_serial_number = 8561

fuel_cell_grid_size = (300, 300)

fuel_cell_grid = zeros(Int, fuel_cell_grid_size)
# Populate fuel cell grid
for x in 1:fuel_cell_grid_size[1]
    for y in 1:fuel_cell_grid_size[2]
        rack_id = x + 10
        power_level = rack_id * y
        power_level += grid_serial_number
        power_level *= rack_id
        if power_level > 99
            power_level = parse(Int, string(power_level)[end-2])
        else
            power_level = 0
        end
        power_level -= 5
        fuel_cell_grid[x,y] = power_level
    end
end

# Part 1 - https://adventofcode.com/2018/day/11#part1
# Generate a dict of top-left fuel cell cordinate -> total power level
total_powers_dict = Dict{Tuple, Int}()
for x in 1:fuel_cell_grid_size[1]-2
    for y in 1:fuel_cell_grid_size[2]-2
        total_powers_dict[x, y] = sum(fuel_cell_grid[x:x+2, y:y+2])
    end
end

println(maximum(values(total_powers_dict)))
println(reduce((x, y) -> total_powers_dict[x] > total_powers_dict[y] ? x : y, keys(total_powers_dict)))

# Part 2 - https://adventofcode.com/2018/day/11#part2
# Generate a dict of top-left fuel cell cordinate, size -> total power level
total_powers_dict = Dict{Tuple, Int}()
for i in 0:299
    println(i)
    for x in 1:fuel_cell_grid_size[1]-i
        for y in 1:fuel_cell_grid_size[2]-i
            total_powers_dict[x, y, i+1] = sum(view(fuel_cell_grid, x:x+i, y:y+i))
        end
    end
end

println(maximum(values(total_powers_dict)))
println(reduce((x, y) -> total_powers_dict[x] > total_powers_dict[y] ? x : y, keys(total_powers_dict)))