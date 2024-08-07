package goas

import (
	"bytes"
	"encoding/json"

	"github.com/elliotchance/orderedmap/v2"
)

const (
	OpenAPIVersion = "3.0.0"

	ContentTypeText        = "text/plain"
	ContentTypeJson        = "application/json"
	ContentTypeOctetStream = "application/octet-stream"
	ContentTypeForm        = "multipart/form-data"
)

type OpenAPIObject struct {
	OpenAPI string         `json:"openapi"` // Required
	Info    InfoObject     `json:"info"`    // Required
	Servers []ServerObject `json:"servers,omitempty"`
	Paths   PathsObject    `json:"paths"` // Required

	Components ComponentsOjbect      `json:"components,omitempty"` // Required for Authorization header
	Security   []map[string][]string `json:"security,omitempty"`

	// Tags
	// ExternalDocs
}

type ServerObject struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`

	// Variables
}

type InfoObject struct {
	Title          string         `json:"title"`
	Description    string         `json:"description,omitempty"`
	TermsOfService string         `json:"termsOfService,omitempty"`
	Contact        *ContactObject `json:"contact,omitempty"`
	License        *LicenseObject `json:"license,omitempty"`
	Version        string         `json:"version"`
}

type ContactObject struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type LicenseObject struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type PathsObject map[string]*PathItemObject

type PathItemObject struct {
	Ref         string           `json:"$ref,omitempty"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Get         *OperationObject `json:"get,omitempty"`
	Post        *OperationObject `json:"post,omitempty"`
	Patch       *OperationObject `json:"patch,omitempty"`
	Put         *OperationObject `json:"put,omitempty"`
	Delete      *OperationObject `json:"delete,omitempty"`
	Options     *OperationObject `json:"options,omitempty"`
	Head        *OperationObject `json:"head,omitempty"`
	Trace       *OperationObject `json:"trace,omitempty"`

	// Servers
	// Parameters
}

type OperationObject struct {
	Responses ResponsesObject `json:"responses"` // Required

	Tags        []string           `json:"tags,omitempty"`
	OperationID string             `json:"operationId,omitempty"`
	Summary     string             `json:"summary,omitempty"`
	Description string             `json:"description,omitempty"`
	Parameters  []ParameterObject  `json:"parameters,omitempty"`
	RequestBody *RequestBodyObject `json:"requestBody,omitempty"`

	// Tags
	// ExternalDocs
	// OperationID
	// Callbacks
	// Deprecated
	// Security
	// Servers
}

type ParameterObject struct {
	Name string `json:"name"` // Required
	In   string `json:"in"`   // Required. Possible values are "query", "header", "path" or "cookie"

	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Example     interface{}   `json:"example,omitempty"`
	Schema      *SchemaObject `json:"schema,omitempty"`

	// Ref is used when ParameterOjbect is as a ReferenceObject
	Ref string `json:"$ref,omitempty"`

	// Deprecated
	// AllowEmptyValue
	// Style
	// Explode
	// AllowReserved
	// Examples
	// Content
}

type ReferenceObject struct {
	Ref string `json:"$ref,omitempty"`
}

type RequestBodyObject struct {
	Content map[string]*MediaTypeObject `json:"content"` // Required

	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`

	// Ref is used when RequestBodyObject is as a ReferenceObject
	Ref string `json:"$ref,omitempty"`
}

type MediaTypeObject struct {
	Schema SchemaObject `json:"schema,omitempty"`
	// Example string       `json:"example,omitempty"`

	// Examples
	// Encoding
}

type Properties struct {
	*orderedmap.OrderedMap[string, *SchemaObject]
}

func NewProperties() *Properties {
	return &Properties{orderedmap.NewOrderedMap[string, *SchemaObject]()}
}

func (p *Properties) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer
	if err := buf.WriteByte('{'); err != nil {
		return nil, err
	}

	encoder := json.NewEncoder(&buf)
	first := true

	for el := p.Front(); el != nil; el = el.Next() {
		if !first {
			if err := buf.WriteByte(','); err != nil {
				return nil, err
			}
		}
		first = false

		// add key
		if err := encoder.Encode(el.Key); err != nil {
			return nil, err
		}

		if err := buf.WriteByte(':'); err != nil {
			return nil, err
		}

		// add value
		if err := encoder.Encode(el.Value); err != nil {
			return nil, err
		}
	}

	if err := buf.WriteByte('}'); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type SchemaObject struct {
	ID                 string              `json:"-"` // For goas
	PkgName            string              `json:"-"` // For goas
	FieldName          string              `json:"-"` // For goas
	DisabledFieldNames map[string]struct{} `json:"-"` // For goas

	Type        string        `json:"type,omitempty"`
	Format      string        `json:"format,omitempty"`
	Required    []string      `json:"required,omitempty"`
	Properties  *Properties   `json:"properties,omitempty"`
	Description string        `json:"description,omitempty"`
	Items       *SchemaObject `json:"items,omitempty"` // use ptr to prevent recursive error
	Example     interface{}   `json:"example,omitempty"`
	Deprecated  bool          `json:"deprecated,omitempty"`
	Nullable    bool          `json:"nullable,omitempty"`
	Minimum     *int64        `json:"minimum,omitempty"`
	Maximum     *int64        `json:"maximum,omitempty"`
	MinLength   *int64        `json:"minLength,omitempty"`
	MaxLength   *int64        `json:"maxLength,omitempty"`
	MinItems    *int64        `json:"minItems,omitempty"`
	MaxItems    *int64        `json:"maxItems,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`

	// Ref is used when SchemaObject is as a ReferenceObject
	Ref string `json:"$ref,omitempty"`

	// Title
	// MultipleOf
	// Maximum
	// ExclusiveMaximum
	// Minimum
	// ExclusiveMinimum
	// MaxLength
	// MinLength
	// Pattern
	// MaxItems
	// MinItems
	// UniqueItems
	// MaxProperties
	// MinProperties
	// Enum
	// AllOf
	// OneOf
	// AnyOf
	// Not
	// AdditionalProperties
	// Description
	// Default
	// Nullable
	// ReadOnly
	// WriteOnly
	// XML
	// ExternalDocs
}

type ResponsesObject map[string]*ResponseObject // [status]ResponseObject

type Headers struct {
	*orderedmap.OrderedMap[string, HeaderObject]
}

func NewHeaders() *Headers {
	return &Headers{orderedmap.NewOrderedMap[string, HeaderObject]()}
}

func (h *Headers) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer
	if err := buf.WriteByte('{'); err != nil {
		return nil, err
	}

	encoder := json.NewEncoder(&buf)
	first := true

	for el := h.Front(); el != nil; el = el.Next() {
		if !first {
			if err := buf.WriteByte(','); err != nil {
				return nil, err
			}
		}
		first = false

		// add key
		if err := encoder.Encode(el.Key); err != nil {
			return nil, err
		}

		if err := buf.WriteByte(':'); err != nil {
			return nil, err
		}

		// add value
		if err := encoder.Encode(el.Value); err != nil {
			return nil, err
		}
	}

	if err := buf.WriteByte('}'); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type ResponseObject struct {
	Description string `json:"description"` // Required

	Headers *Headers                    `json:"headers,omitempty"`
	Content map[string]*MediaTypeObject `json:"content,omitempty"`

	// Ref is for ReferenceObject
	Ref string `json:"$ref,omitempty"`

	// Links
}

// type HeaderObject struct {
// 	Description string `json:"description,omitempty"`
// 	Type        string `json:"type,omitempty"`

// 	// Ref is used when HeaderObject is as a ReferenceObject
// 	Ref string `json:"$ref,omitempty"`
// }

type HeaderObject struct {
	Schema      *SchemaObject `json:"schema,omitempty"`
	Description string        `json:"description,omitempty"`
}

type ComponentsOjbect struct {
	Schemas         map[string]*SchemaObject         `json:"schemas,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeObject `json:"securitySchemes,omitempty"`

	// Responses
	// Parameters
	// Examples
	// RequestBodies
	// Headers
	// Links
	// Callbacks
}

type SecuritySchemeObject struct {
	// Generic fields
	Type        string `json:"type"` // Required
	Description string `json:"description,omitempty"`

	// http
	Scheme string `json:"scheme,omitempty"`

	// apiKey
	In   string `json:"in,omitempty"`
	Name string `json:"name,omitempty"`

	// OpenID
	OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`

	// OAuth2
	OAuthFlows *SecuritySchemeOauthObject `json:"flows,omitempty"`

	// BearerFormat
}

type SecuritySchemeOauthObject struct {
	Implicit              *SecuritySchemeOauthFlowObject `json:"implicit,omitempty"`
	AuthorizationCode     *SecuritySchemeOauthFlowObject `json:"authorizationCode,omitempty"`
	ResourceOwnerPassword *SecuritySchemeOauthFlowObject `json:"password,omitempty"`
	ClientCredentials     *SecuritySchemeOauthFlowObject `json:"clientCredentials,omitempty"`
}

func (s *SecuritySchemeOauthObject) ApplyScopes(scopes map[string]string) {
	if s.Implicit != nil {
		s.Implicit.Scopes = scopes
	}

	if s.AuthorizationCode != nil {
		s.AuthorizationCode.Scopes = scopes
	}

	if s.ResourceOwnerPassword != nil {
		s.ResourceOwnerPassword.Scopes = scopes
	}

	if s.ClientCredentials != nil {
		s.ClientCredentials.Scopes = scopes
	}
}

type SecuritySchemeOauthFlowObject struct {
	AuthorizationUrl string            `json:"authorizationUrl,omitempty"`
	TokenUrl         string            `json:"tokenUrl,omitempty"`
	Scopes           map[string]string `json:"scopes"`
}
