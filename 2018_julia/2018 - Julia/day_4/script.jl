using DelimitedFiles
using Dates

input_array = readdlm("input_2.txt", '\n', String)

regex = r"\[(?<datetime>\d+\-\d+\-\d+ \d+:\d+)] (?<event>.*)"
new_event_regex = r"Guard #(?<id>\d+) begins shift"

# Make a dict with "guard_id", (timestamp, event) tuple
guards = Dict{String, Vector{Tuple{DateTime, String}}}()
# Sort input as strings which will also make it chronologically.
sorted_timestamps = sort(input_array, dims=1)
current_guard_id = ""
for (index, entry) in enumerate(sorted_timestamps)
    global current_guard_id
    m = match(regex, entry)
    timestamp = DateTime(m[:datetime], "yyyy-mm-dd HH:MM")
    event = m[:event]
    new_event_match = match(new_event_regex, event)
    if new_event_match != nothing
        new_guard_id = new_event_match[:id]
        current_guard_id = new_guard_id
    else
        push!(get!(guards, current_guard_id, Vector()), (timestamp, event))
    end
end
guard_tuples = Vector{Tuple{String, Int, Dict}}()
for (guard_id, vector) in guards
    total_minutes_asleep = 0
    # Make a dict that holds each minute mark and a counter.
    minute_marks_asleep = Dict{Int, Int}(i=>0 for i in 0:59)
    last_datetime = vector[1][1]
    last_event = vector[1][2]
    for event in vector[2:length(vector)]
        if event[2] == "wakes up" && last_event == "falls asleep"
            # Iterate over range from last_datetime to current event-1min in minute increments.
            # -1 min as the moment the guard wakes up counts as awake minute.
            for date in last_datetime:Dates.Minute(1):(event[1]-Dates.Minute(1))
                minute_marks_asleep[Dates.value(Dates.Minute(date))] += 1
            end
            total_minutes_asleep += abs(Dates.value(convert(Dates.Minute, (event[1]-last_datetime))))
        end
        last_datetime = event[1]
        last_event = event[2]
    end
    push!(guard_tuples, (guard_id, total_minutes_asleep, minute_marks_asleep))
end


# Part 1 & Part 2
for guard_tuple in guard_tuples
    key_of_largest_sleep_moment = reduce((x,y) -> guard_tuple[3][x] > guard_tuple[3][y] ? x : y, keys(guard_tuple[3]))
    println(guard_tuple[1], " ", guard_tuple[2], " ", key_of_largest_sleep_moment, " ", guard_tuple[3][key_of_largest_sleep_moment])
end