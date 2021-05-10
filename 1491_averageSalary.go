/**

 leetcode 1491 Average Salary Excluding the Minimum and Maximum Salary

 Given an array of unique integers salary where salary[i] is the salary of the employee i.

 Return the average salary of employees excluding the minimum and maximum salary.

 */
package main

func average(salary []int) float64 {
	sum, min, max := salary[0], salary[0], salary[0]
	for i := 1; i < len(salary);i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary [i] < min {
			min = salary[i]
		}
		sum += salary[i]

	}
	return float64(sum - min - max) / float64(len(salary) -2)

}
