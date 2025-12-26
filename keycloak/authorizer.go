package keycloak

type Authorizer interface {
	Register(manifest Manifest) error
}
