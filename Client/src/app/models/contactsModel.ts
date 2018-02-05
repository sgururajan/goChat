export enum UserStatus {
    Online = 'online',
    Ideal = 'ideal',
    Offline = 'offline'
}

export class Contact {
    email: string
    firstName: string
    lastName: string
    avatarUrl: string
    userStatus: UserStatus
}