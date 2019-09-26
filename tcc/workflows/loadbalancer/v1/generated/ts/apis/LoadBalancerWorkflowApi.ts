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
    TetrateApiTccWorkflowsLoadbalancerV1CancelTicketRequest,
    TetrateApiTccWorkflowsLoadbalancerV1CancelTicketRequestFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1CancelTicketRequestToJSON,
    TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequest,
    TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequestFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequestToJSON,
    TetrateApiTccWorkflowsLoadbalancerV1ListTicketsResponse,
    TetrateApiTccWorkflowsLoadbalancerV1ListTicketsResponseFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1ListTicketsResponseToJSON,
    TetrateApiTccWorkflowsLoadbalancerV1PublishTicketRequest,
    TetrateApiTccWorkflowsLoadbalancerV1PublishTicketRequestFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1PublishTicketRequestToJSON,
    TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequest,
    TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequestFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequestToJSON,
    TetrateApiTccWorkflowsLoadbalancerV1TicketStatus,
    TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON,
    TetrateApiTccWorkflowsLoadbalancerV1TicketStatusToJSON,
} from '../models';

export interface ApproveRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequest;
}

export interface AttachRequest {
    tenant: string;
    environment: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequest;
}

export interface CancelRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1CancelTicketRequest;
}

export interface DenyRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequest;
}

export interface DetachRequest {
    tenant: string;
    environment: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequest;
}

export interface GetTicketStatusRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    id: string;
    name?: string;
}

export interface ListPendingTicketsRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    parent?: string;
}

export interface PublishRequest {
    tenant: string;
    environment: string;
    loadbalancer: string;
    id: string;
    body: TetrateApiTccWorkflowsLoadbalancerV1PublishTicketRequest;
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

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling approve.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling approve.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling approve.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling approve.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/requests/{id}/approve`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async approve(requestParameters: ApproveRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1TicketStatus> {
        const response = await this.approveRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async attachRaw(requestParameters: AttachRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling attach.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling attach.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling attach.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling attach.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{id}/attach`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async attach(requestParameters: AttachRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1TicketStatus> {
        const response = await this.attachRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async cancelRaw(requestParameters: CancelRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling cancel.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling cancel.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling cancel.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling cancel.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling cancel.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/requests/{id}/cancel`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1CancelTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async cancel(requestParameters: CancelRequest): Promise<object> {
        const response = await this.cancelRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async denyRaw(requestParameters: DenyRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deny.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling deny.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling deny.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deny.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling deny.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/requests/{id}/deny`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1ResolveTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async deny(requestParameters: DenyRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1TicketStatus> {
        const response = await this.denyRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async detachRaw(requestParameters: DetachRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling detach.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling detach.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling detach.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling detach.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{id}/detach`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1CreateTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async detach(requestParameters: DetachRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1TicketStatus> {
        const response = await this.detachRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getTicketStatusRaw(requestParameters: GetTicketStatusRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getTicketStatus.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling getTicketStatus.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling getTicketStatus.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTicketStatus.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/requests/{id}/status`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1TicketStatusFromJSON(jsonValue));
    }

    /**
     */
    async getTicketStatus(requestParameters: GetTicketStatusRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1TicketStatus> {
        const response = await this.getTicketStatusRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listPendingTicketsRaw(requestParameters: ListPendingTicketsRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listPendingTickets.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling listPendingTickets.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling listPendingTickets.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.parent !== undefined) {
            queryParameters['parent'] = requestParameters.parent;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/pending`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccWorkflowsLoadbalancerV1ListTicketsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listPendingTickets(requestParameters: ListPendingTicketsRequest): Promise<TetrateApiTccWorkflowsLoadbalancerV1ListTicketsResponse> {
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

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling publish.');
        }

        if (requestParameters.loadbalancer === null || requestParameters.loadbalancer === undefined) {
            throw new runtime.RequiredError('loadbalancer','Required parameter requestParameters.loadbalancer was null or undefined when calling publish.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling publish.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling publish.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/environments/{environment}/workflows/loadbalancers/{loadbalancer}/requests/{id}/publish`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"environment"}}`, encodeURIComponent(String(requestParameters.environment))).replace(`{${"loadbalancer"}}`, encodeURIComponent(String(requestParameters.loadbalancer))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccWorkflowsLoadbalancerV1PublishTicketRequestToJSON(requestParameters.body),
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     * LB owner calls this API with additional settings like TLS, to finally expose the service on the load balancer or remove a detached service
     */
    async publish(requestParameters: PublishRequest): Promise<object> {
        const response = await this.publishRaw(requestParameters);
        return await response.value();
    }

}
