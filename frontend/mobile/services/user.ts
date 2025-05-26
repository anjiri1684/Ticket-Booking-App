import { AuthResponse } from "@/types/user";
import { Api } from "./api";

type Credentials = {
     email: string
     password: string
}

async function login(credendtials: Credentials): Promise<AuthResponse>{
     return Api.post("/auth/login", credendtials)
}

async function register(credendtials: Credentials): Promise<AuthResponse>{
     return Api.post("/auth/register", credendtials)
}

const userService = {
     login, register
}

export {userService}