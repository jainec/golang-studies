package main

import "context"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "12324")
	bookHotel(ctx, "Test Hotel")
}

// context comes as first parameter always (good practice)
func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	println(token)
}
