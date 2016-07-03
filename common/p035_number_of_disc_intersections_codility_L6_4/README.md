# p035_number_of_disc_intersections_codility_L6_4
---
http://rafal.io/posts/codility-intersecting-discs.html

1 Attempt 1 - Brute Force
Let’s first try to come up with a brute force solution – this is always a good start.

Given array indices ii,jj (j>ij>i), if the discs are located at (0,i)(0,i) and (0,j)(0,j) respectively, then the discs will intersect iff the following holds:

j−i≤A[i]+A[j]
j−i≤A[i]+A[j]

Given this predicate, simply iterate over array AA, and check for how many subsequent discs the above predicate holds. This gives us an O(n2)O(n2) algorithm.

2 Improving to O(NlogN)O(Nlog⁡N)
We can rewrite the above predicate:

j−i≤A[i]+A[j]A[i]+i≥−(A[j]−j)
j−i≤A[i]+A[j]A[i]+i≥−(A[j]−j)

Now we are getting closer to a solution. We see that the solution is dependent on a ≥≥ predicate where each side is a linear transformation of the input array AA. We can proceed as follows:

Produce two arrays containing A[i]+iA[i]+i and −(A[j]−j)−(A[j]−j).
Sort them - O(Nlog(N))O(Nlog⁡(N))
Compute for how many pairs the above predicate holds.
To do the last step above, simply go through array with sorted values A[i]+iA[i]+i, and for each value xx, check how many values from the sorted array −(A[j]−j)−(A[j]−j) are smaller than xx. This can be done with binary search in O(log(N))O(log⁡(N)).

Crucially, we must subtract what was counted twice and any degenerate solutions – these where the predicate j−i≤A[i]+A[j]j−i≤A[i]+A[j] holds true since j−i≤0j−i≤0. To do this, subtract N∗(N+1)2N∗(N+1)2 from the total.


# note 
There exists O(n) solution

### Input Format 

### Constraints

### Output Format 

### Sample Input
```
```
### Sample Output
```
```
### Explanation
