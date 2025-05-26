import { Stack } from "expo-router"; 
import { View } from "react-native";

export default function TicketsLayout() {
  return (
    <View>
      <Stack screenOptions={{ headerBackTitle: "Tickets" }}>
        <Stack.Screen name="index" />
        <Stack.Screen name="ticket/[id]" />
      </Stack>
    </View>
  );
}