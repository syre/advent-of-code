using DelimitedFiles
using Base.Rounding

input_array = readdlm("input", Int)

function calculate_fuel(mass)
    return Int(round(mass/3, Base.Rounding.RoundToZero))-2
end

# Part 1 https://adventofcode.com/2019/day/1
sum_fuel = 0
for mass in input_array
    global sum_fuel
    sum_fuel += calculate_fuel(mass)
end
println(sum_fuel)

# Part 2 https://adventofcode.com/2019/day/1#part2
sum_fuel = 0
for mass in input_array
    global sum_fuel
    fuel = calculate_fuel(mass)
    sum_fuel += fuel

    added_fuel = calculate_fuel(fuel)
    while added_fuel >= 0
        sum_fuel += added_fuel
        added_fuel = calculate_fuel(added_fuel)
    end
end

println(sum_fuel)