/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (55.83%)
 * Likes:    1292
 * Dislikes: 132
 * Total Accepted:    519.7K
 * Total Submissions: 930.8K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 * 
 * Example 1:
 * 
 * 
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "rat", t = "car"
 * Output: false
 * 
 * 
 * Note:
 * You may assume the string contains only lowercase alphabets.
 * 
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 * 
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }

	sm := map[string]int{}
  	for i := 0; i < len(s); i++ {
      	if _, ok := sm[string(s[i])]; ok {
        	sm[string(s[i])] = sm[string(s[i])] +1
      	}else{
        	sm[string(s[i])] = 1
    	}
  	}

  	for i:=0;i<len(t);i++{ 
    	if _, ok := sm[string(t[i])]; ok {
        	sm[string(t[i])] = sm[string(t[i])] -1
		}else{
        	return false
      	}
  	}

  	for _, v := range sm {
      	if v != 0 {
        	return false
      	}
  	}

	return true
}
// @lc code=end

