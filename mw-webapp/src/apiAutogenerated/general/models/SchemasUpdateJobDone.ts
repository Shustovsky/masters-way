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
 * @interface SchemasUpdateJobDone
 */
export interface SchemasUpdateJobDone {
    /**
     * 
     * @type {string}
     * @memberof SchemasUpdateJobDone
     */
    description?: string;
    /**
     * 
     * @type {number}
     * @memberof SchemasUpdateJobDone
     */
    time?: number;
}

/**
 * Check if a given object implements the SchemasUpdateJobDone interface.
 */
export function instanceOfSchemasUpdateJobDone(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasUpdateJobDoneFromJSON(json: any): SchemasUpdateJobDone {
    return SchemasUpdateJobDoneFromJSONTyped(json, false);
}

export function SchemasUpdateJobDoneFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUpdateJobDone {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': !exists(json, 'description') ? undefined : json['description'],
        'time': !exists(json, 'time') ? undefined : json['time'],
    };
}


export function SchemasUpdateJobDoneToJSON(value?: SchemasUpdateJobDone | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
        'time': value.time,
    };
}

