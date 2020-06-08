package common

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	RequestMethodGET  = "GET"
	//RequestMethodPOST = "POST"

	SignatureMethodHMacSha256 = "HmacSHA256"
)

type Client struct {
	*http.Client

	credential CredentialInterface
	opts       Opts
}

type Opts struct {
	Method          string
	Region          string
	Host            string
	Path            string
	SignatureMethod string
	Schema          string

	Logger *logrus.Logger
}

type CredentialInterface interface {
	GetSecretId() (string, error)
	GetSecretKey() (string, error)

	Values() (CredentialValues, error)
}

type CredentialValues map[string]string

type Credential struct {
	SecretId  string
	SecretKey string
}

func (cred Credential) GetSecretId() (string, error) {
	return cred.SecretId, nil
}

func (cred Credential) GetSecretKey() (string, error) {
	return cred.SecretKey, nil
}

func (cred Credential) Values() (CredentialValues, error) {
	return CredentialValues{}, nil
}

func NewClient(credential CredentialInterface, opts Opts) (*Client, error) {
	if opts.Method == "" {
		opts.Method = RequestMethodGET
	}
	if opts.SignatureMethod == "" {
		opts.SignatureMethod = SignatureMethodHMacSha256
	}
	if opts.Schema == "" {
		opts.Schema = "https"
	}
	if opts.Logger == nil {
		opts.Logger = logrus.New()
	}
	return &Client{
		&http.Client{},
		credential,
		opts,
	}, nil
}
