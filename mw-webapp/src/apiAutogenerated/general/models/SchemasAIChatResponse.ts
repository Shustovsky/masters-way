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
 * @interface SchemasAIChatResponse
 */
export interface SchemasAIChatResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasAIChatResponse
     */
    message: string;
}

/**
 * Check if a given object implements the SchemasAIChatResponse interface.
 */
export function instanceOfSchemasAIChatResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "message" in value;

    return isInstance;
}

export function SchemasAIChatResponseFromJSON(json: any): SchemasAIChatResponse {
    return SchemasAIChatResponseFromJSONTyped(json, false);
}

export function SchemasAIChatResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasAIChatResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'message': json['message'],
    };
}


export function SchemasAIChatResponseToJSON(value?: SchemasAIChatResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'message': value.message,
    };
}

