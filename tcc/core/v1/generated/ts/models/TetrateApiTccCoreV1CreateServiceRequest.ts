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

import { exists, mapValues } from '../runtime';
import {
    TetrateApiTccCoreV1Port,
    TetrateApiTccCoreV1PortFromJSON,
    TetrateApiTccCoreV1PortToJSON,
    TetrateApiTccCoreV1RoutingInfo,
    TetrateApiTccCoreV1RoutingInfoFromJSON,
    TetrateApiTccCoreV1RoutingInfoToJSON,
} from './';

/**
 * 
 * @export
 * @interface TetrateApiTccCoreV1CreateServiceRequest
 */
export interface TetrateApiTccCoreV1CreateServiceRequest {
    /**
     * Internal use only. Auto populated field.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    parent?: string;
    /**
     * Tenant.Id.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    tenant?: string;
    /**
     * Environment.Id.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    environment?: string;
    /**
     * Application.Id.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    application?: string;
    /**
     * If present, this will be used as the id for the created object.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    description?: string;
    /**
     * FQDN hostname of the service.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    hostname?: string;
    /**
     * User identifiable tags associated with this service.
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    labels?: { [key: string]: string; };
    /**
     * 
     * @type {Array<TetrateApiTccCoreV1Port>}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    ports?: Array<TetrateApiTccCoreV1Port>;
    /**
     * 
     * @type {TetrateApiTccCoreV1RoutingInfo}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    routingInfo?: TetrateApiTccCoreV1RoutingInfo;
    /**
     * Namespace where the service is scoped. It should be one of application namespaces. If the application has only one namespace and if this field is omitted, this filed would default to the application namespace. This field cannot be omitted if the application has more than one namespace.
     * @type {string}
     * @memberof TetrateApiTccCoreV1CreateServiceRequest
     */
    namespace?: string;
}

export function TetrateApiTccCoreV1CreateServiceRequestFromJSON(json: any): TetrateApiTccCoreV1CreateServiceRequest {
    return {
        'parent': !exists(json, 'parent') ? undefined : json['parent'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'application': !exists(json, 'application') ? undefined : json['application'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'hostname': !exists(json, 'hostname') ? undefined : json['hostname'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'ports': !exists(json, 'ports') ? undefined : (json['ports'] as Array<any>).map(TetrateApiTccCoreV1PortFromJSON),
        'routingInfo': !exists(json, 'routingInfo') ? undefined : TetrateApiTccCoreV1RoutingInfoFromJSON(json['routingInfo']),
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
    };
}

export function TetrateApiTccCoreV1CreateServiceRequestToJSON(value?: TetrateApiTccCoreV1CreateServiceRequest): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'parent': value.parent,
        'tenant': value.tenant,
        'environment': value.environment,
        'application': value.application,
        'id': value.id,
        'description': value.description,
        'hostname': value.hostname,
        'labels': value.labels,
        'ports': value.ports === undefined ? undefined : (value.ports as Array<any>).map(TetrateApiTccCoreV1PortToJSON),
        'routingInfo': TetrateApiTccCoreV1RoutingInfoToJSON(value.routingInfo),
        'namespace': value.namespace,
    };
}


