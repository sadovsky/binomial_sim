# binomial_sim
Math can be solved with nomenclature and distributions... but it's more fun to do with brute force

Here's the problem:  "A shooter can hit the target 1 out of 10 attempts. How many attempts should he take to ensure that he has 50% chance of hitting the target at least once?"

This can be solved with the help of the Binomial Distribution (https://en.wikipedia.org/wiki/Binomial_distribution)... but that requires math.

So this repo presents a solution in Python, and Go.

Instead of actually doing the math, we're going to run 100,000 simulations (way more than we need) while keeping track of how many shots it takes to hit the target.  Then we're going to find the midway point of that distribution, and discover that it takes 7 attempts to make sure the target is hit half of the time

Both programs get the same results, but go is faster (as, hopefully expected!).  time isn't the best benchmarking tool, but you get the point. 

```
hypermegadata:binomial_sim sadovsky$ time python shoot.py
7.0

real	0m2.348s
user	0m2.261s
sys	0m0.061s

hypermegadata:binomial_sim sadovsky$ time ./shoot
7

real	0m0.058s
user	0m0.050s
sys	0m0.005s
```
