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

import { exists } from '../runtime';
/**
 * Port describes the properties of a specific port of a service on a VM or k8s.
 * @export
 * @interface TetrateApiTccV1Port
 */
export interface TetrateApiTccV1Port {
    /**
     * REQUIRED: A valid non-negative integer port number.
     * @type {number}
     * @memberof TetrateApiTccV1Port
     */
    number?: number;
    /**
     * The protocol exposed on the port. MUST BE one of HTTP|GRPC|HTTP2|TCP|TLS. TLS implies the connection will be routed based on the SNI header to the destination without terminating the TLS connection.
     * @type {string}
     * @memberof TetrateApiTccV1Port
     */
    protocol?: string;
    /**
     * Label assigned to the port.
     * @type {string}
     * @memberof TetrateApiTccV1Port
     */
    name?: string;
    /**
     * The endpoint port to which this service port maps to. For example, service port 80 exposed on the load balancer could map to an endpoint port 9080 on a VM.
     * @type {number}
     * @memberof TetrateApiTccV1Port
     */
    endpointPort?: number;
}

export function TetrateApiTccV1PortFromJSON(json: any): TetrateApiTccV1Port {
    return {
        'number': !exists(json, 'number') ? undefined : json['number'],
        'protocol': !exists(json, 'protocol') ? undefined : json['protocol'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'endpointPort': !exists(json, 'endpointPort') ? undefined : json['endpointPort'],
    };
}

export function TetrateApiTccV1PortToJSON(value?: TetrateApiTccV1Port): any {
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


