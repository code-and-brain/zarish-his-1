// Authentication and Authorization types
import type { GeoPoint } from './common';

export interface User {
    id: string;
    username: string;
    email: string;
    firstName: string;
    lastName: string;
    displayName: string;
    avatar?: string;
    roles: Role[];
    permissions: Permission[];
    status: 'active' | 'inactive' | 'suspended' | 'locked';
    emailVerified: boolean;
    phoneVerified: boolean;
    mfaEnabled: boolean;
    lastLogin?: Date;
    createdAt: Date;
    updatedAt?: Date;
}

export interface Role {
    id: string;
    name: string;
    description: string;
    permissions: Permission[];
    level: number;
    isSystem: boolean;
    createdAt: Date;
}

export interface Permission {
    id: string;
    resource: string;
    action: 'create' | 'read' | 'update' | 'delete' | 'execute';
    scope?: 'own' | 'team' | 'organization' | 'all';
    conditions?: PermissionCondition[];
}

export interface PermissionCondition {
    field: string;
    operator: 'equals' | 'not_equals' | 'in' | 'not_in' | 'greater_than' | 'less_than';
    value: any;
}

export interface AuthToken {
    id: string;
    type: 'access_token' | 'refresh_token' | 'api_key' | 'session_token';
    userId: string;
    token: string;
    scope: string[];
    permissions: string[];
    issuedAt: Date;
    expiresAt: Date;
    lastUsed?: Date;
    status: 'active' | 'expired' | 'revoked' | 'suspended';
}

export interface AuthSession {
    id: string;
    userId: string;
    deviceInfo: DeviceInfo;
    ipAddress: string;
    userAgent: string;
    startTime: Date;
    lastActivity: Date;
    expiresAt: Date;
    status: 'active' | 'expired' | 'terminated' | 'suspicious';
    location?: SessionLocation;
}

export interface DeviceInfo {
    deviceId: string;
    deviceName?: string;
    deviceType: 'mobile' | 'desktop' | 'tablet' | 'server' | 'unknown';
    platform: string;
    browser?: string;
    operatingSystem: string;
    trusted: boolean;
}

export interface SessionLocation {
    country?: string;
    region?: string;
    city?: string;
    coordinates?: GeoPoint;
    timezone: string;
}

export interface MFAConfig {
    enabled: boolean;
    required: boolean;
    methods: MFAMethod[];
    backupCodes: string[];
}

export interface MFAMethod {
    id: string;
    type: 'totp' | 'sms' | 'email' | 'push' | 'hardware_key' | 'biometric';
    name: string;
    enabled: boolean;
    verified: boolean;
    primary: boolean;
    setupDate: Date;
    lastUsed?: Date;
}

export interface AuditLog {
    id: string;
    timestamp: Date;
    userId?: string;
    sessionId?: string;
    action: string;
    resource: string;
    resourceId?: string;
    result: 'success' | 'failure' | 'warning';
    ipAddress: string;
    userAgent: string;
    location?: GeoPoint;
    details?: { [key: string]: any };
}
