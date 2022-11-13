package main

func main() {
	genres := []string{"Action", "Adventure", "Comedy", "Drama", "Fantasy", "Horror", "Mystery", "Romance", "Sci-Fi", "Thriller"}

	for i := 0; i < len(genres); i++ {
		genres[i] = genres[i] + " Movie"
	}
	for i, v := range genres {
		println(i, v)
	}
}
