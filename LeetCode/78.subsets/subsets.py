

def subsets(nums):
	ans = set()
	for p1 in range(len(nums)):
		p2 = p1
		while p2 <= len(nums):
			ans.add(tuple(nums[p1:p2]))
			p2 += 1
	return ans

print(subsets([1, 2, 3]))
print(subsets([0]))