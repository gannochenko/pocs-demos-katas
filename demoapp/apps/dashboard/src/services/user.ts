export enum UserRoleEnum {
    admin = 'admin',
    contributor = 'contributor',
    cicd = 'cicd',
}

export type UserType = {
    id: string;
    roles: string[];
};
