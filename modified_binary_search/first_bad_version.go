package modified_binary_search

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    start, end := 1, n
    var toCheck int
    for start <= end {
        toCheck = (start + end) / 2
        if isBadVersion(toCheck) {
            end = toCheck - 1
        } else {
            start = toCheck + 1
        }
    }
    if isBadVersion(end) {
        return end
    } else {
        return end + 1
    }
}