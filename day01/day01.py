filename = 'day1.txt'
# filename = 'data/day1_ex.txt'
with open(filename) as f:
    rawdata = f.read().splitlines()
a = [i.split('   ') for i in rawdata]
b1 = [int(i[0]) for i in a]
b2 = [int(i[1]) for i in a]
c1,c2 =sorted(b1), sorted(b2)

total_diff=0
for i in range(min(len(c1),len(c2))):
    diff = abs(c1[i]-c2[i])
    total_diff= total_diff + diff
    # print(c1[i],'-',c2[i],'=',diff,'total=',total_diff)
print('part1:',total_diff)


# i was just trying to debug part 1 lol