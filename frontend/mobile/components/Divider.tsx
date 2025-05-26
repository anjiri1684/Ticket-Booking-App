import { defaultShortcuts, ShortcutProps } from "@/styles/shortcuts";
import { View } from "react-native";

interface DividerProps extends ShortcutProps {

}

export function Divider(props: DividerProps) { 
     return (
          <View style={[defaultShortcuts(props), {backgroundColor: 'lightgray', height: 1}]}></View>
     )
}