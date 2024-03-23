package main

func main() {

	client, err := NewHelloServiceClient("tcp", "127.0.0.1:9901")
	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Hello("sane", &reply)
	if err != nil {
		panic(err)
	}

	println(reply)
}
