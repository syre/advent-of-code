using DelimitedFiles

input_array = readdlm("input", '\n', String)

regex = r"Step (?<requirement>[A-Z]{1}) must be finished before step (?<char_step>[A-Z]{1}) can begin."

name_to_requirements_dict = Dict{Char, Vector{Char}}()
completed_steps = Vector{Char}()

# Create dict of Step name -> List of requirement steps
for input in input_array
    m = match(regex, input)
    requirement = m[:requirement][1]
    char_step = m[:char_step][1]
    get!(name_to_requirements_dict, char_step, Vector{Char}())
    get!(name_to_requirements_dict, requirement, Vector{Char}())

    push!(name_to_requirements_dict[char_step], requirement)
end

# Part 1 - https://adventofcode.com/2018/day/7#part1
# Process all steps
while(!isempty(name_to_requirements_dict))
    # Remove the first key from the sorted array of keys in name_to_requirements_dict and add to completed.
    key = popfirst!(sort(collect(keys(filter(p->length(p.second)==0, name_to_requirements_dict)))))
    pop!(name_to_requirements_dict, key)
    push!(completed_steps, key)
    # remove the key from all remaining requirements arrays.
    for requirements_array in values(name_to_requirements_dict)
        if key in requirements_array
            deleteat!(requirements_array, findfirst(x->x==key, requirements_array))
        end
    end
end
# Print answer
println(join(completed_steps))

# Part 2 - https://adventofcode.com/2018/day/7#part2

name_to_requirements_dict = Dict{Char, Vector{Char}}()
completed_steps = Vector{Char}()

# Create dict of Step name -> List of requirement steps
for input in input_array
    m = match(regex, input)
    requirement = m[:requirement][1]
    char_step = m[:char_step][1]
    get!(name_to_requirements_dict, char_step, Vector{Char}())
    get!(name_to_requirements_dict, requirement, Vector{Char}())

    push!(name_to_requirements_dict[char_step], requirement)
end

extra_time_to_complete_step = 60
time_to_complete_dict = Dict{Char, Int}(map(x->(x[1],x[2]+extra_time_to_complete_step), map(reverse, enumerate("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))))
workers_available = 5
time_counter = 0
time_started_dict = Dict{Char, Int}()

while(!isempty(name_to_requirements_dict) || !isempty(time_started_dict))
    global time_counter
    global workers_available

    println(time_counter)
    for (step_key, time_started) in time_started_dict
        if time_counter >= (time_to_complete_dict[step_key] + time_started)
            push!(completed_steps, step_key)
            # remove the key from all remaining requirements arrays.
            for requirements_array in values(name_to_requirements_dict)
                if step_key in requirements_array
                    deleteat!(requirements_array, findfirst(x->x==step_key, requirements_array))
                end
            end
            pop!(time_started_dict, step_key)
            workers_available += 1
        end
    end
    available_keys = sort(collect(keys(filter(p->length(p.second)==0, name_to_requirements_dict))))
    while(workers_available > 0 && !isempty(available_keys))
        workers_available -= 1
        key = popfirst!(available_keys)
        pop!(name_to_requirements_dict, key)
        time_started_dict[key] = time_counter
        println("took ", key)
    end
    time_counter += 1

end

println(join(completed_steps))