// tslint:disable
/**
 * TCC Workflows API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: v1
 * Contact: info@tetrate.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    TetrateApiTccWorkflowsV1LBPublishAction,
    TetrateApiTccWorkflowsV1LBPublishActionFromJSON,
    TetrateApiTccWorkflowsV1LBPublishActionToJSON,
    TetrateApiTccWorkflowsV1LBTicketDetails,
    TetrateApiTccWorkflowsV1LBTicketDetailsFromJSON,
    TetrateApiTccWorkflowsV1LBTicketDetailsToJSON,
    TetrateApiTccWorkflowsV1LBTicketId,
    TetrateApiTccWorkflowsV1LBTicketIdFromJSON,
    TetrateApiTccWorkflowsV1LBTicketIdToJSON,
    TetrateApiTccWorkflowsV1LBTicketResolution,
    TetrateApiTccWorkflowsV1LBTicketResolutionFromJSON,
    TetrateApiTccWorkflowsV1LBTicketResolutionToJSON,
    TetrateApiTccWorkflowsV1LBTicketStatus,
    TetrateApiTccWorkflowsV1LBTicketStatusFromJSON,
    TetrateApiTccWorkflowsV1LBTicketStatusToJSON,
    TetrateApiTccWorkflowsV1ListTicketsResponse,
    TetrateApiTccWorkflowsV1ListTicketsResponseFromJSON,
    TetrateApiTccWorkflowsV1ListTicketsResponseToJSON,
} from '../models';

export interface ApproveRequest {
    tenant: string;
    cluster: string;
    requestid: string;
    body: TetrateApiTccWorkflowsV1LBTicketResolution;
}

export interface AttachRequest {
    tenant: string;
    cluster: string;
    body: TetrateApiTccWorkflowsV1LBTicketDetails;
}

export interface CancelRequest {
    tenant: string;
    cluster: string;
    requestid: string;
    body: TetrateApiTccWorkflowsV1LBTicketId;
}

export interface DenyRequest {
    tenant: string;
    cluster: string;
    requestid: string;
    body: TetrateApiTccWorkflowsV1LBTicketResolution;
}

export interface DetachRequest {
    tenant: string;
    cluster: string;
    body: TetrateApiTccWorkflowsV1LBTicketDetails;
}

export interface GetTicketStatusRequest {
    tenant: string;
    cluster: string;
    requestid: string;
}

export interface ListPendingTicketsRequest {
    tenant: string;
    cluster: string;
}

export interface PublishRequest {
    tenant: string;
    cluster: string;
    body: TetrateApiTccWorkflowsV1LBPublishAction;
}

/**
 * no description
 */
export class LoadBalancerWorkflowApi extends runtime.BaseAPI {

    /**
     */
    async approveRaw(requestParameters: ApproveRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling approve.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling approve.');
        }

        if (requestParameters.requestid === null || requestParameters.requestid === undefined) {
            throw new runtime.RequiredError('requestid','Required parameter requestParameters.requestid was null or undefined when calling approve.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling approve.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/request/{requestid}/approve`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))).replace(`{${"requestid"}}`, encodeURIComponent(String(requestParameters.requestid))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBTicketResolutionToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async approve(requestParameters: ApproveRequest): Promise<TetrateApiTccWorkflowsV1LBTicketStatus> {
        const response = await this.approveRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async attachRaw(requestParameters: AttachRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling attach.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling attach.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling attach.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/attach`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBTicketDetailsToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketIdFromJSON(jsonValue));
    }

    /**
     */
    async attach(requestParameters: AttachRequest): Promise<TetrateApiTccWorkflowsV1LBTicketId> {
        const response = await this.attachRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async cancelRaw(requestParameters: CancelRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling cancel.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling cancel.');
        }

        if (requestParameters.requestid === null || requestParameters.requestid === undefined) {
            throw new runtime.RequiredError('requestid','Required parameter requestParameters.requestid was null or undefined when calling cancel.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling cancel.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/request/{requestid}/cancel`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))).replace(`{${"requestid"}}`, encodeURIComponent(String(requestParameters.requestid))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBTicketIdToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async cancel(requestParameters: CancelRequest): Promise<TetrateApiTccWorkflowsV1LBTicketStatus> {
        const response = await this.cancelRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async denyRaw(requestParameters: DenyRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deny.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling deny.');
        }

        if (requestParameters.requestid === null || requestParameters.requestid === undefined) {
            throw new runtime.RequiredError('requestid','Required parameter requestParameters.requestid was null or undefined when calling deny.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling deny.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/request/{requestid}/deny`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))).replace(`{${"requestid"}}`, encodeURIComponent(String(requestParameters.requestid))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBTicketResolutionToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async deny(requestParameters: DenyRequest): Promise<TetrateApiTccWorkflowsV1LBTicketStatus> {
        const response = await this.denyRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async detachRaw(requestParameters: DetachRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling detach.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling detach.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling detach.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/detach`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBTicketDetailsToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketIdFromJSON(jsonValue));
    }

    /**
     */
    async detach(requestParameters: DetachRequest): Promise<TetrateApiTccWorkflowsV1LBTicketId> {
        const response = await this.detachRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getTicketStatusRaw(requestParameters: GetTicketStatusRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getTicketStatus.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling getTicketStatus.');
        }

        if (requestParameters.requestid === null || requestParameters.requestid === undefined) {
            throw new runtime.RequiredError('requestid','Required parameter requestParameters.requestid was null or undefined when calling getTicketStatus.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/request/{requestid}/status`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))).replace(`{${"requestid"}}`, encodeURIComponent(String(requestParameters.requestid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async getTicketStatus(requestParameters: GetTicketStatusRequest): Promise<TetrateApiTccWorkflowsV1LBTicketStatus> {
        const response = await this.getTicketStatusRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listPendingTicketsRaw(requestParameters: ListPendingTicketsRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listPendingTickets.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling listPendingTickets.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/pending`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1ListTicketsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listPendingTickets(requestParameters: ListPendingTicketsRequest): Promise<TetrateApiTccWorkflowsV1ListTicketsResponse> {
        const response = await this.listPendingTicketsRaw(requestParameters);
        return await response.value();
    }

    /**
     * LB owner calls this API with additional settings like TLS, to finally expose the service on the load balancer or remove a detached service
     */
    async publishRaw(requestParameters: PublishRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling publish.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling publish.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling publish.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/workflows/clusters/{cluster}/loadbalancer/publish`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"cluster"}}`, encodeURIComponent(String(requestParameters.cluster))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsV1LBPublishActionToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsV1LBTicketStatusFromJSON(jsonValue));
    }

    /**
     * LB owner calls this API with additional settings like TLS, to finally expose the service on the load balancer or remove a detached service
     */
    async publish(requestParameters: PublishRequest): Promise<TetrateApiTccWorkflowsV1LBTicketStatus> {
        const response = await this.publishRaw(requestParameters);
        return await response.value();
    }

}
