import { Redirect, Stack } from "expo-router"

export default function AppLayout() {
     //check from context if user is logged in

     const isLoggedIn = true

     if (!isLoggedIn) {
          return <Redirect href="/login" />
     }

     return <Stack screenOptions={{
          headerShown:false
     }}/>
}