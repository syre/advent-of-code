using DelimitedFiles

input_array = readdlm("test_input", '\n', String)
tree_dict = Dict()

metadata_counter = 0

function 

for input in input_array
    input_list = split(input)
    println(input_list[0:2])


end
println(input_array)