import { useCallback, useEffect, useState } from "react";
import { Alert, FlatList, TouchableOpacity,  } from "react-native";
import { eventService } from "@/services/event";
import { VStack } from "@/components/VStack";
import { HStack } from "@/components/HStack";
import { Text } from "@/components/Text";
import { router, useFocusEffect, useNavigation } from "expo-router";
import { useAuth } from "@/context/AuthContext";
import { UserRole } from "@/types/user";
import { TabBarIcon } from "@/components/navigation/TabBarIcon";
import { Divider } from "@/components/Divider";
import { Button } from "@/components/Button";
import { ticketService } from "@/services/ticket";

export default function EventsScreen() {
  const [events, setEvent] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const { user } = useAuth();
  const navigation  = useNavigation()
  const fetchEvents = async () => {
    try {
      setIsLoading(true);

      const response = await eventService.getAll();
      setEvent(response.data);
    } catch (error) {
      Alert.alert("Error", "Failed to fetch");
      console.log(error)
    } finally {
      setIsLoading(false);
    }
  };

  useFocusEffect(
    useCallback(() =>{ fetchEvents()}, []) )

  useEffect(() => {
    navigation.setOptions({
      headerTitle: "Events",
      headerRight: user?.role === UserRole.Manager ? headerRight : null,
    });
  }, [navigation, user]);

  const onGoToEventpage = (id) => {
    if (user?.role === UserRole.Manager) {
      router.push(`/(events)/event/${id}`);
    }
  };

  const buyTicket = async (id) => {
    try {
      await ticketService.createOne(id)
      Alert.alert("Success", "You have succesfully purchasd a ticket")
    } catch (error) {
      Alert.alert("Error", "Failed to buy ticket")
    }
  }

  const options = {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
  };

  return (
    <VStack flex={1} p={20} pb={0} gap={20}>
      <HStack alignItems="center" justifyContent="center">
        <Text fontSize={18} bold>
          {events.length} Events
        </Text>
      </HStack>
      <FlatList
        data={events}
        keyExtractor={({ id }) => id.toString()}
        onRefresh={fetchEvents}
        refreshing={isLoading}
        ItemSeparatorComponent={() => <VStack h={20} />}
        renderItem={({ item: event }) => (
          <VStack
            gap={20}
            p={20}
            style={{ backgroundColor: "white", borderRadius: 20 }}
            key={event.id}
          >
            <TouchableOpacity onPress={() => onGoToEventpage(event.id)}>
              <HStack alignItems="center" justifyContent="space-between">
                <HStack alignItems="center">
                  <Text fontSize={16} bold>
                    {event.name}
                  </Text>
                  <Text fontSize={16} bold>
                    |
                  </Text>
                  <Text fontSize={16} bold>
                    {event.location}
                  </Text>
                </HStack>
                {user?.role === UserRole.Manager && (
                  <TabBarIcon size={24} name="chevron-forward" />
                )}
              </HStack>
            </TouchableOpacity>
            <Divider />
            <HStack justifyContent="space-between">
              <Text bold fontSize={16} color="gray">
                Sold: {event.totalTicketPurchased}
              </Text>
              <Text bold fontSize={16} color="green">
                Entered: {event.totalTicketsEntered}
              </Text>
            </HStack>
            {user?.role === UserRole.Attendee && (
              <VStack>
                <Button
                  variant="outlined"
                  disabled={isLoading}
                  onPress={() => buyTicket(event.id)}
                >
                  Buy Ticket
                </Button>
              </VStack>
            )}
            <Text>
              {new Date(event.date).toLocaleDateString("en-US", options)}
            </Text>
          </VStack>
        )}
      />
    </VStack>
  );
}


const headerRight = () => {
  return (
    <TabBarIcon size={32} name="add-circle-outline" onPress={() => router.push("/(events)/new")}/>
  )
}