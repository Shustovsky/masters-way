// @ts-nocheck
/* eslint-disable */
/**
 * Masters way general API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SchemasUpdateJobTagPayload
 */
export interface SchemasUpdateJobTagPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateJobTagPayload
     */
    color?: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateJobTagPayload
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateJobTagPayload
     */
    name?: string;
}

/**
 * Check if a given object implements the SchemasUpdateJobTagPayload interface.
 */
export function instanceOfSchemasUpdateJobTagPayload(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasUpdateJobTagPayloadFromJSON(json: any): SchemasUpdateJobTagPayload {
    return SchemasUpdateJobTagPayloadFromJSONTyped(json, false);
}

export function SchemasUpdateJobTagPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUpdateJobTagPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'color': !exists(json, 'color') ? undefined : json['color'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}


export function SchemasUpdateJobTagPayloadToJSON(value?: SchemasUpdateJobTagPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'color': value.color,
        'description': value.description,
        'name': value.name,
    };
}

