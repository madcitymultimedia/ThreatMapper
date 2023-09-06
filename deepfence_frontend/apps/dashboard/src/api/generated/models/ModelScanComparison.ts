/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
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
 * @interface ModelScanComparison
 */
export interface ModelScanComparison {
    /**
     * 
     * @type {any}
     * @memberof ModelScanComparison
     */
    deleted: any | null;
    /**
     * 
     * @type {any}
     * @memberof ModelScanComparison
     */
    _new: any | null;
    /**
     * 
     * @type {any}
     * @memberof ModelScanComparison
     */
    updated: any | null;
}

/**
 * Check if a given object implements the ModelScanComparison interface.
 */
export function instanceOfModelScanComparison(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "deleted" in value;
    isInstance = isInstance && "_new" in value;
    isInstance = isInstance && "updated" in value;

    return isInstance;
}

export function ModelScanComparisonFromJSON(json: any): ModelScanComparison {
    return ModelScanComparisonFromJSONTyped(json, false);
}

export function ModelScanComparisonFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelScanComparison {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deleted': json['deleted'],
        '_new': json['new'],
        'updated': json['updated'],
    };
}

export function ModelScanComparisonToJSON(value?: ModelScanComparison | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deleted': value.deleted,
        'new': value._new,
        'updated': value.updated,
    };
}

