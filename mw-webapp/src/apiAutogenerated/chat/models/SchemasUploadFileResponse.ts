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
 * @interface SchemasUploadFileResponse
 */
export interface SchemasUploadFileResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasUploadFileResponse
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUploadFileResponse
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUploadFileResponse
     */
    ownerId: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUploadFileResponse
     */
    previewUrl: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUploadFileResponse
     */
    srcUrl: string;
}

/**
 * Check if a given object implements the SchemasUploadFileResponse interface.
 */
export function instanceOfSchemasUploadFileResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "ownerId" in value;
    isInstance = isInstance && "previewUrl" in value;
    isInstance = isInstance && "srcUrl" in value;

    return isInstance;
}

export function SchemasUploadFileResponseFromJSON(json: any): SchemasUploadFileResponse {
    return SchemasUploadFileResponseFromJSONTyped(json, false);
}

export function SchemasUploadFileResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUploadFileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'ownerId': json['ownerId'],
        'previewUrl': json['previewUrl'],
        'srcUrl': json['srcUrl'],
    };
}


export function SchemasUploadFileResponseToJSON(value?: SchemasUploadFileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'ownerId': value.ownerId,
        'previewUrl': value.previewUrl,
        'srcUrl': value.srcUrl,
    };
}
