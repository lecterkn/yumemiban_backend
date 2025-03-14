package port

type NovelRepository interface {
	GenerateNovel(string) (*string, error)
}
