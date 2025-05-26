import { userService } from "@/services/user";
import { User } from "@/types/user";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import { createContext, PropsWithChildren,  use,  useContext, useEffect, useState } from "react";

interface AuthContextProps {
     isLoggedIn: boolean;
     isLoadingAuth: boolean;
     authenticate: (authMode: "login" | "register", email: string, password: string) => Promise<void>
     logout: VoidFunction;
     user: User | null
}

const AuthContext = createContext({} as AuthContextProps);

//create custome useAuth hook
export function useAuth() {
     return useContext(AuthContext)
}

export function AuthenticationProvider({ children }: PropsWithChildren) {
     const [isLoggedIn, setIsLoggedIn] = useState(false)
     const [isLoadingAuth, setIsLoadingAuth] = useState(false)
     const [user, setUser] = useState<User | null>(null);

     useEffect(() => {
          async function checkIfLoggedIn() {
               const token = await AsyncStorage.getItem("token")
               const user = await AsyncStorage.getItem("user")

               if (token && user) {
                    setIsLoggedIn(true)
                    setUser(JSON.parse(user))
                    router.replace("/(authed)/(tabs)/(events)");

               } else {
                    setIsLoggedIn(false)
               }
          }
          checkIfLoggedIn()
     }, [])


     async function authenticate(authMode: "login" | "register", email: string, password: string): Promise<void> {
          try {
               setIsLoadingAuth(true)

               const responce = await userService[authMode]({ email, password })
               
               if (responce) {
                 const { data } = responce;
                 const { user, token } = data;

                 await AsyncStorage.setItem("token", token);
                 await AsyncStorage.setItem("user", JSON.stringify(user));

                 setUser(user);
                 setIsLoggedIn(true); 

                 router.replace("/(authed)/(tabs)/(events)");
               }
                
          } catch (error) {
               setIsLoggedIn(false)
               console.log(error)
          } finally {
               setIsLoadingAuth(false)
          }
     }

     async function logout() {
          setIsLoggedIn(false)

          await AsyncStorage.removeItem("token")
          await AsyncStorage.removeItem("user")
     }

     return (
          <AuthContext.Provider value={{ isLoadingAuth, isLoggedIn, authenticate,user,logout}}>
               {children}
          </AuthContext.Provider>
     )
}