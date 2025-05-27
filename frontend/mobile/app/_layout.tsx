import { StatusBar } from "react-native";
import { Slot } from "expo-router";
import { AuthenticationProvider } from "@/context/AuthContext";
export default function RootLayout() {
  return (
    <>
      <StatusBar barStyle={"dark-content"} />
      <AuthenticationProvider>
        <Slot />
      </AuthenticationProvider>
    </>
  );
}
