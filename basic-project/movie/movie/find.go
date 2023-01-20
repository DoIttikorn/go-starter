package movie

func Find(imdbID string) string {
	switch imdbID {
	case "tt0111161":
		return "The Shawshank Redemption"
	case "tt0068646":
		return "The Godfather"
	default:
		return "Movie not found"
	}
}
