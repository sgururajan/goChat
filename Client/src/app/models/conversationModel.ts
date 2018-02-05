export class ParticipantInfo {
    firstName: string
    lastName: string
    email: string;
    participantID: string;
}

export class Conversation {
    conversationID: string;
    name: string
    participants: ParticipantInfo[]
}