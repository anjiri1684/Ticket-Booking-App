import { userService } from "@/services/user";
import { User } from "@/types/user";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { router } from "expo-router";
import {
  createContext,
  PropsWithChildren,
  useContext,
  useEffect,
  useState,
} from "react";

interface AuthContextProps {
  isLoggedIn: boolean;
  isLoadingAuth: boolean;
  authenticate: (
    authMode: "login" | "register",
    email: string,
    password: string
  ) => Promise<void>;
  logout: VoidFunction;
  user: User | null;
  isCheckingAuth: boolean;
}

const AuthContext = createContext({} as AuthContextProps);

export function useAuth() {
  return useContext(AuthContext);
}

export function AuthenticationProvider({ children }: PropsWithChildren) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isLoadingAuth, setIsLoadingAuth] = useState(false);
  const [isCheckingAuth, setIsCheckingAuth] = useState(true);
  const [user, setUser] = useState<User | null>(null);
  const [routerReady, setRouterReady] = useState(false);

  // Ensure router is ready before any redirects
  useEffect(() => {
    const timeout = setTimeout(() => {
      setRouterReady(true);
      console.log("[Auth] Router is ready");
    }, 100);
    return () => clearTimeout(timeout);
  }, []);

  // Initial auth check, runs once router is ready
  useEffect(() => {
    if (!routerReady) return;

    async function checkIfLoggedIn() {
      try {
        setIsCheckingAuth(true);
        const token = await AsyncStorage.getItem("token");
        const userString = await AsyncStorage.getItem("user");
        const storedUser = userString ? JSON.parse(userString) : null;

        console.log("[Auth] Checking stored auth info...");
        console.log("[Auth] Token:", token ?? "null");
        console.log("[Auth] User:", storedUser ?? "null");

        if (token && storedUser) {
          setUser(storedUser);
          setIsLoggedIn(true);
          console.log("[Auth] User is authenticated, redirecting...");
          router.replace("/(authed)/(tabs)/(events)");
        } else {
          setUser(null);
          setIsLoggedIn(false);
          console.log("[Auth] No valid auth found.");
        }
      } catch (error) {
        console.error("[Auth] Error checking auth:", error);
        setUser(null);
        setIsLoggedIn(false);
      } finally {
        setIsCheckingAuth(false);
      }
    }
    checkIfLoggedIn();
  }, [routerReady]);

  // Authenticator method for login/register
  async function authenticate(
    authMode: "login" | "register",
    email: string,
    password: string
  ): Promise<void> {
    try {
      setIsLoadingAuth(true);

      const response = await userService[authMode]({ email, password });

      if (!response || !response.data) {
        throw new Error("Invalid response from auth service");
      }

      const { user: authenticatedUser, token } = response.data;

      if (!token || !authenticatedUser) {
        throw new Error("Missing token or user in auth response");
      }

      await AsyncStorage.setItem("token", token);
      await AsyncStorage.setItem("user", JSON.stringify(authenticatedUser));

      setUser(authenticatedUser);
      setIsLoggedIn(true);

      console.log("[Auth] Authentication successful.");
      console.log("[Auth] Token:", token);
      console.log("[Auth] User:", authenticatedUser);

      if (routerReady) {
        router.replace("/(authed)/(tabs)/(events)");
        console.log("[Auth] Redirected post-authentication.");
      }
    } catch (error) {
      setIsLoggedIn(false);
      setUser(null);
      console.error("[Auth] Authentication failed:", error);
    } finally {
      setIsLoadingAuth(false);
    }
  }

  // Logout method
  async function logout() {
    try {
      setIsLoggedIn(false);
      setUser(null);
      await AsyncStorage.removeItem("token");
      await AsyncStorage.removeItem("user");
      console.log("[Auth] User logged out, cleared storage.");
      router.replace("/login"); // Redirect explicitly on logout
    } catch (error) {
      console.error("[Auth] Error during logout:", error);
    }
  }

  return (
    <AuthContext.Provider
      value={{
        isLoadingAuth,
        isLoggedIn,
        authenticate,
        user,
        logout,
        isCheckingAuth,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
