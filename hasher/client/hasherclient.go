package client

type HasherClient interface {
	Hash(s string) (string, error)
}
