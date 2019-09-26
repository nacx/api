// tslint:disable
/**
 * TCC Configuration API
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
    TetrateApiTccCoreV1CreateTeamRequest,
    TetrateApiTccCoreV1CreateTeamRequestFromJSON,
    TetrateApiTccCoreV1CreateTeamRequestToJSON,
    TetrateApiTccCoreV1CreateTenantRequest,
    TetrateApiTccCoreV1CreateTenantRequestFromJSON,
    TetrateApiTccCoreV1CreateTenantRequestToJSON,
    TetrateApiTccCoreV1CreateUserRequest,
    TetrateApiTccCoreV1CreateUserRequestFromJSON,
    TetrateApiTccCoreV1CreateUserRequestToJSON,
    TetrateApiTccCoreV1ListTeamsResponse,
    TetrateApiTccCoreV1ListTeamsResponseFromJSON,
    TetrateApiTccCoreV1ListTeamsResponseToJSON,
    TetrateApiTccCoreV1ListTenantsResponse,
    TetrateApiTccCoreV1ListTenantsResponseFromJSON,
    TetrateApiTccCoreV1ListTenantsResponseToJSON,
    TetrateApiTccCoreV1ListUsersResponse,
    TetrateApiTccCoreV1ListUsersResponseFromJSON,
    TetrateApiTccCoreV1ListUsersResponseToJSON,
    TetrateApiTccCoreV1Team,
    TetrateApiTccCoreV1TeamFromJSON,
    TetrateApiTccCoreV1TeamToJSON,
    TetrateApiTccCoreV1Tenant,
    TetrateApiTccCoreV1TenantFromJSON,
    TetrateApiTccCoreV1TenantToJSON,
    TetrateApiTccCoreV1User,
    TetrateApiTccCoreV1UserFromJSON,
    TetrateApiTccCoreV1UserToJSON,
} from '../models';

export interface CreateTeamRequest {
    tenant: string;
    body: TetrateApiTccCoreV1CreateTeamRequest;
}

export interface CreateTenantRequest {
    body: TetrateApiTccCoreV1CreateTenantRequest;
}

export interface CreateUserRequest {
    tenant: string;
    body: TetrateApiTccCoreV1CreateUserRequest;
}

export interface DeleteTeamRequest {
    tenant: string;
    id: string;
    name?: string;
}

export interface DeleteTenantRequest {
    id: string;
    name?: string;
}

export interface DeleteUserRequest {
    tenant: string;
    id: string;
    name?: string;
}

export interface GetTeamRequest {
    tenant: string;
    id: string;
    name?: string;
}

export interface GetTenantRequest {
    id: string;
    name?: string;
}

export interface GetUserRequest {
    tenant: string;
    id: string;
    name?: string;
}

export interface ListTeamsRequest {
    tenant: string;
    parent?: string;
}

export interface ListUsersRequest {
    tenant: string;
    parent?: string;
}

export interface UpdateTeamRequest {
    tenant: string;
    id: string;
    body: TetrateApiTccCoreV1Team;
}

export interface UpdateTenantRequest {
    id: string;
    body: TetrateApiTccCoreV1Tenant;
}

export interface UpdateUserRequest {
    tenant: string;
    id: string;
    body: TetrateApiTccCoreV1User;
}

/**
 * no description
 */
export class OrganizationApi extends runtime.BaseAPI {

    /**
     */
    async createTeamRaw(requestParameters: CreateTeamRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling createTeam.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createTeam.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/teams`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1CreateTeamRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TeamFromJSON(jsonValue));
    }

    /**
     */
    async createTeam(requestParameters: CreateTeamRequest): Promise<TetrateApiTccCoreV1Team> {
        const response = await this.createTeamRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async createTenantRaw(requestParameters: CreateTenantRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createTenant.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1CreateTenantRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TenantFromJSON(jsonValue));
    }

    /**
     */
    async createTenant(requestParameters: CreateTenantRequest): Promise<TetrateApiTccCoreV1Tenant> {
        const response = await this.createTenantRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling createUser.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/users`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1CreateUserRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1UserFromJSON(jsonValue));
    }

    /**
     */
    async createUser(requestParameters: CreateUserRequest): Promise<TetrateApiTccCoreV1User> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteTeamRaw(requestParameters: DeleteTeamRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deleteTeam.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteTeam.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/teams/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async deleteTeam(requestParameters: DeleteTeamRequest): Promise<object> {
        const response = await this.deleteTeamRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteTenantRaw(requestParameters: DeleteTenantRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteTenant.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async deleteTenant(requestParameters: DeleteTenantRequest): Promise<object> {
        const response = await this.deleteTenantRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling deleteUser.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/users/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.TextApiResponse(response);
    }

    /**
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getTeamRaw(requestParameters: GetTeamRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getTeam.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTeam.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/teams/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TeamFromJSON(jsonValue));
    }

    /**
     */
    async getTeam(requestParameters: GetTeamRequest): Promise<TetrateApiTccCoreV1Team> {
        const response = await this.getTeamRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getTenantRaw(requestParameters: GetTenantRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTenant.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TenantFromJSON(jsonValue));
    }

    /**
     */
    async getTenant(requestParameters: GetTenantRequest): Promise<TetrateApiTccCoreV1Tenant> {
        const response = await this.getTenantRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling getUser.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/users/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1UserFromJSON(jsonValue));
    }

    /**
     */
    async getUser(requestParameters: GetUserRequest): Promise<TetrateApiTccCoreV1User> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listTeamsRaw(requestParameters: ListTeamsRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listTeams.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.parent !== undefined) {
            queryParameters['parent'] = requestParameters.parent;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/teams`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1ListTeamsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listTeams(requestParameters: ListTeamsRequest): Promise<TetrateApiTccCoreV1ListTeamsResponse> {
        const response = await this.listTeamsRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listTenantsRaw(): Promise<runtime.ApiResponse<any>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1ListTenantsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listTenants(): Promise<TetrateApiTccCoreV1ListTenantsResponse> {
        const response = await this.listTenantsRaw();
        return await response.value();
    }

    /**
     */
    async listUsersRaw(requestParameters: ListUsersRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling listUsers.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.parent !== undefined) {
            queryParameters['parent'] = requestParameters.parent;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/tenants/{tenant}/users`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1ListUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async listUsers(requestParameters: ListUsersRequest): Promise<TetrateApiTccCoreV1ListUsersResponse> {
        const response = await this.listUsersRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateTeamRaw(requestParameters: UpdateTeamRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling updateTeam.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateTeam.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateTeam.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/teams/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1TeamToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TeamFromJSON(jsonValue));
    }

    /**
     */
    async updateTeam(requestParameters: UpdateTeamRequest): Promise<TetrateApiTccCoreV1Team> {
        const response = await this.updateTeamRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateTenantRaw(requestParameters: UpdateTenantRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateTenant.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateTenant.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1TenantToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1TenantFromJSON(jsonValue));
    }

    /**
     */
    async updateTenant(requestParameters: UpdateTenantRequest): Promise<TetrateApiTccCoreV1Tenant> {
        const response = await this.updateTenantRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateUserRaw(requestParameters: UpdateUserRequest): Promise<runtime.ApiResponse<any>> {
        if (requestParameters.tenant === null || requestParameters.tenant === undefined) {
            throw new runtime.RequiredError('tenant','Required parameter requestParameters.tenant was null or undefined when calling updateUser.');
        }

        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/tenants/{tenant}/users/{id}`.replace(`{${"tenant"}}`, encodeURIComponent(String(requestParameters.tenant))).replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: TetrateApiTccCoreV1UserToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => TetrateApiTccCoreV1UserFromJSON(jsonValue));
    }

    /**
     */
    async updateUser(requestParameters: UpdateUserRequest): Promise<TetrateApiTccCoreV1User> {
        const response = await this.updateUserRaw(requestParameters);
        return await response.value();
    }

}
