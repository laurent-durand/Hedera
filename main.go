package main
import "fmt"
func main() {
    fmt.Println("--- Hedera: Process Tree Visualizer (Go) ---")
    fmt.Println("PID 1 (systemd) ──┐")
    fmt.Println("                 ├─ PID 42 (bash)")
    fmt.Println("                 └─ PID 101 (hedera)")
}
