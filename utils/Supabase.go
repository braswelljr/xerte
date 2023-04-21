package utils

import (
	"os"

	"github.com/nedpals/supabase-go"
)

// Supabase - Initialize Supabase client
//
//	@param url string
//	@param key string
//	@return void
func Supabase() *supabase.Client {
	// get the supabase url and key from the environment variables
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	var debug bool = true
	Supabase := supabase.CreateClient(url, key, debug)

	return Supabase
}
