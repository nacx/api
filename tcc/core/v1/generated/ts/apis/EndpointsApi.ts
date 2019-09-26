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
    TetrateApiTccV1CreateEndpointRequest,
    TetrateApiTccV1CreateEndpointRequestFromJSON,
    TetrateApiTccV1CreateEndpointRequestToJSON,
    TetrateApiTccV1Endpoint,
    TetrateApiTccV1EndpointFromJSON,
    TetrateApiTccV1EndpointToJSON,
    TetrateApiTccV1ListEndpointResponse,
    TetrateApiTccV1ListEndpointResponseFromJSON,
    TetrateApiTccV1ListEndpointResponseToJSON,
    TetrateApiTccV1UpdateEndpointRequest,
    TetrateApiTccV1UpdateEndpointRequestFromJSON,
    TetrateApiTccV1UpdateEndpointRequestToJSON,
} from '../models';

export interface CreateEndpointRequest {
    tenant: string;
    workspace: string;
    name: string;
    body: TetrateApiTccV1CreateEndpointRequest;
}

export interface DeleteEndpointRequest {
    tenant: string;
    workspace: string;
    name: string;
}

export interface GetEndpointRequest {
    tenant: string;
    workspace: string;
    name: string;
}

export interface ListServiceEndpointRequest {
    tenant: string;
    workspace: string;
    service: string;
}

export interface ListServiceSubsetEndpointRequest {
    tenant: string;
    workspace: string;
    service: string;
    subset: string;
}

export interface ListWorkspaceEndpointRequest {
    tenant: string;
    workspace: string;
}

export interface UpdateEndpointRequest {
    tenant: string;
    workspace: string;
    name: string;
    body: TetrateApiTccV1UpdateEndpointRequest;
}

/**
 * no description
 */
export class EndpointsApi extends runtime.BaseAPI {

    /**
     */
    async createEndpointRaw(requestParameters: CreateEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1Endpoint>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling createEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling createEndpoint.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling createEndpoint.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccV1CreateEndpointRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1EndpointFromJSON(jsonValue));
    }

    /**
     */
    async createEndpoint(requestParameters: CreateEndpointRequest): Promise<TetrateApiTccV1Endpoint> {
        const response = await this.createEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteEndpointRaw(requestParameters: DeleteEndpointRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deleteEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling deleteEndpoint.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling deleteEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async deleteEndpoint(requestParameters: DeleteEndpointRequest): Promise<any> {
        const response = await this.deleteEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getEndpointRaw(requestParameters: GetEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1Endpoint>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling getEndpoint.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling getEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1EndpointFromJSON(jsonValue));
    }

    /**
     */
    async getEndpoint(requestParameters: GetEndpointRequest): Promise<TetrateApiTccV1Endpoint> {
        const response = await this.getEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listServiceEndpointRaw(requestParameters: ListServiceEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1ListEndpointResponse>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listServiceEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling listServiceEndpoint.');
        }

        if (requestParameters.service === null || requestParameters.service === undefined) {
            throw new runtime.RequiredError('service','Required parameter requestParameters.service was null or undefined when calling listServiceEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/service/{service}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"service"}}`, encodeURIComponent(String(requestParameters.service))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1ListEndpointResponseFromJSON(jsonValue));
    }

    /**
     */
    async listServiceEndpoint(requestParameters: ListServiceEndpointRequest): Promise<TetrateApiTccV1ListEndpointResponse> {
        const response = await this.listServiceEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listServiceSubsetEndpointRaw(requestParameters: ListServiceSubsetEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1ListEndpointResponse>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listServiceSubsetEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling listServiceSubsetEndpoint.');
        }

        if (requestParameters.service === null || requestParameters.service === undefined) {
            throw new runtime.RequiredError('service','Required parameter requestParameters.service was null or undefined when calling listServiceSubsetEndpoint.');
        }

        if (requestParameters.subset === null || requestParameters.subset === undefined) {
            throw new runtime.RequiredError('subset','Required parameter requestParameters.subset was null or undefined when calling listServiceSubsetEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/service/{service}/subset/{subset}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"service"}}`, encodeURIComponent(String(requestParameters.service))).replace(`{${"subset"}}`, encodeURIComponent(String(requestParameters.subset))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1ListEndpointResponseFromJSON(jsonValue));
    }

    /**
     */
    async listServiceSubsetEndpoint(requestParameters: ListServiceSubsetEndpointRequest): Promise<TetrateApiTccV1ListEndpointResponse> {
        const response = await this.listServiceSubsetEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listWorkspaceEndpointRaw(requestParameters: ListWorkspaceEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1ListEndpointResponse>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listWorkspaceEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling listWorkspaceEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1ListEndpointResponseFromJSON(jsonValue));
    }

    /**
     */
    async listWorkspaceEndpoint(requestParameters: ListWorkspaceEndpointRequest): Promise<TetrateApiTccV1ListEndpointResponse> {
        const response = await this.listWorkspaceEndpointRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateEndpointRaw(requestParameters: UpdateEndpointRequest): Promise<runtime.ApiResponse<TetrateApiTccV1Endpoint>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling updateEndpoint.');
        }

        if (requestParameters.workspace === null || requestParameters.workspace === undefined) {
            throw new runtime.RequiredError('workspace','Required parameter requestParameters.workspace was null or undefined when calling updateEndpoint.');
        }

        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling updateEndpoint.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateEndpoint.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/endpoints/tenant/{tenant}/workspace/{workspace}/{name}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"workspace"}}`, encodeURIComponent(String(requestParameters.workspace))).replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccV1UpdateEndpointRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccV1EndpointFromJSON(jsonValue));
    }

    /**
     */
    async updateEndpoint(requestParameters: UpdateEndpointRequest): Promise<TetrateApiTccV1Endpoint> {
        const response = await this.updateEndpointRaw(requestParameters);
        return await response.value();
    }

}
