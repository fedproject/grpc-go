func main() {
	output, err := exec.Command("go", "run", "/root/grpc-go/examples/route_guide/client/client.go").Output()
	if err == nil {
		w.Write(output) // write the output with ResponseWriter
	}
}