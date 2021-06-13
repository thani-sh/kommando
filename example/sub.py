#!/usr/bin/env python3

import sys

args = sys.argv[1:]
nums = [a for a in args if not a.startswith("-")]

if "-i" in args:
    nums = nums[::-1]

x = int(nums[0])
y = int(nums[1])
print(x - y)
