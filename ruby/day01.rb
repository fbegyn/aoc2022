# __FILE__ lists the complete path of the current file
# ARGV holds all the arguments after calling the script
input = File.read(ARGV[0])

sum = []

for elf in input.split(/\n\n/)
  sum.push(elf.split(/\n/).map(&:to_i).sum)
end

puts "The elf carryng the most calories, holds %d calories." % sum.max
puts "The top 3 elves hold %d calories." % sum.sort[-3, 3].sum
