//
// Code generated by rdl 1.5.2 DO NOT EDIT.
//

package msd

import (
	"log"

	rdl "github.com/ardielle/ardielle-go/rdl"
)

var schema *rdl.Schema

func init() {
	sb := rdl.NewSchemaBuilder("MSD")
	sb.Version(1)
	sb.Namespace("com.yahoo.athenz.msd")
	sb.Comment("Copyright The Athenz Authors Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms. The Micro Segmentation Defense (MSD) API")

	tSimpleName := rdl.NewStringTypeBuilder("SimpleName")
	tSimpleName.Comment("Copyright The Athenz Authors Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms. Common name types used by several API definitions A simple identifier, an element of compound name.")
	tSimpleName.Pattern("[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tSimpleName.Build())

	tCompoundName := rdl.NewStringTypeBuilder("CompoundName")
	tCompoundName.Comment("A compound name. Most names in this API are compound names.")
	tCompoundName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tCompoundName.Build())

	tDomainName := rdl.NewStringTypeBuilder("DomainName")
	tDomainName.Comment("A domain name is the general qualifier prefix, as its uniqueness is managed.")
	tDomainName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tDomainName.Build())

	tEntityName := rdl.NewStringTypeBuilder("EntityName")
	tEntityName.Comment("An entity name is a short form of a resource name, including only the domain and entity.")
	tEntityName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tEntityName.Build())

	tEntityList := rdl.NewStringTypeBuilder("EntityList")
	tEntityList.Comment("An Entity list is comma separated compound Names")
	tEntityList.Pattern("(([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*,)*([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tEntityList.Build())

	tServiceName := rdl.NewStringTypeBuilder("ServiceName")
	tServiceName.Comment("A service name will generally be a unique subdomain.")
	tServiceName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tServiceName.Build())

	tActionName := rdl.NewStringTypeBuilder("ActionName")
	tActionName.Comment("An action (operation) name.")
	tActionName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tActionName.Build())

	tResourceName := rdl.NewStringTypeBuilder("ResourceName")
	tResourceName.Comment("A resource name Note that the EntityName part is optional, that is, a domain name followed by a colon is valid resource name.")
	tResourceName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*(:([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*)?")
	sb.AddType(tResourceName.Build())

	tYBase64 := rdl.NewStringTypeBuilder("YBase64")
	tYBase64.Comment("The Y-specific URL-safe Base64 variant.")
	tYBase64.Pattern("[a-zA-Z0-9\\._-]+")
	sb.AddType(tYBase64.Build())

	tYEncoded := rdl.NewStringTypeBuilder("YEncoded")
	tYEncoded.Comment("YEncoded includes ybase64 chars, as well as = and %. This can represent a user cookie and URL-encoded values.")
	tYEncoded.Pattern("[a-zA-Z0-9\\._%=-]*")
	sb.AddType(tYEncoded.Build())

	tAuthorityName := rdl.NewStringTypeBuilder("AuthorityName")
	tAuthorityName.Comment("Used as the prefix in a signed assertion. This uniquely identifies a signing authority.")
	tAuthorityName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tAuthorityName.Build())

	tPathElement := rdl.NewStringTypeBuilder("PathElement")
	tPathElement.Comment("A uri-safe path element")
	tPathElement.Pattern("[a-zA-Z0-9-\\._~=+@$,:]*")
	sb.AddType(tPathElement.Build())

	tTransportPolicyEnforcementState := rdl.NewEnumTypeBuilder("Enum", "TransportPolicyEnforcementState")
	tTransportPolicyEnforcementState.Comment("Types of transport policy enforcement states")
	tTransportPolicyEnforcementState.Element("ENFORCE", "")
	tTransportPolicyEnforcementState.Element("REPORT", "")
	sb.AddType(tTransportPolicyEnforcementState.Build())

	tTransportPolicyProtocol := rdl.NewEnumTypeBuilder("Enum", "TransportPolicyProtocol")
	tTransportPolicyProtocol.Comment("Types of transport policy protocols")
	tTransportPolicyProtocol.Element("TCP", "")
	tTransportPolicyProtocol.Element("UDP", "")
	sb.AddType(tTransportPolicyProtocol.Build())

	tTransportPolicySubject := rdl.NewStructTypeBuilder("Struct", "TransportPolicySubject")
	tTransportPolicySubject.Comment("Subject for a transport policy")
	tTransportPolicySubject.Field("domainName", "DomainName", false, nil, "Name of the domain")
	tTransportPolicySubject.Field("serviceName", "EntityName", false, nil, "Name of the service")
	sb.AddType(tTransportPolicySubject.Build())

	tTransportPolicyCondition := rdl.NewStructTypeBuilder("Struct", "TransportPolicyCondition")
	tTransportPolicyCondition.Comment("Transport policy condition. Used to specify additional restrictions for the subject of a transport policy")
	tTransportPolicyCondition.Field("enforcementState", "TransportPolicyEnforcementState", false, nil, "State of transport policy enforcement ( ENFORCE / REPORT )")
	tTransportPolicyCondition.ArrayField("instances", "String", true, "Acts as restrictions. If present, this transport policy should be restricted to only mentioned instances.")
	sb.AddType(tTransportPolicyCondition.Build())

	tTransportPolicyPort := rdl.NewStructTypeBuilder("Struct", "TransportPolicyPort")
	tTransportPolicyPort.Comment("Transport policy port")
	tTransportPolicyPort.Field("port", "Int32", false, nil, "Start port of the port range. port and endPort will have same values for a single port definition.")
	tTransportPolicyPort.Field("endPort", "Int32", false, nil, "End port of the port range. port and endPort will have same values for a single port definition.")
	tTransportPolicyPort.Field("protocol", "TransportPolicyProtocol", false, nil, "Protocol for this transport policy")
	sb.AddType(tTransportPolicyPort.Build())

	tTransportPolicyMatch := rdl.NewStructTypeBuilder("Struct", "TransportPolicyMatch")
	tTransportPolicyMatch.Comment("Selector for the subject of a transport policy")
	tTransportPolicyMatch.Field("athenzService", "TransportPolicySubject", false, nil, "Subject where this transport policy applies")
	tTransportPolicyMatch.ArrayField("conditions", "TransportPolicyCondition", false, "List of additional requirements for restrictions. Requirements are ANDed.")
	sb.AddType(tTransportPolicyMatch.Build())

	tTransportPolicyPeer := rdl.NewStructTypeBuilder("Struct", "TransportPolicyPeer")
	tTransportPolicyPeer.Comment("Source or destination for a transport policy")
	tTransportPolicyPeer.ArrayField("athenzServices", "TransportPolicySubject", false, "List of transport policy subjects")
	tTransportPolicyPeer.ArrayField("ports", "TransportPolicyPort", false, "List of network traffic port part of this transport policy")
	sb.AddType(tTransportPolicyPeer.Build())

	tTransportPolicyEntitySelector := rdl.NewStructTypeBuilder("Struct", "TransportPolicyEntitySelector")
	tTransportPolicyEntitySelector.Comment("Entity to which a transport policy applies. Describes the subject and port(s) for a transport policy.")
	tTransportPolicyEntitySelector.Field("match", "TransportPolicyMatch", false, nil, "Requirements for selecting the subject for this transport policy.")
	tTransportPolicyEntitySelector.ArrayField("ports", "TransportPolicyPort", false, "List of network traffic port of the subject eligible for the transport policy")
	sb.AddType(tTransportPolicyEntitySelector.Build())

	tTransportPolicyIngressRule := rdl.NewStructTypeBuilder("Struct", "TransportPolicyIngressRule")
	tTransportPolicyIngressRule.Comment("Transport policy ingress rule")
	tTransportPolicyIngressRule.Field("id", "Int64", false, nil, "Assertion id associated with this transport policy")
	tTransportPolicyIngressRule.Field("lastModified", "Timestamp", false, nil, "Last modification timestamp of this transport policy")
	tTransportPolicyIngressRule.Field("entitySelector", "TransportPolicyEntitySelector", false, nil, "Describes the entity to which this transport policy applies")
	tTransportPolicyIngressRule.Field("from", "TransportPolicyPeer", false, nil, "Source of network traffic")
	sb.AddType(tTransportPolicyIngressRule.Build())

	tTransportPolicyEgressRule := rdl.NewStructTypeBuilder("Struct", "TransportPolicyEgressRule")
	tTransportPolicyEgressRule.Comment("Transport policy egress rule")
	tTransportPolicyEgressRule.Field("id", "Int64", false, nil, "Assertion id associated with this transport policy")
	tTransportPolicyEgressRule.Field("lastModified", "Timestamp", false, nil, "Last modification timestamp of this transport policy")
	tTransportPolicyEgressRule.Field("entitySelector", "TransportPolicyEntitySelector", false, nil, "Entity to which this transport policy applies")
	tTransportPolicyEgressRule.Field("to", "TransportPolicyPeer", false, nil, "Destination of network traffic")
	sb.AddType(tTransportPolicyEgressRule.Build())

	tTransportPolicyRules := rdl.NewStructTypeBuilder("Struct", "TransportPolicyRules")
	tTransportPolicyRules.Comment("Transport policy containing ingress and egress rules")
	tTransportPolicyRules.ArrayField("ingress", "TransportPolicyIngressRule", false, "List of ingress rules")
	tTransportPolicyRules.ArrayField("egress", "TransportPolicyEgressRule", false, "List of egress rules")
	sb.AddType(tTransportPolicyRules.Build())

	tWorkload := rdl.NewStructTypeBuilder("Struct", "Workload")
	tWorkload.Comment("workload type describing workload associated with an identity")
	tWorkload.Field("domainName", "DomainName", false, nil, "name of the domain, optional for getWorkloadsByService API call")
	tWorkload.Field("serviceName", "EntityName", false, nil, "name of the service, , optional for getWorkloadsByService API call")
	tWorkload.Field("uuid", "String", false, nil, "unique identifier for the workload, usually defined by provider")
	tWorkload.ArrayField("ipAddresses", "String", false, "list of IP addresses associated with the workload, optional for getWorkloadsByIP API call")
	tWorkload.Field("hostname", "String", false, nil, "hostname associated with the workload")
	tWorkload.Field("provider", "String", false, nil, "infrastructure provider e.g. k8s, AWS, Azure, openstack etc.")
	tWorkload.Field("updateTime", "Timestamp", false, nil, "most recent update timestamp in the backend")
	tWorkload.Field("certExpiryTime", "Timestamp", false, nil, "certificate expiry time (ex: getNotAfter)")
	sb.AddType(tWorkload.Build())

	tWorkloads := rdl.NewStructTypeBuilder("Struct", "Workloads")
	tWorkloads.Comment("list of workloads")
	tWorkloads.ArrayField("workloadList", "Workload", false, "list of workloads")
	sb.AddType(tWorkloads.Build())

	mGetTransportPolicyRules := rdl.NewResourceBuilder("TransportPolicyRules", "GET", "/transportpolicies")
	mGetTransportPolicyRules.Comment("API endpoint to get the transport policy rules defined in Athenz")
	mGetTransportPolicyRules.Input("matchingTag", "String", false, "", "If-None-Match", false, nil, "Retrieved from the previous request, this timestamp specifies to the server to return any policies modified since this time")
	mGetTransportPolicyRules.Output("tag", "String", "ETag", false, "The current latest modification timestamp is returned in this header")
	mGetTransportPolicyRules.Auth("", "", true, "")
	mGetTransportPolicyRules.Exception("BAD_REQUEST", "ResourceError", "")
	mGetTransportPolicyRules.Exception("FORBIDDEN", "ResourceError", "")
	mGetTransportPolicyRules.Exception("NOT_FOUND", "ResourceError", "")
	mGetTransportPolicyRules.Exception("TOO_MANY_REQUESTS", "ResourceError", "")
	mGetTransportPolicyRules.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(mGetTransportPolicyRules.Build())

	mGetWorkloadsByService := rdl.NewResourceBuilder("Workloads", "GET", "/domain/{domainName}/service/{serviceName}/workloads")
	mGetWorkloadsByService.Name("getWorkloadsByService")
	mGetWorkloadsByService.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	mGetWorkloadsByService.Input("serviceName", "EntityName", true, "", "", false, nil, "name of the service")
	mGetWorkloadsByService.Input("matchingTag", "String", false, "", "If-None-Match", false, nil, "Retrieved from the previous request, this timestamp specifies to the server to return any workloads modified since this time")
	mGetWorkloadsByService.Output("tag", "String", "ETag", false, "The current latest modification timestamp is returned in this header")
	mGetWorkloadsByService.Auth("", "", true, "")
	mGetWorkloadsByService.Exception("BAD_REQUEST", "ResourceError", "")
	mGetWorkloadsByService.Exception("FORBIDDEN", "ResourceError", "")
	mGetWorkloadsByService.Exception("NOT_FOUND", "ResourceError", "")
	mGetWorkloadsByService.Exception("TOO_MANY_REQUESTS", "ResourceError", "")
	mGetWorkloadsByService.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(mGetWorkloadsByService.Build())

	mGetWorkloadsByIP := rdl.NewResourceBuilder("Workloads", "GET", "/workloads/{ip}")
	mGetWorkloadsByIP.Name("getWorkloadsByIP")
	mGetWorkloadsByIP.Input("ip", "String", true, "", "", false, nil, "ip address to query")
	mGetWorkloadsByIP.Input("matchingTag", "String", false, "", "If-None-Match", false, nil, "Retrieved from the previous request, this timestamp specifies to the server to return any workloads modified since this time")
	mGetWorkloadsByIP.Output("tag", "String", "ETag", false, "The current latest modification timestamp is returned in this header")
	mGetWorkloadsByIP.Auth("", "", true, "")
	mGetWorkloadsByIP.Exception("BAD_REQUEST", "ResourceError", "")
	mGetWorkloadsByIP.Exception("FORBIDDEN", "ResourceError", "")
	mGetWorkloadsByIP.Exception("NOT_FOUND", "ResourceError", "")
	mGetWorkloadsByIP.Exception("TOO_MANY_REQUESTS", "ResourceError", "")
	mGetWorkloadsByIP.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(mGetWorkloadsByIP.Build())

	var err error
	schema, err = sb.BuildParanoid()
	if err != nil {
		log.Fatalf("rdl: schema build failed: %s", err)
	}
}

func MSDSchema() *rdl.Schema {
	return schema
}
