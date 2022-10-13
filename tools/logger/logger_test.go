package logger

import "testing"

func TestLogger(t *testing.T) {

	LogGreen("Tester", "All good!")

	LogYellow("Fake Miner", "Miner Logs some Basic Data!")
	LogYellow("Fake Miner", "Hashing... || Time: Now")
	LogGreen("Fake Miner", "Solution Found!")

	LogBlue("Fake Node", "Loading Node...")
	LogRed("Fake Node", "Loading failed, rebooting...")
	LogGreen("Fake Node", "Node Online!")
}
