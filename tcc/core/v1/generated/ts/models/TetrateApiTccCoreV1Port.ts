// tslint:disable
/**
 * TCC Core API
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
 * 
 * @export
 * @interface TetrateApiTccCoreV1Port
 */
export interface TetrateApiTccCoreV1Port {
    /**
     * A valid non-negative integer port number.
     * @type {number}
     * @memberof TetrateApiTccCoreV1Port
     */
    number?: number;
    /**
     *   The protocol exposed on the port. MUST BE one of HTTP|GRPC|HTTP2|HTTPS|TCP|TLS. TLS implies the connection will be routed based on the SNI header to the destination without terminating the TLS connection.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Port
     */
    protocol?: string;
    /**
     * Name assigned to the port.
     * @type {string}
     * @memberof TetrateApiTccCoreV1Port
     */
    name?: string;
    /**
     * The endpoint port to which this service port maps to. For example, service port 80 exposed on the load balancer could map to an endpoint port 9080 on a VM.
     * @type {number}
     * @memberof TetrateApiTccCoreV1Port
     */
    endpointPort?: number;
}

export function TetrateApiTccCoreV1PortFromJSON(json: any): TetrateApiTccCoreV1Port {
    return {
        'number': !exists(json, 'number') ? undefined : json['number'],
        'protocol': !exists(json, 'protocol') ? undefined : json['protocol'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'endpointPort': !exists(json, 'endpointPort') ? undefined : json['endpointPort'],
    };
}

export function TetrateApiTccCoreV1PortToJSON(value?: TetrateApiTccCoreV1Port): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'number': value.number,
        'protocol': value.protocol,
        'name': value.name,
        'endpointPort': value.endpointPort,
    };
}


