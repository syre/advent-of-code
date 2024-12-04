input_string = open("input") do file
    read(file, String)
end

regex = r"(?<players>[\d]*) players; last marble is worth (?<points>[\d]*) points"

m = match(regex, input_string)
players = parse(Int, m[:players])
points = parse(Int, m[:points])

num_players = players
ends_after_rounds = points

# Push the initial 0 marble
marble_circle = Vector{Int}()
push!(marble_circle, 0)

player_scores = Dict{Int, Int}(x=>0 for x in 1:num_players)
current_index = 1

function print_marble_circle(marble_circle, current_index, current_player)
    print("[", current_player, "]", " ")
    for (index, marble) in enumerate(marble_circle)
        if index == current_index
            print("(", marble, ")", " ")
        else
            print(marble , " ")
        end
    end
    print("\n")
end

# Part 1 - https://adventofcode.com/2018/day/9#part1
for i in 1:ends_after_rounds
    global marble_circle
    global current_index
    current_player = (i % num_players) + 1
    if i % 23 == 0
        player_scores[current_player] += i
        new_current_index = current_index - 7
        if new_current_index < 1
            new_current_index = length(marble_circle) + new_current_index
        end
        player_scores[current_player] += splice!(marble_circle, new_current_index)
        current_index = new_current_index
    else
        new_current_index = ((current_index) % length(marble_circle)) + 2
        insert!(marble_circle, new_current_index, i)
        current_index = new_current_index
    end
    # print_marble_circle(marble_circle, current_index, current_player)
end

println(maximum(values(player_scores)))

# Part 2 - https://adventofcode.com/2018/day/9#part2

# Push the initial 0 marble
marble_circle = Vector{Int}()
push!(marble_circle, 0)

player_scores = Dict{Int, Int}(x=>0 for x in 1:num_players)
current_index = 1
# Times the ends_after_rounds with 100
for i in 1:ends_after_rounds*100
    global marble_circle
    global current_index
    current_player = (i % num_players) + 1
    if i % 23 == 0
        player_scores[current_player] += i
        new_current_index = current_index - 7
        if new_current_index < 1
            new_current_index = length(marble_circle) + new_current_index
        end
        player_scores[current_player] += splice!(marble_circle, new_current_index)
        current_index = new_current_index
    else
        new_current_index = ((current_index) % length(marble_circle)) + 2
        insert!(marble_circle, new_current_index, i)
        current_index = new_current_index
    end
    # print_marble_circle(marble_circle, current_index, current_player)
end

println(maximum(values(player_scores)))