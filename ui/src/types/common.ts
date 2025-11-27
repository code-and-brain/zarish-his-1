// Common types used across the application

export interface GeoPoint {
    latitude: number;
    longitude: number;
    altitude?: number;
    accuracy?: number;
}

export interface Address {
    line1: string;
    line2?: string;
    city: string;
    state: string;
    postalCode: string;
    country: string;
    district?: string;
    subDistrict?: string;
    village?: string;
}

export interface ContactInfo {
    phone?: string;
    mobile?: string;
    email?: string;
    alternatePhone?: string;
}

export interface Period {
    start: Date;
    end?: Date;
}

export interface Reference {
    id: string;
    type: string;
    display?: string;
}

export interface Identifier {
    system: string;
    value: string;
    use?: 'official' | 'temp' | 'secondary';
    period?: Period;
}

export interface Coding {
    system: string;
    code: string;
    display?: string;
    version?: string;
}

export interface CodeableConcept {
    coding: Coding[];
    text?: string;
}

export interface Attachment {
    contentType: string;
    data?: string;
    url?: string;
    size?: number;
    hash?: string;
    title?: string;
    creation?: Date;
}

export interface Annotation {
    authorReference?: Reference;
    authorString?: string;
    time: Date;
    text: string;
}

export interface SyncStatus {
    lastSync: Date;
    syncVersion: number;
    conflicts?: string[];
    status: 'synced' | 'pending' | 'conflict' | 'error';
}

export interface AuditInfo {
    createdBy: string;
    createdAt: Date;
    updatedBy?: string;
    updatedAt?: Date;
    deletedBy?: string;
    deletedAt?: Date;
}

export interface Pagination {
    page: number;
    pageSize: number;
    totalPages: number;
    totalItems: number;
}

export interface SearchParams {
    query?: string;
    filters?: { [key: string]: any };
    sort?: {
        field: string;
        order: 'asc' | 'desc';
    };
    pagination?: {
        page: number;
        pageSize: number;
    };
}

export interface ApiResponse<T> {
    success: boolean;
    data?: T;
    error?: {
        code: string;
        message: string;
        details?: any;
    };
    metadata?: {
        timestamp: Date;
        requestId: string;
        pagination?: Pagination;
    };
}
