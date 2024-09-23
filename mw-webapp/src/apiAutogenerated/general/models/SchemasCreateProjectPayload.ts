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
 * @interface SchemasCreateProjectPayload
 */
export interface SchemasCreateProjectPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateProjectPayload
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateProjectPayload
     */
    ownerId: string;
}

/**
 * Check if a given object implements the SchemasCreateProjectPayload interface.
 */
export function instanceOfSchemasCreateProjectPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "ownerId" in value;

    return isInstance;
}

export function SchemasCreateProjectPayloadFromJSON(json: any): SchemasCreateProjectPayload {
    return SchemasCreateProjectPayloadFromJSONTyped(json, false);
}

export function SchemasCreateProjectPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreateProjectPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': json['name'],
        'ownerId': json['ownerId'],
    };
}


export function SchemasCreateProjectPayloadToJSON(value?: SchemasCreateProjectPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'ownerId': value.ownerId,
    };
}
