package logger

import "testing"

func TestLogger(t *testing.T) {

	log1 := NewLogger("Tester")
	log2 := NewLogger("Miner")
	log3 := NewLogger("Node")

	log1.LogGreen("Tester Logs Green!")

	log2.LogYellow("Miner Logs some Basic Data!")
	log2.LogYellow("Hashing... || Time: Now")
	log2.LogGreen("Solution Found!")

	log3.LogBlue("Loading Node...")
	log3.LogRed("Loading failed, rebooting...")
	log3.LogGreen("Node Online!")
}
