from random import randint
from numpy import median

trials = []

for i in range(100000):
    hit = False
    shots = 0
    while (not hit):
        shots += 1
        if randint(1,10) == 1:
            hit = True
            trials.append(shots)

print(median(trials))
