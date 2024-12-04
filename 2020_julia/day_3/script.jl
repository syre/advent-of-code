using DelimitedFiles

lines = map(collect, readlines("input"))
grid = permutedims(hcat(lines...))
println(size(grid)[1])
println(size(grid)[2])

function generate_indices(grid)
	x_max = size(grid)[1]
	y_max = size(grid)[2]
	indices = Vector{Tuple, 100}
	for i in 2:x_max:
		for j in 4:y_max:3
			if j 
			append!(indices,grid[i,j])
		end
	end
end

#for (x,y) in (zip(range(4, size(grid)[2], step=3), range(2, size(grid)[1], step=1)))
#	println(y,",",x," - ",grid[y,x])
#end
