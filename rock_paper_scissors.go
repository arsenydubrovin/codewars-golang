package main

func Rps(p1, p2 string) (res string) {
	if p1 == p2 {
		res = "Draw!"
	} else if p1 == "scissors" && p2 == "paper" ||
		p1 == "paper" && p2 == "rock" ||
		p1 == "rock" && p2 == "scissors" {
		res = "Player 1 won!"
	} else {
		res = "Player 2 won!"
	}
	return
}
