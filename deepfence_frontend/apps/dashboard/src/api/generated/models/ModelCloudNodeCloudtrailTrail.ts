/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.2.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ModelCloudNodeCloudtrailTrail
 */
export interface ModelCloudNodeCloudtrailTrail {
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeCloudtrailTrail
     */
    account_id?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeCloudtrailTrail
     */
    trail_name?: string;
}

/**
 * Check if a given object implements the ModelCloudNodeCloudtrailTrail interface.
 */
export function instanceOfModelCloudNodeCloudtrailTrail(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelCloudNodeCloudtrailTrailFromJSON(json: any): ModelCloudNodeCloudtrailTrail {
    return ModelCloudNodeCloudtrailTrailFromJSONTyped(json, false);
}

export function ModelCloudNodeCloudtrailTrailFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCloudNodeCloudtrailTrail {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'account_id': !exists(json, 'account_id') ? undefined : json['account_id'],
        'trail_name': !exists(json, 'trail_name') ? undefined : json['trail_name'],
    };
}

export function ModelCloudNodeCloudtrailTrailToJSON(value?: ModelCloudNodeCloudtrailTrail | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'account_id': value.account_id,
        'trail_name': value.trail_name,
    };
}

