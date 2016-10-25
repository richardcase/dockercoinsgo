package client

type RngClient interface {
	GenerateRandomString(length int32) (string, error)
}
