import { StatusBar } from "react-native";
import { Slot } from "expo-router";
export default function RootLayout() {
  return (
    <>
      <StatusBar barStyle={"dark-content"} />;{/* Authentication Provider */}
      <Slot />
    </>
  );
}
