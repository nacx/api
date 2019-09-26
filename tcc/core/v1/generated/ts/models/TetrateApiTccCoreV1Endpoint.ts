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
/**
 * Endpoint defines a network address (IP or hostname) associated with the service.
 * @export
 * @interface TetrateApiTccCoreV1Endpoint
 */
export interface TetrateApiTccCoreV1Endpoint {
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    environment?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    cluster?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    namespace?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    deployment?: string;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    id?: string;
    /**
     * Address associated with the network endpoint without the port. Domain names can be used and must be fully-qualified without wildcards.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    address?: string;
    /**
     * Set of inbound traffic ports associated with the endpoint. The ports must be associated with a port number that was declared as part of the service.
     * @type {{ [key: string]: number; }}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    ports?: { [key: string]: number; };
    /**
     * One or more labels associated with the endpoint.
     * @type {{ [key: string]: string; }}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    labels?: { [key: string]: string; };
    /**
     * The locality associated with the endpoint, in the form country/region/zone. A locality corresponds to a failure domain (country/region/zone).
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    locality?: string;
    /**
     * The load balancing weight associated with the endpoint. Endpoints with higher weights in a pool will receive proportionally higher traffic.
     * @type {number}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    weight?: number;
    /**
     * 
     * @type {string}
     * @memberof TetrateApiTccCoreV1Endpoint
     */
    etag?: string;
}

export function TetrateApiTccCoreV1EndpointFromJSON(json: any): TetrateApiTccCoreV1Endpoint {
    return {
        'name': !exists(json, 'name') ? undefined : json['name'],
        'tenant': !exists(json, 'tenant') ? undefined : json['tenant'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
        'deployment': !exists(json, 'deployment') ? undefined : json['deployment'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'address': !exists(json, 'address') ? undefined : json['address'],
        'ports': !exists(json, 'ports') ? undefined : json['ports'],
        'labels': !exists(json, 'labels') ? undefined : json['labels'],
        'locality': !exists(json, 'locality') ? undefined : json['locality'],
        'weight': !exists(json, 'weight') ? undefined : json['weight'],
        'etag': !exists(json, 'etag') ? undefined : json['etag'],
    };
}

export function TetrateApiTccCoreV1EndpointToJSON(value?: TetrateApiTccCoreV1Endpoint): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'name': value.name,
        'tenant': value.tenant,
        'environment': value.environment,
        'cluster': value.cluster,
        'namespace': value.namespace,
        'deployment': value.deployment,
        'id': value.id,
        'address': value.address,
        'ports': value.ports,
        'labels': value.labels,
        'locality': value.locality,
        'weight': value.weight,
        'etag': value.etag,
    };
}

