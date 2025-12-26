package keycloak

type KeycloakAuthorizer struct {
	client *KeycloakClient
}

func NewKeycloakAuthorizer(client *KeycloakClient) *KeycloakAuthorizer {
	return &KeycloakAuthorizer{client: client}
}

func (a *KeycloakAuthorizer) Register(manifest Manifest) error {
	for _, model := range manifest.Models {
		if err := a.ensureModel(manifest.Service, model); err != nil {
			return err
		}
	}
	return nil
}

func (a *KeycloakAuthorizer) ensureModel(service string, model Model) error {
	// 1. ensure resource
	if err := a.client.EnsureResource(
		model.Name,
		model.URIs,
	); err != nil {
		return err
	}

	// 2. ensure scopes
	for _, scope := range model.Scopes {
		if err := a.client.EnsureScope(scope.Name); err != nil {
			return err
		}
	}

	// 3. link scopes to resource
	for _, scope := range model.Scopes {
		if err := a.client.EnsureScopeLinked(
			model.Name,
			scope.Name,
		); err != nil {
			return err
		}
	}

	return nil
}
