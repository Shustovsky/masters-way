// @ts-nocheck
/* eslint-disable */
/**
 * Masters way chat API
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
 * @interface SchemasCreateRoomPayload
 */
export interface SchemasCreateRoomPayload {
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateRoomPayload
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateRoomPayload
     */
    roomType: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasCreateRoomPayload
     */
    userId?: string | null;
}

/**
 * Check if a given object implements the SchemasCreateRoomPayload interface.
 */
export function instanceOfSchemasCreateRoomPayload(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "roomType" in value;

    return isInstance;
}

export function SchemasCreateRoomPayloadFromJSON(json: any): SchemasCreateRoomPayload {
    return SchemasCreateRoomPayloadFromJSONTyped(json, false);
}

export function SchemasCreateRoomPayloadFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCreateRoomPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'roomType': json['roomType'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
    };
}


export function SchemasCreateRoomPayloadToJSON(value?: SchemasCreateRoomPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'roomType': value.roomType,
        'userId': value.userId,
    };
}

