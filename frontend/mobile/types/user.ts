import { ApiResponse } from "./api"

export enum UserRole {
     Attendee = "attendee",
     Manager = "manager"
}

export type AuthResponse = ApiResponse<{user: User, token: string}>

export type User = {
     id: number,
     email: string,
     role: UserRole,
     CreatedAt: string,
     UpdatedAt: string
}