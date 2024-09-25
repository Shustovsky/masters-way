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
import type { SchemasDefaultWayCollections } from './SchemasDefaultWayCollections';
import {
    SchemasDefaultWayCollectionsFromJSON,
    SchemasDefaultWayCollectionsFromJSONTyped,
    SchemasDefaultWayCollectionsToJSON,
} from './SchemasDefaultWayCollections';
import type { SchemasUserPlainResponse } from './SchemasUserPlainResponse';
import {
    SchemasUserPlainResponseFromJSON,
    SchemasUserPlainResponseFromJSONTyped,
    SchemasUserPlainResponseToJSON,
} from './SchemasUserPlainResponse';
import type { SchemasUserTagResponse } from './SchemasUserTagResponse';
import {
    SchemasUserTagResponseFromJSON,
    SchemasUserTagResponseFromJSONTyped,
    SchemasUserTagResponseToJSON,
} from './SchemasUserTagResponse';
import type { SchemasWayCollectionPopulatedResponse } from './SchemasWayCollectionPopulatedResponse';
import {
    SchemasWayCollectionPopulatedResponseFromJSON,
    SchemasWayCollectionPopulatedResponseFromJSONTyped,
    SchemasWayCollectionPopulatedResponseToJSON,
} from './SchemasWayCollectionPopulatedResponse';
import type { SchemasWayPlainResponse } from './SchemasWayPlainResponse';
import {
    SchemasWayPlainResponseFromJSON,
    SchemasWayPlainResponseFromJSONTyped,
    SchemasWayPlainResponseToJSON,
} from './SchemasWayPlainResponse';

/**
 * 
 * @export
 * @interface SchemasUserPopulatedResponse
 */
export interface SchemasUserPopulatedResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    createdAt: string;
    /**
     * 
     * @type {Array<SchemasWayCollectionPopulatedResponse>}
     * @memberof SchemasUserPopulatedResponse
     */
    customWayCollections: Array<SchemasWayCollectionPopulatedResponse>;
    /**
     * 
     * @type {SchemasDefaultWayCollections}
     * @memberof SchemasUserPopulatedResponse
     */
    defaultWayCollections: SchemasDefaultWayCollections;
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    description: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    email: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SchemasUserPopulatedResponse
     */
    favoriteForUsers: Array<string>;
    /**
     * 
     * @type {Array<SchemasUserPlainResponse>}
     * @memberof SchemasUserPopulatedResponse
     */
    favoriteUsers: Array<SchemasUserPlainResponse>;
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    imageUrl: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasUserPopulatedResponse
     */
    isMentor: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    name: string;
    /**
     * 
     * @type {Array<SchemasUserTagResponse>}
     * @memberof SchemasUserPopulatedResponse
     */
    tags: Array<SchemasUserTagResponse>;
    /**
     * 
     * @type {string}
     * @memberof SchemasUserPopulatedResponse
     */
    uuid: string;
    /**
     * 
     * @type {Array<SchemasWayPlainResponse>}
     * @memberof SchemasUserPopulatedResponse
     */
    wayRequests: Array<SchemasWayPlainResponse>;
}

/**
 * Check if a given object implements the SchemasUserPopulatedResponse interface.
 */
export function instanceOfSchemasUserPopulatedResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "customWayCollections" in value;
    isInstance = isInstance && "defaultWayCollections" in value;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "favoriteForUsers" in value;
    isInstance = isInstance && "favoriteUsers" in value;
    isInstance = isInstance && "imageUrl" in value;
    isInstance = isInstance && "isMentor" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "tags" in value;
    isInstance = isInstance && "uuid" in value;
    isInstance = isInstance && "wayRequests" in value;

    return isInstance;
}

export function SchemasUserPopulatedResponseFromJSON(json: any): SchemasUserPopulatedResponse {
    return SchemasUserPopulatedResponseFromJSONTyped(json, false);
}

export function SchemasUserPopulatedResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasUserPopulatedResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'],
        'customWayCollections': ((json['customWayCollections'] as Array<any>).map(SchemasWayCollectionPopulatedResponseFromJSON)),
        'defaultWayCollections': SchemasDefaultWayCollectionsFromJSON(json['defaultWayCollections']),
        'description': json['description'],
        'email': json['email'],
        'favoriteForUsers': json['favoriteForUsers'],
        'favoriteUsers': ((json['favoriteUsers'] as Array<any>).map(SchemasUserPlainResponseFromJSON)),
        'imageUrl': json['imageUrl'],
        'isMentor': json['isMentor'],
        'name': json['name'],
        'tags': ((json['tags'] as Array<any>).map(SchemasUserTagResponseFromJSON)),
        'uuid': json['uuid'],
        'wayRequests': ((json['wayRequests'] as Array<any>).map(SchemasWayPlainResponseFromJSON)),
    };
}


export function SchemasUserPopulatedResponseToJSON(value?: SchemasUserPopulatedResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'createdAt': value.createdAt,
        'customWayCollections': ((value.customWayCollections as Array<any>).map(SchemasWayCollectionPopulatedResponseToJSON)),
        'defaultWayCollections': SchemasDefaultWayCollectionsToJSON(value.defaultWayCollections),
        'description': value.description,
        'email': value.email,
        'favoriteForUsers': value.favoriteForUsers,
        'favoriteUsers': ((value.favoriteUsers as Array<any>).map(SchemasUserPlainResponseToJSON)),
        'imageUrl': value.imageUrl,
        'isMentor': value.isMentor,
        'name': value.name,
        'tags': ((value.tags as Array<any>).map(SchemasUserTagResponseToJSON)),
        'uuid': value.uuid,
        'wayRequests': ((value.wayRequests as Array<any>).map(SchemasWayPlainResponseToJSON)),
    };
}

