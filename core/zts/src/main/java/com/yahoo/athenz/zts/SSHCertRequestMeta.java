//
// This file generated by rdl 1.5.2. Do not modify!
//

package com.yahoo.athenz.zts;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.yahoo.rdl.*;

//
// SSHCertRequestMeta -
//
@JsonIgnoreProperties(ignoreUnknown = true)
public class SSHCertRequestMeta {
    public String requestor;
    public String origin;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public String clientInfo;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public String sshClientVersion;
    public String certType;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public String athenzService;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public String instanceId;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public Timestamp prevCertValidFrom;
    @RdlOptional
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    public Timestamp prevCertValidTo;

    public SSHCertRequestMeta setRequestor(String requestor) {
        this.requestor = requestor;
        return this;
    }
    public String getRequestor() {
        return requestor;
    }
    public SSHCertRequestMeta setOrigin(String origin) {
        this.origin = origin;
        return this;
    }
    public String getOrigin() {
        return origin;
    }
    public SSHCertRequestMeta setClientInfo(String clientInfo) {
        this.clientInfo = clientInfo;
        return this;
    }
    public String getClientInfo() {
        return clientInfo;
    }
    public SSHCertRequestMeta setSshClientVersion(String sshClientVersion) {
        this.sshClientVersion = sshClientVersion;
        return this;
    }
    public String getSshClientVersion() {
        return sshClientVersion;
    }
    public SSHCertRequestMeta setCertType(String certType) {
        this.certType = certType;
        return this;
    }
    public String getCertType() {
        return certType;
    }
    public SSHCertRequestMeta setAthenzService(String athenzService) {
        this.athenzService = athenzService;
        return this;
    }
    public String getAthenzService() {
        return athenzService;
    }
    public SSHCertRequestMeta setInstanceId(String instanceId) {
        this.instanceId = instanceId;
        return this;
    }
    public String getInstanceId() {
        return instanceId;
    }
    public SSHCertRequestMeta setPrevCertValidFrom(Timestamp prevCertValidFrom) {
        this.prevCertValidFrom = prevCertValidFrom;
        return this;
    }
    public Timestamp getPrevCertValidFrom() {
        return prevCertValidFrom;
    }
    public SSHCertRequestMeta setPrevCertValidTo(Timestamp prevCertValidTo) {
        this.prevCertValidTo = prevCertValidTo;
        return this;
    }
    public Timestamp getPrevCertValidTo() {
        return prevCertValidTo;
    }

    @Override
    public boolean equals(Object another) {
        if (this != another) {
            if (another == null || another.getClass() != SSHCertRequestMeta.class) {
                return false;
            }
            SSHCertRequestMeta a = (SSHCertRequestMeta) another;
            if (requestor == null ? a.requestor != null : !requestor.equals(a.requestor)) {
                return false;
            }
            if (origin == null ? a.origin != null : !origin.equals(a.origin)) {
                return false;
            }
            if (clientInfo == null ? a.clientInfo != null : !clientInfo.equals(a.clientInfo)) {
                return false;
            }
            if (sshClientVersion == null ? a.sshClientVersion != null : !sshClientVersion.equals(a.sshClientVersion)) {
                return false;
            }
            if (certType == null ? a.certType != null : !certType.equals(a.certType)) {
                return false;
            }
            if (athenzService == null ? a.athenzService != null : !athenzService.equals(a.athenzService)) {
                return false;
            }
            if (instanceId == null ? a.instanceId != null : !instanceId.equals(a.instanceId)) {
                return false;
            }
            if (prevCertValidFrom == null ? a.prevCertValidFrom != null : !prevCertValidFrom.equals(a.prevCertValidFrom)) {
                return false;
            }
            if (prevCertValidTo == null ? a.prevCertValidTo != null : !prevCertValidTo.equals(a.prevCertValidTo)) {
                return false;
            }
        }
        return true;
    }
}
