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
 * @interface SchemasGetRoomPreviewResponse
 */
export interface SchemasGetRoomPreviewResponse {
    /**
     * 
     * @type {number}
     * @memberof SchemasGetRoomPreviewResponse
     */
    unreadMessagesAmount: number;
}

/**
 * Check if a given object implements the SchemasGetRoomPreviewResponse interface.
 */
export function instanceOfSchemasGetRoomPreviewResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "unreadMessagesAmount" in value;

    return isInstance;
}

export function SchemasGetRoomPreviewResponseFromJSON(json: any): SchemasGetRoomPreviewResponse {
    return SchemasGetRoomPreviewResponseFromJSONTyped(json, false);
}

export function SchemasGetRoomPreviewResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasGetRoomPreviewResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'unreadMessagesAmount': json['unreadMessagesAmount'],
    };
}


export function SchemasGetRoomPreviewResponseToJSON(value?: SchemasGetRoomPreviewResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'unreadMessagesAmount': value.unreadMessagesAmount,
    };
}

