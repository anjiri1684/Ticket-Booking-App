import { Button } from "@/components/Button";
import { Input } from "@/components/input";
import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { eventService } from "@/services/event";
import { router, useNavigation } from "expo-router";
import { useEffect, useState } from "react";
import { Alert } from "react-native";
import DateTimePicker from "@/components/DateTimePicker";

export default function NewEvent() {
  const [name, setName] = useState("");
  const [location, setLocation] = useState("");
  const [date, setDate] = useState(new Date());
  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigation = useNavigation();

  const onChangeDate = (date?: Date) => {
    setDate(date || new Date)
  }

  async function onSubmit() {
    try {
      setIsSubmitting(true);
      await eventService.createOne(name, location, date.toISOString());
      router.back();
    } catch (error) {
      Alert.alert("Error", "Failed to create Event");
    } finally {
      setIsSubmitting(false);
    }
  }

  useEffect(() => {
    navigation.setOptions({
      headerTitle: "New Event",
    });
  }, []);
  return (
    <VStack m={20} flex={1} gap={30}>
      <VStack gap={5}>
        <Text ml={10} fontSize={14} bold color="gray">
          Name
        </Text>
        <Input
          value={name}
          onChangeText={setName}
          placeholder="Enter event name"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>
      <VStack gap={5}>
        <Text ml={10} fontSize={14} bold color="gray">
          Location
        </Text>
        <Input
          value={location}
          onChangeText={setLocation}
          placeholder="Enter event name"
          placeholderTextColor="darkgray"
          h={48}
          p={14}
        />
      </VStack>
      {/* date time picker */}
      <VStack gap={5}>
        <Text ml={10} fontSize={14} color="gray">
          Date
        </Text>
        <DateTimePicker
                  onChange={onChangeDate}
                  currentDate={date}
                />
      </VStack>
      <Button
        mt="auto"
        isLoading={isSubmitting}
        disabled={isSubmitting}
        onPress={onSubmit}
      >
        Save
      </Button>
    </VStack>
  );
}
