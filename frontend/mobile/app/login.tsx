import { Button } from "@/components/Button";
import { Divider } from "@/components/Divider";
import { HStack } from "@/components/HStack";
import { Input } from "@/components/input";
import { TabBarIcon } from "@/components/navigation/TabBarIcon";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { useAuth } from "@/context/AuthContext";
import { useState } from "react";
import { KeyboardAvoidingView, ScrollView } from "react-native";

export default function Login() {
  const { authenticate, isLoadingAuth } = useAuth();

  async function onAuthenticate() {
    await authenticate(authMode, email, password);
  }

  const [email, setEmail] = useState("");
  const [password, setpassword] = useState("");
  const [authMode, setAuthMode] = useState<"login" | "register">("login");

  function onToggleAuthMode() {
    if (authMode === "login") {
      setAuthMode("register");
    } else {
      setAuthMode("login");
    }
  }
  return (
    <KeyboardAvoidingView behavior="padding" style={{ flex: 1 }}>
      <ScrollView contentContainerStyle={{ flex: 1 }}>
        <VStack
          flex={1}
          justifyContent="center"
          alignItems="center"
          p={40}
          gap={40}
        >
          <HStack gap={10}>
            <Text fontSize={30} bold mb={20}>
              Ticket Booking
            </Text>
            <TabBarIcon name="ticket" size={50} />
          </HStack>
          <VStack w={"100%"} gap={30}>
            <VStack gap={5}>
              <Text ml={10} fontSize={14} color="gray">
                Email
              </Text>
              <Input
                value={email}
                onChangeText={setEmail}
                placeholder="Email"
                placeholderTextColor="darkgray"
                autoCapitalize="none"
                autoCorrect={false}
                h={48}
                p={14}
              />
            </VStack>
            <VStack gap={5}>
              <Text ml={10} fontSize={14} color="gray">
                Password
              </Text>
              <Input
                secureTextEntry
                value={password}
                onChangeText={setpassword}
                placeholder="Password"
                placeholderTextColor="darkgray"
                autoCapitalize="none"
                autoCorrect={false}
                h={48}
                p={14}
              />
            </VStack>
            <Button isLoading={isLoadingAuth} onPress={onAuthenticate}>
              {authMode === "login" ? "Login" : "Register"}
            </Button>
          </VStack>
          <Divider w={"90%"} />
          <Text onPress={onToggleAuthMode} fontSize={16} underline>
            {authMode === "login" ? "Registe new Account" : "Login to account"}
          </Text>
        </VStack>
      </ScrollView>
    </KeyboardAvoidingView>
  );
}
