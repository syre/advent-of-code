import Pkg; Pkg.add("StatsBase")
using StatsBase
using DelimitedFiles

# Part 1 https://adventofcode.com/2019/day/3
input_array = readdlm("input")

board_matrix = Matrix{Char}(undef, 10000, 10000)
fill!(board_matrix, '.')

for wire_sequence in input_array
    x_index = 100
    y_index = 100
    for operation in split(wire_sequence, ",")
        direction = chop(operation, head=0, tail=length(operation)-1)
        number = parse(Int, chop(operation, head=1, tail=0))
        println(direction, " ", number)
        if direction == 'U'
            for i in y_index:y_index-number
                if board_matrix[x_index][i] == '.'
                    board_matrix[x_index][i] = 'X'
                else
                    board_matrix[x_index][i] = '|'
                end
            end
            y_index -= number
        elseif direction == 'R'
            for i in x_index:x_index-number
                if board_matrix[i][y_index] == '.'
                    board_matrix[i][y_index] = 'X'
                else
                    board_matrix[i][y_index] = '-'
                end
            end
            x_index -= number
        elseif direction == 'D'
            for i in y_index:y_index+number
                if board_matrix[x_index][i] == '.'
                    board_matrix[x_index][i] = 'X'
                else
                    board_matrix[x_index][i] = '|'
                end
            end
            y_index -= number
        elseif direction == 'L'
            for i in y_index:x_index+number
                if board_matrix[x_index][i] == '.'
                    board_matrix[x_index][i] = 'X'
                else
                    board_matrix[x_index][i] = '-'
                end
            end
            x_index -= number
        end
    end
end


print(countmap(board_matrix))