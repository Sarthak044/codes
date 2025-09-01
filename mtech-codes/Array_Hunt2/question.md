You are given an integer array and a target value. Your task is to perform a linear search on the array to find the target. Linear search means scanning each element one by one from left to right.

Input Format

The first line contains an integer N — the number of elements in the array.

The second line contains N space-separated integers — the elements of the array.

The third line contains an integer X — the target value to search for.

Constraints

The number of elements, N, is between 1 and 100,000.

Array elements are integers in the range -10^9 to 10^9.

The target value, X, is also an integer in the range -10^9 to 10^9.

Output Format

Print a single integer: -> The index (0-based) of the first occurrence of X in the array, if it exists.

-> -1, if the target value is not found

Sample Input 0

5
2 4 6 8 10
8
Sample Output 0

3
Explanation 0

In this example, the target value 8 is found at index 3 (0-based indexing).

Sample Input 1

6
1 3 5 7 9 11
4
Sample Output 1

-1