# binomial_sim
Math can be solved with nomenclature and distributions... but it's more fun to do with brute force

Here's the problem:  "A shooter can hit the target 1 out of 10 attempts. How many attempts should he take to ensure that he has 50% chance of hitting the target at least once?"

This can be solved with the help of the Binomial Distribution (https://en.wikipedia.org/wiki/Binomial_distribution)... but that requires math.

So this repo presents a solution in Python, and Go (with and without goroutines... this whole repo is just an excuse for me to play with goroutines).

Instead of actually doing the math, we're going to run 10,000 simulations while keeping track of how many shots it takes to hit the target.  Then we're going to find the midway point of that distribution, and discover that it takes 7 attempts to make sure the target is hit half of the time.
