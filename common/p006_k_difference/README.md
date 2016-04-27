# p006_k_difference [geeks4geeks problem](http://www.practice.geeksforgeeks.org/problem-page.php?pid=347)
---
Given an array of N positive integers  a1, a2 ..... an. Find the number of pairs of integers whose difference is equal to a given number K.

NOTE:Current solution isn't optimal O(n2). To be reimplemented

### Input Format 
The first line of the input contains 'T' denoting the total number of testcases.Then follows the description of testcases.
The first line of each testcase contains two space separated positive integers N and K denoting the size of array and the value of K. The second line contains N space separated positive integers denoting the elements of array.

### Constraints
1<=T<=100
1<=N<=100
1<=K<=100
1<=Arr[i]<=1000

### Output Format 
Output the number of pairs having difference equal to K in a separate line.

### Sample Input
```
2
5 2
6 4 5 8 7
4 2
1 3 1 8
```
### Sample Output
```
3
2
```
### Explanation
