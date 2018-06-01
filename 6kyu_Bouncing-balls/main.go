// https://www.codewars.com/kata/5544c7a5cb454edb3c000047

// A child plays with a ball on the nth floor of a big building. The height of this floor is known.
// (float parameter "h" in meters. Condition 1) : h > 0)
// He lets out the ball. The ball bounces for example to two-thirds of its height.
// (float parameter "bounce". Condition 2) : 0 < bounce < 1)
// His mother looks out of a window that is 1.5 meters from the ground.
// (float parameters "window". Condition 3) : window < h).
// How many times will the mother see the ball either falling or bouncing in front of the window?
// If all three conditions above are fulfilled, return a positive integer, otherwise return -1.
// Note
// You will admit that the ball can only be seen if the height of the rebouncing ball is stricty greater than the window parameter.
// Example:
//     h = 3, bounce = 0.66, window = 1.5, result is 3
//     h = 3, bounce = 1, window = 1.5, result is -1 (Condition 2) not fulfilled).

package main

func BouncingBall1(h, bounce, window float64) int {
	if (h > 0) && (0 < bounce && bounce < 1) {
		jumps := -1
		for ; h > window; h *= bounce {
			jumps += 2
		}
		return jumps
	}
	return -1
}

func BouncingBall2(h, bounce, window float64) int {
	if h <= window || bounce <= 0 || bounce >= 1 {
		return -1
	} else {
		return 2 + BouncingBall2((h*bounce), bounce, window)
	}
}

func main() {
	println(BouncingBall1(3, 0.66, 1.5))
	println(BouncingBall2(3, 0.66, 1.5))
}
