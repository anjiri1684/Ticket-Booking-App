import { Text } from "@/components/Text";
import { VStack } from "@/components/VStack";
import { ticketService } from "@/services/ticket";
import { Ticket } from "@/types/tickets";
import { router, useFocusEffect, useLocalSearchParams, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import { Image } from "react-native";

export default function TicketDetailScreen() {
  const navigation = useNavigation()
  const { id } = useLocalSearchParams()
  const [ticket, setticket] = useState<Ticket | null>(null)
  const [qrcode, setQrcode] = useState<string | null>(null)

  const fetchTicket = async () => {
    try {
      
    } catch (error) {
      try {
        const {data} = await ticketService.getOne(Number(id))
        setticket(data.ticket)
        setQrcode(data.qrcode)
      } catch (error) {
        router.back()
      }
    }
  }

  useFocusEffect(useCallback(() => { fetchTicket() }, []))
  
  useEffect(() => {
    navigation.setOptions({headerTitle: ticket?.event.name})
  }, [navigation])

  if (!ticket) return null
     return (
       <VStack alignItems="center" m={20} p={20} gap={20} flex={1} style={{
         backgroundColor: "white",
         borderRadius: 20
       }}>
         <Text fontSize={50} bold>{ticket.event.name}</Text>
         <Text fontSize={20} bold>{ticket.event.location}</Text>
         <Text fontSize={16} color="gray">{new Date(ticket.event.date).toDateString()}</Text>
         <Image style={{borderRadius: 20, width:300}} source={{uri: `data:/image/png;base64,${qrcode}`}} />
       </VStack>
     );
}