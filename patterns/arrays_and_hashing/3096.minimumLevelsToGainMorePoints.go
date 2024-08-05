/* https://leetcode.com/problems/minimum-levels-to-gain-more-points/description/
difficulty: medium

PROBLEM
You are given a binary array possible of length n.

Alice and Bob are playing a game that consists of n levels. Some of the levels in the game are impossible to clear while others can always be cleared. In particular, if possible[i] == 0, then the ith level is impossible to clear for both the players. A player gains 1 point on clearing a level and loses 1 point if the player fails to clear it.

At the start of the game, Alice will play some levels in the given order starting from the 0th level, after which Bob will play for the rest of the levels.

Alice wants to know the minimum number of levels she should play to gain more points than Bob, if both players play optimally to maximize their points.

Return the minimum number of levels Alice should play to gain more points. If this is not possible, return -1.

Note that each player must play at least 1 level

SOLUTION
* Calculate total sum for whole array
* Traverse through the array till len-1 and calculate that leftSum is greater than right sum
*/

package arrays_and_hashing

func minimumLevels(possible []int) int {
	sum := 0
	for idx := 0; idx < len(possible); idx++ {
		sum += getPoints(possible[idx])
	}

	left := 0
	for idx := 0; idx < len(possible)-1; idx++ {
		left += getPoints(possible[idx])
		if sum-left < left {
			return idx + 1
		}
	}
	return -1
}

func getPoints(val int) int {
	if val == 0 {
		return -1
	}
	return val
}