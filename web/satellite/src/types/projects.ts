// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

/**
 * Exposes all project-related functionality.
 */
export interface ProjectsApi {
    /**
     * Creates project.
     *
     * @param createProjectFields - contains project information
     * @throws Error
     */
    create(createProjectFields: ProjectFields): Promise<Project>;
    /**
     * Fetch projects.
     *
     * @returns Project[]
     * @throws Error
     */
    get(): Promise<Project[]>;
    /**
     * Update project name and description.
     *
     * @param projectId - project ID
     * @param name - project name
     * @param description - project description
     * @returns Project[]
     * @throws Error
     */
    update(projectId: string, name: string, description: string): Promise<void>;
    /**
     * Delete project.
     *
     * @param projectId - project ID
     * @throws Error
     */
    delete(projectId: string): Promise<void>;

    /**
     * Get project limits.
     *
     * @param projectId- project ID
     * throws Error
     */
    getLimits(projectId: string): Promise<ProjectLimits>;

    /**
     * Fetch owned projects.
     *
     * @returns ProjectsPage
     * @throws Error
     */
    getOwnedProjects(cursor: ProjectsCursor): Promise<ProjectsPage>;
}

/**
 * MAX_NAME_LENGTH defines maximum amount of symbols for project name.
 */
export const MAX_NAME_LENGTH = 20;

/**
 * MAX_DESCRIPTION_LENGTH defines maximum amount of symbols for project description.
 */
export const MAX_DESCRIPTION_LENGTH = 100;

/**
 * Project is a type, used for creating new project in backend.
 */
export class Project {
    public constructor(
        public id: string = '',
        public name: string = '',
        public description: string = '',
        public createdAt: string = '',
        public ownerId: string = '',
        public isSelected: boolean = false,
        public memberCount: number = 0,
    ) {}

    /**
     * Returns created date as a local string.
     */
    public createdDate(): string {
        const createdAt = new Date(this.createdAt);
        return createdAt.toLocaleString('en-US', {timeZone: 'UTC', timeZoneName: 'short'});
    }
}

/**
 * ProjectFields is a type, used for creating and updating project.
 */
export class ProjectFields {
    public constructor(
        public name: string = '',
        public description: string = '',
        public ownerId: string = '',
    ) {}

    /**
     * checkName checks if project name is valid.
     */
    public checkName(): void {
        try {
            this.nameIsNotEmpty();
            this.nameHasLessThenTwentySymbols();
        } catch (error) {
            throw new Error(error.message);
        }
    }

    /**
     * nameIsNotEmpty checks if project name is not empty.
     */
    private nameIsNotEmpty(): void {
        if (this.name.length === 0) throw new Error('Project name can\'t be empty!');
    }

    /**
     * nameHasLessThenTwentySymbols checks if project name has less then 20 symbols.
     */
    private nameHasLessThenTwentySymbols(): void {
        if (this.name.length > MAX_NAME_LENGTH) throw new Error('Name should be less than 21 character!');
    }
}

/**
 * ProjectLimits is a type, used for describing project limits.
 */
export class ProjectLimits {
    public constructor(
        public bandwidthLimit: number = 0,
        public bandwidthUsed: number = 0,
        public storageLimit: number = 0,
        public storageUsed: number = 0,
    ) {}
}

export class ProjectPage {
    public constructor(
        public projects: Project[] = [],
        public pageCount: number = 0,
        public currentPage: number = 0,
        public totalCount: number = 0,
    ){}
}

/**
 * BucketPage class holds bucket total usages and flag whether more usages available.
 */
export class ProjectsPage {
    public constructor(
        public projects: Project[] = [],
        public limit: number = 0,
        public offset: number = 0,
        public pageCount: number = 0,
        public currentPage: number = 0,
        public totalCount: number = 0,
    ) {}
}

/**
 * BucketCursor class holds cursor for bucket name and limit.
 */
export class ProjectsCursor {
    public constructor(
        public limit: number = 0,
        public page: number = 0,
    ) {}
}
