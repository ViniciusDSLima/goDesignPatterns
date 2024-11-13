package main

type client struct {
	name string
	age  int
}

type clientBuilder struct {
	name string
	age  int
}

func getClientBuilder() clientBuilder {
	return clientBuilder{}
}

func (b *clientBuilder) buildName() {
	b.name = "Vinicius"
}

func (b *clientBuilder) buildAge() {
	b.age = 19
}

func (b *clientBuilder) buildClient() client {
	return client{b.name, b.age}
}

type builder struct {
	client clientBuilder
}

func newBuilder(b clientBuilder) *builder {
	return &builder{
		client: b,
	}
}

func (b *builder) buildClient() client {
	b.client.buildName()
	b.client.buildAge()

	return b.client.buildClient()
}

//
//func main() {
//
//	clientBuilder := getClientBuilder()
//
//	builder := newBuilder(clientBuilder)
//
//	client := builder.buildClient()
//
//	fmt.Println("client age: ", client.age)
//	fmt.Println("client name: ", client.name)
//}
