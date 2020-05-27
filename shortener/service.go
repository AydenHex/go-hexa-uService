package shortener // service.go

type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
