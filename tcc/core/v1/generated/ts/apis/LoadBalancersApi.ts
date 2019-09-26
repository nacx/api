// tslint:disable
/**
 * configproducer.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    TetrateApiTccV1WorkflowsCreateLoadBalancerRequest,
    TetrateApiTccV1WorkflowsCreateLoadBalancerRequestFromJSON,
    TetrateApiTccV1WorkflowsCreateLoadBalancerRequestToJSON,
    TetrateApiTccV1WorkflowsListLoadBalancerResponse,
    TetrateApiTccV1WorkflowsListLoadBalancerResponseFromJSON,
    TetrateApiTccV1WorkflowsListLoadBalancerResponseToJSON,
    TetrateApiTccV1WorkflowsLoadBalancer,
    TetrateApiTccV1WorkflowsLoadBalancerFromJSON,
    TetrateApiTccV1WorkflowsLoadBalancerToJSON,
    TetrateApiTccV1WorkflowsUpdateLoadBalancerRequest,
    TetrateApiTccV1WorkflowsUpdateLoadBalancerRequestFromJSON,
    TetrateApiTccV1WorkflowsUpdateLoadBalancerRequestToJSON,
} from '../models';

export interface CreateLoadBalancerRequest {
    tenant: string;
    workspace: string;
    name: string;
    body: TetrateApiTccV1WorkflowsCreateLoadBalancerRequest;
}

export interface DeleteLoadBalancerRequest {
    tenant: string;
    workspace: string;
    name: string;
}

export interface GetLoadBalancerRequest {
    tenant: string;
    workspace: string;
    name: string;
}

export interface ListTenantLoadBalancerRequest {
    tenant: string;
}

export interface ListWorkspaceLoadBalancerRequest {
    tenant: string;
    workspace: string;
}

export interface UpdateLoadBalancerRequest {
    tenant: string;
    workspace: string;
    name: string;
    body: TetrateApiTccV1WorkflowsUpdateLoadBalancerRequest;
}

/**
 * no description
 */
export class LoadBalancersApi extends runtime.BaseAPI {

    /**
     */
    async createLoadBalancerRaw(requestParameters: CreateLoadBalancerRequest): Promise<runtime.ApiResponse<TetrateApiTccV1WorkflowsLoadBalancer>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling createLoadBalancer.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling createLoadBalancer.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling createLoadBalancer.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccV1WorkflowsCreateLoadBalancerRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1WorkflowsLoadBalancerFromJSON(jsonValue));
    }

    /**
     */
    async createLoadBalancer(requestParameters: CreateLoadBalancerRequest): Promise<TetrateApiTccV1WorkflowsLoadBalancer> {
        const response = await this.createLoadBalancerRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteLoadBalancerRaw(requestParameters: DeleteLoadBalancerRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deleteLoadBalancer.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling deleteLoadBalancer.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling deleteLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async deleteLoadBalancer(requestParameters: DeleteLoadBalancerRequest): Promise<any> {
        const response = await this.deleteLoadBalancerRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getLoadBalancerRaw(requestParameters: GetLoadBalancerRequest): Promise<runtime.ApiResponse<TetrateApiTccV1WorkflowsLoadBalancer>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getLoadBalancer.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling getLoadBalancer.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling getLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1WorkflowsLoadBalancerFromJSON(jsonValue));
    }

    /**
     */
    async getLoadBalancer(requestParameters: GetLoadBalancerRequest): Promise<TetrateApiTccV1WorkflowsLoadBalancer> {
        const response = await this.getLoadBalancerRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listTenantLoadBalancerRaw(requestParameters: ListTenantLoadBalancerRequest): Promise<runtime.ApiResponse<TetrateApiTccV1WorkflowsListLoadBalancerResponse>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listTenantLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1WorkflowsListLoadBalancerResponseFromJSON(jsonValue));
    }

    /**
     */
    async listTenantLoadBalancer(requestParameters: ListTenantLoadBalancerRequest): Promise<TetrateApiTccV1WorkflowsListLoadBalancerResponse> {
        const response = await this.listTenantLoadBalancerRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listWorkspaceLoadBalancerRaw(requestParameters: ListWorkspaceLoadBalancerRequest): Promise<runtime.ApiResponse<TetrateApiTccV1WorkflowsListLoadBalancerResponse>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listWorkspaceLoadBalancer.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling listWorkspaceLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}/workspace/{workspace}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1WorkflowsListLoadBalancerResponseFromJSON(jsonValue));
    }

    /**
     */
    async listWorkspaceLoadBalancer(requestParameters: ListWorkspaceLoadBalancerRequest): Promise<TetrateApiTccV1WorkflowsListLoadBalancerResponse> {
        const response = await this.listWorkspaceLoadBalancerRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateLoadBalancerRaw(requestParameters: UpdateLoadBalancerRequest): Promise<runtime.ApiResponse<TetrateApiTccV1WorkflowsLoadBalancer>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling updateLoadBalancer.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling updateLoadBalancer.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling updateLoadBalancer.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateLoadBalancer.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/loadbalancers/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccV1WorkflowsUpdateLoadBalancerRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1WorkflowsLoadBalancerFromJSON(jsonValue));
    }

    /**
     */
    async updateLoadBalancer(requestParameters: UpdateLoadBalancerRequest): Promise<TetrateApiTccV1WorkflowsLoadBalancer> {
        const response = await this.updateLoadBalancerRaw(requestParameters);
        return await response.value();
    }

}
