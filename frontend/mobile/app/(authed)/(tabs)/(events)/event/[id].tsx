import { Input } from "@/components/input";
import { TabBarIcon } from "@/components/navigation/TabBarIcon";
import { VStack } from "@/components/VStack";
import { eventService } from "@/services/event";
import { Event } from "@/types/event";
import { router, useFocusEffect, useLocalSearchParams, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Button } from "@/components/Button";
import { Text } from "@/components/Text";
import { Alert } from "react-native";
import DateTimePicker from "@/components/DateTimePicker";

export default function EventDetailScreen() {
  const { id } = useLocalSearchParams();
  const navigation = useNavigation();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [eventData, setEventData] = useState<Event | null>(null);

  function updateField(field: keyof Event, value: string | Date) {
    setEventData((prev) => ({
      ...prev!,
      [field]: value,
    }));
  }

  const onDelete = useCallback(() => {
    if (!eventData) return;

    try {
      Alert.alert(
        "Delete Event",
        "Are you sure you want to delete this event?",
        [
          { text: "Cancel" },
          {
            text: "Delete",
            onPress: async () => {
              await eventService.deleteOne(Number(id));
              router.back();
            },
          },
        ]
      );
    } catch (error) {
      Alert.alert("Error", "failed to delete event");
    }
  }, [eventData, id]);

  async function onSubmitChanges() {
    if (!eventData) return;

    try {
      setIsSubmitting(true);
      await eventService.updateOne(
        Number(id),
        eventData.name,
        eventData.location,
        eventData.date
      );
      router.back();
    } catch (error) {
      Alert.alert("Error", "Failed to fetch event");
    } finally {
      setIsSubmitting(false);
    }
  }

  const fetchEvent = useCallback(async () => {
    try {
      const response = await eventService.getOne(Number(id));
      setEventData(response.data);
    } catch (error) {
      router.back();
    }
  }, [id, router]);


  useFocusEffect(useCallback(() => { fetchEvent(); }, []))


  useEffect(() => {
    navigation.setOptions({
      headerTitle: "",
      headerRight: () => headerRight(onDelete),
    });
  }, [navigation, onDelete]);
  return (
    <VStack m={20} flex={1} gap={30}>
      <VStack gap={5}>
        <Text ml={10} fontSize={14} bold color="gray">
          Name
        </Text>
        <Input
          value={eventData?.name}
          onChangeText={(value) => updateField("name", value)}
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
          value={eventData?.location}
          onChangeText={(value) => updateField("location", value)}
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
          onChange={(date) => updateField("date", date || new Date())}
          currentDate={new Date(eventData?.date || new Date())}
        />
      </VStack>
      <Button
        mt="auto"
        isLoading={isSubmitting}
        disabled={isSubmitting}
        onPress={onSubmitChanges}
      >
        Save Changes
      </Button>
    </VStack>
  );
}

const headerRight = (onPress: VoidFunction) => {
  return <TabBarIcon size={30} name="trash" onPress={onPress} />;
};
