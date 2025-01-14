package director

import "fmt"

// gqlFieldsProvider is responsible for generating GraphQL queries that request for all fields for given type
type gqlFieldsProvider struct{}

func (fp *gqlFieldsProvider) Page(item string) string {
	return fmt.Sprintf(`data {
		%s
	}
	pageInfo {%s}
	totalCount
	`, item, fp.ForPageInfo())
}

func (fp *gqlFieldsProvider) ForApplication() string {
	return fmt.Sprintf(`id
		name
		description
		labels
		status {condition timestamp}
		webhooks {%s}
		healthCheckURL
		apis {%s}
		eventAPIs {%s}
		documents {%s}
	`, fp.ForWebhooks(), fp.Page(fp.ForAPIDefinition()), fp.Page(fp.ForEventAPI()), fp.Page(fp.ForDocument()))
}

func (fp *gqlFieldsProvider) ForWebhooks() string {
	return fmt.Sprintf(
		`id
		applicationID
		type
		url
		auth {
		  %s
		}`, fp.ForAuth())
}

func (fp *gqlFieldsProvider) ForAPIDefinition() string {
	return fmt.Sprintf(`		id
		name
		description
		spec {%s}
		targetURL
		group
		auths {%s}
		defaultAuth {%s}
		version {%s}`, fp.ForApiSpec(), fp.ForRuntimeAuth(), fp.ForAuth(), fp.ForVersion())
}

func (fp *gqlFieldsProvider) ForApiSpec() string {
	return fmt.Sprintf(`data
		format
		type
		fetchRequest {%s}`, fp.ForFetchRequest())
}

func (fp *gqlFieldsProvider) ForFetchRequest() string {
	return fmt.Sprintf(`url
		auth {%s}
		mode
		filter
		status {condition timestamp}`, fp.ForAuth())
}

func (fp *gqlFieldsProvider) ForRuntimeAuth() string {
	return fmt.Sprintf(`runtimeID
		auth {%s}`, fp.ForAuth())
}

func (fp *gqlFieldsProvider) ForVersion() string {
	return `value
		deprecated
		deprecatedSince
		forRemoval`
}

func (fp *gqlFieldsProvider) ForPageInfo() string {
	return `startCursor
		endCursor
		hasNextPage`
}

func (fp *gqlFieldsProvider) ForEventAPI() string {
	return fmt.Sprintf(`
			id
			applicationID
			name
			description
			group 
			spec {%s}
			version {%s}
		`, fp.ForEventSpec(), fp.ForVersion())
}

func (fp *gqlFieldsProvider) ForEventSpec() string {
	return fmt.Sprintf(`data
		type
		format
		fetchRequest {%s}`, fp.ForFetchRequest())
}

func (fp *gqlFieldsProvider) ForDocument() string {
	return fmt.Sprintf(`
		id
		applicationID
		title
		displayName
		description
		format
		kind
		data
		fetchRequest {%s}`, fp.ForFetchRequest())
}

func (fp *gqlFieldsProvider) ForAuth() string {
	return fmt.Sprintf(`credential {
				... on BasicCredentialData {
					username
					password
				}
				...  on OAuthCredentialData {
					clientId
					clientSecret
					url
					
				}
			}
			additionalHeaders
			additionalQueryParams
			requestAuth { 
			  csrf {
				tokenEndpointURL
				credential {
				  ... on BasicCredentialData {
				  	username
					password
				  }
				  ...  on OAuthCredentialData {
					clientId
					clientSecret
					url
					
				  }
			    }
				additionalHeaders
				additionalQueryParams
			}
			}
		`)
}

func (fp *gqlFieldsProvider) ForLabel() string {
	return `key
			value`
}

func (fp *gqlFieldsProvider) ForRuntime() string {
	return fmt.Sprintf(`id
		name
		description
		labels 
		status {condition timestamp}
		agentAuth {%s}`, fp.ForAuth())
}
