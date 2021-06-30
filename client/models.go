package client

type Meta struct {
	Created      string `json:"created"`
	LastModified string `json:"lastModified"`
	Location     string `json:"location"`
	ResourceType string `json:"resourceType"`
	Version      string `json:"version"`
}

type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type IDCSCreatedBy struct {
	Ref     string `json:"$ref"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type IDCSLastModifiedBy struct {
	Ref     string `json:"$ref"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type Group struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Groups struct {
	ItemsPerPage int     `json:"itemsPerPage"`
	Resources    []Group `json:"Resources"`
	StartIndex   int     `json:"startIndex"`
	TotalResults int     `json:"totalResults"`
}

type App struct {
	//accessTokenExpiry int
	//accounts
	Active bool `json:"active"`
	//adminRoles
	//aliasApps
	//allowAccessControl bool
	//allowedGrants string
	//allowedOperations string
	//allowedScopes
	//allowedTags
	//allowOffline bool
	//allUrlSchemesAllowed bool
	//appIcon string
	//appSignonPolicy
	//appThumbnail string
	//asOPCService
	//attrRenderingMetadata []struct{}
	//audience string
	//BasedOnTemplate   string `json:"basedOnTemplate"`
	//bypassConsent string
	//callbackServiceUrl string
	//certificates []struct{}
	//clientSecret string
	ClientType string `json:"clientType"`
	//cloudControlProperties []struct[]
	//contactEmailAddress string
	//deleteInProgress bool
	//description string
	DisplayName string `json:"displayName"`
	//editableAttributes []struct{}
	//errorPageUrl string
	//grantedAppRoles []struct{}
	//grants []struct{}
	//homePageUrl string
	//icon string
	ID                 string             `json:"id"`
	IDCSCreatedBy      IDCSCreatedBy      `json:"idcsCreatedBy"`
	IDCSLastModifiedBy IDCSLastModifiedBy `json:"idcsLastModifiedBy"`
	//idcsLastUpgradedInRelease string
	//idcsPreventedOperations string
	//identityProviders []struct{}
	//idpPolicy []struct{}
	//infrastructure bool
	IsAliasApp bool `json:"isAliasApp"`
	//isDatabaseService bool
	//isEnterpriseApp bool
	// isFormFill bool
	//isKerberosRealm bool
	//isLoginTarget bool
	IsManagedApp bool `json:"isManagedApp"`
	//isMobileTarget bool
	//isOAuthClient bool
	//isOAuthResource bool
	//isObligationCapable bool
	IsOPCService bool `json:"isOPCService"`
	//isRadiusApp bool
	//isSamlServiceProvider bool
	//isUnmanagedApp bool
	//isWebTierPolicy bool
	//landingPageUrl string
	//linkingCallbackUrl string
	//loginMechanism(optional): string
	//loginPageUrl(optional): string
	//logoutPageUrl(optional): string
	//logoutUri(optional): string
	//meta(optional): object
	MeterAsOPCService bool `json:"meterAsOPCService"`
	//migrated(optional): boolean
	//name(optional): string
	//postLogoutRedirectUris(optional): string
	//privacyPolicyUrl(optional): string
	//productLogoUrl(optional): string
	//productName(optional): string
	//protectableSecondaryAudiences(optional): array
	//readyToUpgrade(optional): boolean
	//redirectUris(optional): string
	//refreshTokenExpiry(optional): integer
	//samlServiceProvider(optional): object
	Schemas []string `json:"schemas"`
	//scopes(optional): array
	//secondaryAudiences(optional): string
	//serviceParams(optional): array
	ServiceTypeURN string `json:"serviceTypeURN"`
	//serviceTypeVersion(optional): string
	//showInMyApps(optional): boolean
	//signonPolicy(optional): object
	//tags(optional): array
	//termsOfServiceUrl(optional): string
	//termsOfUse(optional): object
	//trustPolicies(optional): array
	//trustScope(optional): string
	//userRoles(optional): array
}

type Apps struct {
	ItemsPerPage int   `json:"itemsPerPage"`
	Resources    []App `json:"Resources"`
	StartIndex   int   `json:"startIndex"`
	TotalResults int   `json:"totalResults"`
}

type AppRole struct {
	AdminRole bool `json:"adminRole"`
	App       struct {
		Ref                       string `json:"$ref"`
		Display                   string `json:"display"`
		Name                      string `json:"name"`
		ServiceInstanceIdentifier string `json:"serviceInstanceIdentifier"`
		Value                     string `json:"value"`
	} `json:"app"`
	AvailableToClients        bool               `json:"availableToClients"`
	AvailableToGroups         bool               `json:"availableToGroups"`
	AvailableToUsers          bool               `json:"availableToUsers"`
	DeleteInProgress          bool               `json:"deleteInProgress"`
	Description               string             `json:"description"`
	DisplayName               string             `json:"displayName"`
	ID                        string             `json:"id"`
	IDCSCreatedBy             IDCSCreatedBy      `json:"idcsCreatedBy"`
	IDCSLastModifiedBy        IDCSLastModifiedBy `json:"idcsLastModifiedBy"`
	IDCSLastUpgradedInRelease string             `json:"idcsLastUpgradedInRelease"`
	IDCSPreventedOperations   string             `json:"idcsPreventedOperations"`
	LegacyGroupName           string             `json:"legacyGroupName"`
	LimitedToOneOrMoreGroups  bool               `json:"limitedToOneOrMoreGroups"`
	Members                   []struct {
		Ref     string `json:"$ref"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		Value   string `json:"value"`
	} `json:"members"`
	Meta       Meta     `json:"meta"`
	Public     bool     `json:"public"`
	Schemas    []string `json:"schemas"`
	Tags       Tags     `json:"tags"`
	UniqueName string   `json:"uniqueName"`
}

type AppRoles struct {
	ItemsPerPage int       `json:"itemsPerPage"`
	Resources    []AppRole `json:"Resources"`
	StartIndex   int       `json:"startIndex"`
	TotalResults int       `json:"totalResults"`
}

type GrantApp struct {
	Ref     string `json:"$ref"`
	Display string `json:"display"`
	Value   string `json:"value"`
}

type Grant struct {
	App GrantApp `json:"app"`
	// App struct {
	// 	Ref     string `json:"$ref"`
	// 	Display string `json:"display"`
	// 	Value   string `json:"value"`
	// } `json:"app"`
	AppEntitlementCollection struct {
		Ref   string `json:"$ref"`
		Value string `json:"value"`
	} `json:"appEntitlementCollection"`
	CompositeKey     string `json:"compositeKey"`
	DeleteInProgress bool   `json:"deleteInProgress"`
	Entitlement      struct {
		AttributeName  string `json:"attributeName"`
		AttributeValue string `json:"attributeValue"`
	} `json:"entitlement"`
	GrantedAttributeValuesJson string `json:"grantedAttributeValuesJson"`
	Grantee                    struct {
		Ref     string `json:"$ref"`
		Display string `json:"display"`
		Type    string `json:"type"`
		Value   string `json:"value"`
	} `json:"grantee"`
	GrantMechanism string `json:"grantMechanism"`
	Grantor        struct {
		Ref     string `json:"$ref"`
		Display string `json:"display"`
		Type    string `json:"type"`
		Value   string `json:"value"`
	} `json:"grantor"`
	ID                        string             `json:"id"`
	IDCSCreatedBy             IDCSCreatedBy      `json:"idcsCreatedBy"`
	IDCSLastModifiedBy        IDCSLastModifiedBy `json:"idcsLastModifiedBy"`
	IDCSLastUpgradedInRelease string             `json:"idcsLastUpgradedInRelease"`
	IDCSPreventedOperations   string             `json:"idcsPreventedOperations"`
	IsFulfilled               bool               `json:"isFulfilled"`
	Meta                      Meta               `json:"meta"`
	Schemas                   []string           `json:"schemas"`
	Tags                      Tags               `json:"tags"`
}

type Grants struct {
	ItemsPerPage int      `json:"itemsPerPage"`
	Resources    []Grant  `json:"Resources"`
	Schemas      []string `json:"schemas"`
	StartIndex   int      `json:"startIndex"`
	TotalResults int      `json:"totalResults"`
}
