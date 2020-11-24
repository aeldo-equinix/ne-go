package api

//ACLTemplate describes Network Edge device ACL template
type ACLTemplate struct {
	UUID              string                   `json:"uuid,omitempty"`
	Name              string                   `json:"name,omitempty"`
	Description       string                   `json:"description,omitempty"`
	MetroCode         string                   `json:"metroCode,omitempty"`
	VirtualDeviceUUID string                   `json:"virtualDeviceUUID,omitempty"`
	DeviceACLStatus   string                   `json:"deviceAclstatus,omitempty"`
	InboundRules      []ACLTemplateInboundRule `json:"inboundRules,omitempty"`
}

//ACLTemplateInboundRule describes inbound ACL rule that is part of
//Network Edge device ACL template
type ACLTemplateInboundRule struct {
	SrcType  string   `json:"srcType,omitempty"`
	Protocol string   `json:"protocol,omitempty"`
	SrcPort  string   `json:"srcPort,omitempty"`
	DstPort  string   `json:"dstPort,omitempty"`
	FQDN     string   `json:"fqdn,omitempty"`
	Subnets  []string `json:"subnets,omitempty"`
	SeqNO    int      `json:"seqNo,omitempty"`
}

//ACLTemplateCreateResponse describes response body for
//ACL template create request
type ACLTemplateCreateResponse struct {
	UUID string `json:"uuid,omitempty"`
}

//ACLTemplatesResponse describes response for a get ACL template collection request
type ACLTemplatesResponse struct {
	TotalCount int           `json:"totalCount,omitempty"`
	PageSize   int           `json:"pageSize,omitempty"`
	Content    []ACLTemplate `json:"content,omitempty"`
}