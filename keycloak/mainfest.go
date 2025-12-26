package keycloak

type Manifest struct {
	Service string
	Models  []Model
}

type Model struct {
	Name   string
	URIs   []string
	Scopes []Scope
}

type Scope struct {
	Name        string
	Description string
}
