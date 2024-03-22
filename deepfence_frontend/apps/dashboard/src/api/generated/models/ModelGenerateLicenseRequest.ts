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
 * @interface ModelGenerateLicenseRequest
 */
export interface ModelGenerateLicenseRequest {
    /**
     * 
     * @type {string}
     * @memberof ModelGenerateLicenseRequest
     */
    company: string;
    /**
     * 
     * @type {string}
     * @memberof ModelGenerateLicenseRequest
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof ModelGenerateLicenseRequest
     */
    first_name: string;
    /**
     * 
     * @type {string}
     * @memberof ModelGenerateLicenseRequest
     */
    last_name: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGenerateLicenseRequest
     */
    resend_email: boolean;
}

/**
 * Check if a given object implements the ModelGenerateLicenseRequest interface.
 */
export function instanceOfModelGenerateLicenseRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "company" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "first_name" in value;
    isInstance = isInstance && "last_name" in value;
    isInstance = isInstance && "resend_email" in value;

    return isInstance;
}

export function ModelGenerateLicenseRequestFromJSON(json: any): ModelGenerateLicenseRequest {
    return ModelGenerateLicenseRequestFromJSONTyped(json, false);
}

export function ModelGenerateLicenseRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelGenerateLicenseRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'company': json['company'],
        'email': json['email'],
        'first_name': json['first_name'],
        'last_name': json['last_name'],
        'resend_email': json['resend_email'],
    };
}

export function ModelGenerateLicenseRequestToJSON(value?: ModelGenerateLicenseRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'company': value.company,
        'email': value.email,
        'first_name': value.first_name,
        'last_name': value.last_name,
        'resend_email': value.resend_email,
    };
}

