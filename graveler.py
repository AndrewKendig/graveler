# Completion Time for 1000000: 0 hours, 1 minutes, 31 seconds

import random
import math
import time
from itertools import repeat

start_time = time.perf_counter()

items = [1,2,3,4]
numbers = [0,0,0,0]
rolls = 0
maxOnes = 0

while numbers[0] < 177 and rolls < 1000000:
    numbers = [0,0,0,0]
    for i in repeat(None, 231):
        roll = random.choice(items)
        numbers[roll-1] = numbers[roll-1] + 1
    rolls = rolls + 1
    if numbers[0] > maxOnes:
        maxOnes = numbers[0]

end_time = time.perf_counter()
duration_seconds = end_time - start_time
hours = duration_seconds // 3600
minutes = (duration_seconds % 3600) // 60
seconds = duration_seconds % 60

print(f"Duration: {int(hours)} hours, {int(minutes)} minutes, {int(seconds)} seconds")

print("Highest Ones Roll:",maxOnes)
print("Number of Roll Sessions: ",rolls)
