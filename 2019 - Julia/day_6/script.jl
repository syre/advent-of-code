using DelimitedFiles

input_array = readdlm("input", '\n', String, ')')
println(input_array)
for input in input_array
    println(input)
end

orbit_dict = Dict{String, Vector{String}}()

for (orbit, orbits) in input_array
    println(orbit, orbits)
    append!(get(orbit_dict, orbit, Vector{String}()), orbits)
end

println(orbit_dict)